package schema

import (
	"fmt"
	"strings"
)

type Person struct {
	Thing
}

func NewPerson(name string) *Person {
	if strings.Contains(name, ", ") {
		parts := strings.SplitN(name, ", ", 2)
		name = fmt.Sprintf("%s %s", parts[1], parts[0])
	}

	return &Person{
		Thing: Thing{
			LinkedData: LinkedData{
				Context: "https://schema.org",
				Type:    "Person",
			},
			Name: name,
		},
	}
}
