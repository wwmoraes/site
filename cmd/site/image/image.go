package image

import (
	"github.com/spf13/cobra"
	"github.com/wwmoraes/site/cmd/site/image/inspect"
	"github.com/wwmoraes/site/cmd/site/image/update"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "image",
		Short: "handles site images",
	}

	cmd.AddCommand(inspect.Command())
	cmd.AddCommand(update.Command())

	return cmd
}
