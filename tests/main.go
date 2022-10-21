package main

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/dice/probability"
)

const (
	Solidarity = iota
	Wealth
	Expansion
	Might
	Administration
	CovertOps
	Economics
	Politics
	Military
	WIL
	INT
	EDU
	CHR
	Leader
	Agent
	TimeFactor_Day
	TimeFactor_Week
	TimeFactor_Month
	TimeFactor_Year
	TASK_Changing_the_Law_Level_of_a_World
	TASK_Changing_the_Might__of_a_world_which_you_do_not_rule
)

func main() {
	l1, _ := createPawn("Abbada Jeremy", Leader)
	fmt.Println(l1)
	fmt.Println("Evaluating: Changing_the_Law_Level_of_a_World")
	tw, err := TaskWeight(TASK_Changing_the_Law_Level_of_a_World, l1)
	fmt.Println(tw, err)
	fmt.Println("Evaluating: Changing_the_Might__of_a_world_which_you_do_not_rule")
	tw2, err2 := TaskWeight(TASK_Changing_the_Might__of_a_world_which_you_do_not_rule, l1)
	fmt.Println(tw2, err2)
}

type Pawn struct {
	name     string
	chars    map[int]int
	skills   map[int]int
	position string
}

func createPawn(name string, pawnType int) (*Pawn, error) {
	p := Pawn{}
	skillPointDice := ""
	skillPointDiceMod := 0
	switch pawnType {
	default:
		return nil, fmt.Errorf("pawn type unreconised")
	case Leader:
		skillPointDice = "2d6"
		p.position = "Leader"
	case Agent:
		skillPointDice = "1d6"
		skillPointDiceMod = 1
		p.position = "Agent"
	}
	p.name = name
	p.chars = make(map[int]int)
	p.skills = make(map[int]int)
	charList := []int{WIL, INT, EDU, CHR}
	skillList := []int{Administration, CovertOps, Economics, Politics, Military}
	dp := dice.New()
	skillPoints := dp.Roll(skillPointDice).DM(skillPointDiceMod).Sum()
	for _, chr := range charList {
		p.chars[chr] = dp.Roll("2d6").Sum()
	}
	for s := 0; s < skillPoints; s++ {
		skl := dp.Roll("1d5").DM(-1).Sum()
		p.skills[skillList[skl]] = p.skills[skillList[skl]] + 1
		if p.skills[skillList[skl]] > 5 {
			p.skills[skillList[skl]] = p.skills[skillList[skl]] - 1
			s--
		}
	}
	return &p, nil
}

func (p *Pawn) String() string {
	r := ""
	r += fmt.Sprintf("%v: %v\n", p.position, p.name)
	r += fmt.Sprintf("WIL %v  INT %v  EDU %v  CHR %v\n", p.chars[WIL], p.chars[INT], p.chars[EDU], p.chars[CHR])
	r += fmt.Sprintf("Administration: %v\n", p.skills[Administration])
	r += fmt.Sprintf("Covert Ops    : %v\n", p.skills[CovertOps])
	r += fmt.Sprintf("Economics     : %v\n", p.skills[Economics])
	r += fmt.Sprintf("Politics      : %v\n", p.skills[Politics])
	r += fmt.Sprintf("Military      : %v", p.skills[Military])
	return r
}

func (p *Pawn) TotalDM(skill, char int) int {
	dm := 0
	switch p.chars[char] {
	case 0:
		dm = dm - 3
	case 1, 2:
		dm = dm - 2
	default:
		dm = p.chars[char]/3 - 2
	}
	switch p.skills[skill] {
	case 0:
		dm = dm - 3
	default:
		dm = dm + p.skills[skill]
	}
	return dm
}

func TaskWeight(task int, p *Pawn) (float64, error) {
	skill, char, diff, _ := taskData(task)
	dm := p.TotalDM(skill, char) + diff
	operand := "+"
	if dm < 0 {
		operand = ""
	}
	code := fmt.Sprintf("2d6%v%v", operand, dm)
	prob, err := probability.Calculate(probability.RESULT_IS_SAME_OR_MORE, code, 8)
	return prob, err
}

func taskData(task int) (int, int, int, int) {
	switch task {
	case TASK_Changing_the_Law_Level_of_a_World:
		return Politics, CHR, -2, TimeFactor_Day
	case TASK_Changing_the_Might__of_a_world_which_you_do_not_rule:
		return Politics, CHR, -4, TimeFactor_Month
	}
	return 0, 0, 0, 0
}
