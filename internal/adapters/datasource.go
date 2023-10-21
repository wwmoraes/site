package adapters

import "context"

type DataSource interface {
	Fetch(context.Context) ([]byte, error)
}
