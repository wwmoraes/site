package update

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/gohugoio/hugo/parser"
	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/wwmoraes/go-rwfs"
	"github.com/wwmoraes/site/internal/blip"
	"github.com/wwmoraes/site/internal/frontmatter"
	"github.com/wwmoraes/site/pkg/hugo"
)

const (
	radarRadius   = 500.0
	quadrantAngle = 90.0
	blipRadius    = 10.0
)

type Blip struct {
	Page         *pageparser.ContentFrontMatter
	Filename     string
	RelPermalink string
}

func GetUpdatedBlips(fsys rwfs.FS, section string, buffer int) (<-chan *Result[Blip], error) {
	// TODO pass radar configuration (size, section zones, rim radius, etc).
	radarParameters, err := blip.NewRadar(blip.RadarOptions{
		Radius:     radarRadius,
		Angle:      quadrantAngle,
		BlipRadius: blipRadius,
		Proportion: [4]float64{
			50.0 / 100.0,
			20.0 / 100.0,
			16.0 / 100.0,
			14.0 / 100.0,
		},
		Tiers:     [4]string{"adopt", "trial", "assess", "hold"},
		Quadrants: [4]string{"techniques", "languages", "tools", "platforms"},
	})
	if err != nil {
		return nil, err
	}

	blips := make(chan *Result[Blip], buffer)

	go func() {
		defer close(blips)

		entries, err := fs.ReadDir(fsys, "/")
		if err != nil {
			blips <- errorResult[Blip](fmt.Errorf("failed to read blips directory: %w", err))

			return
		}

		for index, entry := range entries {
			if !blip.IsSourceFile(entry.Name()) {
				continue
			}

			page, err := blip.Read(fsys, entry.Name())
			if err != nil {
				blips <- errorResult[Blip](fmt.Errorf("failed to read blip %s: %w", entry.Name(), err))

				continue
			}

			err = updateBlipPage(&page, radarParameters, index)
			if err != nil {
				blips <- errorResult[Blip](fmt.Errorf("failed to update blip %s: %w", entry.Name(), err))

				continue
			}

			slug := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))

			blips <- okResult[Blip](&Blip{
				Page:         &page,
				Filename:     entry.Name(),
				RelPermalink: fmt.Sprintf("/%s/%s", section, slug),
			})
		}
	}()

	return blips, nil
}

func updateBlipPage(page *pageparser.ContentFrontMatter, params *blip.RadarParameters, index int) error {
	quadrant, err := hugo.GetString(page, frontmatter.RadarSection)
	if err != nil {
		return err
	}

	tier, err := hugo.GetString(page, frontmatter.RadarTier)
	if err != nil {
		return err
	}

	title, err := hugo.GetString(page, "title")
	if err != nil {
		return err
	}

	position, err := blip.CalculatePosition(params, &blip.BlipParameters{
		Quadrant: quadrant,
		Tier:     tier,
		Title:    title,
	})
	if err != nil {
		return err
	}

	page.FrontMatter[frontmatter.RadarX] = position[0]
	page.FrontMatter[frontmatter.RadarY] = position[1]
	page.FrontMatter[frontmatter.RadarIndex] = index

	return nil
}

func (blip *Blip) WriteFile(fsys rwfs.FS) error {
	fd, err := fsys.OpenFile(blip.Filename, os.O_TRUNC|os.O_WRONLY, 0o640)
	if err != nil {
		return err
	}
	defer fd.Close()

	err = parser.InterfaceToFrontMatter(blip.Page.FrontMatter, blip.Page.FrontMatterFormat, fd)
	if err != nil {
		return err
	}

	_, err = io.Copy(fd, bytes.NewReader(blip.Page.Content))
	if err != nil {
		return err
	}

	return nil
}
