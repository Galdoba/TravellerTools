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
	dp            *dice.Dicepool
	pace          int
	currentTL     int
	baseGen       int
	gen           int
	decline       bool
	tlCap         int
	destinyPoints int
}

func Progress() {
	prog := tlprogression{}
	prog.dp = dice.New()
	prog.baseGen = prog.dp.Sroll("1d6")
	prog.rollPace()
	prog.gen = prog.dp.Sroll("2d6") + 6
	year := 0
	prog.tlCap = tlCap(prog.dp)
	rateMap := rateMap()
	nextEvent := rateMap[prog.currentTL][prog.pace] * prog.gen
	fmt.Println("TL CAP:", prog.tlCap)
	for prog.currentTL < 33 {

		//fmt.Printf("Year %v: ", year)
		if nextEvent > year {
			//	fmt.Printf("next event=%v                      \r", nextEvent)
			year++
			continue
		}
		fmt.Printf("Year %v: ", year)
		nextEvent = year + (rateMap[prog.currentTL][prog.pace] * prog.gen)
		if prog.currentTL == prog.tlCap {
			nextEvent = (rateMap[prog.currentTL][prog.pace] * 100) + year
			switch prog.dp.Sroll("1d6") {
			case 1:
				fmt.Println("Awakening")
				prog.tlCap = 34
			case 2, 3:
				fmt.Println("Regression")
				prog.currentTL = prog.currentTL - prog.dp.Sroll("1d6")
			case 4, 5:
				fmt.Println("Decline")
				prog.decline = true
				prog.tlCap = 34
			case 6:
				fmt.Println("Crisis")
				prog.currentTL = prog.currentTL - prog.dp.Sroll("2d6")
			}
		}
		if prog.advance() {
			break
		}
		fmt.Printf("\nSophont is now TL%v                     \n", prog.currentTL)
		prog.baseGen = prog.dp.Sroll("1d6")
		if prog.baseGen == 6 {
			fmt.Printf("Year %v : ", year)
			switch prog.dp.Sroll("1d6") {
			case 1:
				paus := prog.dp.Sroll("1d6") * prog.gen
				fmt.Printf("event=PAUSE %v years                     \n", paus)
				nextEvent = nextEvent + (paus)
			case 2:
				lost := prog.dp.Sroll("1d6")
				prog.currentTL = prog.currentTL - lost
				fmt.Printf("event=REGRESSION %v levels                      \n", lost)
			case 3:
				prog.pace = letharhic
				fmt.Printf("event=SHIFT PACE TO LETHARGIC                      \n")
			case 4:
				prog.pace = average
				fmt.Printf("event=SHIFT PACE TO AVERAGE                      \n")
			case 5:
				prog.pace = fast
				fmt.Printf("event=SHIFT PACE TO FAST                      \n")
			case 6:
				prog.changeDirection()
				fmt.Printf("event=PROGRESSION DIRECTION CHANGED                      \n")
			}
		}

	}
	fmt.Println("\nSophonts vanished")
}

func (prog *tlprogression) changeDirection() {
	if prog.decline == false {
		prog.decline = true
	} else {
		prog.decline = false
	}
}

func (prog *tlprogression) advance() bool {

	if prog.decline == false {

		prog.currentTL++
		fmt.Printf("Proggressed to TL%v\n", prog.currentTL)
		if prog.currentTL > 21 {
			prog.destinyPoints = prog.destinyPoints + getDestinypoints(prog.dp)
		}
		if prog.currentTL == 33 {
			singularity := prog.dp.Flux() + prog.destinyPoints
			if singularity < -2 {
				fmt.Println("Self-Destruction")
				fmt.Println("Disputies about futere path of society bring Destruction")
				prog.currentTL = -1
				return true
			}
			if singularity > 3 {
				fmt.Println("Pastoralization")
				fmt.Println("Society rejects Technology and focuses on family and social interaction")
				prog.currentTL = prog.dp.Sroll("2d6")
				return false
			}
			switch singularity {
			case -2:
				fmt.Println("Regression")
				fmt.Println("Disputies about futere path of society force a regression to previous levels")
				prog.currentTL = prog.dp.Sroll("1d6")
				return false
			case -1:
				fmt.Println("Paralizys")
				fmt.Println("Society deadlocks on its future paths")
				prog.currentTL = 21 - prog.dp.Sroll("1d6")
				return false
			case 0:
				fmt.Println("Ascent to a higher plane")
				fmt.Println("Society solves a chelenges of technologiey and ascends to a higher plane")
				prog.currentTL = -1
				return true
			case 1:
				fmt.Println("Acceptence")
				fmt.Println("Society accepts technology as a part of its larger social structure")
				prog.currentTL = 21
				return false
			case 2:
				fmt.Println("Divergence")
				fmt.Println("Each of the path of society choses its own futere path")
				fmt.Println(prog.dp.Flux()+prog.destinyPoints, prog.dp.Flux()+prog.destinyPoints, prog.dp.Flux()+prog.destinyPoints)
				prog.currentTL = prog.dp.Sroll("2d6")
				return true
			}
		}
	} else {
		prog.currentTL--
		fmt.Printf("Reggressed to TL%v\n", prog.currentTL)
	}
	if prog.currentTL < 0 {
		fmt.Println("Extinct")
		return true
	}
	return false
}

func getDestinypoints(dp *dice.Dicepool) int {
	fl := dp.Flux()
	switch fl {
	case -5, -4:
		return -2
	case -3, -2:
		return -1
	case -1, 0, 1:
		return 0
	case 3, 2:
		return 1
	case 4, 5:
		return 2
	}
	return -999
}

func (p *tlprogression) rollPace() int {
	switch p.dp.Sroll("1d6") {
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

func tlCap(dp *dice.Dicepool) int {
	r := 10*dp.Sroll("1d6") + dp.Sroll("1d6")
	switch r {
	case 11:
		return 1
	case 12:
		return 2
	case 13:
		return 3
	case 14:
		return 4
	case 15:
		return 5
	case 16:
		return 6
	case 21:
		return 7
	case 22:
		return 8
	case 23:
		return 9
	case 24:
		return 10
	case 25:
		return 11
	case 26:
		return 12
	case 31:
		return 13
	case 32:
		return 14
	case 33:
		return 15
	case 34:
		return 16
	case 35:
		return 17
	case 36:
		return 18
	default:
		return 34
	}
}
