package structure

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
)

type GasGigant struct {
	Nomena       string   `json:"Position Code"`
	Size         string   `json:"Size"`
	Diameter     int      `json:"Diameter"`
	Mass         int      `json:"Mass"`
	MoonQuantity int      `json:"Moon Quantity"`
	MoonSizes    []string `json:"Moon Sizes"`
}

func NewGasGigant(systemData string) *GasGigant {
	dice := dice.New().SetSeed(systemData)
	gg := GasGigant{}
	dm := 0
	r1 := dice.Sroll("1d6") + dm
	r1 = boundInt(r1, 1, 6)
	switch r1 {
	case 1, 2:
		gg.Size = "GS"
	case 3, 4:
		gg.Size = "GM"
	case 5, 6:
		gg.Size = "GL"
	}
	switch gg.Size {
	case "GS":
		gg.Diameter = dice.Sroll("1d3") + dice.Sroll("1d3")
		gg.Mass = 5 * (dice.Sroll("1d6") + 1)
	case "GM":
		gg.Diameter = dice.Sroll("1d6") + 6
		gg.Mass = 20 * (dice.Sroll("3d6") - 1)
	case "GL":
		gg.Diameter = dice.Sroll("2d6") + 6
		gg.Mass = dice.Sroll("1d3") * 50 * (dice.Sroll("3d6") + 4)
		if gg.Mass >= 3000 {
			gg.Mass = 4000 - ((dice.Sroll("2d6") - 2) * 200)
		}
	}
	return &gg
}
