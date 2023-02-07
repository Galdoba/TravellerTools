package cascad

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/calendar"
	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	Created = iota
	Rest
	Eat
	Work
)

type pawn struct {
	Energy      int
	Food        int
	Money       int
	StartAction int64
	duration    int64
}

func NewPawn(e, f, m int) *pawn {
	return &pawn{e, f, m, 0, 0}
}

func (p *pawn) DoSomething(seed int64) error {
	if p.StartAction+p.duration+24 > seed {
		fmt.Println("Rejected: Buzy")
		fmt.Println("Expected to be free at", calendar.ImperialTimeStamp(p.StartAction+p.duration))
		return nil
	}
	if p.Energy < 0 {
		fmt.Println("Exusted: Forced Rest")
		p.Rest()
		return nil
	}
	if p.Food < 0 {
		return fmt.Errorf("pawn starved %v", p)
	}
	dice := dice.New().SetSeed(seed)
	for {
		r := dice.Sroll("1d3")
		dur := int64(dice.Sroll("1d6") * 60)
		p.duration = dur
		p.StartAction = seed
		switch r {
		case Rest:
			fmt.Println("Rest ")
			if p.Energy > 50 {
				fmt.Println("FAILED")
				continue
			}
			p.Rest()
			fmt.Println("SUCCESS")
		case Eat:
			fmt.Println("Eat  ")
			if p.Food > 50 {
				fmt.Println("FAILED")
				continue
			}
			if p.Money < 5 {
				fmt.Println("FAILED")
				continue
			}
			p.Eat()
			fmt.Println("SUCCESS")
		case Work:
			fmt.Println("Work ")
			if p.Energy < 5 {
				fmt.Println("FAILED")
				continue
			}
			if p.Food < 5 {
				fmt.Println("FAILED")
				continue
			}
			p.Work()
			fmt.Println("SUCCESS")
		}
		return nil
	}
}

func (p *pawn) Rest() {
	p.Energy += 5
	p.Food -= 1
}

func (p *pawn) Eat() {
	p.Energy -= 1
	p.Food += 30
	p.Money -= 5
}

func (p *pawn) Work() {
	p.Energy -= 2
	p.Food -= 3
	p.Money += 1
}
