package planets

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

func GenerateLaws(dice *dice.Dicepool, govr, worldtype ehex.Ehex) ehex.Ehex {
	switch worldtype.Code() {
	case "D", "E":
		return ehex.New().Set(0)
	}
	laws := dice.Flux() + govr.Value()
	if laws < 0 {
		laws = 0
	}
	if laws > 18 {
		laws = 18
	}
	return ehex.New().Set(laws)
}
