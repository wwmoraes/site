package update

import (
	"github.com/spf13/cobra"
	"github.com/wwmoraes/go-rwfs"
	"github.com/wwmoraes/site/internal/adapters"
	"github.com/wwmoraes/site/internal/drivers"
	"github.com/wwmoraes/site/internal/usecases"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "refreshes data files",
		RunE:  update,
	}

	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	store := &drivers.HugoDataStore{FS: rwfs.OSDirFS("data")}

	updaters := usecases.DataUpdaters{
		&adapters.DataPipe{
			Source: &adapters.GoodreadsFetcher{
				List:  "138333248-william",
				Shelf: "currently-reading",
			},
			Name:  "goodreads/currently-reading.json",
			Store: store,
		},
		&adapters.DataPipe{
			Source: &adapters.GoodreadsFetcher{
				List:  "138333248-william",
				Shelf: "read",
			},
			Name:  "goodreads/read.json",
			Store: store,
		},
		&adapters.DataPipe{
			Source: &adapters.GoodreadsFetcher{
				List:  "138333248-william",
				Shelf: "to-read",
			},
			Name:  "goodreads/to-read.json",
			Store: store,
		},
	}

	errors := updaters.ExecuteAll(10) //nolint:gomnd

	for err := range errors {
		cmd.PrintErrln(err)
	}

	return nil
}
