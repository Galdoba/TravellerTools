package planets

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

func GeneratePort(dice *dice.Dicepool, pops, worldtype ehex.Ehex) ehex.Ehex {
	switch worldtype.Code() {
	case "E":
		return ehex.New().Set(6)
	}
	portVal := pops.Value() - dice.Sroll("1d6") + 1
	if portVal >= 8 {
		return ehex.New().Set(1)
	}
	switch portVal {
	default:
		return ehex.New().Set(6)
	case 1, 2:
		return ehex.New().Set(5)
	case 3:
		return ehex.New().Set(4)
	case 4, 5:
		return ehex.New().Set(3)
	case 6, 7:
		return ehex.New().Set(2)
	}
}
