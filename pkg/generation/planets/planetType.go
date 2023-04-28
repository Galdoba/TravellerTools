package planets

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

/*
Planet Type:
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

const (
	WORLDTYPE_Hospitable = "A"
	WORLDTYPE_Planetoid  = "B"
	WORLDTYPE_Iceworld   = "C"
	WORLDTYPE_RadWorld   = "D"
	WORLDTYPE_Inferno    = "E"
	WORLDTYPE_BigWorld   = "F"
	WORLDTYPE_Worldlet   = "G"
	WORLDTYPE_InnerWorld = "H"
	WORLDTYPE_Stormworld = "J"
	WORLDTYPE_SSG        = "K"
	WORLDTYPE_LGG        = "L"
	WORLDTYPE_IG         = "M"
)

func GeneratePlanetType(dice *dice.Dicepool, satOrbit, hzVar ehex.Ehex) ehex.Ehex {
	hzVarInt := hzVar.Value() - 10

	zone := HZ_HOSPITABLE
	if hzVarInt < -1 {
		zone = HZ_INNER
	}
	if hzVarInt > 1 {
		zone = HZ_OUTER
	}
	i := dice.Sroll("1d6-1")
	switch zone {
	case HZ_HOSPITABLE:
		return ehex.New().Set("A")
	case HZ_INNER:
		switch satOrbit.Code() {
		default:
			return ehex.New().Set([]string{"E", "H", "F", "J", "D", "A"}[i])
		case "*":
			return ehex.New().Set([]string{"E", "H", "F", "J", "D", "A"}[i])
		}
	case HZ_OUTER:
		switch satOrbit.Code() {
		default:
			return ehex.New().Set([]string{"G", "C", "F", "C", "D", "C"}[i])
		case "*":
			return ehex.New().Set([]string{"G", "C", "F", "J", "D", "C"}[i])
		}
	}
	return nil
}
