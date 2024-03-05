package world

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

func NewHydrographics(dice *dice.Dicepool, size, atmo int) ehex.Ehex {
	dm := 0
	if size < 3 {
		return ehex.New().Set(0)
	}
	if atmo < 2 {
		return ehex.New().Set(0)
	}
	switch atmo {
	case 2, 3, 11, 12:
		dm = -6
	}
	r := dice.Sroll("2d6") + dm
	if r < 0 {
		r = 0
	}
	if r > 10 {
		r = 10
	}
	return ehex.New().Set(r)
}
