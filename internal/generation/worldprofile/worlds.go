package worldprofile

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/internal/ehex"
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
)

/*
NewMain(seed string) string
NewSecondary(ssd survey.SecondSurveyData, otherType int) string



*/
//NewMain - New mainworld UWP
func NewMain(seed string) string {
	dp := dice.New().SetSeed(seed)
	statMap := make(map[int]int)
	///////
	statMap = rollSize(statMap, dp)
	statMap = rollAtmo(statMap, dp)
	statMap = rollHydro(statMap, dp)
	statMap = rollPops(statMap, dp)
	statMap = rollGovr(statMap, dp)
	statMap = rollLaws(statMap, dp)
	statMap = rollPort(statMap, dp)
	statMap = rollTL(statMap, dp)
	statMap = applyEnviromentalLimits(statMap)
	return statMapToString(statMap)
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

func rollSize(statMap map[int]int, dp *dice.Dicepool) map[int]int {
	statMap[size] = dp.Roll("2d6").Sum() - 2
	return statMap
}

func rollAtmo(statMap map[int]int, dp *dice.Dicepool) map[int]int {
	statMap[atmo] = dp.Flux() + statMap[size]
	switch {
	case statMap[atmo] < 0 || statMap[size] == 0:
		statMap[atmo] = 0
	case statMap[atmo] > 15:
		statMap[atmo] = 15
	}
	return statMap
}

func rollHydro(statMap map[int]int, dp *dice.Dicepool) map[int]int {
	dm := 0
	switch statMap[atmo] {
	case 0, 1, 10, 11, 12, 13, 14, 15:
		dm = -4
	}
	statMap[hydro] = dp.Flux() + statMap[atmo] + dm
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

func rollPops(statMap map[int]int, dp *dice.Dicepool) map[int]int {
	statMap[pops] = dp.Roll("2d6").DM(-2).Sum()
	if statMap[pops] == 10 {
		statMap[pops] = dp.Roll("2d6").DM(3).Sum()
	}
	return statMap
}

func rollGovr(statMap map[int]int, dp *dice.Dicepool) map[int]int {
	statMap[govr] = dp.Flux() + statMap[pops]
	switch {
	case statMap[govr] < 0:
		statMap[govr] = 0
	case statMap[govr] > 15:
		statMap[govr] = 15
	}
	return statMap
}

func rollLaws(statMap map[int]int, dp *dice.Dicepool) map[int]int {
	statMap[laws] = dp.Flux() + statMap[govr]
	switch {
	case statMap[laws] < 0:
		statMap[laws] = 0
	case statMap[laws] > 18:
		statMap[laws] = 18
	}
	return statMap
}

func rollPort(statMap map[int]int, dp *dice.Dicepool) map[int]int {
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
	return statMap
}

func rollTL(statMap map[int]int, dp *dice.Dicepool) map[int]int {
	statMap[tl] = dp.Roll("1d6").Sum()
	switch statMap[port] {
	case stA:
		statMap[tl] = statMap[tl] + 6
	case stB:
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

func applyEnviromentalLimits(statMap map[int]int) map[int]int {
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
		fmt.Print(statMapToString(statMap), " -> ")
		statMap[pops] = 0
		statMap[govr] = 0
		statMap[laws] = 0
		statMap[tl] = 0
		statMap[port] = stX
		fmt.Print(statMapToString(statMap), "\n")
	}
	return statMap
}
