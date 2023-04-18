package update

import (
	"github.com/spf13/cobra"
	"github.com/wwmoraes/site/cmd/site/update/goodreads"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "static data update commands",
	}

	cmd.AddCommand(goodreads.New())

	return cmd
}
