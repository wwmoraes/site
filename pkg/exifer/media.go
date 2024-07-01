package exifer

import (
	"errors"
	"fmt"
	"io"

	"github.com/dsoprea/go-exif/v3"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	pngstructure "github.com/dsoprea/go-png-image-structure/v2"
	riimage "github.com/dsoprea/go-utility/v2/image"
	"github.com/gabriel-vasile/mimetype"
)

// MediaParser represents a media format parser that produces an EXIF-enabled
// MediaContext value.
type MediaParser interface {
	Parse(rs io.ReadSeeker, size int) (riimage.MediaContext, error)
}

// MetaMediaParser contains a MediaParser per MIME type. It implements
// MediaParser as well to allow parsing supported types after inspecting the
// descriptor header.
type MetaMediaParser map[string]MediaParser

// ExifMediaReader represents a media type-specific reader that converts a media
// context to an EXIF media value.
type ExifMediaReader interface {
	GetExifMedia(context riimage.MediaContext) (ExifMedia, bool)
}

// MetaMediaReader contains readers for multiple MIME media types. It implements
// ExifMediaReader to try each reader in order, returning the results of the
// first match.
type MetaMediaReader []ExifMediaReader

type ExifMediaReaderFn func(mediaContext riimage.MediaContext) (ExifMedia, bool)

// ExifMedia exports the write and EXIF methods needed to update a media file
// EXIF data and to write it back to storage.
type ExifMedia interface {
	io.WriterTo

	ConstructExifBuilder() (*exif.IfdBuilder, error)
	SetExif(ib *exif.IfdBuilder) error
}

// ExifMediaPng Implements ExifMedia for the underlying PNG ChunkSlice. Its
// original WriteTo implementation does not comply with the io.WriterTo
// interface.
type ExifMediaPng struct {
	*pngstructure.ChunkSlice
}

// ExifMediaJpeg Implements ExifMedia for the underlying JPEG SegmentList. It
// implements a Write method that's a misnomer, as it acts as a WriteTo. Besides
// naming, its signature does not comply with the io.WriterTo interface.
type ExifMediaJpeg struct {
	*jpegstructure.SegmentList
}

// MetaMediaHandler contains both parsers and readers to handle updating EXIF
// metadata on supported media types.
type MetaMediaHandler struct {
	Parser MetaMediaParser
	Reader MetaMediaReader
}

// NewMetaMediaParser creates a MIME-based media parser.
//
// It supports only JPEG and PNG.
func NewMetaMediaParser() MetaMediaParser {
	parser := make(MetaMediaParser, 2)

	parser["image/jpeg"] = jpegstructure.NewJpegMediaParser()
	parser["image/png"] = pngstructure.NewPngMediaParser()

	return parser
}

// NewMetaMediaReader creates an adapter for the dsoprea MediaContext types.
// They don't have a common interface nor do they expose write methods on the
// existing one.
//
// To solve that, a new interface ExifMedia allows internal struct types to
// provide the writing methods to update EXIF.
//
// It supports only JPEG and PNG.
func NewMetaMediaReader() MetaMediaReader {
	reader := make(MetaMediaReader, 0, 2)

	reader = append(reader, ExifMediaReaderFn(func(context riimage.MediaContext) (ExifMedia, bool) {
		media, ok := context.(*pngstructure.ChunkSlice)
		if !ok {
			return nil, false
		}

		return &ExifMediaPng{media}, true
	}))
	reader = append(reader, ExifMediaReaderFn(func(context riimage.MediaContext) (ExifMedia, bool) {
		mediaJpeg, ok := context.(*jpegstructure.SegmentList)
		if !ok {
			return nil, false
		}

		return &ExifMediaJpeg{mediaJpeg}, true
	}))

	return reader
}

func NewMetaMediaHandler() *MetaMediaHandler {
	return &MetaMediaHandler{
		Parser: NewMetaMediaParser(),
		Reader: NewMetaMediaReader(),
	}
}

func (meta MetaMediaParser) Parse(rs io.ReadSeeker, size int) (riimage.MediaContext, error) {
	_, err := rs.Seek(0, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("failed to seek reader to its start: %w", err)
	}

	readerType, err := mimetype.DetectReader(rs)
	if err != nil {
		return nil, fmt.Errorf("failed to detect file type: %w", err)
	}

	mediaParser, ok := meta[readerType.String()]
	if !ok {
		return nil, fmt.Errorf("unable to parse image type %q: %w", readerType.String(), errors.ErrUnsupported)
	}

	_, err = rs.Seek(0, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("failed to seek reader to its start: %w", err)
	}

	return mediaParser.Parse(rs, size)
}

func (fn ExifMediaReaderFn) GetExifMedia(mediaContext riimage.MediaContext) (ExifMedia, bool) {
	return fn(mediaContext)
}

func (meta MetaMediaReader) GetExifMedia(mediaContext riimage.MediaContext) (ExifMedia, bool) {
	for _, extractor := range meta {
		media, ok := extractor.GetExifMedia(mediaContext)
		if ok {
			return media, true
		}
	}

	return nil, false
}

func (media *ExifMediaPng) WriteTo(w io.Writer) (int64, error) {
	return 0, media.ChunkSlice.WriteTo(w)
}

func (media *ExifMediaJpeg) WriteTo(w io.Writer) (int64, error) {
	return 0, media.SegmentList.Write(w)
}

func (handler *MetaMediaHandler) UpdateImageWithEXIF(fd Descriptor, size int, ib *exif.IfdBuilder) error {
	mediaContext, err := handler.Parser.Parse(fd, size)
	if err != nil {
		return fmt.Errorf("failed to parse image: %w", err)
	}

	media, ok := handler.Reader.GetExifMedia(mediaContext)
	if !ok {
		return fmt.Errorf("failed to get EXIF media: %w", errors.ErrUnsupported)
	}

	err = media.SetExif(ib)
	if err != nil {
		return fmt.Errorf("failed to set EXIF: %w", err)
	}

	_, err = fd.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("failed to seek file: %w", err)
	}

	err = fd.Truncate(0)
	if err != nil {
		return fmt.Errorf("failed to truncate file: %w", err)
	}

	_, err = media.WriteTo(fd)
	if err != nil {
		return fmt.Errorf("failed to write image: %w", err)
	}

	return nil
}
