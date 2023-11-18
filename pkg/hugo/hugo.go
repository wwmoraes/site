package hugo

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"

	"github.com/gohugoio/hugo/parser"
	"github.com/gohugoio/hugo/parser/pageparser"
)

// why not use the github.com/gohugoio/hugo/config* libs? It is a huge mess of
// command mixed with IO and configuration. Way too coupled with their CLI.

var (
	KeyNotFoundError   = errors.New("key not set")
	KeyConversionError = errors.New("failed to convert key type")
)

func ParsePage(pagePath string) (pageparser.ContentFrontMatter, error) {
	fd, err := os.OpenFile(pagePath, os.O_RDONLY, 0o640)
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

func ReadPage(r io.ReadCloser) (pageparser.ContentFrontMatter, error) {
	defer r.Close()

	return pageparser.ParseFrontMatterAndContent(r)
}

func WritePage(w io.WriteCloser, page *pageparser.ContentFrontMatter) error {
	defer w.Close()

	err := parser.InterfaceToFrontMatter(page.FrontMatter, page.FrontMatterFormat, w)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bytes.NewReader(page.Content))
	if err != nil {
		return err
	}

	return nil
}
