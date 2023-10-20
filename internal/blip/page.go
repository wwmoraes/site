package blip

import (
	"path/filepath"

	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/wwmoraes/go-rwfs"
	"github.com/wwmoraes/site/pkg/hugo"
)

func Read(fsys rwfs.FS, name string) (pageparser.ContentFrontMatter, error) {
	fd, err := fsys.Open(name)
	if err != nil {
		return pageparser.ContentFrontMatter{}, err
	}
	defer fd.Close()

	return hugo.ReadPage(fd)
}

func Create(fsys rwfs.FS, name string, page *pageparser.ContentFrontMatter) error {
	fd, err := fsys.OpenFile(name, rwfs.O_CREATE|rwfs.O_EXCL|rwfs.O_WRONLY, 0o640)
	if err != nil {
		return err
	}
	defer fd.Close()

	return hugo.WritePage(fd, page)
}

func Update(fsys rwfs.FS, name string, page *pageparser.ContentFrontMatter) error {
	fd, err := fsys.OpenFile(name, rwfs.O_WRONLY|rwfs.O_TRUNC, 0o640)
	if err != nil {
		return err
	}
	defer fd.Close()

	return hugo.WritePage(fd, page)
}

func IsSourceFile(name string) bool {
	if filepath.Ext(name) != ".md" {
		return false
	}

	if filepath.Base(name) == "_index.md" {
		return false
	}

	return true
}
