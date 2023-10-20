package schema

type (
	URL     = string
	Text    = string
	Integer = int
	Boolean = bool
)

type Identifier struct {
	Properties *Thing
	Text       Text
	URL        URL
}

// LinkedData represents the properties any serializable resource must support
//
// https://www.w3.org/TR/json-ld/#syntax-tokens-and-keywords
type LinkedData struct {
	Context string `json:"@context,omitempty"`
	ID      string `json:"@id,omitempty"`
	Type    string `json:"@type,omitempty"`
}

func NewLinkedData(context, typeName, id string) *LinkedData {
	return &LinkedData{
		Context: context,
		Type:    typeName,
		ID:      id,
	}
}
