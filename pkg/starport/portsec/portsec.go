package portsec

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

const (
	org_No               = "No security force"
	org_SmallPT          = "Small, part-time security force"
	org_SmallProf        = "Small professional security force"
	org_Modest           = "Modest-sized professional security force"
	org_Large            = "Large professional security force"
	org_Huge             = "Huge professional security force"
	org_Enormous         = "Enormous security force"
	fund_No              = "No significant funding"
	fund_VeryUnderfunded = "Grossly underfunded"
	fund_Underfunded     = "Underfunded "
	fund_Normal          = "Normal level of funding"
	fund_Well            = "Well-funded"
	fund_Lavishly        = "Lavishly funded"
	equip_None           = "None provided"
	equip_Minimal        = "Minimal"
	equip_Basic          = "Basic"
	equip_Standard       = "Standard"
	equip_Well           = "Well Equipped"
	equip_Heavily        = "Heavily Equipped"
	equip_Lavishly       = "Lavishly Equipped"
)

type StarportSecurityForces struct {
	portName        string
	worldPopulation float64
	basePersonal    float64
	organisation    string
	funding         string
	equipment       string
	competence      string
	response        string
}

type Port interface {
	MW_Name() string
	MW_UWP() string
	PBG() string
}

func GenerateSecurityForces(port Port) (*StarportSecurityForces, error) {
	ssf := StarportSecurityForces{}
	seed := port.MW_Name() + port.MW_UWP() + port.PBG() + port.MW_UWP()
	dp := dice.New().SetSeed(seed)
	ssf.portName = port.MW_Name()
	population := 1.0

	err := fmt.Errorf("err value was not adressed")
	uwpData, err := uwp.FromString(port.MW_UWP())
	if err != nil {
		return &ssf, err
	}
	pops := uwpData.Pops()
	law := uwpData.Laws()
	st := uwpData.Starport()
	for i := 0; i < pops; i++ {
		population = population * 10
	}
	pbg := port.PBG()
	p := string(pbg[0])
	pi, err := strconv.Atoi(p)
	ssf.worldPopulation = population * float64(pi)

	org, mult := defineOrganisation(dp, pops, law)
	if st == "X" {
		org = org_No
	}
	ssf.organisation = org
	ssf.basePersonal = (ssf.worldPopulation / 500) * mult
	fund, equipDM, _ := defineFunding(dp, pops, st)
	equip := defineEquipment(dp, equipDM)
	if org == org_No {
		fund = fund_No
	}
	if fund == fund_No {
		equip = equip_None
	}
	ssf.funding = fund
	ssf.equipment = equip
	fmt.Println("Population :=", ssf.worldPopulation, ssf.basePersonal)

	return &ssf, err
}

func (ssf *StarportSecurityForces) String() string {
	str := fmt.Sprintf("Security Forces at %v:\n", ssf.portName)
	underline := ""
	for len(underline) < 80 {
		underline += "-"
	}
	str += underline + "\n"

	str += fmt.Sprintf("Organisation : %v\n", ssf.organisation)
	str += fmt.Sprintf("Funding      : %v\n", ssf.funding)
	str += fmt.Sprintf("Equipment    : %v\n", ssf.equipment)
	str += fmt.Sprintf("Competence   : %v\n", ssf.competence)
	str += fmt.Sprintf("Response     : %v\n", ssf.response)

	str += underline
	return str
}

func defineOrganisation(dp *dice.Dicepool, pops, laws int) (string, float64) {
	dm := 0
	switch pops {
	default:
		dm += 3
	case 0, 1, 2, 3:
		dm += -2
	case 6, 7:
		dm += 1
	case 4, 5:
		//dm += 0
	}
	switch laws {
	default:
		dm += 3
	case 0:
		dm += -4
	case 1, 2, 3:
		dm += -1
	case 4, 5:
		//dm += 0
	case 6, 7, 8:
		dm += 1
	}
	r := dp.Roll("2d6").DM(dm).Sum()
	switch r {
	default:
		if r < 1 {
			return org_No, 0.0
		}
		return org_Enormous, 5.0
	case 1, 2, 3:
		return org_SmallPT, 0.2
	case 4, 5, 6:
		return org_SmallProf, 0.4
	case 7, 8, 9:
		return org_Modest, 1
	case 10, 11, 12:
		return org_Large, 1.5
	case 13, 14, 15:
		return org_Huge, 2.5

	}
}

func defineFunding(dp *dice.Dicepool, pops int, st string) (string, int, []int) {
	dm := 0
	switch pops {
	case 0, 1, 2:
		dm += -2
	case 3, 4:
		dm += -1
	case 5, 6:
		dm += 0
	case 7, 8:
		dm += 1
	default:
		dm += 2
	}
	switch st {
	case "E":
		dm += -2
	case "D":
		dm += -1
	case "B":
		dm += 1
	case "A":
		dm += 2
	}
	r := dp.Roll("2d6").DM(dm).Sum()
	if r < 2 {
		return fund_No, -1000, []int{100, 0, 0}
	}
	switch r {
	case 2, 3:
		return fund_VeryUnderfunded, -4, []int{100, 0, 0}
	case 4, 5, 6:
		return fund_Underfunded, -2, []int{100, 0, 0}
	case 7, 8, 9:
		return fund_Normal, 0, []int{90, 9, 1}
	case 10, 11, 12:
		return fund_Well, 2, []int{80, 15, 5}
	default:
		return fund_Lavishly, 4, []int{70, 20, 10}
	}
	return "", 0, []int{}
}

func defineEquipment(dp *dice.Dicepool, dm int) string {
	r := dp.Roll("2d6").DM(dm).Sum()
	if r < 1 {
		return equip_None
	}
	switch r {
	case 1, 2, 3:
		return equip_Minimal
	case 4, 5, 6:
		return equip_Basic
	case 7, 8, 9:
		return equip_Standard
	case 10, 11, 12:
		return equip_Well
	case 13, 14, 15:
		return equip_Heavily
	default:
		return equip_Lavishly

	}
}
