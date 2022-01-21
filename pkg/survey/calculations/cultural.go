package calculations

import (
	"strings"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/internal/ehex"
)

func Cultural(uwp, seed string, ix int) string {

	hex := strings.Split(uwp, "")
	pop := ehex.New().Set(hex[4])
	tl := ehex.New().Set(hex[8])
	if pop.Value() == 0 {
		return "[0000]"
	}
	h, a, st, sy := 1, 1, 1, 1
	d := dice.New().SetSeed(seed)
	h = pop.Value() + d.Flux()
	if h < 1 {
		h = 1
	}
	a = pop.Value() + ix
	if a < 1 {
		a = 1
	}
	st = d.Flux() + 5
	if st < 1 {
		st = 1
	}
	sy = d.Flux() + tl.Value()
	if sy < 1 {
		sy = 1
	}
	return "[" + ehex.New().Set(h).String() + ehex.New().Set(a).String() + ehex.New().Set(st).String() + ehex.New().Set(sy).String() + "]"
}

func CxValid(cx, uwp string) bool {
	culturalInvalid := []string{"[????]", "", "----", "[]"}
	for _, val := range culturalInvalid {
		if cx == val {
			return false
		}
	}
	if len(cx) != 6 {
		return false
	}
	hexu := strings.Split(uwp, "")
	pop := ehex.New().Set(hexu[4])
	if pop.Value() == 0 && cx != "[0000]" {
		return false
	}
	hex := strings.Split(cx, "")
	if !hetValid(ehex.New().Set(hex[1]).Value()) {
		return false
	}
	if !accValid(ehex.New().Set(hex[2]).Value()) {
		return false
	}
	if !strValid(ehex.New().Set(hex[3]).Value()) {
		return false
	}
	if !symValid(ehex.New().Set(hex[4]).Value()) {
		return false
	}
	return true
}

// func reRollHeterogenity(d *dice.Dicepool, pop ehex.DataRetriver) int {
// 	h := pop.Value() + d.FluxNext()
// 	if h < 1 {
// 		h = 1
// 	}
// 	return h
// }

func hetValid(h int) bool {
	return h <= 20
}

func accValid(a int) bool {
	return a <= 21
}

func reRollAcceptance(d *dice.Dicepool, ix int, pop ehex.Ehex) int {
	return 0
}

func strValid(st int) bool {
	return st <= 11
}

func symValid(sy int) bool {
	if sy > 23 {
		return false
	}
	return true
}
