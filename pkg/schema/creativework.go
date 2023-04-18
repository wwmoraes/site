package schema

// CreativeWork The most generic kind of creative work, including books, movies,
// photographs, software programs, etc
//
// https://schema.org/CreativeWork
type CreativeWork struct {
	Thing

	// The author of this content or rating. Please note that author is special in that HTML 5 provides a special mechanism for indicating authorship via the rel tag. That is equivalent to this and may be used interchangeably.
	Author []*Person `json:"author,omitempty"`

	// A thumbnail image relevant to the Thing.
	ThumbnailURL URL `json:"thumbnailUrl,omitempty"`

	// The location where the CreativeWork was created, which may not be the same as the location depicted in the CreativeWork.
	LocationCreated *Place `json:"locationCreated,omitempty"`

	// The publisher of the creative work.
	Publisher *Organization `json:"publisher,omitempty"` // TODO Person

	// Date of first broadcast/publication.
	DatePublished string `json:"datePublished,omitempty"`
}
