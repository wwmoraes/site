package blip

import "fmt"

const (
	Adopt  Tier = "adopt"
	Assess Tier = "assess"
	Hold   Tier = "hold"
	Trial  Tier = "trial"
)

type Tier string

func (t Tier) String() string {
	return string(t)
}

func (t *Tier) Set(v string) error {
	switch v {
	case "adopt", "assess", "hold", "trial":
		*t = Tier(v)
	default:
		return fmt.Errorf("Tier must be either adopt, assess, hold or trial")
	}

	return nil
}

func (t *Tier) Type() string {
	return "Tier"
}
