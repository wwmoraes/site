package adapters

type DataSource interface {
	Fetch() ([]byte, error)
}
