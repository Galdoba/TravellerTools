package family

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/t4/core/task/pocketempire/calendar"
)

const (
	def = iota
	spouse
	child
)

type Family struct {
	FamilyName string
	Membrs     map[int]*Member
	dice       *dice.Dicepool
}

func New(name string) *Family {
	fm := Family{}
	fm.dice = dice.New()
	fm.Membrs = make(map[int]*Member)
	fm.FamilyName = name
	fm.Membrs[0] = Archont("0")
	return &fm
}

func (m *Member) String() string {
	upp := ""
	for _, c := range m.StatsCurrent {
		upp += ehex.New().Set(c).Code()
	}
	str := fmt.Sprintf("%v [%v]", m.Name, upp)
	if m.dead {
		str += " (dead)"
	}
	return str
}

func (fm *Family) AddMember(relative *Member, currentDate calendar.Date) {
	m := Member{}
	m.Name = fm.FamilyName + fmt.Sprintf(" %v", len(fm.Membrs))
	if relative.spouse == nil {
		m.spouse = relative
		m.Birthdate = calendar.Add(currentDate, fm.dice.Sroll("1d365")*-1, fm.dice.Flux()-18)
		m.generation = relative.generation
		m.rollUPP()
		relative.spouse = &m
		fm.Membrs[len(fm.Membrs)] = &m
		fmt.Println(relative, m)
		return
	}
	if len(relative.children) < 3 {
		m.Birthdate = currentDate
		m.generation = relative.generation + 1
		m.rollUPP()
		fm.Membrs[len(fm.Membrs)] = &m
		fmt.Println(relative, m)
		return
	}
}

func (fm *Family) Grow(currentDate calendar.Date) {
	for _, mem := range fm.Membrs {
		age := calendar.After(currentDate, mem.Birthdate)
		fmt.Println(mem, age.String())
		if age.Year() > 80 {
			fmt.Println("----------")
			if fm.dice.Sroll("3d6") == 18 {
				mem.dead = true
				mem.deathdate = currentDate
			}
		}
		if age.Year() < 21 {
			continue
		}
		if age.Year() < 41 {
			fmt.Println("+++++++++++")
			r := fm.dice.Sroll("3d6")
			if r == 18 {
				fm.AddMember(mem, currentDate)
			}
		}
	}
}

type Member struct {
	Name         string
	Birthdate    calendar.Date
	Birthworld   string
	Homeworld    string
	StatsCurrent []int
	StatsGenetic []int
	comments     string
	generation   int
	dead         bool
	deathdate    calendar.Date
	spouse       *Member
	father       *Member
	mother       *Member
	children     []*Member
}

func Archont(name string) *Member {
	m := Member{}
	m.generation = 0
	m.Birthdate = *calendar.New().Set(1, 1)
	m.rollUPP()
	return &m
}

func (mem *Member) rollUPP() {
	dp := dice.New()
	for i := 0; i < 6; i++ {
		r1 := dp.Sroll("1d6")
		rg := dp.Sroll("1d6")
		if i < 4 {
			mem.StatsGenetic = append(mem.StatsGenetic, rg)
		}
		val := r1 + rg
		if i == 5 && val < 8 {
			val = 8
		}
		mem.StatsCurrent = append(mem.StatsCurrent, val)
	}
}
