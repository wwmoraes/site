package openlibrary

import (
	"fmt"
	"io"
	"net/http"

	"github.com/goccy/go-json"
)

type BooksDetails map[string]*BookDetails

type BookDetails struct {
	BibKey       string  `json:"bib_key,omitempty"`
	InfoURL      string  `json:"info_url,omitempty"`
	Preview      string  `json:"preview,omitempty"`
	PreviewURL   string  `json:"preview_url,omitempty"`
	ThumbnailURL string  `json:"thumbnail_url,omitempty"`
	Details      Details `json:"details,omitempty"`
}

type Details struct {
	Authors       []*Author `json:"authors,omitempty"`
	Key           string    `json:"key,omitempty"`
	NumberOfPages int       `json:"number_of_pages,omitempty"`
	PublishDate   string    `json:"publish_date,omitempty"`
	Publishers    []string  `json:"publishers,omitempty"`
	PublishPlaces []string  `json:"publish_places,omitempty"`
	Subtitle      string    `json:"subtitle,omitempty"`
	Title         string    `json:"title,omitempty"`

	//nolint:dupword
	// Identifiers        Identifiers            `json:"identifiers"`
	// TableOfContents    []TableOfContent       `json:"table_of_contents"`
	// Weight             string                 `json:"weight"`
	// Covers             []int64                `json:"covers"`
	// PhysicalFormat     string                 `json:"physical_format"`
	// Classifications    map[string]interface{} `json:"classifications"`
	// SourceRecords      []string               `json:"source_records"`
	// Languages          []Type                 `json:"languages"`
	// Works              []Type                 `json:"works"`
	// Type               Type                   `json:"type"`
	// Uris               []string               `json:"uris"`
	// PhysicalDimensions string                 `json:"physical_dimensions"`
	// Ocaid              string                 `json:"ocaid"`
	// Isbn10             []string               `json:"isbn_10"`
	// Isbn13             []string               `json:"isbn_13"`
	// Pagination         string                 `json:"pagination"`
	// LcClassifications  []string               `json:"lc_classifications"`
	// OclcNumbers        []string               `json:"oclc_numbers"`
	// LocalID            []string               `json:"local_id"`
	// LatestRevision     int64                  `json:"latest_revision"`
	// Revision           int64                  `json:"revision"`
	// Created            Created                `json:"created"`
	// LastModified       Created                `json:"last_modified"`
}

// type Created struct {
// 	Type  string `json:"type"`
// 	Value string `json:"value"`
// }

// type Identifiers struct {
// 	Google       []string `json:"google"`
// 	AmazonDeAsin []string `json:"amazon.de_asin"`
// 	Librarything []string `json:"librarything"`
// 	Goodreads    []string `json:"goodreads"`
// }

// type Type struct {
//nolint:dupword
// 	Key Key `json:"key"`
// }

// type TableOfContent struct {
// 	Level int64   `json:"level"`
// 	Label string  `json:"label"`
// 	Title *string `json:"title,omitempty"`
//nolint:dupword
// 	Type  Type    `json:"type"`
// }

func GetBookDetails(client *http.Client, isbn string) (*BookDetails, error) {
	resp, err := client.Get(
		fmt.Sprintf("https://openlibrary.org/api/books?bibkeys=ISBN:%s&jscmd=details&format=json", isbn),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	results := make(BooksDetails, 1)

	err = json.Unmarshal(data, &results)
	if err != nil {
		return nil, err
	}

	meta := results[fmt.Sprintf("ISBN:%s", isbn)]

	return meta, nil
}
