package adapters

type DataPipe struct {
	Source DataSource
	Store  DataStore
	Name   string
}

func (pipe *DataPipe) Update() error {
	data, err := pipe.Source.Fetch()
	if err != nil {
		return err
	}

	return pipe.Store.Update(pipe.Name, data)
}
