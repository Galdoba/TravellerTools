package tech

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	letharhic = iota
	average
	fast
)

type tlprogression struct {
	pace      int
	currentTL int
	baseGen   int
	gen       int
}

func Progress() {
	prog := tlprogression{}
	dp := dice.New()
	prog.baseGen = dp.Sroll("1d6")
	prog.pace = rollPace(dp)
	prog.gen = dp.Sroll("5d6")
	year := 0
	nextEvent := 0
	rateMap := rateMap()
	for prog.currentTL < 33 {
		fmt.Printf("Year %v: ", year)
		if nextEvent > year {
			fmt.Printf("next event=%v                      \r", nextEvent)
			year++
			continue
		}
		nextEvent = year + (rateMap[prog.currentTL][prog.pace] * prog.gen)

	}
}

func rollPace(dp *dice.Dicepool) int {
	switch dp.Sroll("1d6") {
	case 1, 2:
		return letharhic
	case 3, 4:
		return average
	case 5, 6:
		return fast
	}
	return -1
}

func rateMap() map[int][]int {
	rm := make(map[int][]int)
	for i := 0; i < 34; i++ {
		switch i {
		case 0, 1:
			rm[i] = []int{10000, 1000, 100}
		case 2, 3:
			rm[i] = []int{1000, 100, 10}
		case 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21:
			rm[i] = []int{100, 10, 1}
		case 22, 23, 24:
			rm[i] = []int{80, 8, 1}
		case 25, 26, 27:
			rm[i] = []int{60, 6, 1}
		case 28, 29, 30:
			rm[i] = []int{40, 4, 1}
		case 31, 32:
			rm[i] = []int{20, 2, 1}
		case 33:
			rm[i] = []int{10, 1, 1}
		}
	}
	return rm
}
