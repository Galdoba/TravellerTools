package orbit

import (
	"fmt"
	"math"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/generation/stellar"
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

type SystemOrbits struct {
	orb map[string]*orbit
}

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

func Generate(dice *dice.Dicepool, stellarCode string) SystemOrbits {
	sysOrb := SystemOrbits{}

	orbMap := make(map[string]*orbit)
	starTypes := stellar.Parse(stellarCode)
	for star := 0; star < len(starTypes); star++ {
		for body := -1; body < 21; body++ {
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
				sOrb.au = orbMap[k].au
				orbMap[sOrb.bodyname] = sOrb
			}
		}
	}
	sysOrb.orb = orbMap
	return sysOrb
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
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20:
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

	o.au = decimalOrbit2(o.bodyNum, dice.Flux())
	fmt.Println(o.au, "----")
}

// func decimalOrbit(dice *dice.Dicepool, body int) float64 {
// 	flux := dice.Flux()
// 	orbitBase := []float64{0.2, 0.4, 0.7, 1.0, 1.6, 2.8, 5.2, 10.0, 20.0, 40.0, 77.0, 154.0, 308.0, 615.0, 1230.0, 2458.0, 4916.0, 9830.0}
// 	switch body {
// 	default:
// 		return -1.0
// 	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17:
// 		orbit := orbitBase[body]
// 		switch flux {
// 		case -5, -4, -3, -2, -1:
// 			orbit = orbit * ((100.0 + (5.0 * float64(flux))) / 100.0)
// 		case 5, 4, 3, 2, 1:
// 			orbit = orbit * ((100.0 + (10.0 * float64(flux))) / 100.0)
// 		}
// 		return utils.RoundFloat64(orbit, 2)
// 	}
// }

func decimalOrbit2(orb int, flux int) float64 {
	if orb < 0 || orb > 20 {
		return -1
	}
	do := [][]float64{
		[]float64{0.107, 0.114, 0.12, 0.127, 0.133, 0.14, 0.153, 0.166, 0.18, 0.193, 0.206}, //0
		[]float64{0.206, 0.219, 0.232, 0.245, 0.257, 0.27, 0.295, 0.321, 0.346, 0.372, 0.397},
		[]float64{0.397, 0.422, 0.447, 0.471, 0.496, 0.52, 0.569, 0.618, 0.667, 0.716, 0.765},
		[]float64{0.765, 0.812, 0.859, 0.906, 0.953, 1, 1.094, 1.189, 1.283, 1.377, 1.477},
		[]float64{1.477, 1.574, 1.666, 1.757, 1.849, 1.94, 2.123, 2.306, 2.488, 2.671, 2.868},
		[]float64{2.868, 3.06, 3.237, 3.415, 3.592, 3.77, 4.125, 4.481, 4.836, 5.191, 5.567}, //5
		[]float64{5.567, 5.932, 6.277, 6.621, 6.966, 7.31, 7.999, 8.688, 9.377, 10.07, 10.8},
		[]float64{10.8, 11.52, 12.18, 12.85, 13.52, 14.19, 15.53, 16.86, 18.2, 19.54, 20.96},
		[]float64{20.96, 22.35, 23.65, 24.94, 26.24, 27.54, 30.14, 32.73, 35.33, 37.92, 40.69},
		[]float64{40.69, 43.38, 45.9, 48.42, 50.94, 53.46, 58.5, 63.54, 68.57, 73.61, 78.98},
		[]float64{78.98, 84.2, 89.09, 93.97, 98.86, 103.8, 113.5, 123.3, 133.1, 142.9, 153.3}, //10
		[]float64{153.3, 163.4, 172.9, 182.4, 191.9, 201.4, 220.3, 239.3, 258.3, 277.3, 297.5},
		[]float64{297.5, 317.2, 335.6, 354, 372.4, 390.8, 427.7, 464.5, 501.3, 538.1, 577.4},
		[]float64{577.4, 615.6, 651.3, 687.1, 722.8, 758.6, 830, 901.5, 973, 1044, 1121},
		[]float64{1121, 1195, 1264, 1334, 1403, 1472, 1611, 1750, 1888, 2027, 2175},
		[]float64{2175, 2319, 2454, 2588, 2723, 2857, 3127, 3396, 3665, 3935, 4222}, //15
		[]float64{4222, 4501, 4762, 5023, 5285, 5546, 6069, 6591, 7114, 7636, 8194},
		[]float64{8194, 8736, 9243, 9750, 10257, 10764, 11778, 12793, 13807, 14821, 15903},
		[]float64{15903, 16955, 17939, 18923, 19908, 20892, 22861, 24829, 26798, 28766, 30866},
		[]float64{30866, 32907, 34817, 36728, 38638, 40549, 44370, 48190, 52011, 55832, 59907},
		[]float64{59907, 63868, 67576, 71284, 74992, 78700, 86116, 93532, 100948, 108364, 115780}, //20
	}
	fmt.Println("---", orb, flux)
	return do[orb][flux+5]
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

// //////////
// TODO: создать карту орбит для звезды
func AddStar(star string, pos int) *orbit {
	//return New(orb.parentStar, orb.starNum, orb.bodyNum, satellite)
	return nil
}

func SetupMap() map[string]*orbit {
	return make(map[string]*orbit)
}

/*
0
1
2
0	0.16	0.17	0.18	0.19	0.20	0.22	0.23	0.24	0.25	0.26	0.27
1	0.27	0.29	0.31	0.32	0.34	0.36	0.38	0.40	0.41	0.43	0.45
2	0.45	0.48	0.51	0.54	0.57	0.60	0.63	0.66	0.69	0.72	0.75
3	0.75	0.80	0.85	0.90	0.95	1.00	1.05	1.10	1.15	1.20	1.25
4	1.25	1.31	1.38	1.44	1.50	1.56	1.63	1.69	1.75	1.81	1.89
5
*/
