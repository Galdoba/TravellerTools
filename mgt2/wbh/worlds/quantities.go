package worlds

import (
	"github.com/Galdoba/TravellerTools/mgt2/wbh/helper"
	"github.com/Galdoba/TravellerTools/mgt2/wbh/star"
	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func GasGigantsQuantity(dice *dice.Dicepool, systemStars map[string]star.Star) int {
	r1 := dice.Sroll("2d6")
	if r1 > 9 {
		return 0
	}
	dm := 0
	if len(systemStars) == 1 && systemStars["Aa"].Class == star.ClassV {
		dm = dm + 1
	}
	if len(systemStars) > 3 {
		dm = dm - 1
	}
	switch systemStars["Aa"].Class {
	case star.ClassD, star.Pulsar, star.NeutronStar, star.BlackHole, star.ClassBD:
		dm = dm - 2
	}
	dm = dm - (1 * forAnyPostStellarObject(systemStars))
	r2 := helper.EnsureMinMax(dice.Sroll("2d6")+dm, 4, 13)
	switch r2 {
	case 4:
		return 1
	case 5, 6:
		return 2
	case 7, 8:
		return 3
	case 9, 10, 11:
		return 4
	case 12:
		return 5
	default:
		return 6
	}
}

func PlanetoidBeltsQuantity(dice *dice.Dicepool, systemStars map[string]star.Star, ggNum int) int {
	r1 := dice.Sroll("2d6")
	if r1 < 8 {
		return 0
	}
	dm := 0
	if ggNum > 0 {
		dm = dm + 1
	}
	switch systemStars["Aa"].Class {
	case star.Protostar:
		dm = dm + 3
	case star.ClassD, star.Pulsar, star.NeutronStar, star.BlackHole:
		dm = dm + 1
	}
	dm = dm + (1 * forAnyPostStellarObject(systemStars))
	if systemStars["Aa"].Specialcase == star.Primordial {
		dm = dm + 2
	}
	r2 := helper.EnsureMinMax(dice.Sroll("2d6")+dm, 6, 12)
	switch r2 {
	case 6:
		return 1
	case 12:
		return 3
	default:
		return 2
	}
}

func forAnyPostStellarObject(systemStars map[string]star.Star) int {
	dm := 0
	for _, st := range systemStars {
		switch st.Class {
		case star.ClassD, star.Pulsar, star.NeutronStar, star.BlackHole:
			dm++
		}
	}
	return dm
}

func TerrestialPlanetsQuantity(dice *dice.Dicepool, systemStars map[string]star.Star) int {
	dm := (-1 * forAnyPostStellarObject(systemStars))
	r1 := dice.Sroll("2d6") - 2 + dm
	if r1 >= 3 {
		return r1 + dice.Sroll("1d3") - 1
	}
	return dice.Sroll("1d3") + 2
}
