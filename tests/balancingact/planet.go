package balancingact

import (
	"fmt"
	"math"
	"strings"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/internal/ehex"
	"github.com/Galdoba/TravellerTools/pkg/astrogation"
)

type planet struct {
	name        string
	currrentUWP string
	starport    string
	atrib       map[int]int
	coordX      int
	coordY      int
}

type SurveyData interface {
	MW_Name() string
	MW_UWP() string
	CoordX() int
	CoordY() int
}

func createPlanet(name, uwp string, x, y int) (*planet, error) {
	p := planet{}
	p.name = name
	p.currrentUWP = uwp
	p.coordX, p.coordY = x, y
	creator := dice.New().SetSeed(fmt.Sprintf("n%vu%vx%vy%v", name, uwp, x, y))
	p.atrib = make(map[int]int)
	for _, key := range []int{Solidarity, Wealth, Expansion, Might, Development} {
		atr := creator.Roll("2d6").DM(-2).Sum()
		if key == Development && atr > p.atrib[Wealth] {
			atr = p.atrib[Wealth]
		}
		p.atrib[key] = atr
	}
	p.updateUWP(p.currrentUWP)
	return &p, nil
}

func (pl *planet) updateUWP(newUWP string) {
	pl.currrentUWP = newUWP
	uwpData := strings.Split(pl.currrentUWP, "")
	for i, val := range uwpData {
		switch i {
		default:
			continue
		case 0:
			switch val {
			default:
				pl.atrib[Starport] = UnlistedValue
			case "A":
				pl.atrib[Starport] = STPRT_A
			case "B":
				pl.atrib[Starport] = STPRT_B
			case "C":
				pl.atrib[Starport] = STPRT_C
			case "D":
				pl.atrib[Starport] = STPRT_D
			case "E":
				pl.atrib[Starport] = STPRT_E
			case "X":
				pl.atrib[Starport] = STPRT_X
			}
		case 4:
			pl.atrib[Population] = ehex.New().Set(val).Value()
		case 5:
			pl.atrib[Govr] = ehex.New().Set(val).Value()
		case 6:
			pl.atrib[Law] = ehex.New().Set(val).Value()
		}
	}
}

func (pl *planet) String() string {
	str := fmt.Sprintf("Planet %v - %v (%v,%v)\n", pl.name, pl.currrentUWP, pl.coordX, pl.coordY)
	str += fmt.Sprintf("SOL %v  WLT %v  EXP %v  MGT %v  DEV %v\n", pl.atrib[Solidarity], pl.atrib[Wealth], pl.atrib[Expansion], pl.atrib[Might], pl.atrib[Development])
	str += fmt.Sprintf("Global Planetary Product: %v/%v\n", pl.GPP(), roundFl64(pl.GPP()/52.0, 1))
	str += fmt.Sprintf("Planetary Market : %v/%v\n", pl.PM(), roundFl64(pl.PM()/52, 1))
	str += fmt.Sprintf("Goverment Ship Building Budget : %v\n", pl.ShipBuildingBudget())
	str += fmt.Sprintf("Trade Factor : %v\n", pl.TradeFactor())
	return str
}

func ImportPlanet(sd SurveyData) (*planet, error) {
	pl, err := createPlanet(sd.MW_Name(), sd.MW_UWP(), sd.CoordX(), sd.CoordY())
	return pl, err
}

//GPP - returns yearly Global Planetary Product in MCr
func (pl *planet) GPP() float64 {
	gpp := pl.atrib[Wealth] * pl.atrib[Solidarity] * pl.atrib[Population]
	return roundFl64(float64(gpp), 1)
}

//PM - returns yearly Planetary Market in MCr
func (pl *planet) PM() float64 {
	pm := pl.atrib[Development] * pl.atrib[Solidarity] * pl.atrib[Population]
	return roundFl64(float64(pm), 1)
}

//ShipBuildingBudget - returns yearly Planetary Market in MCr
func (pl *planet) ShipBuildingBudget() float64 {
	sbb := pl.GPP() * 0.1 * float64(pl.atrib[Expansion]) / 52
	return roundFl64(sbb, 1)
}

//TradeFactor - returns yearly Planetary TradeFactor in MCr
func (pl *planet) TradeFactor() float64 {
	ll := float64(pl.atrib[Law])
	if ll == 0 {
		ll = 0.95
	}
	stp := 0.0
	switch pl.atrib[Starport] {
	case STPRT_A:
		stp = 3
	case STPRT_B:
		stp = 1.5
	case STPRT_C:
		stp = 1
	case STPRT_D:
		stp = 0.5
	case STPRT_E:
		stp = 0.25
	case STPRT_X:
		stp = 0.0
	}
	tf := (pl.PM() * stp) / ll
	tf = roundFl64(tf, 2)
	return tf
}

//TradeRouteValue - return yearly Trade Route Value
func TradeRouteValue(pl1, pl2 *planet) float64 {
	dist := astrogation.DistanceRaw(pl1.coordX, pl1.coordY, pl2.coordX, pl2.coordY)
	return (pl1.TradeFactor() + pl2.TradeFactor()) / float64(dist)
}

//////////////
//roundFl64 - округляет float64 до требуемого кол-ва разрядов
func roundFl64(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
