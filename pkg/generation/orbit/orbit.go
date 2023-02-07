package orbit

import (
	"fmt"
	"math"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/generation/stellar"
	"github.com/Galdoba/utils"
)

const (
	Position_INNER        = "Inner"
	Position_OUTER        = "Outer"
	Position_REMOTE       = "Remote"
	Temperature_INFERNO   = "Inferno"
	Temperature_HOT       = "Hot"
	Temperature_HABITABLE = "Habitable"
	Temperature_COLD      = "Cold"
	Temperature_FROZEN    = "Frozen"
)

type orbit struct {
	parentStar      string
	starNum         int
	bodyNum         int
	satNumb         int
	systemPosition  string
	temperatureZone string
	comment         string
	au              float64
	bodyname        string
}

func (orb *orbit) TemperatureZone() string {
	return orb.temperatureZone
}

func New(parentstar string, star, body, satelite int) *orbit {
	orb := orbit{}
	orb.parentStar = parentstar
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
	orb.calculateTemperatureZone()
	return &orb
}

func (orb *orbit) calculateTemperatureZone() {
	zone := stellar.HabitableOrbitByCode(orb.parentStar) - orb.bodyNum
	switch {
	default:
		orb.temperatureZone = "UNDEFINED"
		fmt.Println("star", orb.parentStar)
		fmt.Println("")
	case strings.Contains(orb.parentStar, "O"):
		orb.temperatureZone = defineZone(zone, 5, 3)
	case strings.Contains(orb.parentStar, "B"):
		orb.temperatureZone = defineZone(zone, 4, 3)
	case strings.Contains(orb.parentStar, "A"):
		orb.temperatureZone = defineZone(zone, 3, 3)
	case strings.Contains(orb.parentStar, "F"):
		orb.temperatureZone = defineZone(zone, 2, 3)
	case strings.Contains(orb.parentStar, "G"):
		orb.temperatureZone = defineZone(zone, 1, 4)
	case strings.Contains(orb.parentStar, "K"):
		orb.temperatureZone = defineZone(zone, 1, 3)
	case strings.Contains(orb.parentStar, "M"):
		orb.temperatureZone = defineZone(zone, 1, 2)
	case strings.Contains(orb.parentStar, "D"):
		orb.temperatureZone = defineZone(zone, 5, 1)
	case strings.Contains(orb.parentStar, "T"):
		orb.temperatureZone = defineZone(zone, 5, 1)
	case strings.Contains(orb.parentStar, "Y"):
		orb.temperatureZone = defineZone(zone, 5, 1)
	case strings.Contains(orb.parentStar, "L"):
		orb.temperatureZone = defineZone(zone, 5, 1)
	}
}

func defineZone(z int, hot, cold int) string {
	if z < 0 {
		if math.Abs(float64(z)) > math.Abs(float64(hot)) {
			return Temperature_INFERNO
		}
		return Temperature_HOT
	}
	if z > 0 {
		if math.Abs(float64(z)) > math.Abs(float64(cold)) {
			return Temperature_FROZEN
		}
		return Temperature_COLD
	}
	return Temperature_HABITABLE
}

func Generate(dice *dice.Dicepool, stellarCode string) map[string]*orbit {
	orbMap := make(map[string]*orbit)
	starTypes := stellar.Parse(stellarCode)
	for star := 0; star < len(starTypes); star++ {
		for body := -1; body < 18; body++ {
			orb := New(starTypes[star], star, body, -1)
			orb.CalculateDecimalOrbit(dice)
			orbMap[orb.bodyname] = orb

		}
	}

	switch len(starTypes) {
	case 0:
		for k, o := range orbMap {
			if o.starNum != -1 && o.bodyNum != -1 && o.satNumb != -1 {
				delete(orbMap, k)
			}
		}
	default:
		for k, o := range orbMap {
			if o.starNum+1 > len(starTypes) || o.starNum == -1 {
				delete(orbMap, k)
				continue
			}
			orbMap[k].parentStar = starTypes[o.starNum]
			for sat := 0; sat < 26; sat++ {
				sOrb := AddSatellite(orbMap[k], sat)
				orbMap[sOrb.bodyname] = sOrb
			}
		}
	}

	return orbMap
}

func AddSatellite(orb *orbit, satellite int) *orbit {
	return New(orb.parentStar, orb.starNum, orb.bodyNum, satellite)
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
	TemperatureZone() string
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

////////////
//TODO: создать карту орбит для звезды
func AddStar(star string, pos int) *orbit {
	//return New(orb.parentStar, orb.starNum, orb.bodyNum, satellite)
	return nil
}

func SetupMap() map[string]*orbit {
	return make(map[string]*orbit)
}
