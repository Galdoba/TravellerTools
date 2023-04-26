package starsystem

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/generation/belts"
	"github.com/Galdoba/TravellerTools/pkg/generation/gasgigants"
	"github.com/Galdoba/TravellerTools/pkg/generation/orbit"
	"github.com/Galdoba/TravellerTools/pkg/generation/pbg"
	"github.com/Galdoba/TravellerTools/pkg/generation/planets"
	"github.com/Galdoba/TravellerTools/pkg/generation/star"
	"github.com/Galdoba/TravellerTools/pkg/generation/stellar"
)

type starSystem struct {
	stellar       stellar.Stellar
	stars         []string
	primaryStar   star.StarBody
	closeStar     star.StarBody
	nearStar      star.StarBody
	farStar       star.StarBody
	starPlanetSys orbit.StarPlanetSystem
	pbg           pbg.PBG
	reservedFor   map[int]planets.PlanetaryBody
	//planets []planets.planet
}

func New(dice *dice.Dicepool) *starSystem {
	sts := starSystem{}
	sts.reservedFor = make(map[int]planets.PlanetaryBody)
	stel, _ := stellar.ConstructNew(stellar.CONSTRUCTOR_PARADIGM_T5, dice)
	sts.stellar = stel
	starMap := stellar.StarMap(sts.stellar)
	for i := 1; i <= 4; i++ {
		pair, err := star.NewPair(starMap[(i*2)-1], starMap[(i*2)])
		if err != nil {
			continue
		}
		if pair == nil {
			continue
		}
		switch i {
		case 1:
			sts.primaryStar = pair
		case 2:
			sts.closeStar = pair
		case 3:
			sts.nearStar = pair
		case 4:
			sts.farStar = pair
		}
		for k, v := range orbit.StarPlanetOrbits(dice, pair) {
			if v.Distance() >= pair.InnerLimit() {
				sts.reservedFor[(20000*i)-10000+(k*100)] = nil
			}
		}
		starOrbit := star.RollStarOrbit(dice, i-1) * 100
		sts.reservedFor[10000+starOrbit] = planets.SetupStar(pair.Class())
		for k := range sts.reservedFor {
			if k > (20000*i)-10000+(starOrbit-300) && k > 29999 {
				delete(sts.reservedFor, k)
			}
		}
	}
	sts.pbg = pbg.New(dice)
	sts.PlaceGasGigants(dice)
	sts.PlaceBelts(dice)
	sts.PlacePlanets(dice)
	return &sts
}

func (sts *starSystem) PlaceStars(dice *dice.Dicepool) {

}

func (sts *starSystem) PlacePlanets(dice *dice.Dicepool) {
	starCode := 100
	diceCode := fmt.Sprintf("2d6")
	expectedPlanets := dice.Sroll(diceCode)
	fol := sts.freeOrbitsLeft()
	if expectedPlanets > fol {
		expectedPlanets = fol
	}
	for i := 0; i < expectedPlanets; i++ {
	plan:
		for {
			star := ""
			for i := 0; i < dice.Sroll("2d6"); i++ {
				starCode = pickStar(sts, dice)
			}
			star = sts.starByCode(starCode)
			starHZ := stellar.HabitableOrbitByCode(star)
			n := -999
			switch i {
			case expectedPlanets - 1:
				n = planets.OfferWorldOrbit2(dice.Sroll("2d6"))
			default:
				n = planets.OfferWorldOrbit(dice.Sroll("2d6"))
			}
			testOrbits := expandOrbitNum(n)
			for _, o := range testOrbits {
				oCode := o * 100
				if sts.orbitIsFree(starCode + oCode) {
					sts.reservedFor[starCode+oCode] = planets.PhysicalData_T5(dice, o-starHZ, star)
					break plan
				}
			}
			// if bod, ok := sts.reservedFor[starCode+n]; ok {
			// 	if bod == nil {
			// 		sts.reservedFor[starCode+n] = planets.PhysicalData_T5(dice, n-starHZ, star)
			// 		break

			// 	}
			// 	continue plan

			// }
			continue plan
		}
	}
}

func expandOrbitNum(n int) []int {
	exp := []int{n}
	for i := 1; i < 21; i++ {
		exp = append(exp, n-i)
		exp = append(exp, n+i)
	}
	filtered := []int{}
	for _, o := range exp {
		if o >= 0 && o <= 20 {
			filtered = append(filtered, o)
		}
	}
	return filtered
}

func (sts *starSystem) orbitIsFree(code int) bool {
	if bod, ok := sts.reservedFor[code]; ok {
		if bod == nil {
			return true
		}
	}
	return false
}

func (sts *starSystem) starByCode(code int) string {
	switch {
	default:
		return sts.primaryStar.Class()

	case code > 69900:
		return sts.farStar.Class()

	case code > 49900:
		return sts.nearStar.Class()
	case code > 29900:
		return sts.closeStar.Class()
	}
}

func (sts *starSystem) nextStarCode(current int) int {
	for k, v := range sts.reservedFor {
		if v == nil {
			continue
		}
		if v.PlType() == planets.Star && k != current {
			return k
		}
	}
	return current
}

func (sts *starSystem) PlaceGasGigants(dice *dice.Dicepool) {
	starCode := 100
	for i := 0; i < sts.pbg.GasGigants().Value(); i++ {
		gg := gasgigants.New(dice)
	thisGG:
		for {
			star := ""
			for i := 0; i < dice.Sroll("2d6"); i++ {
				starCode = pickStar(sts, dice)
			}
			star = sts.starByCode(starCode)
			hz := stellar.HabitableOrbitByCode(star)
			n := gasgigants.OfferOrbit(dice, gg.GGType(), hz)
			exp := expandOrbitNum(n)
			for _, o := range exp {
				oCode := o * 100
				if !sts.orbitIsFree(starCode + oCode) {
					continue
				}
				sts.reservedFor[starCode+oCode] = planets.SetupGasGigant(star, gg.Size(), gg.GGType(), o-hz)
				break thisGG
			}
		}
	}
}

func (sts *starSystem) PlaceBelts(dice *dice.Dicepool) {
	starCode := 100

	for i := 0; i < sts.pbg.Belts().Value(); i++ {
	belt:
		for {
			star := ""
			for i := 0; i < dice.Sroll("2d6"); i++ {
				starCode = pickStar(sts, dice)
			}
			star = sts.starByCode(starCode)
			hz := stellar.HabitableOrbitByCode(star)

			n := belts.OfferBeltOrbit(dice.Flux() + 5)
			exp := expandOrbitNum(n)
			for _, o := range exp {
				oCode := o * 100
				if !sts.orbitIsFree(starCode + oCode) {
					continue
				}
				sts.reservedFor[starCode+oCode] = planets.SetupBelt(star, o-hz)
				break belt
			}

		}
	}
}

func primary(stel string) string {
	stars := stellar.Parse(stel)
	return stars[0]
}

func (sts *starSystem) freeOrbitsLeft() int {
	fol := 0
	for _, v := range sts.reservedFor {
		if v == nil {
			fol++
		}
	}
	return fol
}

func pickStar(sts *starSystem, dice *dice.Dicepool) int {
	avail := []int{}
	for i, pair := range []star.StarBody{sts.primaryStar, sts.closeStar, sts.nearStar, sts.farStar} {
		if pair == nil {
			continue
		}
		avail = append(avail, i)
	}
	s := avail[dice.Sroll(fmt.Sprintf("1d%v", len(avail)))-1]
	return s*20000 + 10000
}
