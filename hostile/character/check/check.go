package check

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/hostile/character/characteristic"
	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func Passed_Basic(dice *dice.Dicepool, tn int) bool {
	r := dice.Sroll("2d6")
	return r >= tn
}

func ReEnlist(dice *dice.Dicepool, tn int) (bool, bool) {
	r := dice.Sroll("2d6")
	passed := false
	locked := false
	passed = (r >= tn)
	locked = (r == 12)
	return passed, locked
}

func Char(code string, dice *dice.Dicepool, charSet *characteristic.CharSet) bool {
	charCode, tn, err := ParseCode(code)
	if err != nil {
		panic(err.Error())
	}
	if charCode == characteristic.BASIC {
		return Passed_Basic(dice, tn)
	}
	chr := charSet.Chars[charCode]
	dm := chr.Mod()
	r := dice.Sroll("2d6") + dm
	return r >= tn
}

func Char_WithSOC(code string, dice *dice.Dicepool, charSet *characteristic.CharSet) bool {
	charCode, tn, err := ParseCode(code)
	if err != nil {
		panic(err.Error())
	}
	if charCode == characteristic.BASIC {
		return Passed_Basic(dice, tn)
	}
	chr := charSet.Chars[charCode]
	dm := chr.Mod() + charSet.Chars[characteristic.SOC].Mod()
	r := dice.Sroll("2d6") + dm
	return r >= tn
}

func ParseCode(code string) (int, int, error) {
	charCode := 1
	data := strings.Fields(code)
	if len(data) != 2 {
		return charCode, -1, fmt.Errorf("can't parse code '%v'", code)
	}
	switch data[0] {
	default:
		return -1, -1, fmt.Errorf("can't parse code '%v'", code)
	case "STR":
		charCode = characteristic.STR
	case "DEX":
		charCode = characteristic.DEX
	case "END":
		charCode = characteristic.END
	case "INT":
		charCode = characteristic.INT
	case "EDU":
		charCode = characteristic.EDU
	case "SOC":
		charCode = characteristic.SOC
	case "Basic":
		charCode = characteristic.BASIC
	}
	data[1] = strings.TrimSuffix(data[1], "+")
	tn, err := strconv.Atoi(data[1])
	if err != nil {
		return charCode, tn, fmt.Errorf("can't parse code '%v': %v", code, err)
	}
	return charCode, tn, nil
}
