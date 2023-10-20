package schema

const (
	AudioBookFormat    BookFormatType = "AudiobookFormat"
	EBookFormat        BookFormatType = "EBook"
	GraphicNovelFormat BookFormatType = "GraphicNovel"
	HardcoverFormat    BookFormatType = "Hardcover"
	PaperbackFormat    BookFormatType = "Paperback"
)

type BookFormatType string

// Book represents a book in any format: physical, digital, audio, etc.
type Book struct {
	CreativeWork

	// Abridged indicates whether the book is an abridged edition.
	Abridged Boolean `json:"abridged,omitempty"`

	// BookEdition the edition of the book.
	BookEdition Text `json:"bookEdition,omitempty"`

	// BookFormat the format of the book.
	BookFormat BookFormatType `json:"bookFormat,omitempty"`

	// Illustrator the illustrator of the book.
	// Illustrator *Person `json:"illustrator,omitempty"`

	// ISBN the International Standard Book Number of the book.
	ISBN Text `json:"isbn,omitempty"`

	// NumberOfPages the number of pages in the book.
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
