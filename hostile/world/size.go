package world

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

func NewSize(dice *dice.Dicepool) ehex.Ehex {
	code := ""
	switch dice.Sroll("1d6") {
	default:
		code = "2d6-2"
	case 6:
		code = "2d6+4"
	}
	return ehex.New().Set(dice.Sroll(code))
}
