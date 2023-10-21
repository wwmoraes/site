package drivers

import (
	"context"

	"github.com/goccy/go-json"
)

type GoodreadsSource struct {
	List  string
	Shelf string
}

func (source *GoodreadsSource) Fetch(ctx context.Context) ([]byte, error) {
	collector := goodreadsCollector{}

	err := collector.Collect(source.List, source.Shelf)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(collector.Books, "", "  ")
}
