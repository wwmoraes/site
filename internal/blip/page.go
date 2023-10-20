package blip

import (
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
	fd, err := fsys.OpenFile(name, rwfs.O_CREATE|rwfs.O_EXCL|rwfs.O_WRONLY, 0640)
	if err != nil {
		return err
	}
	defer fd.Close()

	return hugo.WritePage(fd, page)
}

func Update(fsys rwfs.FS, name string, page *pageparser.ContentFrontMatter) error {
	fd, err := fsys.OpenFile(name, rwfs.O_WRONLY|rwfs.O_TRUNC, 0640)
	if err != nil {
		return err
	}
	defer fd.Close()

	return hugo.WritePage(fd, page)
}
