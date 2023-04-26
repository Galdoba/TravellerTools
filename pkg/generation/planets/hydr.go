package planets

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

func GenerateHydr(dice *dice.Dicepool, size, atmo, habzone, worldtype ehex.Ehex) ehex.Ehex {
	if size.Value() <= 1 {
		return ehex.New().Set(0)
	}
	if habzone.Value() < 8 {
		return ehex.New().Set(0)
	}
	dm := 0
	if atmo.Value() <= 1 {
		dm -= 4
	}
	if atmo.Value() >= 10 {
		dm -= 4
	}
	if habzone.Value() == 9 {
		dm -= 2
	}
	if habzone.Value() == 8 {
		dm -= 6
	}
	if worldtype.Code() == "H" || worldtype.Code() == "J" {
		dm -= 4
	}
	hydr := dice.Sroll("2d6-7") + size.Value() + dm
	if hydr > 10 {
		hydr = 10
	}
	if hydr < 0 {
		hydr = 0
	}
	return ehex.New().Set(hydr)
}
