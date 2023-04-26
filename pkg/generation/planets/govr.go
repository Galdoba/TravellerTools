package planets

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

func GenerateGoverment(dice *dice.Dicepool, pops, worldtype ehex.Ehex) ehex.Ehex {
	switch worldtype.Code() {
	case "D", "E":
		return ehex.New().Set(0)
	}
	govr := dice.Flux() + pops.Value()
	if govr < 0 {
		govr = 0
	}
	if govr > 15 {
		govr = 15
	}
	return ehex.New().Set(govr)
}
