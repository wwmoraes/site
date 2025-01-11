package update

import (
	"fmt"
	"io"
	"reflect"

	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	exifundefined "github.com/dsoprea/go-exif/v3/undefined"
	"github.com/goccy/go-json"
)

//nolint:tagliatelle
type Tags struct {
	Artist           string         `json:"Artist,omitempty"`
	Copyright        string         `json:"Copyright,omitempty"`
	ImageDescription string         `json:"ImageDescription,omitempty"`
	ImageUniqueID    string         `json:"ImageUniqueID,omitempty"`
	Make             string         `json:"Make,omitempty"`
	DocumentName     string         `json:"DocumentName,omitempty"`
	UserComment      map[string]any `json:"UserComment,omitempty"`
}

func (tags *Tags) ExifIFDBuilder() (*exif.IfdBuilder, error) {
	mapping, err := exifcommon.NewIfdMappingWithStandard()
	if err != nil {
		return nil, fmt.Errorf("failed to create IFD mapping: %w", err)
	}

	tagIndex := exif.NewTagIndex()
	if err := exif.LoadStandardTags(tagIndex); err != nil {
		return nil, fmt.Errorf("failed to load standard tags: %v", err)
	}

	builder := exif.NewIfdBuilder(
		mapping,
		tagIndex,
		exifcommon.IfdStandardIfdIdentity,
		exifcommon.EncodeDefaultByteOrder,
	)

	ifd0, err := exif.GetOrCreateIbFromRootIb(builder, exifcommon.IfdStandardIfdIdentity.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to create IFD0: %w", err)
	}

	// exifIfd, err := exif.GetOrCreateIbFromRootIb(builder, "IFD/Exif")
	exifIfd, err := exif.GetOrCreateIbFromRootIb(builder, exifcommon.IfdExifStandardIfdIdentity.String())
	if err != nil {
		return nil, fmt.Errorf("failed to create IFD/Exif: %w", err)
	}

	setStandardWithNameIfNotEmpty(ifd0, "Artist", tags.Artist)
	setStandardWithNameIfNotEmpty(ifd0, "Copyright", tags.Copyright)
	setStandardWithNameIfNotEmpty(ifd0, "DocumentName", tags.DocumentName)
	setStandardWithNameIfNotEmpty(ifd0, "ImageDescription", tags.ImageDescription)
	setStandardWithNameIfNotEmpty(exifIfd, "ImageUniqueID", tags.ImageUniqueID)

	if len(tags.UserComment) > 0 {
		userComment, err := json.Marshal(tags.UserComment)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal user comment: %w", err)
		}

		exifIfd.SetStandardWithName("UserComment", exifundefined.Tag9286UserComment{
			EncodingType:  exifundefined.TagUndefinedType_9286_UserComment_Encoding_ASCII,
			EncodingBytes: userComment,
		})
	}

	return builder, nil
}

func loadTags(r io.ReadCloser) (*Tags, error) {
	defer r.Close()

	tags := &Tags{}

	err := json.NewDecoder(r).Decode(tags)

	return tags, err
}

func setStandardWithNameIfNotEmpty(ifd *exif.IfdBuilder, tagName string, value any) error {
	if reflect.ValueOf(value).IsZero() {
		return nil
	}

	return ifd.SetStandardWithName(tagName, value)
}
