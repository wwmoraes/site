package adapters

import (
	"github.com/goccy/go-json"
)

type GoodreadsFetcher struct {
	List  string
	Shelf string
}

func (fetcher *GoodreadsFetcher) Fetch() ([]byte, error) {
	collector := goodreadsCollector{}

	err := collector.Collect(fetcher.List, fetcher.Shelf)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(collector.Books, "", "  ")
}
