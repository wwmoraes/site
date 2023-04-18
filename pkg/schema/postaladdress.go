package schema

type PostalAddress struct {
	ContactPoint

	//// The country. For example, USA. You can also provide the two-letter ISO 3166-1 alpha-2 country code.
	AddressCountry Text `json:"addressCountry,omitempty"` // TODO Country support

	//// The locality in which the street address is, and which is in the region. For example, Mountain View.
	AddressLocality Text `json:"addressLocality,omitempty"`

	//// The region in which the locality is, and which is in the country. For example, California or another appropriate first-level Administrative division.
	AddressRegion Text `json:"addressRegion,omitempty"`

	//// The post office box number for PO box addresses.
	PostOfficeBoxNumber Text `json:"postOfficeBoxNumber,omitempty"`

	//// The postal code. For example, 94043.
	PostalCode Text `json:"postalCode,omitempty"`

	//// The street address. For example, 1600 Amphitheatre Pkwy.
	StreetAddress Text `json:"streetAddress,omitempty"`
}
