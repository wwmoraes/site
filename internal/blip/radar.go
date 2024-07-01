package blip

import (
	"fmt"
	"math"

	"github.com/gohugoio/hugo/tpl/crypto"
)

const (
	Quadrants = 4
	Tiers     = 4
)

var orientations = [Quadrants]Orientation{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}}

type (
	Orientation = [2]float64
	Percentage  = float64
	Position    = Orientation
)

type RadarOptions struct {
	// Tiers names of tiers from the innermost to the outermost
	Tiers [4]string

	// Quadrants names of quadrants in a clockwise order starting with the top-left one
	Quadrants [4]string

	// Proportion of each tier rim relative to the radar radius. It must sum to a 100%
	Proportion [4]Percentage

	// Angle maximum angle the blips will end up
	Angle float64

	// BlipRadius the radius of blips plotted on the radar
	BlipRadius float64

	// Radius the radar maximum radius
	Radius float64
}

type RadarParameters struct {
	QuadrantOrientation map[string]Orientation
	TierRadius          map[string]float64
	TierOffset          map[string]float64
	Angle               float64
	BlipRadius          float64
	Radius              float64
}

type BlipParameters struct {
	Quadrant string
	Tier     string
	Title    string
}

func NewRadar(options RadarOptions) (*RadarParameters, error) {
	params := &RadarParameters{
		Angle:               options.Angle,
		BlipRadius:          options.BlipRadius,
		QuadrantOrientation: make(map[string]Orientation, Quadrants),
		Radius:              options.Radius,
		TierOffset:          make(map[string]float64, Tiers),
		TierRadius:          make(map[string]float64, Tiers),
	}

	for index, name := range options.Quadrants {
		params.QuadrantOrientation[name] = orientations[index]
	}

	offset := 0.0

	for index, name := range options.Tiers {
		size := options.Proportion[index] * options.Radius

		params.TierOffset[name] = offset
		params.TierRadius[name] = size

		offset += size
	}

	return params, nil
}

func CalculatePosition(radar *RadarParameters, blip *BlipParameters) (Position, error) {
	if radar == nil {
		return Position{}, fmt.Errorf("radar parameters must not be empty")
	}

	if blip == nil {
		return Position{}, fmt.Errorf("blip parameters must not be empty")
	}

	fnv, err := crypto.New().FNV32a(blip.Title)
	if err != nil {
		return Position{}, err
	}

	angle := math.Mod(float64(fnv), radar.Angle)
	cos := math.Abs(math.Cos(angle))
	sin := math.Abs(math.Sin(angle))

	// adjust the offset and width of band to avoid drawing on top of boundaries
	width := radar.TierRadius[blip.Tier] - radar.BlipRadius*2 //nolint:mnd
	offset := radar.TierOffset[blip.Tier] + radar.BlipRadius
	radius := offset + math.Mod(float64(fnv), width)

	x := math.Max(radar.BlipRadius, math.Round(cos*radius)) * radar.QuadrantOrientation[blip.Quadrant][0]
	y := math.Max(radar.BlipRadius, math.Round(sin*radius)) * radar.QuadrantOrientation[blip.Quadrant][1]

	return Position{x, y}, nil
}
