package orbit

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/generation/stellar"
	"github.com/Galdoba/utils"
)

type orbit struct {
	parentStar     string
	starNum        int
	bodyNum        int
	satNumb        int
	systemPosition string
	comment        string
	au             float64
	bodyname       string
}

func New(star, body, satelite int) *orbit {
	orb := orbit{}
	orb.starNum = star
	orb.bodyNum = body
	orb.satNumb = satelite
	orb.bodyname = bodyname(orb.starNum, orb.bodyNum, orb.satNumb)
	switch orb.bodyNum {
	case 0, 1, 2, 3, 4, 5:
		orb.systemPosition = "Inner"
	case 6, 7, 8, 9, 10, 11, 12:
		orb.systemPosition = "Outer"
	case 13, 14, 15, 16, 17:
		orb.systemPosition = "Remote"
	}
	return &orb
}

func Generate(dice *dice.Dicepool, stellarCode string) map[string]*orbit {
	orbMap := make(map[string]*orbit)
	for star := 0; star < 5; star++ {
		for body := -1; body < 18; body++ {
			orb := New(star, body, -1)
			orb.CalculateDecimalOrbit(dice)
			orbMap[orb.bodyname] = orb

		}
	}

	stars := stellar.Parse(stellarCode)
	switch len(stars) {
	case 0:
		for k, o := range orbMap {
			if o.starNum != -1 && o.bodyNum != -1 && o.satNumb != -1 {
				delete(orbMap, k)
			}
		}
	default:
		for k, o := range orbMap {
			if o.starNum+1 > len(stars) || o.starNum == -1 {
				delete(orbMap, k)
				continue
			}
			orbMap[k].parentStar = stars[o.starNum]
			for sat := 0; sat < 26; sat++ {
				sOrb := AddSatellite(orbMap[k], sat)
				orbMap[sOrb.bodyname] = sOrb
			}
		}
	}

	return orbMap
}

func AddSatellite(orb *orbit, satellite int) *orbit {
	sOrb := orbit{}
	sOrb.starNum = orb.starNum
	sOrb.bodyNum = orb.bodyNum
	sOrb.satNumb = satellite
	sOrb.bodyname = bodyname(sOrb.starNum, sOrb.bodyNum, sOrb.satNumb)
	sOrb.parentStar = orb.parentStar
	sOrb.systemPosition = orb.systemPosition
	sOrb.au = orb.au
	return &sOrb
}

func bodyname(star, body, satellite int) string {
	bodyName := greekLetter(star) + " " + orbitCode(body) + englishLetter(satellite)
	bodyName = strings.TrimSuffix(bodyName, " ")
	bodyName = strings.TrimPrefix(bodyName, " ")
	return bodyName
}

func greekLetter(i int) string {
	switch i {
	default:
		return ""
	case 0:
		return "Alpha"
	case 1:
		return "Beta"
	case 2:
		return "Gamma"
	case 3:
		return "Delta"
	case 4:
		return "Epsilon"
	case 5:
		return "Zeta"
	case 6:
		return "Eta"
	case 7:
		return "Theta"
	case 8:
		return "Iota"
	case 9:
		return "Kappa"
	case 10:
		return "Lambda"
	case 11:
		return "Mu"
	case 12:
		return "Nu"
	case 13:
		return "Xi"
	case 14:
		return "Omicron"
	case 15:
		return "Pi"
	case 16:
		return "Rho"
	case 17:
		return "Sigma"
	case 18:
		return "Tau"
	case 19:
		return "Upsilon"
	case 20:
		return "Phi"
	case 21:
		return "Chi"
	case 22:
		return "Psi"
	case 23:
		return "Omega"
	}
}
func orbitCode(i int) string {
	switch i {
	default:
		return ""
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
		return fmt.Sprintf("%v", i)
	}
}
func englishLetter(i int) string {
	switch i {
	default:
		return ""
	case 0:
		return "a"
	case 1:
		return "b"
	case 2:
		return "c"
	case 3:
		return "d"
	case 4:
		return "e"
	case 5:
		return "f"
	case 6:
		return "g"
	case 7:
		return "h"
	case 8:
		return "i"
	case 9:
		return "j"
	case 10:
		return "k"
	case 11:
		return "l"
	case 12:
		return "m"
	case 13:
		return "n"
	case 14:
		return "o"
	case 15:
		return "p"
	case 16:
		return "q"
	case 17:
		return "r"
	case 18:
		return "s"
	case 19:
		return "t"
	case 20:
		return "u"
	case 21:
		return "v"
	case 22:
		return "w"
	case 23:
		return "x"
	case 24:
		return "y"
	case 25:
		return "z"

	}
}
func (o *orbit) CalculateDecimalOrbit(dice *dice.Dicepool) {
	o.au = decimalOrbit(dice, o.bodyNum)
}

func decimalOrbit(dice *dice.Dicepool, body int) float64 {
	flux := dice.Flux()
	orbitBase := []float64{0.2, 0.4, 0.7, 1.0, 1.6, 2.8, 5.2, 10.0, 20.0, 40.0, 77.0, 154.0, 308.0, 615.0, 1230.0, 2458.0, 4916.0, 9830.0}
	switch body {
	default:
		return -1.0
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17:
		orbit := orbitBase[body]
		switch flux {
		case -5, -4, -3, -2, -1:
			orbit = orbit * ((100.0 + (5.0 * float64(flux))) / 100.0)
		case 5, 4, 3, 2, 1:
			orbit = orbit * ((100.0 + (10.0 * float64(flux))) / 100.0)
		}
		return utils.RoundFloat64(orbit, 2)
	}
}

type Orbiter interface {
	SystemPosition() (int, int, int)
	BodyName() string
	AU() float64
	ParentStar() string
}

func (o *orbit) SystemPosition() (int, int, int) {
	return o.starNum, o.bodyNum, o.satNumb
}
func (o *orbit) BodyName() string {
	return o.bodyname
}
func (o *orbit) AU() float64 {
	return o.au
}
func (o *orbit) ParentStar() string {
	return o.bodyname
}
