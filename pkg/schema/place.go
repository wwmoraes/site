package schema

type Place struct {
	Thing

	// Physical address of the item.
	Address *PostalAddress `json:"address,omitempty"`
}

func NewPlace(name string) *Place {
	return &Place{
		Thing: Thing{
			LinkedData: LinkedData{
				Context: "https://schema.org",
				Type:    "Place",
			},
			Name: name,
		},
	}
}
