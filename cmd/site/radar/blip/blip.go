package blip

import (
	"github.com/spf13/cobra"
	"github.com/wwmoraes/site/cmd/site/radar/blip/create"
	"github.com/wwmoraes/site/cmd/site/radar/blip/update"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blip",
		Short: "manages radar blips (entries)",
		Long:  "creates and updates blip entries",
	}

	cmd.AddCommand(create.Command())
	cmd.AddCommand(update.Command())

	return cmd
}
