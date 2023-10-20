package radar

import (
	"github.com/spf13/cobra"
	"github.com/wwmoraes/site/cmd/site/radar/blip"
	"github.com/wwmoraes/site/cmd/site/radar/update"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "radar",
		Short: "updates radar blips",
		Long:  "calculates the position of each radar entry and generate their blips in the SVG",
	}

	flags := cmd.PersistentFlags()
	flags.StringP("content", "c", "content", "filesystem path to the Hugo content directory")
	flags.StringP("section", "s", "radar", "Hugo section name that contains the blips files and the radar SVG")

	cmd.AddCommand(blip.Command())
	cmd.AddCommand(update.Command())

	return cmd
}
