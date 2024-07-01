package update

import (
	"io/fs"
	"log"
	"sync"

	"github.com/spf13/cobra"
	"github.com/wwmoraes/go-rwfs"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "generates the radar SVG",
		Long:  "calculates the position of each radar entry and generate their blips in the SVG",
		RunE:  update,
	}

	flags := cmd.Flags()
	flags.StringP("template", "t", "radar.svg.tmpl", "filename of the SVG template within the section folder")
	flags.StringP("output", "o", "radar.svg", "output filename of the generated SVG")
	flags.IntP("radius", "r", 500, "SVG radius in user units (viewbox)")
	flags.Int("size", 800, "SVG size in px (will apply to both width and height; must be <= radius * 2)")

	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	settings, err := getSettings(cmd)
	if err != nil {
		return err
	}

	cmd.SilenceUsage = true

	log.Println("retrieving blips")

	blips, err := GetUpdatedBlips(settings.SectionFS, settings.SectionName, 2) //nolint:mnd
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for blip := range blips {
		if blip.Error != nil {
			log.Println(blip.Error)

			continue
		}

		settings.TemplateData.Blips = append(settings.TemplateData.Blips, blip.Value)

		wg.Add(1)

		go func(blip *Blip) {
			defer wg.Done()

			log.Println("updating blip:", blip.Filename)

			err := blip.WriteFile(settings.SectionFS)
			if err != nil {
				log.Println(err)
			}
		}(blip.Value)
	}

	templateData, err := fs.ReadFile(settings.SectionFS, settings.TemplateFilename)
	if err != nil {
		return err
	}

	fd, err := settings.SectionFS.OpenFile(settings.OutputFilename, rwfs.O_CREATE|rwfs.O_TRUNC|rwfs.O_WRONLY, 0o640)
	if err != nil {
		return err
	}

	log.Println("generating radar")

	err = generateRadar(string(templateData), fd, &settings.TemplateData)
	if err != nil {
		return err
	}

	wg.Wait()

	return nil
}
