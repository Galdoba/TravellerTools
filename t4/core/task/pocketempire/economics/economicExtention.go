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
	Labor() int
	Infrastructure() int
	Culture() int
	String() string
	RecalculateRA(*dice.Dicepool) error
	StatBlock() string
}

type economicValue struct {
	val float64
}

func (ev *economicValue) Code() string {
	return ehex.New().Set(int(ev.val)).Code()
}

func (ev *economicValue) Value() int {
	return int(ev.val)
}

func (ev *economicValue) ValueFl64() float64 {
	return ev.val
}

func setEconomicValue(i int) *economicValue {
	return &economicValue{float64(i)}
}

type economicPower struct {
	baseRolls1     []int
	baseRolls2     []int
	pbgStash       []int
	eventMods      map[string][]eventMod
	resource       *economicValue
	labor          *economicValue
	infrastructure *economicValue
	culture        *economicValue
	///world Data
	uwp               uwp.UWP
	tradeCodes        []string
	totalDemand       float64
	resourceAvailable float64
	excess            float64
	deficit           float64
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

func minFl64(fl1, fl2 float64) float64 {
	if fl1 <= fl2 {
		return fl1
	}
	return fl2
}

func (ep *economicPower) RecalculateRA(dice *dice.Dicepool) error {
	baseDemand := 0.0
	ep.excess = 0.0
	ep.deficit = 0.0
	ep.resourceAvailable = 0.0
	r := dice.Sroll("2d6")
	switch {
	case ep.infrastructure.Value()-ep.resource.Value() <= 0:
		switch {
		case ep.uwp.Pops() >= 4:
			baseDemand = ep.resource.ValueFl64()
		case ep.uwp.Pops() < 4:
			baseDemand = float64(ep.uwp.Pops())
		}
		ep.totalDemand = totalDemandTable(r, int(baseDemand), ep.uwp.Pops(), ep.Culture())
		ep.resourceAvailable = minFl64(ep.totalDemand, ep.resource.ValueFl64())
		ep.excess = ep.resource.ValueFl64() - ep.totalDemand
	case ep.infrastructure.Value()-ep.resource.Value() > 0:
		switch {
		case ep.uwp.Pops() >= 4:
			baseDemand = ep.infrastructure.ValueFl64()
		case ep.uwp.Pops() < 4:
			baseDemand = float64(ep.uwp.Pops())
		}
		ep.totalDemand = totalDemandTable(r, int(baseDemand), ep.uwp.Pops(), ep.Culture())
		switch {
		//case 1
		case ep.totalDemand <= ep.resource.ValueFl64():
			ep.resourceAvailable = ep.totalDemand
			ep.excess = ep.resource.ValueFl64() - ep.totalDemand
		//case 2
		case ep.totalDemand > ep.resource.ValueFl64():
			ep.resourceAvailable = ep.resource.ValueFl64()
			ep.deficit = ep.resource.ValueFl64() - ep.totalDemand
			ep.resourceAvailable = ep.resourceAvailable - ep.deficit
			//case 2.5
			if ep.resourceAvailable < 0 {
				ep.infrastructure.val = ep.infrastructure.val - (ep.deficit / 10)
				ep.resourceAvailable = 0
			}
		}
	}
	ep.excess = utils.RoundFloat64(ep.excess, 1)
	ep.deficit = utils.RoundFloat64(ep.deficit, 1)
	ep.resourceAvailable = utils.RoundFloat64(ep.resourceAvailable, 1)
	ep.totalDemand = utils.RoundFloat64(ep.totalDemand, 1)
	ep.ResourceTrade(dice)
	return nil
}

func (ep *economicPower) ResourceTrade(dice *dice.Dicepool) {
	if ep.excess > 0 {
		resExport := ep.resource.val - ep.totalDemand
		ep.resourceAvailable = ep.resourceAvailable + (resExport * exportBenefit(dice.Sroll("2d6")))
	}
	if ep.deficit > 0 {
		resImport := ep.totalDemand - ep.resource.val
		ep.resourceAvailable = ep.resourceAvailable + (resImport * exportBenefit(dice.Sroll("2d6")))
	}
	ep.resourceAvailable = utils.RoundFloat64(ep.resourceAvailable, 1)
}

func exportBenefit(i int) float64 {
	return []float64{0.3, 0.3, 0.4, 0.4, 0.5, 0.5, 0.5, 0.5, 0.6, 0.6, 0.7}[i-2]
}

func importBenefit(i int) float64 {
	return []float64{0.2, 0.2, 0.3, 0.3, 0.4, 0.4, 0.4, 0.4, 0.5, 0.5, 0.6}[i-2]
}

func totalDemandTable(r, baseDemandVal, pop, cult int) float64 {
	dm := 0
	baseDemandVal = utils.BoundInt(baseDemandVal, 0, 15)
	switch pop {
	case 0, 1:
		dm = dm - 3
	case 2, 3:
		dm = dm - 2
	case 4, 5:
		dm = dm - 1
	case 6:
		dm = dm - 0
	case 7, 8:
		dm = dm + 1
	case 9, 10:
		dm = dm + 2
	default:
		dm = dm + 3
	}
	switch cult {
	case 0, 1:
		dm = dm - 3
	case 2, 3:
		dm = dm - 2
	case 4, 5:
		dm = dm - 1
	case 6, 7:
		dm = dm + 0
	case 8, 9, 10:
		dm = dm + 1
	case 11, 12, 13:
		dm = dm + 2
	case 14, 15:
		dm = dm + 3
	}
	roll := r + dm
	roll = utils.BoundInt(roll, 0, 15)
	baseDemand := make(map[int][]float64)
	baseDemand[0] = []float64{0, 0, 0, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7}
	baseDemand[1] = []float64{0, 0, 0, 1, 2, 2, 3, 4, 5, 5, 6, 6, 8, 8, 9, 9}
	baseDemand[2] = []float64{0, 0, 1, 1, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10, 11, 11}
	baseDemand[3] = []float64{0, 0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	baseDemand[4] = []float64{0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	baseDemand[5] = []float64{0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	baseDemand[6] = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	baseDemand[7] = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	baseDemand[8] = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	baseDemand[9] = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	baseDemand[10] = []float64{1, 2, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	baseDemand[11] = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 15, 16, 17}
	baseDemand[12] = []float64{1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16, 17, 18}
	baseDemand[13] = []float64{1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 16, 17, 18, 19}
	baseDemand[14] = []float64{1, 2, 3, 5, 7, 8, 9, 10, 11, 12, 13, 14, 17, 18, 19, 20}
	baseDemand[15] = []float64{1, 2, 4, 5, 7, 8, 9, 11, 12, 13, 15, 16, 18, 19, 21, 22}
	return baseDemand[baseDemandVal][roll]
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
			ep.resource = setEconomicValue(base)
		case 1:
			ep.labor = setEconomicValue(base)
		case 2:
			ep.infrastructure = setEconomicValue(base)
		case 3:
			ep.culture = setEconomicValue(base)
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

func (ep *economicPower) StatBlock() string {
	s := "-----------\n"
	s += fmt.Sprintf("EconPower=%v\n", ep.String())
	s += fmt.Sprintf("totalDemand=%v\n", ep.totalDemand)
	s += fmt.Sprintf("excess=%v\n", ep.excess)
	s += fmt.Sprintf("deficit=%v\n", ep.deficit)
	s += fmt.Sprintf("resourceAvailable=%v\n", ep.resourceAvailable)
	return s
}
