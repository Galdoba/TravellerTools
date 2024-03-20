package characteristic

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/decidion"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

const (
	STR = iota
	DEX
	END
	INT
	EDU
	SOC
	INST
	BASIC
	max = 100
)

type CharSet struct {
	Chars map[int]*Char
}

type Char struct {
	Name       string //STR/DEX/END/INT/EDU/SOC
	Type       string //Physical/Mental/Social
	Maximum    ehex.Ehex
	Current    int
	RollCode   string
	UpperLimit int
}

func New(code int) (*Char, error) {
	chr := Char{}
	switch code {
	default:
		return nil, fmt.Errorf("can't create char with code %v", code)
	case STR:
		chr.Name = "STR"
		chr.Type = "Physical"
	case DEX:
		chr.Name = "DEX"
		chr.Type = "Physical"
	case END:
		chr.Name = "END"
		chr.Type = "Physical"
	case INT:
		chr.Name = "INT"
		chr.Type = "Mental"
	case INST:
		chr.Name = "INST"
		chr.Type = "Mental"
	case EDU:
		chr.Name = "EDU"
		chr.Type = "Mental"
	case SOC:
		chr.Name = "SOC"
		chr.Type = "Social"
	}
	chr.RollCode = "2d6"
	chr.UpperLimit = 15
	return &chr, nil

}

func (chr *Char) SetRollCode(code string) {
	chr.RollCode = code
}

func (chr *Char) SetUpperLimit(ul int) {
	chr.UpperLimit = ul
}
func (chr *Char) Roll(dice *dice.Dicepool) {
	r := dice.Sroll(chr.RollCode)
	chr.Maximum = ehex.New().Set(r)
	chr.Current = r
}

func (chr *Char) ChangeMaximumBy(i int) {
	v := chr.Maximum.Value() + i
	if v > chr.UpperLimit {
		v = chr.UpperLimit
	}
	if v < 0 {
		v = 0
	}
	chr.Maximum = ehex.New().Set(v)
	chr.Current = chr.Maximum.Value()
}

func (chr *Char) Damage(i int) int {
	chr.Current -= i
	if chr.Current < 0 {
		v := -1 * chr.Current
		chr.Current = 0
		return v
	}
	return 0
}

func (chr *Char) Heal(i int) {
	v := chr.Current + i
	if v > chr.Maximum.Value() {
		v = chr.Maximum.Value()
	}
	chr.Current = v
}

func (c *Char) Mod() int {
	return charMod(c.Current)
}

func charMod(i int) int {
	switch i {
	default:
		if i <= 2 {
			return -2
		}
		if i >= 15 {
			return 3
		}
	case 3, 4, 5:
		return -1
	case 6, 7, 8:
		return 0
	case 9, 10, 11:
		return 1
	case 12, 13, 14:
		return 2
	}
	return -999
}

// func (cs *CharSet) Strength() int {
// 	return cs.chars[STR]
// }

// func (cs *CharSet) Dexterity() int {
// 	return cs.chars[DEX]
// }

// func (cs *CharSet) Endurance() int {
// 	return cs.chars[END]
// }

// func (cs *CharSet) Inteligence() int {
// 	return cs.chars[INT]
// }

// func (cs *CharSet) Education() int {
// 	return cs.chars[EDU]
// }

// func (cs *CharSet) Social() int {
// 	return cs.chars[SOC]
// }

func Human() []int {
	return []int{STR, DEX, END, INT, EDU, SOC}
}

func NewCharSet(charCodes ...int) (*CharSet, error) {
	chrSet := CharSet{}
	chrSet.Chars = make(map[int]*Char)
	// chrSet.rollCodes = make(map[int]string)
	for _, code := range charCodes {
		chr, err := New(code)
		if err != nil {
			return nil, err
		}
		chrSet.Chars[code] = chr
		// chrSet.chars[code] = New()
		// chrSet.rollCodes[code] = "2d6"
	}
	return &chrSet, nil
}

func (cs *CharSet) Roll(dice *dice.Dicepool) error {
	for i := STR; i <= INST; i++ {
		if _, ok := cs.Chars[i]; ok {
			cs.Chars[i].Roll(dice)
		}
	}
	return nil
}

func (cs *CharSet) Mod(i int) int {
	switch i {
	default:
		return -99
	case STR, DEX, END, INT, EDU, SOC:
		return (cs.Chars[i].Current / 3) - 2
	}
	return -3
}

func (cs *CharSet) SetupMax() {
	for i := STR; i <= SOC; i++ {
		cs.Chars[max+i] = cs.Chars[i]
	}
}

func (cs *CharSet) String() string {
	str := ""

	for i := STR; i <= INST; i++ {
		if chr, ok := cs.Chars[i]; ok {
			str += ehex.New().Set(chr.Current).Code()
		}
	}
	return str
}

func FromText(text string) (int, int) {
	chrID := -1
	chrVal := 0
	data := strings.Split(text, " ")
	for _, d := range data {
		switch d {
		case "STR":
			chrID = STR
		case "DEX":
			chrID = DEX
		case "END":
			chrID = END
		case "INT":
			chrID = INT
		case "EDU":
			chrID = EDU
		case "SOC":
			chrID = SOC
		case "INST":
			chrID = INST
		default:
			v, err := strconv.Atoi(d)
			if err == nil {
				chrVal = v
			}
		}
	}
	return chrID, chrVal
}

func (cs *CharSet) AgingRoll(dice *dice.Dicepool, dm int, manual bool) (string, error) {
	r := dice.Sroll("2d6") - dm
	if r > 1 {
		r = 1
	}
	if r < -6 {
		r = -6
	}
	msg := "Aging:"
	vals := []int{}
	switch r {
	case -6:
		vals = []int{-2, -2, -2, -1}
	case -5:
		vals = []int{-2, -2, -2}
	case -4:
		vals = []int{-2, -2, -1}
	case -3:
		vals = []int{-2, -1, -1}
	case -2:
		vals = []int{-1, -1, -1}
	case -1:
		vals = []int{-1, -1}
	case 0:
		vals = []int{-1}
	case 1:
		vals = []int{}
		msg += " No Effect"
		return msg, nil
	}
	charCodes := pickCharacteristic(len(vals), dice, manual)
	for i, chCode := range charCodes {
		cs.Chars[chCode].ChangeMaximumBy(vals[i])
		msg += fmt.Sprintf(" %v reduced by %v,", cs.Chars[chCode].Name, vals[i]*-1)
		if cs.Chars[i].Current == 0 {
			msg = "Character died of old age."
			return "", fmt.Errorf("Character died of old age")
		}
	}
	msg = strings.TrimSuffix(msg, ",") + "."
	return msg, nil
}

func pickCharacteristic(n int, dice *dice.Dicepool, manual bool) []int {
	options := []string{"STR", "DEX", "END"}
	picked := ""
	result := []int{}
	for i := 0; i < n; i++ {
		if i == 4 {
			options = []string{"INT", "EDU"}
		}
		switch manual {
		case false:
			picked, options = decidion.Random_One_Exclude(dice, options...)
		case true:
			panic("not implemented")
		}
		switch picked {
		case "STR":
			result = append(result, STR)
		case "DEX":
			result = append(result, DEX)
		case "END":
			result = append(result, END)
		case "INT":
			result = append(result, INT)
		case "EDU":
			result = append(result, EDU)
		}
	}
	return result
}
