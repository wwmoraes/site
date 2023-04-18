package openlibrary

type Key = string

type Author struct {
	Key  string `json:"key,omitempty"`
	URL  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
}

type Publisher struct {
	Name string `json:"name,omitempty"`
}

type Place struct {
	Name string `json:"name,omitempty"`
}
