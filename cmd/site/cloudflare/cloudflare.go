package cloudflare

import (
	"github.com/spf13/cobra"
	"github.com/wwmoraes/site/cmd/site/cloudflare/tidy"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloudflare",
		Aliases: []string{"cf"},
		Short:   "manages Cloudflare resources",
	}

	cmd.AddCommand(tidy.Command())

	return cmd
}
