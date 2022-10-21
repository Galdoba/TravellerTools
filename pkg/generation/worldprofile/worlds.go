package worldprofile

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

const (
	port = iota
	size
	atmo
	hydro
	pops
	govr
	laws
	tl
	stA
	stB
	stC
	stD
	stE
	stX
	stF
	stG
	stH
	stY
	Hospitable
	Planetoid
	IceWorld
	RadWorld
	Inferno
	BigWorld
	Worldlet
	InnerWorld
	StormWorld
	SGG
	LGG
	IG
	PlanetaryRings
	AsteroidBelt
	validUWPregexp = "[ABCDEFGHXY][0123456789ABCDEF][0123456789ABCDEF][0123456789A][0123456789ABCDEF][0123456789ABCDEF][0123456789ABCDEFHGJ]-[0123456789ABCDEFGHJKL]"
)

/*
NewMain(seed string) string
NewSecondary(ssd survey.SecondSurveyData, otherType int) string



*/
//NewMain - New mainworld UWP

func ByMask(mask string, seed string) (string, error) {
	maskParts := strings.Split(mask, "")

	if err := maskIsValid(mask); err != nil {
		return "", err
	}

	return maskParts[0], fmt.Errorf("not implemented")
}

func maskIsValid(mask string) error {
	mp := strings.Split(mask, "")
	validUWPregexp := []string{
		"[ABCDEFGHXY]",            //port
		"[0123456789ABCDEF]",      //size
		"[0123456789ABCDEF]",      //atm
		"[0123456789A]",           //hydr
		"[0123456789ABCDEF]",      //pop
		"[0123456789ABCDEF]",      //govr
		"[0123456789ABCDEFHGJ]",   //law
		"-",                       //separator
		"[0123456789ABCDEFGHJKL]", //tl
	}
	for i, uwpSegment := range validUWPregexp {
		if m, err := regexp.MatchString(uwpSegment, mp[i]); err != nil {
			switch i {
			case 0:
				return fmt.Errorf("starport data invalid %v", m)
			case 1:
				return fmt.Errorf("size data invalid %v", m)
			case 2:
				return fmt.Errorf("atmospere data invalid %v", m)
			case 3:
				return fmt.Errorf("hydrospgere data invalid %v", m)
			case 4:
				return fmt.Errorf("population data invalid %v", m)
			case 5:
				return fmt.Errorf("goverment data invalid %v", m)
			case 6:
				return fmt.Errorf("laws data invalid %v", m)
			case 7:
				return fmt.Errorf("separator data invalid %v", m)
			case 8:
				return fmt.Errorf("tl data invalid %v", m)
			}
			fmt.Printf("segment %v with data %v = valid (%v)", i, uwpSegment, m)
		}

	}
	return nil
}

func NewMain(seed string) string {
	dp := dice.New().SetSeed(seed)
	statMap := make(map[int]int)
	///////
	statMap = rollSize(statMap, dp, 0)
	statMap = rollAtmo(statMap, dp, 0)
	statMap = rollHydro(statMap, dp, 0)
	statMap = rollPops(statMap, dp, 0)
	statMap = rollGovr(statMap, dp, 0)
	statMap = rollLaws(statMap, dp, 0)
	statMap = rollPort(statMap, dp, 0)
	statMap = rollTL(statMap, dp, 0)
	statMap = applyEnviromentalLimits(statMap, 0)
	return statMapToString(statMap)
}

type SurveyDataRetriver interface {
	MW_UWP() string
	GenerationSeed() string
}

func NewSecondary(ssd SurveyDataRetriver, worldType int, orbitalSuffix string) string {
	if worldType == LGG {
		return "Large Gas Gigant"
	}
	if worldType == SGG {
		return "Small Gas Gigant"
	}
	if worldType == IG {
		return "Ice Gigant"
	}
	if worldType == PlanetaryRings {
		return "Planetary Ring"
	}
	if worldType == AsteroidBelt {
		return "Asteroid Belt"
	}
	mwUWP := ssd.MW_UWP()
	mwStats := stringToStatMap(mwUWP)
	swStats := make(map[int]int)
	dp := dice.New().SetSeed(ssd.GenerationSeed() + "_" + orbitalSuffix)
	swStats = rollSize(swStats, dp, worldType)
	swStats = rollAtmo(swStats, dp, worldType)
	swStats = rollHydro(swStats, dp, worldType)
	swStats = rollPops(swStats, dp, worldType)
	if swStats[pops] > mwStats[pops]-1 {
		swStats[pops] = mwStats[pops] - 1
	}
	if swStats[pops] < 0 {
		swStats[pops] = 0
	}
	swStats = rollGovr(swStats, dp, worldType)
	swStats = rollLaws(swStats, dp, worldType)
	swStats = rollTL(swStats, dp, worldType)
	if swStats[tl] > mwStats[tl]-1 {
		swStats[tl] = mwStats[tl] - 1
	}
	if mwStats[tl] < 7 {
		swStats[pops] = 0
		swStats[govr] = 0
		swStats[laws] = 0
		swStats[tl] = 0
	}
	if swStats[tl] < 0 {
		swStats[tl] = 0
	}
	swStats = rollPort(swStats, dp, worldType)
	swStats = applyEnviromentalLimits(swStats, worldType)
	return statMapToString(swStats)
}

func statMapToString(statMap map[int]int) string {
	res := ""
	switch statMap[port] {
	case stA:
		res = "A"
	case stB:
		res = "B"
	case stC:
		res = "C"
	case stD:
		res = "D"
	case stE:
		res = "E"
	case stX:
		res = "X"
	case stF:
		res = "F"
	case stG:
		res = "G"
	case stH:
		res = "H"
	case stY:
		res = "Y"
	}
	res += ehex.New().Set(statMap[size]).Code()
	res += ehex.New().Set(statMap[atmo]).Code()
	res += ehex.New().Set(statMap[hydro]).Code()
	res += ehex.New().Set(statMap[pops]).Code()
	res += ehex.New().Set(statMap[govr]).Code()
	res += ehex.New().Set(statMap[laws]).Code()
	res += "-"
	res += ehex.New().Set(statMap[tl]).Code()
	return res
}

func stringToStatMap(uwp string) map[int]int {
	statMap := make(map[int]int)
	uwpData := strings.Split(uwp, "")
	switch uwpData[0] {
	case "A":
		statMap[port] = stA
	case "B":
		statMap[port] = stB
	case "C":
		statMap[port] = stC
	case "D":
		statMap[port] = stD
	case "E":
		statMap[port] = stE
	case "X":
		statMap[port] = stX
	case "F":
		statMap[port] = stF
	case "G":
		statMap[port] = stG
	case "H":
		statMap[port] = stH
	case "Y":
		statMap[port] = stY
	}
	statMap[size] = ehex.New().Set(uwpData[1]).Value()
	statMap[atmo] = ehex.New().Set(uwpData[2]).Value()
	statMap[hydro] = ehex.New().Set(uwpData[3]).Value()
	statMap[pops] = ehex.New().Set(uwpData[4]).Value()
	statMap[govr] = ehex.New().Set(uwpData[5]).Value()
	statMap[laws] = ehex.New().Set(uwpData[6]).Value()
	statMap[tl] = ehex.New().Set(uwpData[8]).Value()
	return statMap
}

func rollSize(statMap map[int]int, dp *dice.Dicepool, worldType int) map[int]int {
	switch worldType {
	default:
		statMap[size] = dp.Roll("2d6").Sum() - 2
	case BigWorld:
		statMap[size] = dp.Roll("2d6").Sum() + 7
	case RadWorld, StormWorld:
		statMap[size] = dp.Roll("2d6").Sum()
	case Inferno:
		statMap[size] = dp.Roll("1d6").Sum() + 6
	case Worldlet:
		statMap[size] = dp.Roll("1d6").Sum() - 3
	case Planetoid:
		statMap[size] = 0
	}
	if statMap[size] < 0 {
		statMap[size] = 0
	}
	return statMap
}

func rollAtmo(statMap map[int]int, dp *dice.Dicepool, worldType int) map[int]int {
	switch worldType {
	default:
		statMap[atmo] = dp.Flux() + statMap[size]
	case Planetoid:
		statMap[atmo] = 0
	case Inferno:
		statMap[atmo] = 11
	case StormWorld:
		statMap[atmo] = dp.Flux() + statMap[size] + 4
	}
	switch {
	case statMap[atmo] < 0 || statMap[size] == 0:
		statMap[atmo] = 0
	case statMap[atmo] > 15:
		statMap[atmo] = 15
	}
	return statMap
}

func rollHydro(statMap map[int]int, dp *dice.Dicepool, worldType int) map[int]int {
	dm := 0
	switch statMap[atmo] {
	case 0, 1, 10, 11, 12, 13, 14, 15:
		dm = -4
	}
	switch worldType {
	default:
		statMap[hydro] = dp.Flux() + statMap[atmo] + dm
	case Planetoid, Inferno:
		statMap[hydro] = 0
	case InnerWorld, StormWorld:
		statMap[hydro] = dp.Flux() + statMap[atmo] + dm - 4
	}
	if statMap[size] < 2 {
		statMap[hydro] = 0
	}
	if statMap[hydro] < 0 {
		statMap[hydro] = 0
	}
	if statMap[hydro] > 10 {
		statMap[hydro] = 10
	}
	return statMap
}

func rollPops(statMap map[int]int, dp *dice.Dicepool, worldType int) map[int]int {
	switch worldType {
	default:
		statMap[pops] = dp.Roll("2d6").Sum() - 2
		if statMap[pops] == 10 {
			statMap[pops] = dp.Roll("2d6").Sum() + 3
		}
	case IceWorld, StormWorld:
		statMap[pops] = dp.Roll("2d6").Sum() - 6
	case InnerWorld:
		statMap[pops] = dp.Roll("2d6").Sum() - 4
	case RadWorld, Inferno:
		statMap[pops] = 0
	}
	return statMap
}

func rollGovr(statMap map[int]int, dp *dice.Dicepool, worldType int) map[int]int {
	switch worldType {
	default:
		statMap[govr] = dp.Flux() + statMap[pops]
	case RadWorld, Inferno:
		statMap[govr] = 0
	}
	switch {
	case statMap[govr] < 0:
		statMap[govr] = 0
	case statMap[govr] > 15:
		statMap[govr] = 15
	}
	return statMap
}

func rollLaws(statMap map[int]int, dp *dice.Dicepool, worldType int) map[int]int {
	switch worldType {
	default:
		statMap[laws] = dp.Flux() + statMap[govr]
	case RadWorld, Inferno:
		statMap[laws] = 0
	}
	switch {
	case statMap[laws] < 0:
		statMap[laws] = 0
	case statMap[laws] > 18:
		statMap[laws] = 18
	}
	return statMap
}

func rollPort(statMap map[int]int, dp *dice.Dicepool, worldType int) map[int]int {
	switch worldType {
	default:
		portIndex := statMap[pops] - dp.Roll("1d6").Sum()
		switch {
		case portIndex >= 4:
			statMap[port] = stF
		case portIndex == 3:
			statMap[port] = stG
		case portIndex == 2 || portIndex == 1:
			statMap[port] = stH
		case portIndex <= 0:
			statMap[port] = stY
		}
	case Inferno:
		statMap[port] = stY
	case 0: //Mainworld
		stDM := 0
		switch {
		case statMap[pops] == 8 || statMap[pops] == 9:
			stDM = 1
		case statMap[pops] > 9:
			stDM = 2
		case statMap[pops] < 3:
			stDM = -2
		case statMap[pops] == 3 || statMap[pops] == 4:
			stDM = -1
		}
		statMap[port] = stX
		stR := dp.Roll("2d6").DM(stDM).Sum()
		switch stR {
		case 3, 4:
			statMap[port] = stE
		case 5, 6:
			statMap[port] = stD
		case 7, 8:
			statMap[port] = stC
		case 9, 10:
			statMap[port] = stB
		}
		if stR > 10 {
			statMap[port] = stA
		}
	}
	return statMap
}

func rollTL(statMap map[int]int, dp *dice.Dicepool, worldType int) map[int]int {
	statMap[tl] = dp.Roll("1d6").Sum()
	switch statMap[port] {
	case stA:
		statMap[tl] = statMap[tl] + 6
	case stB, stF:
		statMap[tl] = statMap[tl] + 4
	case stC:
		statMap[tl] = statMap[tl] + 2
	case stX:
		statMap[tl] = statMap[tl] - 4
	}
	switch statMap[size] {
	case 0, 1:
		statMap[tl] = statMap[tl] + 2
	case 2, 3, 4:
		statMap[tl] = statMap[tl] + 1
	}
	switch statMap[atmo] {
	case 0, 1, 2, 3, 10, 11, 12, 13, 14, 15:
		statMap[tl] = statMap[tl] + 1
	}
	switch statMap[hydro] {
	case 9:
		statMap[tl] = statMap[tl] + 1
	case 10:
		statMap[tl] = statMap[tl] + 2
	}
	switch statMap[pops] {
	case 1, 2, 3, 4, 5:
		statMap[tl] = statMap[tl] + 1
	case 9:
		statMap[tl] = statMap[tl] + 2
	case 10, 11, 12, 13, 14, 15:
		statMap[tl] = statMap[tl] + 4
	}
	switch statMap[govr] {
	case 0, 5:
		statMap[tl] = statMap[tl] + 1
	case 13:
		statMap[tl] = statMap[tl] - 2
	}
	if statMap[tl] < 0 {
		statMap[tl] = 0
	}
	return statMap
}

func applyEnviromentalLimits(statMap map[int]int, worldType int) map[int]int {
	min := 0
	current := statMap[tl]
	switch statMap[atmo] {
	case 0, 1:
		min = 8
	case 2, 3:
		min = 5
	case 4, 7, 9:
		min = 3
	case 10:
		min = 8
	case 11:
		min = 9
	case 12:
		min = 10
	case 13, 14:
		min = 5
	case 15:
		min = 8
	}
	if current < min {
		statMap[pops] = 0
		statMap[govr] = 0
		statMap[laws] = 0
		statMap[tl] = 0
		switch worldType {
		default:
			statMap[port] = stY
		case 0:
			statMap[port] = stX
		}

	}
	return statMap
}
