package adapters

import "context"

type DataPipe struct {
	Source DataSource
	Store  DataStore
	Name   string
}

func (pipe *DataPipe) Update(ctx context.Context) error {
	data, err := pipe.Source.Fetch(ctx)
	if err != nil {
		return err
	}

	return pipe.Store.Update(pipe.Name, data)
}
