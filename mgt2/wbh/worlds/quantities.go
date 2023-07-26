package worlds

import (
	"github.com/Galdoba/TravellerTools/mgt2/wbh/helper"
	"github.com/Galdoba/TravellerTools/mgt2/wbh/stars"
	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func GasGigantsQuantity(dice *dice.Dicepool, star map[string]stars.Star) int {
	r1 := dice.Sroll("2d6")
	if r1 > 9 {
		return 0
	}
	dm := 0
	if len(star) == 1 && star["Aa"].Class == stars.ClassV {
		dm = dm + 1
	}
	if len(star) > 3 {
		dm = dm - 1
	}
	switch star["Aa"].Class {
	case stars.ClassD, stars.Pulsar, stars.NeutronStar, stars.BlackHole, stars.ClassBD:
		dm = dm - 2
	}
	for _, st := range star {
		switch st.Class {
		case stars.ClassD, stars.Pulsar, stars.NeutronStar, stars.BlackHole:
			dm = dm - 1
		}
	}
	r2 := helper.EnsureMinMax(dice.Sroll("2d6")+dm, 4, 13)

	switch r2 {
	default:
		return -1
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
	case 13:
		return 6
	}
}

func PlanetoidBeltsQuantity(dice *dice.Dicepool, star map[string]stars.Star, ggNum int) int {
	r1 := dice.Sroll("2d6")
	if r1 < 8 {
		return 0
	}
	dm := 0
	if ggNum > 0 {
		dm = dm + 1
	}
	///////////////
	// if len(star) == 1 && star["Aa"].Class == stars.ClassV {
	// 	dm = dm + 1
	// }
	// if len(star) > 3 {
	// 	dm = dm - 1
	// }
	// switch star["Aa"].Class {
	// case stars.ClassD, stars.Pulsar, stars.NeutronStar, stars.BlackHole, stars.ClassBD:
	// 	dm = dm - 2
	// }
	// for _, st := range star {
	// 	switch st.Class {
	// 	case stars.ClassD, stars.Pulsar, stars.NeutronStar, stars.BlackHole:
	// 		dm = dm - 1
	// 	}
	// }
	// r2 := helper.EnsureMinMax(dice.Sroll("2d6")+dm, 4, 13)

	// switch r2 {
	// default:
	// 	return -1
	// case 4:
	// 	return 1
	// case 5, 6:
	// 	return 2
	// case 7, 8:
	// 	return 3
	// case 9, 10, 11:
	// 	return 4
	// case 12:
	// 	return 5
	// case 13:
	// 	return 6
	// }
	return -1
}
