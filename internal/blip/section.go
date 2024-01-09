package blip

import "fmt"

const (
	Languages  Section = "languages"
	Platforms  Section = "platforms"
	Techniques Section = "techniques"
	Tools      Section = "tools"
)

type Section string

func (s Section) String() string {
	return string(s)
}

func (s *Section) Set(v string) error {
	switch v {
	case "languages", "platforms", "techniques", "tools":
		*s = Section(v)
	default:
		return fmt.Errorf("Section must be either languages, platforms, techniques or tools")
	}

	return nil
}

func (s *Section) Type() string {
	return "Section"
}
