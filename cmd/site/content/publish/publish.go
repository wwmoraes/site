package publish

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/wwmoraes/go-rwfs"
	"github.com/wwmoraes/site/internal/blip"
	"github.com/wwmoraes/site/internal/frontmatter"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "publish",
		Short: "updates a page publish date",
		RunE:  publish,
	}

	return cmd
}

func publish(cmd *cobra.Command, args []string) error {
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

		date, err := time.Parse(time.RFC3339Nano, page.FrontMatter[frontmatter.Date].(string))
		if err != nil {
			return err
		}

		page.FrontMatter[frontmatter.Date] = date.Local()
		page.FrontMatter[frontmatter.PublishDate] = time.Now().Local()
		delete(page.FrontMatter, frontmatter.Draft)

		err = blip.Update(fsys, filename, page)
		if err != nil {
			return err
		}
	}

	return nil
}
