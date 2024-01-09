package create

import (
	"fmt"
	"path"
	"strings"

	"github.com/charmbracelet/huh"
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
	name    string
	section blip.Section
	tier    blip.Tier
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "create a new radar entry",
		Long:    "generates a blank radar Markdown with the minimal metadata",
		PreRunE: preCreate,
		RunE:    create,
		Args:    cobra.MaximumNArgs(1),
	}

	flags := cmd.Flags()
	flags.VarP(&section, "quadrant", "q", "blip section/quadrant (languages, platforms, techniques or tools)")
	flags.VarP(&tier, "tier", "t", "blip usage stage (adopt, assess, hold or trial)")

	return cmd
}

func preCreate(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		name = args[0]
	}

	err := huh.NewForm(
		huh.NewGroup(
			newHuhSelectEnum(
				"Quadrant",
				&section,
				blip.Languages,
				blip.Platforms,
				blip.Techniques,
				blip.Tools,
			),
			newHuhSelectEnum(
				"Tier",
				&tier,
				blip.Adopt,
				blip.Trial,
				blip.Assess,
				blip.Hold,
			),
			huh.NewInput().Title("Name").Value(&name),
		),
	).Run()
	if err != nil {
		return err
	}

	if section.String() == "" {
		return fmt.Errorf("section not set")
	}

	if tier.String() == "" {
		return fmt.Errorf("tier not set")
	}

	if len(name) <= 0 {
		return fmt.Errorf("name not set")
	}

	return nil
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

	slug := strings.ReplaceAll(strings.ToLower(name), " ", "-")
	filename := slug + ".md"
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

type selectEnum interface {
	comparable
	fmt.Stringer
}

func newHuhSelectEnum[T selectEnum](title string, value *T, options ...T) *huh.Select[T] {
	huhOptions := make([]huh.Option[T], len(options))
	for index, option := range options {
		huhOptions[index] = huh.NewOption[T](option.String(), option).Selected(*value == option)
	}

	return huh.NewSelect[T]().Title(title).Value(value).Options(huhOptions...)
}
