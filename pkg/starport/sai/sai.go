package sai

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

const (
	tonnage_Bulk = iota
	tonnage_Large
	tonnage_Medium
	tonnage_Small
	tonnage_Minor
)

type ShippingActivity struct {
	portName   string   //имя планеты
	portUWP    string   //UWP планеты для расчета активноси
	tradecodes []string //торговые коды
	traffic    []int    // значение прилетающих/улетающих/транзитных показателей груза.
	///////////

	averageShips   int
	minmumShips    int
	maximumShips   int
	shipsByTonnage map[int]int
	traffMult      int
	traffDm        int
}

type Port interface {
	MW_Name() string
	MW_UWP() string
	MW_Remarks() string
	TravelZone() string
}

type port struct {
	name string
	uwp  string
	rem  string
	tz   string
}

func (p *port) MW_Name() string {
	return p.name
}
func (p *port) MW_UWP() string {
	return p.uwp
}
func (p *port) MW_Remarks() string {
	return p.rem
}
func (p *port) TravelZone() string {
	return p.tz
}

func NewShippingActivity(port Port, traffic []int) (*ShippingActivity, error) {
	//arrive+depart+transit+transit
	err := fmt.Errorf("error value was not adressed")
	sa := ShippingActivity{}
	sa.portName = port.MW_Name()
	sa.portUWP = port.MW_UWP()
	sa.averageShips = -1000
	sa.traffMult = -1000
	sa.traffDm = -1000
	sa.minmumShips = -1000
	sa.maximumShips = -1000
	sa.shipsByTonnage = make(map[int]int)
	baseSAI := traffic[0] + traffic[1] + (2 * traffic[2])
	portUWP, err := uwp.FromString(port.MW_UWP())
	if err != nil {
		return &sa, err
	}
	st := portUWP.Starport()

	switch st {
	case "A":
		sa.traffMult = 3
		sa.traffDm = 0
		sa.shipsByTonnage[tonnage_Bulk] = 5
		sa.shipsByTonnage[tonnage_Large] = 10
		sa.shipsByTonnage[tonnage_Medium] = 20
		sa.shipsByTonnage[tonnage_Small] = 30
		sa.shipsByTonnage[tonnage_Minor] = 35
	case "B":
		sa.traffMult = 2
		sa.traffDm = 0
		sa.shipsByTonnage[tonnage_Bulk] = 0
		sa.shipsByTonnage[tonnage_Large] = 5
		sa.shipsByTonnage[tonnage_Medium] = 10
		sa.shipsByTonnage[tonnage_Small] = 20
		sa.shipsByTonnage[tonnage_Minor] = 65
	case "C":
		sa.traffMult = 1
		sa.traffDm = 0
		sa.shipsByTonnage[tonnage_Bulk] = 0
		sa.shipsByTonnage[tonnage_Large] = 0
		sa.shipsByTonnage[tonnage_Medium] = 5
		sa.shipsByTonnage[tonnage_Small] = 10
		sa.shipsByTonnage[tonnage_Minor] = 85
	case "D":
		sa.traffMult = 1
		sa.traffDm = -1
		baseSAI += -2
		sa.shipsByTonnage[tonnage_Bulk] = 0
		sa.shipsByTonnage[tonnage_Large] = 0
		sa.shipsByTonnage[tonnage_Medium] = 5
		sa.shipsByTonnage[tonnage_Small] = 10
		sa.shipsByTonnage[tonnage_Minor] = 85
	case "E":
		sa.traffMult = 1
		sa.traffDm = -2
		baseSAI += -4
		sa.shipsByTonnage[tonnage_Bulk] = 0
		sa.shipsByTonnage[tonnage_Large] = 0
		sa.shipsByTonnage[tonnage_Medium] = 0
		sa.shipsByTonnage[tonnage_Small] = 5
		sa.shipsByTonnage[tonnage_Minor] = 95
	case "X":
		sa.traffMult = 1
		sa.traffDm = -3
		baseSAI += -8
		sa.shipsByTonnage[tonnage_Bulk] = 0
		sa.shipsByTonnage[tonnage_Large] = 0
		sa.shipsByTonnage[tonnage_Medium] = 0
		sa.shipsByTonnage[tonnage_Small] = 0
		sa.shipsByTonnage[tonnage_Minor] = 100
	}
	for _, tc := range sa.tradecodes {
		switch tc {
		case "Ag", "In", "Ri", "Hi":
			baseSAI += 2
		case "Na", "Ni", "Po", "Lo":
			baseSAI += -2
		case "As":
			baseSAI += 1
		case "Ga", "Cx", "Cs":
			baseSAI += 3
		}
	}
	if port.TravelZone() == "R" {
		baseSAI += -5
	}
	if port.TravelZone() == "A" {
		baseSAI += -3
	}
	dices := baseSAI * sa.traffMult
	dp := dice.New().SetSeed(port.MW_Name() + port.MW_Remarks() + port.MW_UWP() + port.TravelZone())
	sa.minmumShips = dp.Roll(strconv.Itoa(dices) + "d1").DM(dices * sa.traffDm).Sum()
	sa.maximumShips = dp.Roll(strconv.Itoa(dices) + "d1").DM((dices * sa.traffDm) + (5 * dices)).Sum()
	if sa.minmumShips > sa.maximumShips {
		sa.minmumShips, sa.maximumShips = sa.maximumShips, sa.minmumShips
	}
	sa.averageShips = (sa.minmumShips + sa.maximumShips) / 2
	if sa.minmumShips < 0 {
		sa.minmumShips = 0
	}
	if sa.averageShips < 0 {
		sa.averageShips = 0
	}
	if sa.maximumShips < 0 {
		sa.maximumShips = 0
	}

	return &sa, err
}

func (sa *ShippingActivity) String() string {
	str := fmt.Sprintf("Shipping Activity on %v:\n", sa.portName)
	str += fmt.Sprintf("At any given moment port expected have %v-%v ships (%v average). Of them:\n", sa.minmumShips, sa.maximumShips, sa.averageShips)
	for i, ton := range []int{tonnage_Bulk, tonnage_Large, tonnage_Medium, tonnage_Small, tonnage_Minor} {
		ships := sa.averageShips * sa.shipsByTonnage[ton] / 100
		if ships > 0 {
			sType := ""
			switch i {
			case 0:
				sType = "Bulk ships   (50,000+ tons)      : "
			case 1:
				sType = "Large ships  (10,000-49,999 tons): "
			case 2:
				sType = "Medium ships (5,000-9,999 tons)  : "
			case 3:
				sType = "Small ships  (1,000-4,999 tons)  : "
			case 4:
				sType = "Minor ships  (100-999 tons)      : "
			}
			str += fmt.Sprintf("%v%v\n", sType, ships)
		}

	}
	wt := 0
	wDM := -4
	wInc := 0
	d := strings.Split(sa.portUWP, "")
	st := d[0]
	switch st {
	case "A":
		wInc = 35
	case "B":
		wInc = 20
	case "C":
		wInc = 15
	case "D":
		wInc = 7
	case "E":
		wInc = 4
	case "X":
		wInc = 0
	}
	wt = (7 + wDM) * wInc
	str += fmt.Sprintf("Average waiting time for Minor ships is %v minutes\n", wt)
	str += "--------------------------------------------------------------------------------"
	return str
}
