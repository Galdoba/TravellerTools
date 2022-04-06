package balancingact

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/internal/dice"
)

type pawn struct {
	name         string
	chars        map[int]int
	skills       map[int]int
	position     int
	organisation int
	subordinates []*pawn
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
		p.skills[skillList[toUp]]++
	}
	return &p, nil
}
