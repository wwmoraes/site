package adapters

type DataStore interface {
	Update(string, []byte) error
}
