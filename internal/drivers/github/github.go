package github

import (
	"io"

	"github.com/goccy/go-json"
	"github.com/wwmoraes/go-rwfs"
)

type Variables map[string]interface{}

func writeFile(fsys rwfs.FS, name string, res any) error {
	fd, err := fsys.OpenFile(name, rwfs.O_WRONLY|rwfs.O_CREATE|rwfs.O_TRUNC, 0o640)
	if err != nil {
		return err
	}

	defer fd.Close()

	_, err = marshal(res, fd)

	return err
}

func marshal(v any, w io.Writer) (int64, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return 0, err
	}

	n, err := w.Write(data)

	return int64(n), err
}
