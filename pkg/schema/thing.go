package schema

// Thing represents the most generic type of item
//
// https://schema.org/Thing
type Thing struct {
	LinkedData

	// An additional type for the item, typically used for adding more specific
	// types from external vocabularies in microdata syntax. This is a
	// relationship between something and a class that the thing is in. In RDFa
	// syntax, it is better to use the native RDFa syntax - the 'typeof'
	// attribute - for multiple types. Schema.org tools may have only weaker
	// understanding of extra types, in particular those defined externally.
	// additionalType URL

	// An alias for the item.
	// alternateName Text

	// A description of the item.
	Description Text `json:"description,omitempty"`

	// A sub property of description. A short description of the item used to
	// disambiguate from other, similar items. Information from other properties
	// (in particular, name) may be necessary for the description to be useful
	// for disambiguation.
	// disambiguatingDescription Text

	// The identifier property represents any kind of identifier for any kind of
	// Thing, such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated
	// properties for representing many of these, either as textual strings or as
	// URL (URI) links. See background notes for more details.
	// identifier *Identifier

	//// An image of the item. This can be a URL or a fully described ImageObject.
	// image ImageObject | URL

	// Indicates a page (or other CreativeWork) for which this thing is the main
	// entity being described. See background notes for details. Inverse property:
	// mainEntity
	// mainEntityOfPage CreativeWork | URL

	// The name of the item.
	Name Text `json:"name,omitempty"`

	// Indicates a potential Action, which describes an idealized action in which
	// this thing would play an 'object' role.
	// potentialAction Action

	// URL of a reference Web page that unambiguously indicates the item's
	// identity. E.g. the URL of the item's Wikipedia page, Wikidata entry, or
	// official website.
	// sameAs URL

	// A CreativeWork or Event about this Thing. Inverse property: about
	// subjectOf CreativeWork | Event

	// URL of the item.
	URL URL `json:"url,omitempty"`
}

func NewThing(parent *LinkedData, name string) *Thing {
	return &Thing{
		LinkedData: *parent,
		Name:       name,
	}
}
