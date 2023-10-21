package data

import (
	"github.com/spf13/cobra"
	"github.com/wwmoraes/site/cmd/site/data/update"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "data",
		Short: "handles site data",
	}

	cmd.AddCommand(update.Command())

	return cmd
}
