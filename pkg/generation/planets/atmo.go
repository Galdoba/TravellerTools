package planets

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

func GenerateAtmo(dice *dice.Dicepool, size, worldtype ehex.Ehex) ehex.Ehex {
	dm := 0
	switch worldtype.Code() {
	case "J": //Stormworld
		dm += 4
	case "B": //Planetoid
		return ehex.New().Set(0)
	}
	s := size.Value()
	atmo := dice.Sroll("2d6") - 7 + s
	switch s {
	case 0:
		atmo = 0
	case 1:
		atmo = atmo - 5
	case 3, 4:
		switch atmo {
		case 4, 5, 8, 9:
			atmo = atmo - 2
		case 7:
			atmo = 4
		case 6:
			atmo = 5
		case 2, 3:
			atmo = 1

		}
	}
	if atmo < 0 {
		atmo = 0
	}
	if atmo > 15 {
		atmo = 15
	}
	return ehex.New().Set(atmo)
}
