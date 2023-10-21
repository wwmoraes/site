package update

import (
	"github.com/spf13/cobra"
	"github.com/wwmoraes/site/cmd/site/update/github"
	"github.com/wwmoraes/site/cmd/site/update/goodreads"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "static data update commands",
	}

	cmd.AddCommand(github.Command())
	cmd.AddCommand(goodreads.Command())

	return cmd
}
