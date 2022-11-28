package economics

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/calendar"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"github.com/Galdoba/utils"
)

const (
	RESOURCE_CHR        = "Resource"
	LABOR_CHR           = "Labor"
	INFRASTRUCTRURE_CHR = "Infrastructure"
	CULTURE_CHR         = "Culture"
)

type EconomicPower interface {
	Resources() int
	String() string
}

type economicPower struct {
	baseRolls1     []int
	baseRolls2     []int
	pbgStash       []int
	eventMods      map[string][]eventMod
	resource       ehex.Ehex
	labor          ehex.Ehex
	infrastructure ehex.Ehex
	culture        ehex.Ehex
	///world Data
	uwp        uwp.UWP
	tradeCodes []string
}

type eventMod struct {
	startDate         calendar.Date
	endDate           calendar.Date
	degradationPeriod int
	description       string
	mod               int
}

type World interface {
	UWP() uwp.UWP
	TradeCodes() []string
	PBG() string
}

func GenerateInitialEconomicPower(wrld World, dice *dice.Dicepool) *economicPower {
	ep := economicPower{}
	ep.setupBase(wrld, dice)
	return &ep
}

func (ep *economicPower) Resources() int {
	return ep.getBase(RESOURCE_CHR) // add event modifiers
}

func (ep *economicPower) Labor() int {
	return ep.getBase(LABOR_CHR) // add event modifiers
}

func (ep *economicPower) Infrastructure() int {
	return ep.getBase(INFRASTRUCTRURE_CHR) // add event modifiers
}

func (ep *economicPower) Culture() int {
	return ep.getBase(CULTURE_CHR) // add event modifiers
}

func (ep *economicPower) Update(wrld World) {
	ep.uwp = wrld.UWP()
	ep.tradeCodes = wrld.TradeCodes()
}

func (ep *economicPower) setupBase(wrld World, dice *dice.Dicepool) {
	ep.baseRolls1 = nil
	ep.baseRolls2 = nil
	ep.pbgStash = nil
	for i := 0; i < 4; i++ {
		ep.baseRolls1 = append(ep.baseRolls1, dice.Sroll("1d6"))
		ep.baseRolls2 = append(ep.baseRolls2, dice.Sroll("1d6"))
	}
	ep.eventMods = make(map[string][]eventMod)
	ep.pbgStash = extractPBG(wrld.PBG())
	ep.Update(wrld)
	for i, vType := range []string{RESOURCE_CHR, LABOR_CHR, INFRASTRUCTRURE_CHR, CULTURE_CHR} {
		base := ep.getBase(vType)
		if base < 0 {
			base = 0
		}
		if base > 15 {
			base = 15
		}
		switch i {
		case 0:
			ep.resource = ehex.New().Set(base)
		case 1:
			ep.labor = ehex.New().Set(base)
		case 2:
			ep.infrastructure = ehex.New().Set(base)
		case 3:
			ep.culture = ehex.New().Set(base)
		}
	}
}

func (ep *economicPower) String() string {
	return fmt.Sprintf("%v%v%v%v", ep.resource.Code(), ep.labor.Code(), ep.infrastructure.Code(), ep.culture.Code())
}

func (ep *economicPower) getBase(vType string) int {
	base := -1
	switch vType {
	default:
		return -1
	case RESOURCE_CHR:
		base = ep.resourceBase()
	case LABOR_CHR:
		base = ep.laborBase()
	case INFRASTRUCTRURE_CHR:
		base = ep.infrastructureBase()
	case CULTURE_CHR:
		base = ep.cultureBase()
	}
	return base
}

func (ep *economicPower) resourceBase() int {
	r := 0
	switch listContains(ep.tradeCodes, "As", "Ba", "Po") {
	case false:
		r = ep.baseRolls1[0] + ep.baseRolls2[0] - 2
		if ep.uwp.TL() >= 8 {
			r += ep.pbgStash[1] + ep.pbgStash[2]
		}
	case true:
		r = ep.baseRolls1[0] - 1
	}
	r += resourceModsTC(ep.tradeCodes, ep.uwp)
	return r
}

func resourceModsTC(actualTC []string, u uwp.UWP) int {
	r := 0
	for _, code := range actualTC {
		switch code {
		case "In":
			r = r + 2
		case "Ag", "Hi", "Ri":
			r = r + 1
		case "De", "Fl", "Na":
			r = r - 1
		case "Va":
			if listContains(actualTC, "As") {
				continue
			}
			r = r - 1
		}
	}
	switch u.Starport() {
	case "A":
		r = r + 2
	case "B":
		r = r + 1
	}
	return r
}

func (ep *economicPower) laborBase() int {
	return utils.Max(0, ep.uwp.Pops()-1)
}

func (ep *economicPower) infrastructureBase() int {
	r := 0
	switch listContains(ep.tradeCodes, "Ba") {
	case false:
		r = ep.baseRolls1[2] + ep.baseRolls2[2] - 2
	case true:
		return 0
	}
	if ep.uwp.Pops() == 0 {
		r = ep.baseRolls1[2] / 3
	}
	r += infrastructureModsTC(ep.tradeCodes, ep.uwp)
	return r
}

func infrastructureModsTC(actualTC []string, u uwp.UWP) int {
	in := 0
	for _, code := range actualTC {
		switch code {
		case "Po":
			in = in - 2
		case "As", "Lo", "Wa":
			in = in - 1
		case "Hi":
			in = in + 1
		case "In", "Ri":
			in = in + 2
		}
	}
	switch u.Starport() {
	case "A":
		in = in + 4
	case "B":
		in = in + 3
	case "C":
		in = in + 2
	case "D":
		in = in + 1
	}
	return in
}

func (ep *economicPower) cultureBase() int {
	r := 0
	switch listContains(ep.tradeCodes, "Ba") {
	case true:
		return 0
	case false:
		r = ep.baseRolls1[3] + ep.baseRolls2[3]
	}
	if ep.uwp.Pops() == 0 {
		r = ep.baseRolls1[3] / 5
	}
	r += cultureModsTC(ep.tradeCodes, ep.uwp)
	return r
}

func cultureModsTC(actualTC []string, u uwp.UWP) int {
	cu := 0
	for _, code := range actualTC {
		switch code {
		case "As", "De", "Fl", "Ic", "Po", "Ri":
			cu = cu + 1
		case "Ag", "Na", "Ni":
			cu = cu - 1
		case "Va":
			if listContains(actualTC, "As") {
				continue
			}
			cu = cu + 1
		}
	}
	return cu
}

func listContains(sl []string, el ...string) bool {
	for _, elem := range el {
		if utils.ListContains(sl, elem) {
			return true
		}
	}
	return false
}

func extractPBG(pbg string) []int {
	dt := strings.Split(pbg, "")
	pbgInt := []int{}
	for i, d := range dt {
		switch i {
		case 0, 1, 2:
			val, err := strconv.Atoi(d)
			if err != nil {
				val = -1000
			}
			pbgInt = append(pbgInt, val)
		}
	}
	return pbgInt
}
