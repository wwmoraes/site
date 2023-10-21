package adapters

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/gocolly/colly"
	"github.com/wwmoraes/site/pkg/openlibrary"
	"github.com/wwmoraes/site/pkg/schema"
)

var whitespaces = regexp.MustCompile("[ ]+")

type goodreadsCollector struct {
	Books []*schema.Book
}

func (goodreads *goodreadsCollector) Collect(list, shelf string) error {
	collector := colly.NewCollector(
		colly.UserAgent(
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",
		),
	)

	goodreads.Books = []*schema.Book{}

	booksInput := make(chan *schema.Book, 10)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		for book := range booksInput {
			goodreads.Books = append(goodreads.Books, book)
		}
	}()

	collector.OnHTML("table#books tr.bookalike.review", func(elem *colly.HTMLElement) {
		isbn13 := elem.ChildText("td.field.isbn13 .value")
		title := cleanTitle(elem.ChildText("td.field.title .value"))
		author := strings.TrimSpace(strings.TrimSuffix(elem.ChildText("td.field.author .value"), "*"))

		book := schema.NewBook(title, isbn13)
		book.ID = fmt.Sprintf("https://openlibrary.org/isbn/%s", isbn13)
		book.Author = append(book.Author, schema.NewPerson(author))
		book.ThumbnailURL = elem.ChildAttr("td.field.cover img", "src")

		meta, err := openlibrary.GetBookData(http.DefaultClient, isbn13)
		if err != nil {
			return
		}

		if meta == nil {
			return
		}

		err = meta.AugmentBook(book)
		if err != nil {
			return
		}

		booksInput <- book
	})

	collector.OnScraped(func(r *colly.Response) {
		close(booksInput)

		wg.Wait()
	})

	return collector.Visit(
		fmt.Sprintf("https://www.goodreads.com/review/list/%s?shelf=%s&order=d&sort=date_read", list, shelf),
	)
}

func cleanTitle(title string) string {
	title = strings.TrimSpace(title)
	title = strings.ReplaceAll(title, "\r", "")
	title = strings.ReplaceAll(title, "\n", "")
	title = strings.ReplaceAll(title, "\t", " ")
	title = whitespaces.ReplaceAllString(title, " ")

	return title
}
