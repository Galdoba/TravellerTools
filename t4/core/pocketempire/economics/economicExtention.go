package economics

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/calendar"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"github.com/Galdoba/TravellerTools/t4/core/task"
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
	GWP() int
	String() string
	DescridetoryTax() float64
	Process(*dice.Dicepool, InterstellarDemand) error
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

type Pawn interface {
	//берет статы от персонажа
}

type MilitaryUnits interface {
	//берет статы от армии
	Maintainance() float64
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
	baseGWP           float64
	finalGWP          *economicValue
	actionLog         []string
	tradePartners     int
	tradeMultipler    float64
	descridetoryTax   float64
	govermentalBudget float64
	civilExpenses     float64
	militaryExpenses  float64
	//other
	administrator Pawn
	units         []MilitaryUnits
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

func (ep *economicPower) Process(dice *dice.Dicepool, demand InterstellarDemand) error {
	for _, err := range []error{
		ep.determinePlanetaryDemand(dice),
		ep.engageInResourceTrade(dice),
		ep.computeBaseGrossWorldProduct(),
		ep.calculateFinalGrossWorldProduct(demand),
		ep.calculateGovermentalBudget(),
		ep.calculateExpenses(dice),
	} {
		if err != nil {
			return err
		}
	}

	return nil
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

func (ep *economicPower) GWP() int {
	return ep.finalGWP.Value()
}

func (ep *economicPower) calculateExpenses(dice *dice.Dicepool) error {
	c := ep.culture.val
	i := ep.infrastructure.val
	l := float64(ep.uwp.Laws())
	a := ep.administratorMult(dice)
	g := ep.govermentalExpenceFactor()
	gb := ep.govermentalBudget
	ce := ((c + i + l) / 100.0) * a * g * gb
	ep.civilExpenses = utils.RoundFloat64(ce, 1)
	ep.actionLog = append(ep.actionLog, fmt.Sprintf("civil expences: %v RU", ep.civilExpenses))
	m := ep.militaryCost()
	me := m * a * g
	ep.militaryExpenses = utils.RoundFloat64(me, 1)
	ep.actionLog = append(ep.actionLog, fmt.Sprintf("military expences: %v RU", ep.militaryExpenses))
	return nil
}

func (ep *economicPower) militaryCost() float64 {
	summ := 0.0
	for _, unit := range ep.units {
		summ += unit.Maintainance()
	}
	return summ
}

func (ep *economicPower) administratorMult(dice *dice.Dicepool) float64 {
	a := 1.1
	tn := 0
	if ep.administrator == nil {
		randomTaskNumber(dice)
	}
	tsk := task.New(tn, task.TaskFormidable)
	tsk.SetResolver(dice)
	switch tsk.Resolve() {
	case task.SpectacularSuccess:
		a = 0.9
	case task.Success:
		a = 1.0
	case task.Failure:
		a = 1.1
	case task.SpectacularFailure:
		a = 1.2
	default:
		a = 1.1
	}
	return a
}

func (ep *economicPower) govermentalExpenceFactor() float64 {
	switch ep.uwp.Govr() {
	case 0:
		return 0.95
	case 1, 14:
		return 1.10
	case 2:
		return 1.40
	case 3, 8:
		return 1.30
	case 4:
		return 1.15
	case 5, 9:
		return 1.35
	case 10, 13:
		return 1.05
	case 11:
		return 1.00
	case 12:
		return 1.25
	case 15:
		return 1.20
	}
	return 1.18
}

func randomTaskNumber(dice *dice.Dicepool) int {
	i := dice.Sroll("1d6-1")
	return []int{18, 16, 14, 12, 10, 8}[i]
}

func (ep *economicPower) baseTax() float64 {
	switch ep.uwp.Govr() {
	case 0:
		return 0.05
	case 1, 3:
		return 0.20
	case 4, 5, 11:
		return 0.25
	case 2, 10, 12:
		return 0.30
	case 9, 14, 15:
		return 0.35
	case 8, 13:
		return 0.40
	}
	//0.26 - это среднее ориентироваться по хозяину
	return 0.26
}

func (ep *economicPower) socialTax() float64 {
	return float64(ep.uwp.Laws()+ep.Culture()) / 100
}

func (ep *economicPower) TotalTax() float64 {
	return ep.baseTax() + ep.socialTax() + ep.descridetoryTax
}

func (ep *economicPower) DescridetoryTax() float64 {
	return ep.descridetoryTax
}

func (ep *economicPower) determinePlanetaryDemand(dice *dice.Dicepool) error {
	baseDemand := 0.0
	ep.excess = 0.0
	ep.deficit = 0.0
	ep.resourceAvailable = 0.0
	r := dice.Sroll("2d6")
	switch {
	case ep.infrastructure.Value()-ep.resource.Value() <= 0:
		baseDemand = ep.resource.ValueFl64()
		if ep.uwp.Pops() < 4 {
			baseDemand = float64(ep.uwp.Pops())
		}
		ep.totalDemand = totalDemandTable(r, int(baseDemand), ep.uwp.Pops(), ep.Culture())
		switch {
		case ep.totalDemand <= ep.resource.val:
			ep.resourceAvailable = ep.totalDemand
			ep.excess = ep.resource.ValueFl64() - ep.totalDemand
		case ep.totalDemand > ep.resource.val:
			ep.resourceAvailable = ep.resource.val
			ep.deficit = ep.totalDemand - ep.resource.val
		}
	case ep.infrastructure.Value()-ep.resource.Value() > 0:
		baseDemand = ep.infrastructure.ValueFl64()
		if ep.uwp.Pops() < 4 {
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
			ep.deficit = ep.totalDemand - ep.resource.ValueFl64()
			ep.resourceAvailable = ep.resourceAvailable - ep.deficit
			//case 2.5
			if ep.resourceAvailable < 0 {
				ep.infrastructure.val = ep.infrastructure.val - (ep.deficit / 10)
				ep.resourceAvailable = 0
			}
		}
	}

	return nil
}

func (ep *economicPower) roundValues() {
	ep.excess = utils.RoundFloat64(ep.excess, 1)
	ep.deficit = utils.RoundFloat64(ep.deficit, 1)
	ep.resourceAvailable = utils.RoundFloat64(ep.resourceAvailable, 1)
	ep.totalDemand = utils.RoundFloat64(ep.totalDemand, 1)
	ep.baseGWP = utils.RoundFloat64(ep.baseGWP, 1)
	ep.finalGWP.val = utils.RoundFloat64(ep.finalGWP.val, 1)
}

func (ep *economicPower) engageInResourceTrade(dice *dice.Dicepool) error {
	if ep.excess != 0 && ep.deficit != 0 {
		return fmt.Errorf("invalid ep.excess/ep.deficit data (%v/%v)", ep.excess, ep.deficit)
	}
	if ep.excess > 0 {
		resExport := ep.resource.val - ep.totalDemand
		benefit := exportBenefit(dice.Sroll("2d6"))
		ep.resourceAvailable = ep.resourceAvailable + (resExport * benefit)
		ep.actionLog = append(ep.actionLog, fmt.Sprintf("exported %vx%v=%v resource points", resExport, benefit, resExport*benefit))
	}
	if ep.deficit > 0 {
		resImport := ep.totalDemand - ep.resource.val
		benefit := importBenefit(dice.Sroll("2d6"))
		ep.resourceAvailable = ep.resourceAvailable + (resImport * benefit)
		ep.actionLog = append(ep.actionLog, fmt.Sprintf("imported %vx%v=%v resource points", resImport, benefit, resImport*benefit))
	}
	ep.roundValues()
	return nil
}

func (ep *economicPower) computeBaseGrossWorldProduct() error {
	re := utils.RoundFloat64(float64(ep.uwp.TL())*0.1*ep.resourceAvailable, 1) //resouces exploitable rounded
	lf := laborFactor(ep.Labor()) * float64(ep.pbgStash[0])
	i := ep.infrastructure.val
	c := ep.culture.val
	ep.baseGWP = (re * lf * i) / (c + 1.0)
	if ep.baseGWP < 0 {
		return fmt.Errorf("negative base GWP")
	}
	ep.roundValues()
	credits := utils.RoundFloat64(1000000/(re*i), 1)
	ep.actionLog = append(ep.actionLog, fmt.Sprintf("Base GWP estimated: %v RU", ep.baseGWP))
	ep.actionLog = append(ep.actionLog, fmt.Sprintf("                  : %v MCr", credits))
	return nil
}

func (ep *economicPower) calculateFinalGrossWorldProduct(id InterstellarDemand) error {
	fgt := 1.0
	port := ep.uwp.Starport()
	switch port {
	case "A":
		fgt = 2.0 - (2.0 / math.Sqrt(float64(ep.tradePartners)+3.0))
	case "B":
		fgt = 1.7 - (1.7 / math.Sqrt(float64(ep.tradePartners)+4.89796))
	case "C":
		fgt = 1.4 - (1.4 / math.Sqrt(float64(ep.tradePartners)+11.25))
	case "D":
		fgt = 1.1 - (1.1 / math.Sqrt(float64(ep.tradePartners)+120.0))
	case "E":
		fgt = 1.01 - (1.01 / math.Sqrt(float64(ep.tradePartners)+10200.0))
	default:
		fgt = 1.000
	}
	fgt = fgt * id.Demand()
	fgt = utils.RoundFloat64(fgt, 3)
	ep.finalGWP.val = ep.baseGWP * fgt
	ep.roundValues()
	//ep.actionLog = append(ep.actionLog, fmt.Sprintf("FGTM for Starport %v with %v worlds is %v", port, ep.tradePartners, fgt))
	ep.actionLog = append(ep.actionLog, fmt.Sprintf("Final GWP         : %v RU", ep.finalGWP.val))
	return nil
}

func (ep *economicPower) determineinterstellarDemandMultipler(id InterstellarDemand) error {
	return nil
}

func (ep *economicPower) calculateGovermentalBudget() error {
	gb := utils.RoundFloat64(ep.finalGWP.val*ep.TotalTax(), 1)
	ep.govermentalBudget = gb
	ep.actionLog = append(ep.actionLog, fmt.Sprintf("Govermental Budget: %v RU", gb))
	return nil
}

func finishedGoodsTradeTable(ep *economicPower) float64 {
	fgt := 0.0
	tmaMap := make(map[int][]float64)
	tmaMap[1] = []float64{1.000, 1.000, 1.000, 1.000}
	tmaMap[2] = []float64{1.106, 1.053, 1.015, 1.000}
	tmaMap[3] = []float64{1.184, 1.095, 1.029, 1.001}
	tmaMap[4] = []float64{1.244, 1.130, 1.041, 1.001}
	tmaMap[5] = []float64{1.293, 1.160, 1.053, 1.002}
	tp := ep.tradePartners
	if ep.tradePartners > 5 {
		tp = 5
	}
	i := 3
	switch ep.uwp.Starport() {
	case "A":
		i = 0
	case "B":
		i = 1
	case "C":
		i = 2
	case "D":
		i = 3
	}
	fgt = tmaMap[tp][i]
	return fgt
}

func laborFactor(pop int) float64 {
	lf := -999.9
	switch pop {
	case 0:
		lf = 0.0000001
	case 1:
		lf = 0.000001
	case 2:
		lf = 0.00001
	case 3:
		lf = 0.0001
	case 4:
		lf = 0.001
	case 5:
		lf = 0.01
	case 6:
		lf = 0.1
	case 7:
		lf = 1
	case 8:
		lf = 10
	case 9:
		lf = 100
	case 10:
		lf = 1000
	case 11:
		lf = 10000
	case 12:
		lf = 100000
	case 13:
		lf = 1000000
	case 14:
		lf = 10000000
	case 15:
		lf = 100000000
	}
	return lf
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
	ep.finalGWP = setEconomicValue(0)
	ep.tradePartners = 1
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
	s += fmt.Sprintf("LOG:\n")
	for _, v := range ep.actionLog {
		s += fmt.Sprintf("%v\n", v)
	}
	return s
}

type interstellarDemand struct {
	//history []float64 - потенциально пишем в файл
	//TODO: написать аналитику шансов на ближайшее изменение
	demand float64
	dice   *dice.Dicepool
}

type InterstellarDemand interface {
	Demand() float64
	Aggregate()
}

func (id *interstellarDemand) Demand() float64 {
	return id.demand
}

func NewAggregatedDemand(seed ...string) *interstellarDemand {
	dice := *dice.New()
	switch len(seed) {
	default:
		dice.SetSeed(seed[0])
	case 0:
	}
	id := interstellarDemand{1.00, &dice}
	return &id
}

func (id *interstellarDemand) Aggregate() {

	dm := 0
	switch {
	case inRangeFl64(id.demand, -999.9, 0.49):
		dm = 5
	case inRangeFl64(id.demand, 0.5, 0.69):
		dm = 4
	case inRangeFl64(id.demand, 0.70, 0.79):
		dm = 3
	case inRangeFl64(id.demand, 0.8, 0.89):
		dm = 2
	case inRangeFl64(id.demand, 0.9, 0.95):
		dm = 1
	case inRangeFl64(id.demand, 1.06, 1.10):
		dm = -1
	case inRangeFl64(id.demand, 1.11, 1.20):
		dm = -2
	case inRangeFl64(id.demand, 1.21, 1.30):
		dm = -3
	case inRangeFl64(id.demand, 1.31, 1.4):
		dm = -4
	case inRangeFl64(id.demand, 1.41, 999.9):
		dm = -5
	}
	r := id.dice.Sroll("2d6") + dm
	r = utils.BoundInt(r, 2, 12) - 2
	changeBy := []float64{-0.06, -0.04, -0.03, -0.02, -0.01, 0, 0.01, 0.02, 0.03, 0.04, 0.06}
	id.demand = id.demand + changeBy[r]
	id.demand = utils.RoundFloat64(id.demand, 2)
}

func inRangeFl64(fl, min, max float64) bool {
	if fl < min {
		return false
	}
	if fl > max {
		return false
	}
	return true
}

func (id *interstellarDemand) String() string {
	return fmt.Sprintf("Aggregated Interstellar Demand: %v", id.demand)
}

func (id *interstellarDemand) Bar() string {
	s := fmt.Sprintf("Interstellar Demand: ")
	for i := 0; i < 200; i++ {
		if i <= int(id.demand*100) {
			s += "|"
		} else {
			s += " "
		}
	}
	return s
}

func GEPcode(ru int) string {
	val := 0
	switch {
	case ru == 0:
		val = 0
	case ru == 1:
		val = 1
	case utils.InRange(ru, 2, 3):
		val = 2
	case utils.InRange(ru, 4, 10):
		val = 3
	case utils.InRange(ru, 11, 30):
		val = 4
	case utils.InRange(ru, 31, 100):
		val = 5
	case utils.InRange(ru, 101, 250):
		val = 6
	case utils.InRange(ru, 251, 1000):
		val = 7
	case utils.InRange(ru, 1001, 2500):
		val = 8
	case utils.InRange(ru, 2501, 10000):
		val = 9
	case utils.InRange(ru, 10001, 25000):
		val = 10
	case utils.InRange(ru, 25001, 60000):
		val = 11
	case utils.InRange(ru, 60001, 180000):
		val = 12
	case utils.InRange(ru, 180001, 600000):
		val = 13
	case utils.InRange(ru, 600001, 1800000):
		val = 14
	case utils.InRange(ru, 1800001, 1800000000):
		val = 15
	}

	return ehex.ToCode(val)
}
