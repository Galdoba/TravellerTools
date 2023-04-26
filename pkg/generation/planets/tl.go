package planets

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

func GenerateTL(dice *dice.Dicepool, port, size, atmo, hydr, pops, govr ehex.Ehex) ehex.Ehex {
	tl := dice.Sroll("1d6")
	switch port.Code() {
	case "A":
		tl += 6
	case "B":
		tl += 4
	case "C":
		tl += 2
	case "F":
		tl += 1
	case "X":
		tl -= 4
	}
	switch size.Value() {
	case 0, 1:
		tl += 2
	case 2, 3, 4:
		tl += 1
	}
	switch atmo.Value() {
	case 0, 1, 2, 3, 10, 11, 12, 13, 14, 15:
		tl += 1
	}
	switch hydr.Value() {
	case 9:
		tl += 1
	case 10:
		tl += 2
	}
	switch pops.Value() {
	case 1, 2, 3, 4, 5:
		tl += 1
	case 6, 7, 8:
	case 9:
		tl += 2
	default:
		tl += 4
	}
	switch govr.Value() {
	case 0, 5:
		tl += 1
	case 13:
		tl -= 2
	}
	if tl < 0 {
		tl = 0
	}
	return ehex.New().Set(tl)
}
