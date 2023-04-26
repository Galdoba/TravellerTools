package planets

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

func GeneratePops(dice *dice.Dicepool, worldtype, limit ehex.Ehex) ehex.Ehex {
	dm := 0
	switch worldtype.Code() {
	case "C", "J":
		dm = -6
	case "H":
		dm = -4
	}
	pop := dice.Sroll("2d6") + dm
	if pop == 12 {
		r := dice.Sroll("1d6")
		for r == 6 {
			pop++
			r = dice.Sroll("1d6")
		}
	}
	if pop > limit.Value() {
		pop = limit.Value()
	}
	if pop < 0 {
		pop = 0
	}
	return ehex.New().Set(pop)
}

func PopulationDigit(dice *dice.Dicepool) ehex.Ehex {
	return ehex.New().Set(dice.Sroll("1d9"))
}
