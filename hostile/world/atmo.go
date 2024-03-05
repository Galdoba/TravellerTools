package world

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

func NewAtmosphere(dice *dice.Dicepool, size int) ehex.Ehex {
	dm := 0
	switch size {
	case 0, 1, 2:
		return ehex.New().Set(0)
	case 3, 4:
		dm = -5
	default:
		if size >= 8 {
			dm = 2
		}
	}
	r := dice.Sroll("2d6") + dm
	if r < 0 {
		r = 0
	}
	if r > 12 {
		r = 12
	}
	return ehex.New().Set(r)
}
