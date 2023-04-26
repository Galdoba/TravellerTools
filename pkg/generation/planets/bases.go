package planets

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

func GenerateBases(dice *dice.Dicepool, port ehex.Ehex) ehex.Ehex {
	m := dice.Sroll("2d6")
	n := dice.Sroll("2d6")
	s := dice.Sroll("2d6")

	basePresence := 0
	switch port.Code() {
	case "A":
		if m >= 8 {
			basePresence += 1
		}
		if n >= 8 {
			basePresence += 2
		}
		if s >= 10 {
			basePresence += 4
		}
	case "B":
		if m >= 8 {
			basePresence += 1
		}
		if n >= 8 {
			basePresence += 2
		}
		if s >= 9 {
			basePresence += 4
		}
	case "C", "F":
		if m >= 8 {
			basePresence += 1
		}
		if n >= 8 {
			basePresence += 2
		}
		if s >= 9 {
			basePresence += 4
		}
	case "D", "G":
		if s >= 8 {
			basePresence += 4
		}
	}
	return ehex.New().Set(basePresence)
}
