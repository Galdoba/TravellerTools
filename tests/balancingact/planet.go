package balancingact

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/internal/ehex"
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
)

type planet struct {
	name        string
	currrentUWP string
	starport    string
	atrib       map[int]int
	hex         hexagon.Hexagon
	coordX      int
	coordY      int
	coordQ      int
	coordR      int
	coordS      int
	iX          string
	eX          string
	cX          string
	pbg         string
	/*
		Solidarity - measure of how united the populance (T5 = -Heterogenety)
		Wealth = Amount of resources and amount of economic power (T5 = Resources+Labor)
		Expantion = A measure of desire to expand (T5 = Strangeness or Importance)
		Might = measure of how strong militarises and passionate populance is (T5 = -Acceptance)
	*/
}

type SurveyData interface {
	MW_Name() string
	MW_UWP() string
	MW_Importance() string
	MW_Economic() string
	MW_Cultural() string
	PBG() string
	// CoordX() int
	// CoordY() int
	// CoordQ() int
	// CoordR() int
	// CoordS() int
	hexagon.Hex
	hexagon.Cube
}

func createPlanet(name, uwp, iX, eX, cX, pbg string, x, y int) (*planet, error) {
	p := planet{}
	p.name = name
	p.currrentUWP = uwp
	p.iX = iX
	p.eX = eX
	p.cX = cX
	p.pbg = pbg
	p.hex, _ = hexagon.New(hexagon.Feed_HEX, x, y)
	//p.coordX, p.coordY = x, y
	creator := dice.New().SetSeed(fmt.Sprintf("n%vu%vx%vy%v", name, uwp, x, y))
	p.atrib = make(map[int]int)
	for _, key := range []int{ /*Solidarity, Wealth, Expansion, Might, Development, */ Solidarity, Wealth, Expansion, Might, Development} {
		atr := creator.Roll("2d6").DM(-2).Sum()
		switch key {
		case Solidarity:
			atr = p.heterogenity()
		case Wealth:
			atr = p.wealth()
		case Expansion:
			atr = 3 + p.expansion()
		case Might:
			atr = p.might()
		case Development:
			atr = p.development()
		}
		p.atrib[key] = atr

		if string(uwp[4]) == "0" {
			p.atrib[key] = 0
		}

	}
	if p.atrib[Development] > p.atrib[Wealth] {
		p.atrib[Development] = p.atrib[Wealth]
	}
	p.updateUWP(p.currrentUWP)
	return &p, nil
}

func (p *planet) heterogenity() int {
	het := ehex.New().Set(string(p.cX[1]))
	return 20 - het.Value()
}

func (p *planet) wealth() int {
	res := ehex.New().Set(string(p.eX[1]))
	lab := ehex.New().Set(string(p.eX[2]))
	return res.Value() + lab.Value()
}

func (p *planet) expansion() int {
	data := strings.Fields(p.iX)
	i, _ := strconv.Atoi(data[1])
	return i
}

func (p *planet) might() int {
	acc := ehex.New().Set(string(p.cX[2]))
	return acc.Value()
}

func (p *planet) development() int {
	inf := ehex.New().Set(string(p.eX[3]))
	lab := ehex.New().Set(string(p.eX[2]))
	return inf.Value() + lab.Value()
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
	//str += fmt.Sprintf("///\nSOL %v  WLT %v  EXP %v  MGT %v\n///\n", pl.atrib[T5_Solidarity], pl.atrib[T5_Wealth], pl.atrib[T5_Expansion], pl.atrib[T5_Might])
	str += fmt.Sprintf("Global Planetary Product: %v/%v\n", pl.GPP(), roundFl64(pl.GPP()/52.0, 1))
	str += fmt.Sprintf("Planetary Market : %v/%v\n", pl.PM(), roundFl64(pl.PM()/52, 1))
	str += fmt.Sprintf("Goverment Ship Building Budget : %v\n", pl.ShipBuildingBudget())
	str += fmt.Sprintf("Trade Factor : %v\n", pl.TradeFactor())
	return str
}

func ImportPlanet(sd SurveyData) (*planet, error) {
	pl, err := createPlanet(sd.MW_Name(), sd.MW_UWP(), sd.MW_Importance(), sd.MW_Economic(), sd.MW_Cultural(), sd.PBG(), sd.CoordX(), sd.CoordY())
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
	dist := hexagon.Distance(pl1.hex, pl2.hex)
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
