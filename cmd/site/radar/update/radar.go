package update

import (
	"io"
	"text/template"
)

type RadarTemplate struct {
	Blips  []*Blip
	Radius int
	Size   int
}

func generateRadar(templateString string, out io.Writer, values *RadarTemplate) error {
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"div": func(a, b int) int { return a / b },
		"mul": func(a, b int) int { return a * b },
	}).Parse(templateString)
	if err != nil {
		return err
	}

	return tmpl.Execute(out, values)
}
