package update

import (
	"fmt"
	"path"

	"github.com/spf13/cobra"
	"github.com/wwmoraes/go-rwfs"
	"github.com/wwmoraes/site/internal/blip"
	"github.com/wwmoraes/site/internal/frontmatter"
)

var (
	section blip.Section
	tier    blip.Tier
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update an existing radar entry",
		Long:  "updates a radar Markdown entry metadata",
		RunE:  update,
		Args:  cobra.ExactArgs(1),
	}

	flags := cmd.Flags()
	flags.VarP(&section, "quadrant", "q", "blip section/quadrant (languages, platforms, techniques or tools)")
	flags.VarP(&tier, "tier", "t", "blip usage stage (adopt, assess, hold or trial)")

	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	contentDir, err := cmd.Flags().GetString("content")
	if err != nil {
		return err
	}

	sectionName, err := cmd.Flags().GetString("section")
	if err != nil {
		return err
	}

	if section.String() == "" {
		return fmt.Errorf("section not set")
	}

	if tier.String() == "" {
		return fmt.Errorf("tier not set")
	}

	name := args[0]
	filename := name + ".md"
	fsys := rwfs.OSDirFS(path.Join(contentDir, sectionName))

	cmd.SilenceUsage = true

	cmd.Printf("updating blip @ %s", path.Join(contentDir, sectionName, filename))

	page, err := blip.Read(fsys, filename)
	if err != nil {
		return err
	}

	page.FrontMatter[frontmatter.RadarSection] = section.String()
	page.FrontMatter[frontmatter.RadarTier] = tier.String()

	return blip.Update(fsys, filename, page)
}
