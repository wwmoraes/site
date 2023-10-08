package update

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/gohugoio/hugo/parser"
	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/gohugoio/hugo/tpl/crypto"
	"github.com/wwmoraes/go-rwfs"
	"github.com/wwmoraes/site/pkg/hugo"
)

const (
	maxAngle     = 90.0
	baseRadius   = 500.0
	blipDiameter = 10 * 2
)

var (
	cartesianFactors = map[string][]float64{
		"techniques": {-1, -1},
		"languages":  {1, -1},
		"platforms":  {-1, 1},
		"tools":      {1, 1},
	}

	offsetFactors = map[string]float64{
		"adopt":  0.0,
		"trial":  0.5,
		"assess": 0.70,
		"hold":   0.86,
	}

	radiusFactors = map[string]float64{
		"adopt":  0.5,
		"trial":  0.3,
		"assess": 0.14,
		"hold":   0.06,
	}
)

type Blip struct {
	Page         *pageparser.ContentFrontMatter
	Filename     string
	RelPermalink string
}

// TODO pass radar configuration (size, section zones, rim radius, etc)
func GetUpdatedBlips(fsys fs.FS, section string, buffer int) (<-chan *Result[Blip], error) {
	blips := make(chan *Result[Blip], buffer)

	go func() {
		defer close(blips)

		entries, err := fs.ReadDir(fsys, "/")
		if err != nil {
			blips <- errorResult[Blip](fmt.Errorf("failed to read blips directory: %w", err))
			return
		}

		for index, entry := range entries {
			if filepath.Ext(entry.Name()) != ".md" {
				continue
			}

			if filepath.Base(entry.Name()) == "_index.md" {
				continue
			}

			fd, err := fsys.Open(entry.Name())
			if err != nil {
				blips <- errorResult[Blip](fmt.Errorf("failed to open blip %s: %w", entry.Name(), err))
				continue
			}
			defer fd.Close()

			page, err := pageparser.ParseFrontMatterAndContent(fd)
			if err != nil {
				blips <- errorResult[Blip](fmt.Errorf("failed to parse blip %s: %w", entry.Name(), err))
				continue
			}
			fd.Close() // early close to avoid hanging descriptors during the loop

			title, err := hugo.GetString(&page, "title")
			if err != nil {
				blips <- errorResult[Blip](fmt.Errorf("failed to get title of blip %s: %w", entry.Name(), err))
				continue
			}

			radarSection, err := hugo.GetString(&page, "radarSection")
			if err != nil {
				blips <- errorResult[Blip](fmt.Errorf("failed to get section of blip %s: %w", entry.Name(), err))
				continue
			}

			radarTier, err := hugo.GetString(&page, "radarTier")
			if err != nil {
				blips <- errorResult[Blip](fmt.Errorf("failed to get tier of blip %s: %w", entry.Name(), err))
				continue
			}

			radiusFactor, ok := radiusFactors[radarTier]
			if !ok {
				blips <- errorResult[Blip](fmt.Errorf("unknown blip '%s' tier '%s'", entry.Name(), radarTier))
				continue
			}

			offsetFactor, ok := offsetFactors[radarTier]
			if !ok {
				blips <- errorResult[Blip](fmt.Errorf("unknown blip '%s' tier '%s'", entry.Name(), radarTier))
				continue
			}

			fnv, err := crypto.New().FNV32a(title)
			if err != nil {
				blips <- errorResult[Blip](fmt.Errorf("failed to hash blip %s: %w", entry.Name(), err))
				continue
			}

			name := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
			angle := math.Mod(float64(fnv), maxAngle)
			radius := (baseRadius * offsetFactor) + math.Max(blipDiameter, math.Mod(float64(fnv), baseRadius*radiusFactor))
			x := math.Max(blipDiameter, math.Round(math.Abs(math.Cos(angle))*radius)) * cartesianFactors[radarSection][0]
			y := math.Max(blipDiameter, math.Round(math.Abs(math.Sin(angle))*radius)) * cartesianFactors[radarSection][1]

			page.FrontMatter["radarX"] = x
			page.FrontMatter["radarY"] = y
			page.FrontMatter["radarIndex"] = index

			blips <- okResult[Blip](&Blip{
				Page:         &page,
				Filename:     entry.Name(),
				RelPermalink: fmt.Sprintf("/%s/%s", section, name),
			})
		}
	}()

	return blips, nil
}

func (blip *Blip) WriteFile(fsys rwfs.FS) error {
	fd, err := fsys.OpenFile(blip.Filename, os.O_TRUNC|os.O_WRONLY, 0640)
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
