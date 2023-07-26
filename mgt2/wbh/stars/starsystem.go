package stars

import (
	"fmt"
	"strings"

	orbitns "github.com/Galdoba/TravellerTools/mgt2/wbh/orbits"
	"github.com/Galdoba/TravellerTools/pkg/dice"
)

type StarSystem struct {
	starGenerationMethod int
	TypeTableVariant     int
	//primary              Star
	Star map[string]Star
	age  float64
}

func NewStarSystem(dice *dice.Dicepool, starGenerationMethod, tableVariant int) (*StarSystem, error) {
	ss := StarSystem{}
	switch starGenerationMethod {
	case GenerationMethodUnusual, GenerationMethodSpecial:
	default:
		return &ss, fmt.Errorf("starGenerationMethod unknown (%v)", starGenerationMethod)
	}
	ss.Star = make(map[string]Star)

	primary, err := NewStar(dice, tableVariant, starGenerationMethod, designationPrimary, determinationPrimary)
	if err != nil {
		return &ss, err
	}
	ss.Star[designationPrimary] = primary
	designations := defineStarPresence(ss.Star[designationPrimary], dice)
	for _, desig := range designations {
		if _, ok := ss.Star[desig]; ok {
			continue
		}
		determ, context := defineStarDetermination(primary, desig, dice)
		star, err := NewStar(dice, tableVariant, starGenerationMethod, desig, determ, ss.Star[context])
		if err != nil {
			return &ss, fmt.Errorf("secondary star %v creation: %v", desig, err.Error())
		}

		ss.Star[desig] = star
	}
	ss.ageResetIfRequired(dice)

	//TODO: dm for eccentricity
	dm := 0
	for desig, st := range ss.Star {
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
	return &ss, nil
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
				au := (AUof(star2) + 2*AUof(star1)) / 2
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
				au := AUof(star2)
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
	case ClassIa, ClassIb, ClassII, ClassIII, ClassIV, ClassV, ClassVI, ClassBD:
		for _, v := range ss.Star {
			switch v.Class {
			case ClassD, Pulsar, NeutronStar, BlackHole:
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
	for _, desig := range designationCodes() {
		if star, ok := ss.Star[desig]; ok {
			switch desig {
			case "Aa":
				prf += shortStarDescription(star) //star.stType + star.subType + " " + star.Class
				prf += "-" + fmt.Sprintf("%v", star.Mass)
				prf += "-" + fmt.Sprintf("%v", star.Diameter)
				prf += "-" + fmt.Sprintf("%v", star.Luminocity)
				prf += "-" + fmt.Sprintf("%v", star.Age)
			default:
				prf += ":" + fmt.Sprintf("%v", desig)
				prf += "-" + fmt.Sprintf("%v", star.Orbit.OrbitNum)
				prf += "-" + fmt.Sprintf("%v", star.Orbit.Eccentricity)
				prf += "-" + shortStarDescription(star)
				prf += "-" + fmt.Sprintf("%v", star.Mass)
				prf += "-" + fmt.Sprintf("%v", star.Diameter)
				prf += "-" + fmt.Sprintf("%v", star.Luminocity)
				//prf += "&" + fmt.Sprintf("%v", star.Orbit.Period)
			}
		}
	}
	prf = strings.TrimPrefix(prf, "1-")
	return prf
}
