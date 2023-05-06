package goodreads

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/leveldbcache"
	"github.com/spf13/cobra"
	"github.com/wwmoraes/site/pkg/openlibrary"
	"github.com/wwmoraes/site/pkg/schema"
)

var whitespaces *regexp.Regexp

func init() {
	var err error

	whitespaces, err = regexp.Compile("[ ]+")

	if err != nil {
		panic(err)
	}
}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goodreads",
		Short: "fetch Goodreads shelf data",
		RunE:  cmdFetch,
	}

	flags := cmd.Flags()
	flags.String("list", "", "Goodreads list name")
	flags.String("shelves", "", "Goodreads shelf names (comma-separated)")

	return cmd
}

func cmdFetch(cmd *cobra.Command, args []string) error {
	list, err := cmd.Flags().GetString("list")
	if err != nil {
		return err
	}
	if len(list) <= 0 {
		return fmt.Errorf("List name must be provided")
	}

	rawShelves, err := cmd.Flags().GetString("shelves")
	if err != nil {
		return err
	}

	shelves := strings.Split(rawShelves, ",")

	if len(shelves) <= 0 {
		return fmt.Errorf("at least one shelf name must be provided")
	}

	path := "data/goodreads"
	err = os.MkdirAll(path, 0750)
	if err != nil {
		log.Println(err)
		return err
	}

	cache, err := leveldbcache.New("tmp/update-cache")
	if err != nil {
		log.Println(err)
		return err
	}

	var wg sync.WaitGroup
	var errors []error
	for _, shelf := range shelves {
		wg.Add(1)
		go func(shelf string) {
			defer wg.Done()

			err = fetch(cache, path, list, shelf)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to fetch shelf %s data: %w", shelf, err))
			}
		}(shelf)
	}
	wg.Wait()

	if len(errors) > 0 {
		var sb strings.Builder
		sb.WriteString("failed to fetch shelves:\n")
		for _, err := range errors {
			sb.WriteString(fmt.Sprintf("%s\n", err.Error()))
		}
		return fmt.Errorf(sb.String())
	}

	return nil
}

func fetch(cache *leveldbcache.Cache, path, list, shelf string) error {
	collector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"),
	)
	transport := httpcache.NewTransport(cache)
	client := transport.Client()
	collector.WithTransport(transport)

	books := []*schema.Book{}

	collector.OnHTML("table#books tr.bookalike.review", func(elem *colly.HTMLElement) {
		isbn13 := elem.ChildText("td.field.isbn13 .value")
		title := CleanTitle(elem.ChildText("td.field.title .value"))
		author := strings.TrimSpace(strings.TrimSuffix(elem.ChildText("td.field.author .value"), "*"))

		log.Println("collecting book", title)

		book := schema.NewBook(title, isbn13)
		book.ID = fmt.Sprintf("https://openlibrary.org/isbn/%s", isbn13)
		book.Author = append(book.Author, schema.NewPerson(author))
		book.ThumbnailURL = elem.ChildAttr("td.field.cover img", "src")

		// enrich metadata using the Open Library information
		// https://openlibrary.org/isbn/9780596514983.json
		meta, err := openlibrary.GetBookData(client, isbn13)
		if err != nil {
			log.Println("failed to fetch metadata from Open Library:", err)
		} else if meta == nil {
			log.Println("book not found in Open Library:", book.Name, book.ISBN)
		} else {
			if book.ThumbnailURL == "" {
				if meta.Cover != nil {
					if meta.Cover.Large != "" {
						book.ThumbnailURL = meta.Cover.Large
					} else if meta.Cover.Medium != "" {
						book.ThumbnailURL = meta.Cover.Medium
					} else if meta.Cover.Small != "" {
						book.ThumbnailURL = meta.Cover.Small
					}
				}
			}

			if meta.PublishPlaces != nil && len(meta.PublishPlaces) > 0 {
				book.LocationCreated = schema.NewPlace(meta.PublishPlaces[0].Name)
				parts := strings.Split(book.LocationCreated.Name, ", ")
				if len(parts) == 3 {
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
		}

		books = append(books, book)

		log.Println("collected book:", book.Name)
	})

	collector.OnScraped(func(r *colly.Response) {
		log.Println("scraped", r.Request.URL)

		data, err := json.MarshalIndent(books, "", "  ")
		if err != nil {
			log.Println(err)
			return
		}

		err = os.WriteFile(filepath.Join(path, fmt.Sprintf("%s.json", shelf)), data, 0640)
		if err != nil {
			log.Println(err)
		}
	})

	collector.OnRequest(func(r *colly.Request) {
		log.Println("requesting", r.URL)
	})

	return collector.Visit(fmt.Sprintf("https://www.goodreads.com/review/list/%s?shelf=%s&order=d&sort=date_read", list, shelf))
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

func CleanTitle(title string) string {
	title = strings.TrimSpace(title)
	title = strings.ReplaceAll(title, "\r", "")
	title = strings.ReplaceAll(title, "\n", "")
	title = strings.ReplaceAll(title, "\t", " ")
	title = whitespaces.ReplaceAllString(title, " ")

	return title
}
