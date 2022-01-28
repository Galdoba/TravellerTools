package calculations

import (
	"strings"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/internal/ehex"
)

func UWPvalid(uwp string) bool {

	hex := strings.Split(uwp, "")
	for len(hex) < 9 {
		hex = append(hex, "")
	}
	if !starportValid(ehex.New().Set(hex[0])) {
		return false
	}
	if !sizeValid(ehex.New().Set(hex[1])) {
		return false
	}
	if !atmoValid(ehex.New().Set(hex[2])) {
		return false
	}
	if !hydrValid(ehex.New().Set(hex[3])) {
		return false
	}
	if !popsValid(ehex.New().Set(hex[4])) {
		return false
	}
	if !govrValid(ehex.New().Set(hex[5])) {
		return false
	}
	if !lawsValid(ehex.New().Set(hex[6])) {
		return false
	}
	if !tlValid(ehex.New().Set(hex[8])) {
		return false
	}
	////////////
	return true
}

func FixUWP(uwp string, seed string) string {
	d := dice.New().SetSeed(seed)
	switch uwp {
	case "X233000-X":
		return "X233000-0"
	case "X420000-X":
		return "X420000-0"
	case "X400000-X":
		return "X400000-0"
	case "X100000-X":
		return "X100000-0"
	case "X7A6000-X":
		return "X7A6000-0"
	case "X424000-X":
		return "X424000-0"
	case "X411000-X":
		return "X411000-0"
	case "X110000-X":
		return "X110000-0"
	case "X000000-X":
		return "X000000-0"
	case "X439000-X":
		return "X439000-0"
	case "X000XXX-X":
		return "X000000-0"
	case "B453889-X":
		uwp = "B453889-?"
	case "X200000-X":
		return "X200000-0"
	case "X484XXX-X":
		return "X484000-0"
	case "C857360-N":
		uwp = "C857360-?"
	case "A6VV997-D":
		uwp = "A6??997-D"
	case "XXXXXXX-X":
		uwp = "???????-?"
	case "":
		uwp = "???????-?"
	}
	hex := strings.Split(uwp, "")
	if !starportValid(ehex.New().Set(hex[0])) {
		hex[0] = reRollStarport(d).Code()
	}
	if !sizeValid(ehex.New().Set(hex[1])) {
		hex[1] = reRollSize(d).Code()
	}
	if !atmoValid(ehex.New().Set(hex[2])) {
		hex[2] = reRollAtmo(d, ehex.New().Set(hex[1])).Code()
	}
	if !hydrValid(ehex.New().Set(hex[3])) {
		hex[3] = reRollHydr(d, ehex.New().Set(hex[1]), ehex.New().Set(hex[2])).Code()
	}
	if !popsValid(ehex.New().Set(hex[4])) {
		hex[4] = reRollPops(d).Code()
	}
	if !govrValid(ehex.New().Set(hex[5])) {
		hex[5] = reRollGovr(d, ehex.New().Set(hex[4])).Code()
	}
	if !lawsValid(ehex.New().Set(hex[6])) {
		hex[6] = reRollLaws(d, ehex.New().Set(hex[5])).Code()
	}
	if !tlValid(ehex.New().Set(hex[8])) {
		hex[8] = reRollTL(d, hex).Code()
	}
	uwp = ""
	for _, h := range hex {
		uwp += h
	}
	return uwp
}

func starportValid(eh ehex.Ehex) bool {
	switch eh.Code() {
	default:
		return false
	case "A", "B", "C", "D", "E", "X", "Y", "F", "G", "H":
	}
	return true
}

func reRollStarport(d *dice.Dicepool) ehex.Ehex {
	stpt := "X"
	switch d.Roll("2d6").Sum() {
	case 2, 3, 4:
		stpt = "A"
	case 5, 6:
		stpt = "B"
	case 7, 8:
		stpt = "C"
	case 9:
		stpt = "D"
	case 10, 11:
		stpt = "E"
	}
	return ehex.New().Set(stpt)
}

func sizeValid(eh ehex.Ehex) bool {
	switch eh.Code() {
	default:
		return false
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F":
	}
	return true
}

func reRollSize(d *dice.Dicepool) ehex.Ehex {
	size := d.Roll("2d6").Sum() - 2
	if size == 10 {
		size = d.Roll("1d6").Sum() + 9
	}
	return ehex.New().Set(size)
}

func atmoValid(eh ehex.Ehex) bool {
	switch eh.Code() {
	default:
		return false
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F":
	}
	return true
}

func reRollAtmo(d *dice.Dicepool, size ehex.Ehex) ehex.Ehex {
	atmo := d.Flux() + size.Value()
	switch {
	case size.Value() == 0:
		atmo = 0
	case atmo < 0:
		atmo = 0
	case atmo > 15:
		atmo = 15
	}
	return ehex.New().Set(atmo)
}

func hydrValid(eh ehex.Ehex) bool {
	switch eh.Code() {
	default:
		return false
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A":
	}
	return true
}

func reRollHydr(d *dice.Dicepool, size, atmo ehex.Ehex) ehex.Ehex {
	mod := 0
	if atmo.Value() < 2 || atmo.Value() > 9 {
		mod = -4
	}
	hydr := d.Flux() + atmo.Value() + mod
	switch {
	case size.Value() < 2:
		hydr = 0
	case hydr < 0:
		hydr = 0
	case hydr > 10:
		hydr = 10
	}
	return ehex.New().Set(hydr)
}

func popsValid(eh ehex.Ehex) bool {
	switch eh.Code() {
	default:
		return false
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F":
	}
	return true
}

func reRollPops(d *dice.Dicepool) ehex.Ehex {
	pops := d.Roll("2d6").Sum() - 2
	if pops == 10 {
		pops = d.Roll("2d6").Sum() + 3
	}
	return ehex.New().Set(pops)
}

func govrValid(eh ehex.Ehex) bool {
	switch eh.Code() {
	default:
		return false
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "U", "W", "X":
	}
	return true
}

func reRollGovr(d *dice.Dicepool, pops ehex.Ehex) ehex.Ehex {
	govr := d.Flux() + pops.Value()
	switch {
	case govr < 0:
		govr = 0
	case govr > 15:
		govr = 15
	}
	return ehex.New().Set(govr)
}

func lawsValid(eh ehex.Ehex) bool {
	switch eh.Code() {
	default:
		return false
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L", "S":
	}
	return true
}

func reRollLaws(d *dice.Dicepool, govr ehex.Ehex) ehex.Ehex {
	laws := d.Flux() + govr.Value()
	switch {
	case laws < 0:
		laws = 0
	case laws > 18:
		laws = 18
	}
	return ehex.New().Set(laws)
}

func tlValid(eh ehex.Ehex) bool {
	switch eh.Code() {
	default:
		return false
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L":
	}
	return true
}

func reRollTL(d *dice.Dicepool, hex []string) ehex.Ehex {
	tl := d.Roll("1d6").Sum()
	switch hex[0] {
	case "A":
		tl += 6
	case "B":
		tl += 4
	case "C":
		tl += 2
	case "X":
		tl -= 4
	}
	switch hex[1] {
	case "0", "1":
		tl += 2
	case "2", "3", "4":
		tl += 1
	}
	switch hex[2] {
	case "0", "1", "2", "3", "A", "B", "C", "D", "E", "F":
		tl += 1
	}
	switch hex[3] {
	case "9":
		tl += 1
	case "A":
		tl += 2
	}
	switch hex[4] {
	case "1", "2", "3", "4", "5":
		tl += 1
	case "9":
		tl += 2
	case "A":
		tl += 4
	}
	switch hex[5] {
	case "0", "5":
		tl += 1
	case "D":
		tl -= 2
	}
	if tl < 0 {
		tl = 0
	}
	return ehex.New().Set(tl)
}
