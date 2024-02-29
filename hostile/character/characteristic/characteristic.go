package characteristic

import "github.com/Galdoba/TravellerTools/pkg/dice"

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
