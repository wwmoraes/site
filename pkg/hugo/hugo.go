package hugo

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/config/allconfig"
	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/hugofs"
	"github.com/gohugoio/hugo/hugolib"
	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/spf13/afero"
)

// why not use the github.com/gohugoio/hugo/config* libs? It is a huge mess of
// command mixed with IO and configuration. Way too coupled with their CLI.

var (
	KeyNotFoundError   = errors.New("key not set")
	KeyConversionError = errors.New("failed to convert key type")
)

func New() (*hugolib.HugoSites, error) {
	environment := "development"
	configDir := "config"

	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	provider, _, err := config.LoadConfigFromDir(afero.NewOsFs(), configDir, environment)
	if err != nil {
		return nil, err
	}

	provider.Set("renderToDisk", true)
	provider.Set("workingDir", pwd)
	provider.Set("publishDir", "public")

	fs := hugofs.NewFromSourceAndDestination(hugofs.Os, hugofs.Os, provider)
	configs, err := allconfig.LoadConfig(allconfig.ConfigSourceDescriptor{
		Flags:       provider,
		Fs:          hugofs.Os,
		Filename:    "config.yaml",
		ConfigDir:   configDir,
		Environment: environment,
	})
	if err != nil {
		return nil, err
	}

	sites, err := hugolib.NewHugoSites(deps.DepsCfg{
		Configs: configs,
		Fs:      fs,
	})
	if err != nil {
		return nil, err
	}

	return sites, sites.Init()
}

func ParsePage(pagePath string) (pageparser.ContentFrontMatter, error) {
	fd, err := os.OpenFile(pagePath, os.O_RDONLY, 0640)
	if err != nil {
		return pageparser.ContentFrontMatter{}, err
	}

	return pageparser.ParseFrontMatterAndContent(fd)
}

func GetString(matter *pageparser.ContentFrontMatter, key string) (string, error) {
	v, ok := matter.FrontMatter[key]
	if !ok {
		return "", KeyNotFoundError
	}

	vv, ok := v.(string)
	if !ok {
		return "", KeyConversionError
	}

	return vv, nil
}

func GetBool(matter *pageparser.ContentFrontMatter, key string) (bool, error) {
	v, ok := matter.FrontMatter[key]
	if !ok {
		return false, KeyNotFoundError
	}

	vv, ok := v.(bool)
	if !ok {
		return false, KeyConversionError
	}

	return vv, nil
}

func GetContentPath(content string) (string, error) {
	stat, err := os.Stat(content)

	if errors.Is(err, fs.ErrPermission) {
		return "", err
	}

	if errors.Is(err, fs.ErrNotExist) {
		return contentFile(fmt.Sprintf("%s.md", content))
	}

	if stat.IsDir() {
		return contentDir(content)
	}

	if err != nil {
		return "", err
	}

	// passthrough in case content is a valid path to a file
	return content, nil
}

func contentFile(content string) (string, error) {
	stat, err := os.Stat(content)

	if errors.Is(err, fs.ErrPermission) {
		return "", err
	}

	if errors.Is(err, fs.ErrNotExist) {
		return "", err
	}

	if !stat.Mode().IsRegular() {
		return "", fmt.Errorf("%w: %s ins't a regular file", fs.ErrInvalid, content)
	}

	return content, nil
}

func contentDir(content string) (string, error) {
	filePath, err := contentFile(path.Join(content, "index.md"))
	if len(filePath) > 0 {
		return filePath, err
	}

	return contentFile(path.Join(content, "_index.md"))
}
