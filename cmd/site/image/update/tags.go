package update

import (
	"fmt"
	"io"

	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"github.com/goccy/go-json"
)

//nolint:tagliatelle
type Tags struct {
	Artist           string `json:"Artist,omitempty"`
	Copyright        string `json:"Copyright,omitempty"`
	ImageUniqueID    string `json:"ImageUniqueID,omitempty"`
	ImageDescription string `json:"ImageDescription,omitempty"`
	Make             string `json:"Make,omitempty"`
}

func (tags *Tags) ExifIFDBuilder() (*exif.IfdBuilder, error) {
	mapping, err := exifcommon.NewIfdMappingWithStandard()
	if err != nil {
		return nil, fmt.Errorf("failed to create IFD mapping: %w", err)
	}

	builder := exif.NewIfdBuilder(
		mapping,
		exif.NewTagIndex(),
		exifcommon.IfdStandardIfdIdentity,
		exifcommon.EncodeDefaultByteOrder,
	)

	ifd0, err := exif.GetOrCreateIbFromRootIb(builder, "IFD0")
	if err != nil {
		return nil, fmt.Errorf("failed to create IFD0: %w", err)
	}

	exifIfd, err := exif.GetOrCreateIbFromRootIb(builder, "IFD/Exif")
	if err != nil {
		return nil, fmt.Errorf("failed to create IFD/Exif: %w", err)
	}

	ifd0.SetStandardWithName("Artist", tags.Artist)
	ifd0.SetStandardWithName("Copyright", tags.Copyright)
	ifd0.SetStandardWithName("ImageDescription", tags.ImageDescription)
	ifd0.SetStandardWithName("Make", tags.Make)

	exifIfd.SetStandardWithName("ImageUniqueID", tags.ImageUniqueID)

	return builder, nil
}

func loadTags(r io.ReadCloser) (*Tags, error) {
	defer r.Close()

	tags := &Tags{}

	err := json.NewDecoder(r).Decode(tags)

	return tags, err
}
