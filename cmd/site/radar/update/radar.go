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
		"div": func(a, b int) int { return a / b },
		"mul": func(a, b int) int { return a * b },
	}).Parse(templateString)
	if err != nil {
		return err
	}

	return tmpl.Execute(out, values)
}

func (radar *RadarTemplate) Tier1Radius() int {
	return radar.Radius / 2
}

func (radar *RadarTemplate) Tier2Radius() int {
	return radar.Radius*70/100 - radar.Tier1Radius()
}

func (radar *RadarTemplate) Tier3Radius() int {
	return radar.Radius*86/100 - radar.Tier2Radius()
}

func (radar *RadarTemplate) Tier4Radius() int {
	return radar.Radius - radar.Tier3Radius()
}
