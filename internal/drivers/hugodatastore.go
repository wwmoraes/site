package drivers

import (
	"bytes"
	"io"

	"github.com/wwmoraes/go-rwfs"
)

type HugoDataStore struct {
	rwfs.FS
}

func (store *HugoDataStore) Update(name string, data []byte) error {
	fd, err := store.FS.OpenFile(name, rwfs.O_CREATE|rwfs.O_WRONLY|rwfs.O_TRUNC, 0o640)
	if err != nil {
		return err
	}
	defer fd.Close()

	_, err = io.Copy(fd, bytes.NewReader(data))

	return err
}
