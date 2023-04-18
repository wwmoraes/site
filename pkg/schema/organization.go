package schema

type Organization struct {
	Thing
}

func NewOrganization(name string) *Organization {
	return &Organization{
		Thing: Thing{
			LinkedData: LinkedData{
				Context: "https://schema.org",
				Type:    "Organization",
			},
			Name: name,
		},
	}
}
