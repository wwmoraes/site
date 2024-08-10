package hugolite

import (
	"errors"
	"io"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

func InterfaceToFrontMatter(in any, format Format, w io.Writer) error {
	if in == nil {
		return errors.New("input was nil")
	}

	switch format {
	case YAML:
		_, err := w.Write([]byte(yamlDelimiter))
		if err != nil {
			return err
		}

		err = InterfaceToConfig(in, format, w)
		if err != nil {
			return err
		}

		_, err = w.Write([]byte(yamlDelimiter))

		return err

	case TOML:
		_, err := w.Write([]byte(tomlDelimiter))
		if err != nil {
			return err
		}

		err = InterfaceToConfig(in, format, w)
		if err != nil {
			return err
		}

		_, err = w.Write([]byte(tomlDelimiter))

		return err

	default:
		return errors.New("unsupported Format provided")
	}
}

func InterfaceToConfig(in any, format Format, w io.Writer) error {
	if in == nil {
		return errors.New("input was nil")
	}

	switch format {
	case YAML:
		enc := yaml.NewEncoder(w)
		enc.SetIndent(2) //nolint:mnd

		return enc.Encode(in)

	case TOML:
		enc := toml.NewEncoder(w)
		enc.Indent = ""

		return enc.Encode(in)

	default:
		return errors.New("unsupported Format provided")
	}
}
