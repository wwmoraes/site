package openlibrary

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"github.com/wwmoraes/site/pkg/schema"
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

func (meta *BookData) getThumbnailURL(current string) string {
	if current != "" || meta.Cover == nil {
		return current
	}

	if meta.Cover.Large != "" {
		return meta.Cover.Large
	} else if meta.Cover.Medium != "" {
		return meta.Cover.Medium
	} else if meta.Cover.Small != "" {
		return meta.Cover.Small
	}

	return current
}

func (meta *BookData) AugmentBook(book *schema.Book) error {
	if meta == nil {
		return fmt.Errorf("book not found in Open Library (%s, %s)", book.Name, book.ISBN)
	}

	book.ThumbnailURL = meta.getThumbnailURL(book.ThumbnailURL)

	if meta.PublishPlaces != nil && len(meta.PublishPlaces) > 0 {
		book.LocationCreated = schema.NewPlace(meta.PublishPlaces[0].Name)

		parts := strings.Split(book.LocationCreated.Name, ", ")
		if len(parts) == 3 { //nolint:mnd
			book.LocationCreated.Address = &schema.PostalAddress{
				AddressLocality: parts[0],
				AddressRegion:   parts[1],
				AddressCountry:  parts[2],
			}
		}
	}

	if len(meta.Publishers) > 0 {
		book.Publisher = schema.NewOrganization(meta.Publishers[0].Name)
	}

	if len(meta.PublishDate) > 0 {
		parsed, err := ParseDate(meta.PublishDate)
		if err == nil {
			book.DatePublished = parsed.UTC().Format(time.RFC3339)
		}
	}

	book.Author = make([]*schema.Person, len(meta.Authors))
	for index, author := range meta.Authors {
		book.Author[index] = schema.NewPerson(author.Name)
	}

	return nil
}

func ParseDate(value string) (*time.Time, error) {
	var result time.Time

	var err error

	layouts := []string{
		"January 2, 2006",
		"Jan 02, 2006",
		"2006",
		"2006-01",
		"January 2006",
	}

	for _, layout := range layouts {
		result, err = time.Parse(layout, value)
		if err == nil {
			return &result, nil
		}
	}

	return nil, err
}
