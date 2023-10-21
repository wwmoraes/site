package goodreads

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/goccy/go-json"
	"github.com/gocolly/colly/v2"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/leveldbcache"
	"github.com/spf13/cobra"
	"github.com/wwmoraes/go-rwfs"
	"github.com/wwmoraes/site/pkg/openlibrary"
	"github.com/wwmoraes/site/pkg/schema"
)

const (
	dataPath  = "data/goodreads"
	cachePath = "tmp/books-fetch-cache"
)

var whitespaces = regexp.MustCompile("[ ]+")

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goodreads",
		Short: "fetch Goodreads shelf data",
		RunE:  cmdFetch,
	}

	flags := cmd.Flags()
	flags.String("list", "", "Goodreads list name")
	flags.StringSlice("shelves", []string{}, "Goodreads shelf names (comma-separated)")

	return cmd
}

func cmdFetch(cmd *cobra.Command, args []string) error {
	list, err := cmd.Flags().GetString("list")
	if err != nil {
		return err
	}

	if len(list) <= 0 {
		return fmt.Errorf("List name must be set")
	}

	shelves, err := cmd.Flags().GetStringSlice("shelves")
	if err != nil {
		return err
	}

	if len(shelves) <= 0 {
		return fmt.Errorf("at least one shelf name must be set")
	}

	err = os.MkdirAll(dataPath, 0o750)
	if err != nil {
		return err
	}

	err = os.MkdirAll(cachePath, 0o750)
	if err != nil {
		return err
	}

	cache, err := leveldbcache.New(cachePath)
	if err != nil {
		return err
	}

	fsys := rwfs.OSDirFS(dataPath)

	errors := fetchShelves(cache, fsys, list, shelves)

	failed := false

	var fail sync.Once

	for err := range errors {
		fail.Do(func() {
			failed = true
		})

		log.Println(err)
	}

	if failed {
		return fmt.Errorf("failed to fetch book shelves")
	}

	return nil
}

func fetchShelves(cache *leveldbcache.Cache, fsys rwfs.FS, list string, shelves []string) <-chan error {
	errors := make(chan error, 2)

	go func() {
		defer close(errors)

		var wg sync.WaitGroup

		for _, shelf := range shelves {
			wg.Add(1)

			go func(shelf string) {
				defer wg.Done()

				err := fetch(cache, fsys, list, shelf)
				if err != nil {
					errors <- err
				}
			}(shelf)
		}

		wg.Wait()
	}()

	return errors
}

func fetch(cache *leveldbcache.Cache, fsys rwfs.FS, list, shelf string) error {
	collector := colly.NewCollector(
		colly.UserAgent(
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",
		),
	)
	transport := httpcache.NewTransport(cache)
	client := transport.Client()
	collector.WithTransport(transport)

	books := []*schema.Book{}
	booksInput := make(chan *schema.Book, 10)

	go func() {
		for book := range booksInput {
			books = append(books, book)
		}
	}()

	collector.OnHTML("table#books tr.bookalike.review", func(elem *colly.HTMLElement) {
		if err := collectBook(elem, booksInput, client); err != nil {
			log.Println(fmt.Errorf("failed to collect book: %w", err))
		}
	})

	collector.OnScraped(func(r *colly.Response) {
		close(booksInput)

		err := writeBooks(fsys, shelf, books)
		if err != nil {
			log.Println(err)
		}
	})

	return collector.Visit(
		fmt.Sprintf("https://www.goodreads.com/review/list/%s?shelf=%s&order=d&sort=date_read", list, shelf),
	)
}

func collectBook(elem *colly.HTMLElement, books chan<- *schema.Book, client *http.Client) error {
	isbn13 := elem.ChildText("td.field.isbn13 .value")
	title := CleanTitle(elem.ChildText("td.field.title .value"))
	author := strings.TrimSpace(strings.TrimSuffix(elem.ChildText("td.field.author .value"), "*"))

	log.Println("collecting book", title)

	book := schema.NewBook(title, isbn13)
	book.ID = fmt.Sprintf("https://openlibrary.org/isbn/%s", isbn13)
	book.Author = append(book.Author, schema.NewPerson(author))
	book.ThumbnailURL = elem.ChildAttr("td.field.cover img", "src")

	// enrich metadata using the Open Library information
	meta, err := openlibrary.GetBookData(client, isbn13)
	if err != nil {
		return fmt.Errorf("failed to fetch metadata from Open Library: %w", err)
	}

	if meta == nil {
		return fmt.Errorf("book not found in Open Library (%s, %s)", book.Name, book.ISBN)
	}

	err = meta.AugmentBook(book)
	if err != nil {
		return fmt.Errorf("failed to update book metadata: %w", err)
	}

	books <- book

	log.Println("collected book:", book.Name)

	return nil
}

func writeBooks(fsys rwfs.FS, shelf string, books []*schema.Book) error {
	data, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return err
	}

	fd, err := fsys.OpenFile(fmt.Sprintf("%s.json", shelf), rwfs.O_WRONLY|rwfs.O_TRUNC|rwfs.O_CREATE, 0o640)
	if err != nil {
		return err
	}
	defer fd.Close()

	_, err = fd.Write(data)

	return err
}

func CleanTitle(title string) string {
	title = strings.TrimSpace(title)
	title = strings.ReplaceAll(title, "\r", "")
	title = strings.ReplaceAll(title, "\n", "")
	title = strings.ReplaceAll(title, "\t", " ")
	title = whitespaces.ReplaceAllString(title, " ")

	return title
}
