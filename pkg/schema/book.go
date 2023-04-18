package schema

const (
	AudioBookFormat    BookFormatType = "AudiobookFormat"
	EBookFormat        BookFormatType = "EBook"
	GraphicNovelFormat BookFormatType = "GraphicNovel"
	HardcoverFormat    BookFormatType = "Hardcover"
	PaperbackFormat    BookFormatType = "Paperback"
)

type BookFormatType string

// Book represents a book in any format: physical, digital, audio, etc
type Book struct {
	CreativeWork

	//// Indicates whether the book is an abridged edition.
	Abridged Boolean `json:"abridged,omitempty"`

	//// The edition of the book.
	BookEdition Text `json:"bookEdition,omitempty"`

	//// The format of the book.
	BookFormat BookFormatType `json:"bookFormat,omitempty"`

	//// The illustrator of the book.
	// illustrator *Person `json:"illustrator,omitempty"`

	//// The ISBN of the book.
	ISBN Text `json:"isbn,omitempty"`

	//// The number of pages in the book.
	NumberOfPages Integer `json:"numberOfPages,omitempty"`
}

func NewBook(title, isbn string) *Book {
	return &Book{
		CreativeWork: CreativeWork{
			Thing: Thing{
				LinkedData: LinkedData{
					Context: "https://schema.org",
					Type:    "Book",
				},
				Name: title,
			},
		},
		ISBN: isbn,
	}
}
