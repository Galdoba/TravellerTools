package gasgigants

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

const (
	LGG = "Large Gas Gigant"
	SGG = "Small Gas Gigant"
	IG  = "Ice Gigant"
)

type gasGigant struct {
	size     ehex.Ehex
	diameter int
	gType    string
	//sateliteOrbits int
	//ringOrbits     int
	//orbitData      orbit.Orbiter
}

func (gg *gasGigant) GGType() string {
	return gg.gType
}

func (gg *gasGigant) Size() ehex.Ehex {
	return gg.size
}

func New(dice *dice.Dicepool) *gasGigant {
	gg := gasGigant{}
	gg.size = ehex.New().Set(dice.Sroll("2d6+19"))
	switch gg.size.Value() {
	case 20, 21, 22:
		switch dice.Sroll("1d2") {
		case 1:
			gg.gType = SGG
		case 2:
			gg.gType = IG
		}
	default:
		gg.gType = LGG
	}
	return &gg
}

func OfferOrbit(dice *dice.Dicepool, ggtype string, hz int) int {
	n := dice.Sroll("2d6")
	switch ggtype {
	case LGG:
		return offerLGGorbit(n) + hz
	case SGG:
		return offerSGGorbit(n) + hz
	case IG:
		return offerIGorbit(n) + hz

	}
	return -1
}

func offerLGGorbit(i int) int {
	return []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7}[i]
}

func offerSGGorbit(i int) int {
	return []int{-4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8}[i]
}

func offerIGorbit(i int) int {
	return []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}[i]
}
