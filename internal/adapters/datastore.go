package adapters

type DataStore interface {
	Update(name string, data []byte) error
}
