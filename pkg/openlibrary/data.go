package openlibrary

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BooksData map[string]*BookData

type BookData struct {
	Authors       []*Author    `json:"authors,omitempty"`
	Key           string       `json:"key,omitempty"`
	NumberOfPages int          `json:"number_of_pages,omitempty"`
	PublishDate   string       `json:"publish_date,omitempty"`
	Publishers    []*Publisher `json:"publishers,omitempty"`
	PublishPlaces []*Place     `json:"publish_places,omitempty"`
	Subtitle      string       `json:"subtitle,omitempty"`
	Title         string       `json:"title,omitempty"`
	URL           string       `json:"url,omitempty"`
	Cover         *Cover       `json:"cover,omitempty"`

	// Pagination         string                 `json:"pagination"`
}

type Cover struct {
	Small  string `json:"small,omitempty"`
	Medium string `json:"medium,omitempty"`
	Large  string `json:"large,omitempty"`
}

func GetBookData(client *http.Client, isbn string) (*BookData, error) {
	resp, err := client.Get(fmt.Sprintf("https://openlibrary.org/api/books?bibkeys=ISBN:%s&jscmd=data&format=json", isbn))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	results := make(BooksData, 1)
	err = json.Unmarshal(data, &results)
	if err != nil {
		return nil, err
	}

	meta := results[fmt.Sprintf("ISBN:%s", isbn)]

	return meta, nil
}
