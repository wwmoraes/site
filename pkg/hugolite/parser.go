package hugolite

import (
	"bytes"
	"io"

	"github.com/BurntSushi/toml"
	"github.com/adrg/frontmatter"
	"gopkg.in/yaml.v3"
)

// Format defines a supported frontmatter type in Hugo.
type Format string

const (
	// ORG  Format = "org"
	// JSON Format = "json"

	TOML Format = "toml"
	YAML Format = "yaml"

	// orgDelimiter  = "#+"
	// jsonDelimiter = "{"

	yamlDelimiter   = "---\n"
	tomlDelimiter   = "+++\n"
	delimiterLength = 4
)

type ContentFrontMatter struct {
	Content           []byte
	FrontMatter       map[string]any
	FrontMatterFormat Format
}

func ParseFrontMatterAndContent(r io.Reader) (*ContentFrontMatter, error) {
	frontMatter := make(map[string]any)

	var start [delimiterLength]byte

	_, err := io.ReadAtLeast(r, start[:], delimiterLength)
	if err != nil {
		return nil, err
	}

	format := detectFrontMatterFormat(string(start[:]))

	r = io.MultiReader(bytes.NewReader(start[:]), r)

	rest, err := frontmatter.Parse(r, frontMatter,
		frontmatter.NewFormat("---", "---", yaml.Unmarshal),
		frontmatter.NewFormat("+++", "+++", toml.Unmarshal),
	)
	if err != nil {
		return nil, err
	}

	return &ContentFrontMatter{
		Content:           rest,
		FrontMatter:       frontMatter,
		FrontMatterFormat: format,
	}, nil
}

// func InterfaceToFrontMatter(in any, format Format, w io.Writer) error {
// 	return parser.InterfaceToFrontMatter(in, metadecoders.Format(format), w)
// }

func detectFrontMatterFormat(str string) Format {
	switch str[:delimiterLength] {
	case yamlDelimiter:
		return YAML
	case tomlDelimiter:
		return TOML
	default:
		return ""
	}
}
