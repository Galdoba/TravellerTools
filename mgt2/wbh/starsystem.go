package wbh

import (
	"fmt"
	"strings"

	orbitns "github.com/Galdoba/TravellerTools/mgt2/wbh/orbits"
	"github.com/Galdoba/TravellerTools/mgt2/wbh/star"
	"github.com/Galdoba/TravellerTools/mgt2/wbh/worlds"
	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	defauiltVal = iota
	GenerationMethodSpecial
	GenerationMethodUnusual
	TypeVariantTraditional
	TypeVariantRealistic

	designationPrimary          = "Aa"
	designationPrimaryCompanion = "Ab"
	designationClose            = "Ba"
	designationCloseCompanion   = "Bb"
	designationNear             = "Ca"
	designationNearCompanion    = "Cb"
	designationFar              = "Da"
	designationFarCompanion     = "Db"
	determinationPrimary        = "Prime"
	determinationRandom         = "Random"
	determinationLesser         = "Lesser"
	determinationSibling        = "Sibling"
	determinationTwin           = "Twin"
	determinationOther          = "Other"
)

type StarSystem struct {
	starGenerationMethod int
	TypeTableVariant     int
	//primary              Star
	Star map[string]star.Star
	age  float64
}

func NewStarSystem(dice *dice.Dicepool, starGenerationMethod, tableVariant int) (*StarSystem, error) {
	ss := StarSystem{}
	switch starGenerationMethod {
	case GenerationMethodUnusual, GenerationMethodSpecial:
	default:
		return &ss, fmt.Errorf("starGenerationMethod unknown (%v)", starGenerationMethod)
	}
	ss.Star = make(map[string]star.Star)

	primary, err := star.New(dice, tableVariant, starGenerationMethod, designationPrimary, determinationPrimary)
	if err != nil {
		return &ss, err
	}
	ss.Star[designationPrimary] = primary
	designations := star.DefineStarPresence(ss.Star[designationPrimary], dice)
	for _, desig := range designations {
		if _, ok := ss.Star[desig]; ok {
			continue
		}
		determ, context := star.DefineStarDetermination(primary, desig, dice)
		star, err := star.New(dice, tableVariant, starGenerationMethod, desig, determ, ss.Star[context])
		if err != nil {
			return &ss, fmt.Errorf("secondary star %v creation: %v", desig, err.Error())
		}

		ss.Star[desig] = star
	}
	ss.ageResetIfRequired(dice)

	//TODO: dm for eccentricity
	dm := 0
	for _, desig := range star.DesignationCodes() {
		if _, ok := ss.Star[desig]; !ok {
			continue
		}
		st := ss.Star[desig]
		orbN, err := orbitns.DetermineStarOrbit(dice, desig)
		if err != nil {
			return &ss, fmt.Errorf("orbitns.DetermineStarOrbit: %v", err.Error())
		}
		st.Orbit = orbitns.New(orbN)
		st.Orbit.DetermineEccentrisity(dice, dm)
		//st.normalizeValues()
		ss.Star[desig] = st
	}
	ss.CalculateOP()
	ss.PlaceWorlds(dice)

	return &ss, nil
}

func systemStars(ss *StarSystem) []star.Star {
	stars := []star.Star{}
	for _, desig := range star.DesignationCodes() {
		if st, ok := ss.Star[desig]; ok {
			stars = append(stars, st)
		}
	}
	return stars
}

func (ss *StarSystem) CalculateOP() {
	for _, code := range []string{"Ba", "Bb", "Ca", "Cb", "Da", "Db"} {
		switch {
		case strings.Contains(code, "Aa"):
			continue
		case strings.Contains(code, "a"):
			far := strings.TrimSuffix(code, "a")
			if _, ok := ss.Star[far+"a"]; ok {
				star1 := ss.Star[far+"a"]
				star2 := ss.Star[far+"b"]
				star3 := ss.Star["Aa"]
				star4 := ss.Star["Ab"]
				au := (star.AUof(star2) + 2*star.AUof(star1)) / 2
				m1 := star1.Mass + star2.Mass
				m2 := star3.Mass + star4.Mass
				orbPer := orbitns.CalculateOrbitalPeriod(au, m1, m2)
				//star1.Orbit.Period = orbPer
				ss.Star[far+"a"].Orbit.Period = orbPer
			}
		case strings.Contains(code, "b"):
			far := strings.TrimSuffix(code, "b")
			if _, ok := ss.Star[far+"b"]; ok {
				star1 := ss.Star[far+"a"]
				star2 := ss.Star[far+"b"]
				au := star.AUof(star2)
				m1 := star1.Mass
				m2 := star2.Mass
				orbPer := orbitns.CalculateOrbitalPeriod(au, m1, m2)
				ss.Star[far+"b"].Orbit.Period = orbPer
			}
		}
	}

}

func (ss *StarSystem) ageResetIfRequired(dice *dice.Dicepool) {
	switch ss.Star["Aa"].Class {
	case star.ClassIa, star.ClassIb, star.ClassII, star.ClassIII, star.ClassIV, star.ClassV, star.ClassVI, star.ClassBD:
		for _, v := range ss.Star {
			switch v.Class {
			case star.ClassD, star.Pulsar, star.NeutronStar, star.BlackHole:
				primary := ss.Star["Aa"]
				primary.Age = v.Age
				ss.Star["Aa"] = primary

				// primary.Age = starFinalAge(v.Mass, dice)
				// if primary.Age < was {
				// 	fmt.Println("set new age")
				// 	primary.Age = was
				// }
				// if primary.Age > 13.5 {
				// 	fmt.Println("set age border", primary)

				// 	primary.Age = 13.5
				// }
				// ss.Star["Aa"] = primary
			}
		}
	}
	ss.age = ss.Star["Aa"].Age
}

func (ss *StarSystem) String() string {
	prf := fmt.Sprintf("%v-", len(ss.Star))
	for _, desig := range star.DesignationCodes() {
		if st, ok := ss.Star[desig]; ok {
			switch desig {
			case "Aa":
				prf += star.ShortStarDescription(st) //star.stType + star.subType + " " + star.Class
				prf += "-" + fmt.Sprintf("%v", st.Mass)
				prf += "-" + fmt.Sprintf("%v", st.Diameter)
				prf += "-" + fmt.Sprintf("%v", st.Luminocity)
				prf += "-" + fmt.Sprintf("%v", st.Age)
			default:
				prf += ":" + fmt.Sprintf("%v", desig)
				prf += "-" + fmt.Sprintf("%v", st.Orbit.OrbitNum)
				prf += "-" + fmt.Sprintf("%v", st.Orbit.Eccentricity)
				prf += "-" + star.ShortStarDescription(st)
				prf += "-" + fmt.Sprintf("%v", st.Mass)
				prf += "-" + fmt.Sprintf("%v", st.Diameter)
				prf += "-" + fmt.Sprintf("%v", st.Luminocity)
				//prf += "&" + fmt.Sprintf("%v", st.MAO)
			}
		}
	}
	prf = strings.TrimPrefix(prf, "1-")
	return prf
}

func (ss *StarSystem) PlaceWorlds(dice *dice.Dicepool) {
	gasGigantsNum := worlds.GasGigantsQuantity(dice, ss.Star)
	beltsNum := worlds.PlanetoidBeltsQuantity(dice, ss.Star, gasGigantsNum)
	rockyWorldsNum := worlds.TerrestialPlanetsQuantity(dice, ss.Star)
	totalWorlds := gasGigantsNum + beltsNum + rockyWorldsNum
	////
	fmt.Println("total Worlds:", totalWorlds)
	ss.Star = star.CalculateAllowableOrbits(ss.Star)
	ss.Star = star.DefineHZCO(ss.Star)
	allowed := star.AllocateWorldlimitsByStars(totalWorlds, ss.Star)

	sstmBaseNmbr := systemBaseNumber(ss, dice, totalWorlds)
	fmt.Println(sstmBaseNmbr, allowed)
	panic("STOPPED HERE")
	//determineBaselineOrbit
	//	a) BN > 1 && BN < totalWorlds
	//	b) BN < 1
	//	c) BN > totalWorlds
	//EmptyOrbits
	//SystemSpread
	//PlaceOrbits
	//AddAnomalousPlanets
	//PlaceWorlds
	//DetermineEccentricity
}

func systemBaseNumber(syst *StarSystem, dice *dice.Dicepool, totalWorlds int) int {
	r := dice.Sroll("2d6")
	if _, ok := syst.Star["Ab"]; !ok {
		r = r - 2
	}
	primary := syst.Star["Aa"]
	switch primary.Class {
	case star.ClassIa, star.ClassIb, star.ClassII:
		r = r + 3
	case star.ClassIII:
		r = r + 2
	case star.ClassIV:
		r = r + 1
	case star.ClassVI:
		r = r - 1
	case star.ClassD, star.Pulsar, star.NeutronStar, star.BlackHole:
		r = r - 2
	}
	switch totalWorlds {
	default:
		if totalWorlds < 6 {
			r = r - 4
		}
		if totalWorlds > 20 {
			r = r + 2
		}
	case 6, 7, 8, 9:
		r = r - 3
	case 10, 11, 12:
		r = r - 2
	case 13, 14, 15:
		r = r - 1
	case 18, 19, 20:
		r = r + 1
	}
	return r
}
