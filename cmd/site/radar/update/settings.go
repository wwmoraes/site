package update

import (
	"path"

	"github.com/spf13/cobra"
	"github.com/wwmoraes/go-rwfs"
)

type Settings struct {
	SectionFS        rwfs.FS
	OutputFilename   string
	SectionName      string
	TemplateFilename string
	TemplateData     RadarTemplate
}

func getSettings(cmd *cobra.Command) (*Settings, error) {
	result := &Settings{}
	flags := cmd.Flags()

	var err error

	result.OutputFilename, err = flags.GetString("output")
	if err != nil {
		return nil, err
	}

	result.SectionName, err = flags.GetString("section")
	if err != nil {
		return nil, err
	}

	result.TemplateFilename, err = flags.GetString("template")
	if err != nil {
		return nil, err
	}

	result.TemplateData.Radius, err = flags.GetInt("radius")
	if err != nil {
		return nil, err
	}

	result.TemplateData.Size, err = flags.GetInt("size")
	if err != nil {
		return nil, err
	}

	contentPath, err := flags.GetString("content")
	if err != nil {
		return nil, err
	}

	result.SectionFS = rwfs.OSDirFS(path.Join(contentPath, result.SectionName))

	return result, nil
}
