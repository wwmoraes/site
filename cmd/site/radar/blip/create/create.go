package create

import (
	"fmt"
	"path"

	"github.com/gohugoio/hugo/parser/metadecoders"
	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/spf13/cobra"
	"github.com/wwmoraes/go-rwfs"
	"github.com/wwmoraes/site/internal/blip"
	"github.com/wwmoraes/site/internal/frontmatter"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	section blip.Section
	tier    blip.Tier
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create a new radar entry",
		Long:  "generates a blank radar Markdown with the minimal metadata",
		RunE:  create,
		Args:  cobra.ExactArgs(1),
	}

	flags := cmd.Flags()
	flags.VarP(&section, "quadrant", "q", "blip section/quadrant (languages, platforms, techniques or tools)")
	flags.VarP(&tier, "tier", "t", "blip usage stage (adopt, assess, hold or trial)")

	return cmd
}

func create(cmd *cobra.Command, args []string) error {
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

	cmd.Printf("creating blip @ %s", path.Join(contentDir, sectionName, filename))

	page := pageparser.ContentFrontMatter{
		Content: []byte("\nTODO justification\n"),
		FrontMatter: map[string]any{
			frontmatter.Build: map[string]any{
				frontmatter.List:   true,
				frontmatter.Render: false,
			},
			frontmatter.Description: "lorem ipsum",
			// frontmatter.Draft:           true,
			frontmatter.RadarSection:    section.String(),
			frontmatter.RadarTier:       tier.String(),
			frontmatter.TableOfContents: false,
			frontmatter.Title:           cases.Title(language.AmericanEnglish).String(name),
		},
		FrontMatterFormat: metadecoders.YAML,
	}

	return blip.Create(fsys, filename, &page)
}
