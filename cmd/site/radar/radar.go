package radar

import (
	"github.com/spf13/cobra"
	"github.com/wwmoraes/site/cmd/site/radar/update"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "radar",
		Short: "updates radar blips",
		Long:  "calculates the position of each radar entry and generate their blips in the SVG",
	}

	flags := cmd.PersistentFlags()
	flags.StringP("content", "c", "content", "filesystem path to content directory")
	flags.StringP("section", "s", "radar", "section name that contains the blips files and the radar SVG")
	flags.StringP("template", "t", "radar.svg.tmpl", "filename of the SVG template within the section folder")
	flags.StringP("output", "o", "radar.svg", "output filename of the generated SVG")
	flags.IntP("radius", "r", 500, "SVG radius in user units (viewbox)")
	flags.Int("size", 800, "SVG size in px (will apply to both width and height; must be <= radius * 2)")

	cmd.AddCommand(update.Command())

	return cmd
}
