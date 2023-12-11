package uwp

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

const (
	SET_TO_NONE = "set to 0"
)

func HostileUWP(data map[string]ehex.Ehex) string {
	s := data["port"].Code()
	s += data["size"].Code()
	s += data["atmo"].Code()
	s += data["hydr"].Code()
	s += data["pops"].Code()
	s += data["govr"].Code()
	s += data["laws"].Code()
	s += "-" + data["tl"].Code()
	return s
}

func GenerateMainWorldUWP(dice *dice.Dicepool, ruleset map[string]string) map[string]ehex.Ehex {
	data := make(map[string]ehex.Ehex)
	data["size"] = hostileSize(dice, ruleset, data)
	data["atmo"] = hostileAtmo(dice, ruleset, data)
	data["hydr"] = hostileHydr(dice, ruleset, data)
	data["pops"] = hostilePops(dice, ruleset, data)
	data["govr"] = hostileGovr(dice, ruleset, data)
	data["laws"] = hostileLaws(dice, ruleset, data)
	data["tl"] = hostileTL(dice, ruleset, data)
	data["port"] = hostilePort(dice, ruleset, data)
	return data
}

func hostileSize(dice *dice.Dicepool, ruleset map[string]string, data map[string]ehex.Ehex) ehex.Ehex {
	r1 := dice.Sroll("1d6")
	if r1 != 6 {
		return ehex.New().Set(r1)
	}
	r2 := 0
	switch ruleset["super-earth allowed"] {
	default:
		r2 = dice.Sroll("2d6-2")
	case "yes":
		r2 = dice.Sroll("2d6+4")
	}
	return ehex.New().Set(r2)
}

func hostileAtmo(dice *dice.Dicepool, ruleset map[string]string, data map[string]ehex.Ehex) ehex.Ehex {
	if ruleset["atmo"] == SET_TO_NONE {
		return ehex.New().Set(0)
	}
	dm := 0
	switch data["size"].Value() {
	case 0, 1, 2:
		dm = -99
	case 3, 4:
		dm = -5
	case 5, 6, 7:
		dm = 0
	default:
		dm = 2
	}
	r1 := dice.Sroll("2d6") + dm
	if r1 < 0 {
		r1 = 0
	}
	if r1 > 12 {
		r1 = 12
	}
	return ehex.New().Set(r1)
}

func hostileHydr(dice *dice.Dicepool, ruleset map[string]string, data map[string]ehex.Ehex) ehex.Ehex {
	dm := 0
	switch data["size"].Value() {
	case 0, 1, 2:
		return ehex.New().Set(0)
	}
	switch data["atmo"].Value() {
	case 0, 1:
		return ehex.New().Set(0)
	case 2, 3, 11, 12:
		dm = -6
	}
	r1 := dice.Sroll("2d6") + dm
	if r1 < 0 {
		r1 = 0
	}
	if r1 > 10 {
		r1 = 10
	}
	return ehex.New().Set(r1)
}

func hostilePops(dice *dice.Dicepool, ruleset map[string]string, data map[string]ehex.Ehex) ehex.Ehex {
	dm := 0
	if data["size"].Value() <= 2 {
		dm = dm - 1
	}
	if data["size"].Value() >= 10 {
		dm = dm - 1
	}
	if data["atmo"].Value() >= 11 {
		dm = dm - 1
	}
	if data["atmo"].Value() != 5 && data["atmo"].Value() != 6 && data["atmo"].Value() != 8 {
		dm = dm - 1
	}
	if data["atmo"].Value() == 5 && data["atmo"].Value() == 6 && data["atmo"].Value() == 8 {
		dm = dm + 1
	}

	r1 := dice.Sroll("2d6") + dm - 5
	if r1 < 0 {
		r1 = 0
	}
	if r1 > 8 {
		r1 = 8
	}
	return ehex.New().Set(r1)
}

func hostileGovr(dice *dice.Dicepool, ruleset map[string]string, data map[string]ehex.Ehex) ehex.Ehex {
	if data["pops"].Value() == 0 {
		return ehex.New().Set(0)
	}
	dm := data["pops"].Value()
	r1 := dice.Sroll("2d6") + dm - 7
	if r1 < 0 {
		r1 = 0
	}
	if r1 > 13 {
		r1 = 13
	}
	return ehex.New().Set(r1)
}

func hostileLaws(dice *dice.Dicepool, ruleset map[string]string, data map[string]ehex.Ehex) ehex.Ehex {
	if data["pops"].Value() == 0 {
		return ehex.New().Set(0)
	}
	dm := data["govr"].Value()
	r1 := dice.Sroll("2d6") + dm - 7
	if r1 < 0 {
		r1 = 0
	}
	return ehex.New().Set(r1)
}

func hostileTL(dice *dice.Dicepool, ruleset map[string]string, data map[string]ehex.Ehex) ehex.Ehex {
	if data["pops"].Value() == 0 {
		return ehex.New().Set(0)
	}
	return ehex.New().Set(12)
}

func hostilePort(dice *dice.Dicepool, ruleset map[string]string, data map[string]ehex.Ehex) ehex.Ehex {
	r1 := dice.Sroll("1d3")
	if data["pops"].Value() == 0 {
		return ehex.New().Set("X")
	}
	val := r1 + data["pops"].Value()
	switch val {
	case 1, 2, 3, 4, 5:
		return ehex.New().Set("E")
	case 6, 7:
		return ehex.New().Set("D")
	case 8, 9:
		return ehex.New().Set("C")
	case 10, 11:
		return ehex.New().Set("B")
	default:
		panic("этого не должно быть: val=" + fmt.Sprintf("%v + %v", r1, data["pops"].Value()))
	}
}
