package orbit

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/generation/star"
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

type starPlanetOrbits struct {
	stellar    string
	starLayout []int
	orb        map[int]*orbitPlanetary
}

type StarPlanetSystem interface {
	Stellar() string
	Layout() []int
	ByCode(int) *orbitPlanetary
}

func (po *starPlanetOrbits) Stellar() string {
	return po.stellar
}
func (po *starPlanetOrbits) Layout() []int {
	return po.starLayout
}
func (po *starPlanetOrbits) ByCode(code int) *orbitPlanetary {
	if val, ok := po.orb[code]; ok {
		return val
	}
	return nil
}

type orbitPlanetary struct {
	parentStar     string
	positionNumber int
	habitableZone  int
	systemPosition string
	au             float64
	comment        string
	//temperatureZone string
}

type planetOrbit struct {
	parent       star.StarBody
	positionCode int
	distance     float64
	info         string
}

func (po *planetOrbit) Position() int {
	return po.positionCode
}
func (po *planetOrbit) Distance() float64 {
	return po.distance
}
func (po *planetOrbit) Info() string {
	return po.info
}

type PlanetOrbit interface {
	Position() int
	Distance() float64
	Info() string
}

// type OrbitPlanetary interface {
// 	ParentStar() string
// 	PositionNumber() int
// 	HabitableZone() int
// 	SystemPosition() string
// 	AU() float64
// 	Comment() string
// }

// func (op *orbitPlanetary) ParentStar() string {
// 	return op.parentStar
// }
// func (op *orbitPlanetary) PositionNumber() int {
// 	return op.positionNumber
// }
// func (op *orbitPlanetary) HabitableZone() int {
// 	return op.habitableZone
// }
// func (op *orbitPlanetary) SystemPosition() string {
// 	return op.systemPosition
// }
// func (op *orbitPlanetary) AU() float64 {
// 	return op.au
// }
// func (op *orbitPlanetary) Comment() string {
// 	return op.comment
// }

func newStarPlanetOrbit(parentstar string, positionNumber int) *orbitPlanetary {
	orb := orbitPlanetary{}
	orb.parentStar = parentstar
	orb.positionNumber = positionNumber
	hz := stellar.HabitableOrbitByCode(parentstar)
	orb.habitableZone = positionNumber - hz
	switch orb.positionNumber {
	case 0, 1, 2, 3, 4, 5:
		orb.systemPosition = "Inner"
	case 6, 7, 8, 9, 10, 11, 12:
		orb.systemPosition = "Outer"
	case 13, 14, 15, 16, 17, 18, 19, 20:
		orb.systemPosition = "Remote"
	}
	//orb.temperatureZone = calculateTemperatureZone(orb.parentStar, orb.positionNumber)
	return &orb
}

func calculateTemperatureZone(parentStar string, body int) string {
	zone := stellar.HabitableOrbitByCode(parentStar) - body
	temperatureZone := ""
	switch {
	default:
		temperatureZone = "UNDEFINED"
		fmt.Println("star", parentStar)
		fmt.Println("")
	case strings.Contains(parentStar, "O"):
		temperatureZone = defineZone(zone, 5, 3)
	case strings.Contains(parentStar, "B"):
		temperatureZone = defineZone(zone, 4, 3)
	case strings.Contains(parentStar, "A"):
		temperatureZone = defineZone(zone, 3, 3)
	case strings.Contains(parentStar, "F"):
		temperatureZone = defineZone(zone, 2, 3)
	case strings.Contains(parentStar, "G"):
		temperatureZone = defineZone(zone, 1, 4)
	case strings.Contains(parentStar, "K"):
		temperatureZone = defineZone(zone, 1, 3)
	case strings.Contains(parentStar, "M"):
		temperatureZone = defineZone(zone, 1, 2)
	case strings.Contains(parentStar, "D"):
		temperatureZone = defineZone(zone, 5, 1)
	case strings.Contains(parentStar, "T"):
		temperatureZone = defineZone(zone, 5, 1)
	case strings.Contains(parentStar, "Y"):
		temperatureZone = defineZone(zone, 5, 1)
	case strings.Contains(parentStar, "L"):
		temperatureZone = defineZone(zone, 5, 1)
	}
	return temperatureZone
}

func defineZone(z int, hot, cold int) string {
	if z < 0 {
		if math.Abs(float64(z)) > math.Abs(float64(hot)) {
			return Temperature_FROZEN
		}
		return Temperature_COLD
	}
	if z > 0 {
		if math.Abs(float64(z)) > math.Abs(float64(cold)) {
			return Temperature_INFERNO
		}
		return Temperature_HOT
	}
	return Temperature_HABITABLE
}

func StarPlanetOrbits0(dice *dice.Dicepool, stellar stellar.Stellar) *starPlanetOrbits {
	sysOrb := starPlanetOrbits{}
	//sysOrb.dice = dice
	sysOrb.stellar = stellar.String()
	sysOrb.starLayout = stellar.Layout()[2:]
	orbMap := make(map[int]*orbitPlanetary)
	starTypes := stellar.Stars()
	for starNum := 0; starNum < len(starTypes); starNum++ {
		orbMap[(starNum+1)*-1] = nil
		for body := 0; body < 21; body++ {
			orb := newStarPlanetOrbit(starTypes[starNum], body)
			orb.calculateDecimalOrbit(dice)
			orb.comment = "???"
			strData := star.New(starTypes[starNum])
			if orb.au < strData.InnerLimit() {
				orb.comment = "vaporized"
				continue
			}
			if orb.au > strData.OuterLimit()*6 {
				orb.comment = "rogue"
			}
			orbID := sysOrb.starLayout[starNum]
			orbMap[(100*(orbID))+body] = orb
		}
	}

	sysOrb.orb = orbMap
	sysOrb.PlaceStars(dice)
	return &sysOrb
}

type StarPlanetOrbitsMap struct {
	ByCode map[int]orbitPlanetary
}

func StarPlanetOrbits(dice *dice.Dicepool, pair star.StarBody) map[int]*planetOrbit {
	spoMap := make(map[int]*planetOrbit)
	for i := 0; i <= 20; i++ {
		orbit := planetOrbit{}
		orbit.parent = pair
		orbit.positionCode = i
		orbit.distance = decimalOrbit2(i, dice.Flux())

		if orbit.distance < pair.InnerLimit() {
			orbit.info = "vaporized"
		}
		if orbit.distance >= pair.InnerLimit() && orbit.distance < pair.HabitableLow()/3*2 {
			orbit.info = "inferno"
		}
		if orbit.distance > pair.HabitableLow()/2 && orbit.distance < pair.HabitableLow() {
			orbit.info = "hot"
		}
		if orbit.distance > pair.HabitableLow() && orbit.distance < pair.HabitableHigh() {
			orbit.info = "habitable"
		}
		if orbit.distance > pair.HabitableHigh() {
			orbit.info = "cold"
		}
		if orbit.distance > pair.HabitableHigh()/3*5 {
			orbit.info = "frozen"
		}
		if orbit.distance > pair.OuterLimit()*8 {
			orbit.info = "rogue"
		}
		spoMap[orbit.positionCode] = &orbit
	}

	return spoMap
}

func setupplanetaryOrbits(dice *dice.Dicepool, parentStar string) []orbitPlanetary {
	return []orbitPlanetary{}
}

func comment(star, body int) string {
	bodyName := greekLetter(star) + " " + orbitCode(body)
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
func (o *orbitPlanetary) calculateDecimalOrbit(dice *dice.Dicepool) {
	o.au = decimalOrbit2(o.positionNumber, dice.Flux())
}

func decimalOrbit2(orb int, flux int) float64 {
	if orb < 0 || orb > 20 {
		return -1
	}
	do := [][]float64{
		{0.107, 0.114, 0.12, 0.127, 0.133, 0.14, 0.153, 0.166, 0.18, 0.193, 0.206}, //0
		{0.206, 0.219, 0.232, 0.245, 0.257, 0.27, 0.295, 0.321, 0.346, 0.372, 0.397},
		{0.397, 0.422, 0.447, 0.471, 0.496, 0.52, 0.569, 0.618, 0.667, 0.716, 0.765},
		{0.765, 0.812, 0.859, 0.906, 0.953, 1, 1.094, 1.189, 1.283, 1.377, 1.477},
		{1.477, 1.574, 1.666, 1.757, 1.849, 1.94, 2.123, 2.306, 2.488, 2.671, 2.868},
		{2.868, 3.06, 3.237, 3.415, 3.592, 3.77, 4.125, 4.481, 4.836, 5.191, 5.567}, //5
		{5.567, 5.932, 6.277, 6.621, 6.966, 7.31, 7.999, 8.688, 9.377, 10.07, 10.8},
		{10.8, 11.52, 12.18, 12.85, 13.52, 14.19, 15.53, 16.86, 18.2, 19.54, 20.96},
		{20.96, 22.35, 23.65, 24.94, 26.24, 27.54, 30.14, 32.73, 35.33, 37.92, 40.69},
		{40.69, 43.38, 45.9, 48.42, 50.94, 53.46, 58.5, 63.54, 68.57, 73.61, 78.98},
		{78.98, 84.2, 89.09, 93.97, 98.86, 103.8, 113.5, 123.3, 133.1, 142.9, 153.3}, //10
		{153.3, 163.4, 172.9, 182.4, 191.9, 201.4, 220.3, 239.3, 258.3, 277.3, 297.5},
		{297.5, 317.2, 335.6, 354, 372.4, 390.8, 427.7, 464.5, 501.3, 538.1, 577.4},
		{577.4, 615.6, 651.3, 687.1, 722.8, 758.6, 830, 901.5, 973, 1044, 1121},
		{1121, 1195, 1264, 1334, 1403, 1472, 1611, 1750, 1888, 2027, 2175},
		{2175, 2319, 2454, 2588, 2723, 2857, 3127, 3396, 3665, 3935, 4222}, //15
		{4222, 4501, 4762, 5023, 5285, 5546, 6069, 6591, 7114, 7636, 8194},
		{8194, 8736, 9243, 9750, 10257, 10764, 11778, 12793, 13807, 14821, 15903},
		{15903, 16955, 17939, 18923, 19908, 20892, 22861, 24829, 26798, 28766, 30866},
		{30866, 32907, 34817, 36728, 38638, 40549, 44370, 48190, 52011, 55832, 59907},
		{59907, 63868, 67576, 71284, 74992, 78700, 86116, 93532, 100948, 108364, 115780}, //20
	}
	return do[orb][flux+5]
}

func (plOrb *starPlanetOrbits) PlaceStars(dice *dice.Dicepool) {

	for _, l := range plOrb.starLayout {
		fmt.Println(":l", l)
		starpos := stellar.SetStarOrbit(dice, l)
		if starpos == -1 {
			continue
		}
		switch l {
		case 2, 4, 6, 8:

			orbID := ((l) * 100)
			for {
				if _, ok := plOrb.orb[orbID]; ok == false {
					orbID++
				} else {
					break
				}
			}
			plOrb.orb[orbID].comment = "star"
			for id := (l) * 100; id < ((l)*100 + 25); id++ {
				delete(plOrb.orb, id)
			}
		case 3:
			if starpos == 0 && plOrb.orb[100].comment == "star" { //TODO: возможна проблема
				plOrb.orb[1].comment = "star"
			} else {
				plOrb.orb[100+starpos].comment = "star"
			}
			plOrb.orb[100+starpos].comment = "star"
			for id := (l)*100 + (starpos - 2); id < ((l)*100 + 25); id++ {
				delete(plOrb.orb, id)
			}
		case 5, 7:
			plOrb.orb[100+starpos].comment = "star"
			for id := (l)*100 + (starpos - 2); id < ((l)*100 + 25); id++ {
				delete(plOrb.orb, id)
			}
		}

	}
}

func AllOrbitalCodes() []int {
	codes := []int{}
	for _, starPosition := range []int{0, 10000, 30000, 50000, 70000} {
		for _, planetPosition := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 50} {
			for _, satelitePosition := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26} {
				code := starPosition + (planetPosition * 100) + satelitePosition
				if code > 0 && code < 10000 && code != 5000 {
					continue
				}
				if code > 35000 && code < 50000 {
					continue
				}
				if code > 55000 && code < 70000 {
					continue
				}
				if code > 75000 {
					continue
				}
				codes = append(codes, code)
				fmt.Println(code)
				time.Sleep(time.Second)
			}
		}
	}
	return codes
}
