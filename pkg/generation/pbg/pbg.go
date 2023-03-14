package pbg

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

type PBG struct {
	popDigit ehex.Ehex
	belts    ehex.Ehex
	ggigs    ehex.Ehex
}

func New(dice *dice.Dicepool) PBG {
	p := PBG{}
	p.popDigit = ehex.New().Set(-1)
	b := dice.Sroll("1d6-3")
	if dice.Sroll("2d6") == 12 {
		b++
	}
	if b < 0 {
		b = 0
	}

	p.belts = ehex.New().Set(b)
	g := dice.Sroll("2d6")/2 - 2
	if g < 0 {
		g = 0
	}
	p.ggigs = ehex.New().Set(g)
	return p
}

func (p *PBG) String() string {
	return p.popDigit.Code() + p.belts.Code() + p.ggigs.Code()
}

func (p *PBG) GasGigants() ehex.Ehex {
	return p.ggigs
}
func (p *PBG) Belts() ehex.Ehex {
	return p.belts
}
