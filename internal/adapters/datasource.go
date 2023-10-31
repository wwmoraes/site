package adapters

import "context"

type DataSource interface {
	Fetch(ctx context.Context) ([]byte, error)
}
