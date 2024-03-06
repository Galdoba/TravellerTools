package characteristic

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

const (
	Unknown = iota
	STR
	DEX
	END
	INT
	EDU
	SOC
	max = 100
)

type CharSet struct {
	chars map[int]int
}

type Char struct {
	Name    string //STR/DEX/END/INT/EDU/SOC
	Type    string //Physical/Mental/Social
	Maximum ehex.Ehex
	Current int
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
	case EDU:
		chr.Name = "EDU"
		chr.Type = "Mental"
	case SOC:
		chr.Name = "SOC"
		chr.Type = "Social"
	}
	return &chr, nil
}

func (chr *Char) Roll(dice *dice.Dicepool) {
	r := dice.Sroll("2d6")
	chr.Maximum = ehex.New().Set(r)
	chr.Current = r
}

func (chr *Char) ChangeMaximumBy(i int) {
	v := chr.Maximum.Value() + i
	if v > 15 {
		v = 15
	}
	if v < 0 {
		v = 0
	}
	chr.Maximum = ehex.New().Set(v)
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

func (cs *CharSet) Strength() int {
	return cs.chars[STR]
}

func (cs *CharSet) Dexterity() int {
	return cs.chars[DEX]
}

func (cs *CharSet) Endurance() int {
	return cs.chars[END]
}

func (cs *CharSet) Inteligence() int {
	return cs.chars[INT]
}

func (cs *CharSet) Education() int {
	return cs.chars[EDU]
}

func (cs *CharSet) Social() int {
	return cs.chars[SOC]
}

func NewCharSet(dice *dice.Dicepool) *CharSet {
	chrSet := CharSet{}
	chrSet.chars = make(map[int]int)
	for i := STR; i <= SOC; i++ {
		chrSet.chars[i] = dice.Sroll("2d6")
	}
	return &chrSet
}

func (cs *CharSet) Mod(i int) int {
	switch i {
	default:
		return -99
	case STR, DEX, END, INT, EDU, SOC:
		return (cs.chars[i] / 3) - 2
	}
	return -3
}

func (cs *CharSet) SetupMax() {
	for i := STR; i <= SOC; i++ {
		cs.chars[max+i] = cs.chars[i]
	}
}
