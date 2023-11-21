package touch

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/wwmoraes/go-rwfs"
	"github.com/wwmoraes/site/internal/blip"
	"github.com/wwmoraes/site/internal/frontmatter"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "touch",
		Short: "updates a page date",
		RunE:  touch,
	}

	return cmd
}

func touch(cmd *cobra.Command, args []string) error {
	contentDir, err := cmd.Flags().GetString("content")
	if err != nil {
		return err
	}

	fsys := rwfs.OSDirFS(contentDir)

	cmd.SilenceUsage = true

	for _, filename := range args {
		page, err := blip.Read(fsys, filename)
		if err != nil {
			return err
		}

		page.FrontMatter[frontmatter.Date] = time.Now().Local()

		err = blip.Update(fsys, filename, &page)
		if err != nil {
			return err
		}
	}

	return nil
}
