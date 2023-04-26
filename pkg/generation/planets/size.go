package planets

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

/*
	Hospitable -    A
    Planetoid -     B
    Iceworld -      C
    RadWorld -      D
    Inferno -       E
    BigWorld -      F
    Worldlet -      G
    Inner World -   H
    Stormworld -    J
	SSG -			K
	LGG - 			L
	IG -			M
*/

func GenerateSize(dice *dice.Dicepool, planetType ehex.Ehex) ehex.Ehex {
	diceCode := ""
	switch planetType.Code() {
	default:
		return nil
	case "A", "C", "H":
		diceCode = "2d6-2"
	case "B":
		return ehex.New().Set(0)
	case "D", "J":
		diceCode = "2d6"
	case "E":
		diceCode = "1d6+6"
	case "F":
		diceCode = "2d6+7"
	case "G":
		diceCode = "1d6-3"
	}
	s := dice.Sroll(diceCode)
	if s < 0 {
		s = 0
	}
	return ehex.New().Set(s)
}
