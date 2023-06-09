package career

import (
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/t5/pawn/skill"
)

const (
	NoCareer = iota
	Craftsman
	Scholar
	Entertainer
	Citizen
	Scout
	Merchant
	Spacer
	Solder
	Rogue
	Agent
	Noble
	Marine
	Functionary
)

type career struct {
	rank         string
	title        string
	toBeginTest  []beginTest
	controlChars []string
	name         string
	descr        string
	ID           int
	forbidden    bool
	haveRetry    bool
}

type beginTest struct {
	rank    string
	require string
}

type careerHistory struct {
	career []*career
}

func newCareer(i int) *career {
	c := career{}
	c.ID = i
	c.name = careerName(i)
	c.haveRetry = true
	switch i {
	case Scholar:
		c.haveRetry = false
	case Entertainer:
		c.haveRetry = false

	}
	return &c
}

func careerName(i int) string {
	switch i {
	case Craftsman:
		return "Craftsman"
	case Scholar:
		return "Scholar"
	case Entertainer:
		return "Entertainer"
	case Citizen:
		return "Citizen"
	case Scout:
		return "Scout"
	case Merchant:
		return "Merchant"
	case Spacer:
		return "Spacer"
	case Solder:
		return "Solder"
	case Rogue:
		return "Rogue"
	case Agent:
		return "Agent"
	case Noble:
		return "Noble"
	case Marine:
		return "Marine"
	case Functionary:
		return "Functionary"
	}
	return "UNDEFINED"
}

type Worker interface {
	Profile() profile.Profile
}

func careerOptions(wk Worker, ch *careerHistory) []int {
	options := []int{}
	prf := wk.Profile()
	for _, check := range []int{Craftsman, Scholar, Entertainer, Citizen, Scout, Merchant, Spacer, Solder, Rogue, Agent, Noble, Marine, Functionary} {
		switch check {
		case Craftsman:
			if len(ch.career) == 0 {
				continue
			}
			if craftSkl := prf.Data(skill.NameByID(Craftsman)); craftSkl != nil {
				if craftSkl.Value() < 1 {
					continue
				}
			}
			foundSkills := []int{}
			for _, sklID := range skill.AllID() {
				if skl := prf.Data(skill.NameByID(sklID)); skl != nil {
					if skl.Value() > 5 {
						foundSkills = append(foundSkills, sklID)
					}
				}
			}
			if len(foundSkills) > 1 {
				options = append(options, check)
			}
		case Scholar:

		}
	}
	return options
}
