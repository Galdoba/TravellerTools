package family

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/calendar"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
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

func Create(name string, foundationDate calendar.Date) *Family {
	fm := Family{}
	fm.dice = dice.New().SetSeed(name + foundationDate.String())
	fm.Membrs = make(map[int]*Member)
	fm.FamilyName = name
	fm.Membrs[0] = &Member{
		Name:       "Adam",
		Birthworld: "[Unknown]",
		Homeworld:  "[Unset]",
		FamilyName: name,
		birthDate:  foundationDate,
		deathDate:  calendar.Date{},
		Age:        20,
		generation: 0,
		dead:       false,
		spouse:     nil,
		father:     nil,
		mother:     nil,
		children:   nil,
	}
	fm.Membrs[0].rollUPP(fm.dice)

	return &fm
}

func (m *Member) String() string {
	upp := ""
	for _, c := range m.StatsCurrent {
		upp += ehex.New().Set(c).Code()
	}
	str := fmt.Sprintf("%v [%v] %v", m.Name+" "+m.FamilyName, upp, m.Age)
	if m.dead {
		str += " (dead)"
	}
	return str
}

func (fm *Family) AddMember(relative *Member, currentDate calendar.Date, newMemberStatus int) {
	m := Member{}
	m.Name = fm.FamilyName + fmt.Sprintf(" %v", len(fm.Membrs))
	m.birthDate = currentDate
	switch newMemberStatus {
	case spouse:
		m.spouse = relative
		m.generation = relative.generation
		m.rollUPP(fm.dice)
		relative.spouse = &m
		fm.Membrs[len(fm.Membrs)] = &m
	case child:
		m.generation = relative.generation + 1
		m.rollUPPbyGenetics(fm.dice, relative.StatsGenetic, relative.spouse.StatsGenetic)
		fm.Membrs[len(fm.Membrs)] = &m

	}

}

func (m *Member) AgeNow(now calendar.Date) int {
	switch m.dead {
	case true:
	case false:
	}
	return -1
}

func (fm *Family) Grow(currentDate calendar.Date) {
	for _, mem := range fm.Membrs {
		if mem.dead {
			panic(mem)
		}
		age := mem.AgeNow(currentDate)
		if age > 80 {
			//fmt.Println("----------")
			//CHECK FOR DEATH
			fmt.Println("CHECK FOR DEATH")
		}
		// if age.Year() < 21 {
		// 	continue
		// }
		// if age.Year() < 41 {
		// 	//fmt.Println("+++++++++++")
		// 	r := fm.dice.Sroll("4d6")
		// 	if r == 4 {
		// 		fm.AddMember(mem, currentDate)
		// 	}
		// }
	}
}

type Member struct {
	Name         string
	Birthworld   string
	Homeworld    string
	FamilyName   string
	StatsCurrent []int
	StatsGenetic []int
	birthDate    calendar.Date
	deathDate    calendar.Date
	Age          int
	comments     string
	generation   int
	dead         bool
	spouse       *Member
	father       *Member
	mother       *Member
	children     []*Member
}

func Archont(name string, dice *dice.Dicepool) *Member {
	m := Member{}
	m.generation = 0
	m.rollUPP(dice)
	m.Name = name
	return &m
}

func (mem *Member) rollUPP(dp *dice.Dicepool) {
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

func (mem *Member) rollUPPbyGenetics(dp *dice.Dicepool, gen1, gen2 []int) {
	for i := 0; i < 6; i++ {
		switch i {
		case 0, 1, 2, 3:
			r := dp.Sroll("1d2")
			add := 0
			switch r {
			case 1:
				add = gen1[i]
			case 2:
				add = gen2[i]
			}
			mem.StatsGenetic = append(mem.StatsGenetic, add)
			mem.StatsCurrent = append(mem.StatsCurrent, add+dp.Sroll("1d6"))
		case 4, 5:
			mem.StatsCurrent = append(mem.StatsCurrent, dp.Sroll("2d6"))
		}

	}
}

func (m *Member) GeneticsCard() string {
	s := "--------------------------------------------------------------------------------\n"
	s += fmt.Sprintf("Family Member Name: %v\n", m.Name)
	s += fmt.Sprintf("Birthdate         : %v\n", m.birthDate.String())
	s += fmt.Sprintf("Birthworld        : %v\n", m.Birthworld)
	s += fmt.Sprintf("UPP (Current)     : %v%v%v%v%v%v\n", ehex.ToCode(m.StatsCurrent[0]), ehex.ToCode(m.StatsCurrent[1]), ehex.ToCode(m.StatsCurrent[2]), ehex.ToCode(m.StatsCurrent[3]), ehex.ToCode(m.StatsCurrent[4]), ehex.ToCode(m.StatsCurrent[5]))
	s += fmt.Sprintf("UPP (Genetic)     : %v%v%v%vxx\n", ehex.ToCode(m.StatsGenetic[0]), ehex.ToCode(m.StatsGenetic[1]), ehex.ToCode(m.StatsGenetic[2]), ehex.ToCode(m.StatsGenetic[3]))
	s += "--------------------------------------------------------------------------------"
	return s
}
