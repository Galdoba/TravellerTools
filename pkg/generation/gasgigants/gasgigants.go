package gasgigants

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

type gasGigant struct {
	size           ehex.Ehex
	diameter       int
	gType          string
	sateliteOrbits int
	ringOrbits     int
	parentStar     string
	star           int
	orbit          int
}

func Generate(dice *dice.Dicepool) []*gasGigant {
	ggNum := (dice.Sroll("2d6") / 2) - 2
	if ggNum <= 0 {
		return nil
	}
	ggData := []*gasGigant{}
	for i := 0; i < ggNum; i++ {
		gg := gasGigant{}

		gg.size = ehex.New().Set(dice.Sroll("2d6+19"))

		switch gg.size.Value() {
		case 20, 21, 22:
			switch dice.Sroll("1d2") {
			case 1:
				gg.gType = "Small Gas Gigant"
			case 2:
				gg.gType = "Ice Gigant"
			}
		default:
			gg.gType = "Large Gas Gigant"
		}

		satelitesPlased := false
		for !satelitesPlased {
			roll := dice.Sroll("1d6-1")
			switch roll {
			case 0:
				gg.ringOrbits++
			case 1, 2, 3, 4, 5:
				gg.sateliteOrbits = roll
				satelitesPlased = true
			}
		}
		ggData = append(ggData, &gg)
	}
	return ggData
}

func (gg *gasGigant) SystemPosition() (int, int, int) {
	return gg.star, gg.orbit, -1
}

func (gg *gasGigant) SetStar(i int) {
	gg.star = i
}

func (gg *gasGigant) SetOrbit(i int) {
	gg.orbit = i
}
