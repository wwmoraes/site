package github

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"

	"github.com/spf13/cobra"
	"github.com/wwmoraes/go-rwfs"
)

type Variables map[string]interface{}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "github",
		Short: "fetch GitHub user data",
		RunE:  cmdFetch,
	}

	flags := cmd.Flags()
	flags.Int("count", 10, "maximum number of entries for each data source")
	flags.String("path", "data/github/recents", "base directory where the JSON results will be stored at")

	return cmd
}

func cmdFetch(cmd *cobra.Command, args []string) error {
	count, err := cmd.Flags().GetInt("count")
	if err != nil {
		return err
	}

	basePath, err := cmd.Flags().GetString("path")
	if err != nil {
		return err
	}

	err = os.MkdirAll(basePath, 0o750)
	if err != nil {
		return err
	}

	fsys := rwfs.OSDirFS(basePath)

	handler, err := NewHandler(cmd.Context())
	if err != nil {
		return err
	}

	errors := handler.HandleAll(cmd.Context(), count, fsys)

	for err := range errors {
		log.Println(err)
	}

	return nil
}

func handleFetch[T any](
	wg *sync.WaitGroup,
	errors chan<- error,
	ctx context.Context,
	fsys rwfs.FS,
	name string,
	count int,
	handle func(context.Context, int) ([]T, error),
) {
	defer wg.Done()

	data, err := handle(ctx, count)
	if err != nil {
		errors <- err

		return
	}

	err = writeFile(fsys, name, data)
	if err != nil {
		errors <- err
	}
}

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
