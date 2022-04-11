package balancingact

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/internal/dice"
)

type pawn struct {
	ID         int
	name       string
	chars      map[int]int
	skills     map[int]int
	position   int // leader or agent
	location   *planet
	activeTask string //состояние активного задания
}

func (p *pawn) String() string {
	str := "Name: " + p.name + "\n"
	str += fmt.Sprintf("WIL %v  INT %v  EDU %v  CHR %v\n", p.chars[WIL], p.chars[INT], p.chars[EDU], p.chars[CHR])

	str += fmt.Sprintf("Administration %v\n", p.skills[Administration])
	str += fmt.Sprintf("CovertOps %v\n", p.skills[CovertOps])
	str += fmt.Sprintf("Economics %v\n", p.skills[Economics])
	str += fmt.Sprintf("Politics %v\n", p.skills[Politics])
	str += fmt.Sprintf("Military %v", p.skills[Military])
	return str
}

func createPawn(name string, pawnType int) (*pawn, error) {
	p := pawn{}
	p.name = name
	switch pawnType {
	default:
		return nil, fmt.Errorf("pawn type unknown (%v)", pawnType)
	case Leader, Agent:
		p.position = pawnType
	}
	roller := dice.New()
	p.chars = make(map[int]int)
	for _, chr := range []int{WIL, INT, EDU, CHR} {
		p.chars[chr] = roller.Roll("2d6").Sum()
	}
	skillPoints := 0
	switch p.position {
	case Leader:
		skillPoints = roller.Roll("2d6").Sum() / 2
	case Agent:
		skillPoints = roller.Roll("1d6").Sum() + 1
	}
	p.skills = make(map[int]int)
	skillList := []int{Administration, CovertOps, Economics, Politics, Military}
	for _, skill := range []int{Administration, CovertOps, Economics, Politics, Military} {
		p.skills[skill] = 0
	}
	for s := 0; s < skillPoints; s++ {
		toUp := roller.Roll("1d5").DM(-1).Sum()
		if p.skills[skillList[toUp]] >= 5 {
			s--
			continue
		}
		p.skills[skillList[toUp]]++
	}
	return &p, nil
}
