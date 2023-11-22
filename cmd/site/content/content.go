package content

import (
	"github.com/spf13/cobra"
	"github.com/wwmoraes/site/cmd/site/content/publish"
	"github.com/wwmoraes/site/cmd/site/content/touch"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "content",
		Short: "handles site content",
	}

	flags := cmd.PersistentFlags()
	flags.StringP("content", "c", "content", "filesystem path to the Hugo content directory")

	cmd.AddCommand(publish.Command())
	cmd.AddCommand(touch.Command())

	return cmd
}
