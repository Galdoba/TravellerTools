package worldcharacter

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/mgt2trade/traffic/tradecodes"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"github.com/Galdoba/TravellerTools/t4/core/task/pocketempire/economics"
)

const (
	selfDetermination_Roll   = 0
	progression_Roll         = 1
	planning_Roll            = 2
	advancement_Roll         = 3
	growth_Roll              = 4
	militancy_Roll           = 5
	unity_Roll               = 6
	tolerance_Roll           = 7
	PROGRESSION_STAT         = "Progression"
	PLANNING_STAT            = "Planning"
	ADVANCEMENT_STAT         = "Advancement"
	GROWTH_STAT              = "Growth"
	MILITANCY_STAT           = "Militancy"
	UNITY_STAT               = "Unity"
	TOLERANCE_STAT           = "Tolerance"
	RADICAL_Progression      = "Radical"
	PROGRESSIVE_Progression  = "Progressive"
	CONSERVATIVE_Progression = "Conservative"
	REACTIONARY_Progression  = "Reactionary"
	VERY_SHORT_TERM_Planning = "Very Short Term (1 year)"
	SHORT_TERM_Planning      = "Short Term (2-5 years)"
	MEDIUM_TERM_Planning     = "Medium Term (6-10 years)"
	LONG_TERM_Planning       = "Long Term (11-50 years)"
	VERY_LONG_TERM_Planning  = "Very Long Term (51-100 years)"
	FAR_FUTURE_Planning      = "Far Future (>100 years)"
	ENTERPRISING_Advancement = "Enterprising"
	ADVANCING_Advancement    = "Advancing"
	INDIFFIRENT_Advancement  = "Indiffirent"
	STAGNANT_Advancement     = "Stagnant"
	MILITANT_Militancy       = "Militant"
	NEUTRAL_Militancy        = "Neutral"
	PEACEABLE_Militancy      = "Peaceable"
	CONCILIATORY_Militancy   = "Conciliatory"
	EXPANSIONIST_Growth      = "Expansionist"
	COMPETITIVE_Growth       = "Competitive"
	UNAGRESSIVE_Growth       = "Unagressive"
	PASSIVE_Growth           = "Passive"
	MONOLITHIC_Unity         = "Monolithic"
	HARMONIOUS_Unity         = "Harmonious"
	DISCORDANT_Unity         = "Discordant"
	FRAGMENTED_Unity         = "Fragmented"
	XENOPHILIC_Tolerance     = "Xenophilic"
	FRIENDLY_Tolerance       = "Friendly"
	NEUTRAL_Tolerance        = "Neutral"
	ALOOF_Tolerance          = "Aloof"
	XENOPHOBIC_Tolerance     = "Xenophobic"
)

type world struct {
	name              string
	uwp               uwp.UWP
	pbg               string
	tradecodes        []string
	tradeGoods        []string
	econEx            economics.EconomicPower
	selfDetermination ehex.Ehex //0-10
	localPopularity   ehex.Ehex //0-15
	progression       int
	factions          []int //распределение по фракциям суммарно 100%
	hex               hexagon.Hexagon
	baseRolls         []int //всего 7 Progression/Planning/Advancment/Grown/Militancy/Unity/Tolerance
}

type World interface {
	Attitude(string) string
	SelfDetermination() int
	Progression() int
	Planning() int
	Advancement() int
	Growth() int
	Militancy() int
	Unity() int
	Tolerance() int
}

func WorldCharacter(worldName, uwpStr, pbgStr string, x, y int) (*world, error) {
	dice := dice.New().SetSeed(worldName + uwpStr + pbgStr)
	wc := world{}
	wc.name = worldName
	wc.uwp = uwp.Inject(uwpStr)
	wc.pbg = pbgStr
	hex, err := hexagon.New(hexagon.Feed_HEX, x, y)
	if err != nil {
		return &wc, err
	}
	wc.hex = hex
	wc.tradecodes = setupTradeCodes(wc.uwp, dice)
	//wc.tradeGoods = availableResources(wc.uwp, wc.tradecodes, dice)
	wc.econEx = economics.GenerateInitialEconomicPower(&wc, dice)

	wc.setupBaseRolls(dice)
	wc.selfDetermination = ehex.New().Set(wc.baseRolls[selfDetermination_Roll] - 2)
	return &wc, nil
}

func (wc *world) setupBaseRolls(dice *dice.Dicepool) {
	wc.baseRolls = nil
	for i := 0; i < 8; i++ {
		wc.baseRolls = append(wc.baseRolls, dice.Sroll("2d6"))
	}
}

func (wc *world) Attitude(stat string) string {
	switch stat {
	default:
		return "Unknown Stat"
	case PROGRESSION_STAT:
		return progressionAttutude(wc.Progression())
	case PLANNING_STAT:
		return planningAttutude(wc.Planning())
	case ADVANCEMENT_STAT:
		return advancementAttitude(wc.Advancement())
	case GROWTH_STAT:
		return growthAttitude(wc.Growth())
	case MILITANCY_STAT:
		return militancyAttitude(wc.Militancy())
	case UNITY_STAT:
		return unityAttitude(wc.Unity())
	case TOLERANCE_STAT:
		return toleranceAttitude(wc.Tolerance())
	}
}

func progressionAttutude(val int) string {
	if val <= 3 {
		return RADICAL_Progression
	}
	if val <= 7 {
		return PROGRESSIVE_Progression
	}
	if val <= 11 {
		return CONSERVATIVE_Progression
	}
	return REACTIONARY_Progression
}

func planningAttutude(val int) string {
	if val <= 3 {
		return VERY_SHORT_TERM_Planning
	}
	if val <= 5 {
		return SHORT_TERM_Planning
	}
	if val <= 7 {
		return MEDIUM_TERM_Planning
	}
	if val <= 9 {
		return LONG_TERM_Planning
	}
	if val <= 11 {
		return VERY_LONG_TERM_Planning
	}
	return FAR_FUTURE_Planning
}

func advancementAttitude(val int) string {
	if val <= 5 {
		return ENTERPRISING_Advancement
	}
	if val <= 9 {
		return ADVANCING_Advancement
	}
	if val <= 12 {
		return INDIFFIRENT_Advancement
	}
	return STAGNANT_Advancement
}

func growthAttitude(val int) string {
	if val <= 3 {
		return EXPANSIONIST_Growth
	}
	if val <= 6 {
		return COMPETITIVE_Growth
	}
	if val <= 10 {
		return UNAGRESSIVE_Growth
	}
	return PASSIVE_Growth
}

func militancyAttitude(val int) string {
	if val <= 4 {
		return MILITANT_Militancy
	}
	if val <= 8 {
		return NEUTRAL_Militancy
	}
	if val <= 11 {
		return PEACEABLE_Militancy
	}
	return CONCILIATORY_Militancy
}

func unityAttitude(val int) string {
	if val <= 3 {
		return MONOLITHIC_Unity
	}
	if val <= 7 {
		return HARMONIOUS_Unity
	}
	if val <= 11 {
		return DISCORDANT_Unity
	}
	return FRAGMENTED_Unity
}

func toleranceAttitude(val int) string {
	if val <= 3 {
		return XENOPHILIC_Tolerance
	}
	if val <= 6 {
		return FRIENDLY_Tolerance
	}
	if val <= 9 {
		return NEUTRAL_Tolerance
	}
	if val <= 11 {
		return ALOOF_Tolerance
	}
	return XENOPHOBIC_Tolerance
}

func (wc *world) Progression() int {
	dm := 0
	if wc.uwp.Pops() >= 6 {
		dm++
	}
	if wc.uwp.Pops() >= 9 {
		dm++
	}
	if wc.uwp.Laws() >= 10 {
		dm++
	}
	if wc.econEx.Culture() <= 3 {
		dm--
	}
	if wc.econEx.Culture() >= 8 {
		dm++
	}
	return wc.baseRolls[progression_Roll] + dm
}

func (wc *world) Planning() int {
	dm := 0
	if wc.Attitude(PROGRESSION_STAT) == CONSERVATIVE_Progression {
		dm = dm + 2
	}
	if wc.Attitude(PROGRESSION_STAT) == REACTIONARY_Progression {
		dm = dm + 2
	}
	if wc.Attitude(PROGRESSION_STAT) == RADICAL_Progression {
		dm = dm - 2
	}
	return wc.baseRolls[planning_Roll] + dm
}

func (wc *world) Advancement() int {
	dm := 0
	if wc.uwp.Laws() >= 10 {
		dm = dm + 1
	}
	if wc.Attitude(PROGRESSION_STAT) == CONSERVATIVE_Progression {
		dm = dm + 3
	}
	if wc.Attitude(PROGRESSION_STAT) == REACTIONARY_Progression {
		dm = dm + 6
	}
	return wc.baseRolls[advancement_Roll] + dm
}

func (wc *world) Growth() int {
	dm := 0
	if wc.uwp.Laws() >= 10 {
		dm = dm + 1
	}
	if wc.econEx.Culture() <= 3 {
		dm = dm - 1
	}
	if wc.econEx.Culture() >= 8 {
		dm = dm + 1
	}
	return wc.baseRolls[growth_Roll] + dm
}

func (wc *world) Militancy() int {
	dm := 0
	if wc.uwp.Laws() >= 10 {
		dm = dm + 1
	}
	if wc.Attitude(GROWTH_STAT) == EXPANSIONIST_Growth {
		dm = dm - 2
	}
	if wc.Attitude(GROWTH_STAT) == COMPETITIVE_Growth {
		dm = dm - 1
	}
	if wc.Attitude(GROWTH_STAT) == PASSIVE_Growth {
		dm = dm + 2
	}
	return wc.baseRolls[militancy_Roll] + dm
}

func (wc *world) Unity() int {
	dm := 0
	if wc.uwp.Laws() <= 4 {
		dm = dm + 1
	}
	if wc.uwp.Laws() >= 10 {
		dm = dm - 1
	}
	if wc.uwp.Govr() <= 2 {
		dm = dm + 1
	}
	if wc.uwp.Govr() == 7 {
		dm = dm + 3
	}
	if wc.uwp.Govr() == 15 {
		dm = dm - 1
	}
	if wc.Attitude(GROWTH_STAT) == PASSIVE_Growth {
		dm = dm + 2
	}
	return wc.baseRolls[unity_Roll] + dm
}

func (wc *world) Tolerance() int {
	dm := 0
	if wc.uwp.Starport() == "A" {
		dm = dm - 2
	}
	if wc.uwp.Starport() == "B" {
		dm = dm - 1
	}
	if wc.uwp.Starport() == "D" {
		dm = dm + 1
	}
	if wc.uwp.Starport() == "E" {
		dm = dm + 3
	}
	if wc.Attitude(PROGRESSION_STAT) == CONSERVATIVE_Progression {
		dm = dm + 2
	}
	if wc.Attitude(PROGRESSION_STAT) == REACTIONARY_Progression {
		dm = dm + 4
	}
	if wc.uwp.Laws() >= 10 {
		dm = dm + 1
	}
	return wc.baseRolls[tolerance_Roll] + dm
}

func (wc *world) SelfDetermination() int {
	return wc.selfDetermination.Value()
}

func setupTradeCodes(u uwp.UWP, dice *dice.Dicepool) []string {
	tc, _ := tradecodes.FromUWP(u)
	switch rollTemp(u, dice) {
	case 1:
		tc = append(tc, "Fr")
	case 2:
		tc = append(tc, "Co")
	case 4:
		tc = append(tc, "Ho")
	case 5:
		tc = append(tc, "Bo")
	}
	nl := nativeLifeRoll(u, dice, tc)
	if nl > 0 {
		tc = append(tc, fmt.Sprintf("NL%v", nl))
	}
	return tc
}

func nativeLifeRoll(u uwp.UWP, dice *dice.Dicepool, tc []string) int {
	dm := 0
	switch u.Atmo() {
	case 0:
		dm = dm - 3
	case 4, 5, 6, 7, 8, 9:
		dm = dm + 4
	}
	switch u.Hydr() {
	case 0:
		dm = dm - 2
	case 2, 3, 4, 5, 6, 7, 8:
		dm = dm + 1
	}
	for _, t := range tc {
		switch t {
		case "Fr", "Co", "Ho", "Bo":
			dm = dm - 1
		}
	}
	fmt.Println("NL dm:", dm)
	return dice.Sroll("2d6") + dm - 10
}

func rollTemp(u uwp.UWP, dice *dice.Dicepool) int {
	dm := 0
	temperature := 0
	extreme := false
	size := ehex.ToCode(u.Size())
	switch size {
	case "0", "1":
		extreme = true
	case "2", "3":
		dm = -2
	case "4", "5", "E":
		dm = -1
	case "6", "7":
		dm = 0
	case "8", "9":
		dm = 1
	case "A", "D", "F":
		dm = 2
	case "B", "C":
		dm = 6
	}
	r := dice.Sroll("2d6") + dm
	switch r {
	case 3, 4:
		temperature = 2
		if extreme {
			temperature = 1
		}
	case 5, 6, 7, 8, 9:
		temperature = 3
		if extreme {
			switch dice.Sroll("1d2") {
			case 1:
				temperature = 1
			case 2:
				temperature = 5
			}
		}
	case 10, 11:
		temperature = 4
		if extreme {
			temperature = 5
		}
	}
	if r <= 2 {
		temperature = 1
	}
	if r >= 12 {
		temperature = 5
	}
	return temperature
}

func (wc *world) UWP() uwp.UWP {
	return wc.uwp
}

func (wc *world) TradeCodes() []string {
	return wc.tradecodes
}

func (wc *world) PBG() string {
	return wc.pbg
}

func (wc *world) Descr() string {
	s := fmt.Sprintf("World: %v\n", wc.name)
	s += fmt.Sprintf("UWP  : %v-%v\n", wc.uwp.String(), wc.econEx.String())
	s += fmt.Sprintf("HEX  : %v\n", wc.hex.String())
	s += fmt.Sprintf("PBG: : %v\n", wc.pbg)
	s += fmt.Sprintf("SD   : %v\n", wc.selfDetermination.Code())
	s += fmt.Sprintf("TC   : ")
	for _, tc := range wc.tradecodes {
		s += tc + " "
	}
	s += "\n"
	s += "DEBUG INFO:-------------\n"
	s += fmt.Sprintf("%v\n", wc.baseRolls)
	s += fmt.Sprintf("Progression: %v (%v)\n", wc.Progression(), wc.Attitude(PROGRESSION_STAT))
	s += fmt.Sprintf("Planning   : %v (%v)\n", wc.Planning(), wc.Attitude(PLANNING_STAT))
	s += fmt.Sprintf("Advancement: %v (%v)\n", wc.Advancement(), wc.Attitude(ADVANCEMENT_STAT))
	s += fmt.Sprintf("Growth     : %v (%v)\n", wc.Growth(), wc.Attitude(GROWTH_STAT))
	s += fmt.Sprintf("Militancy  : %v (%v)\n", wc.Militancy(), wc.Attitude(MILITANCY_STAT))
	s += fmt.Sprintf("Unity      : %v (%v)\n", wc.Unity(), wc.Attitude(UNITY_STAT))
	s += fmt.Sprintf("Tolerance  : %v (%v)\n", wc.Tolerance(), wc.Attitude(TOLERANCE_STAT))
	return s
}
