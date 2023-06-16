package career

import (
	"github.com/Galdoba/TravellerTools/internal/counter"
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

type rank struct {
	code           string
	title          string
	nextCode       string
	autoSkill      int
	require        string
	test           string
	tenureRequired bool
	promotion      []string
	commision      string
	publications   int
}

func setRankMap(career int) map[string]rank {
	rankByCode := make(map[string]rank)
	switch career {
	case Scholar:
		rankByCode["X"] = rank{
			code:      "X",
			title:     "Non-Traditional Scholar",
			autoSkill: skill.ID_NONE,
			require:   "C5=Tra",
			test:      "Edu 8+",
		}
		rankByCode["0"] = rank{
			code:      "0",
			title:     "Amateur",
			nextCode:  "1",
			autoSkill: skill.ID_NONE,
			test:      "Edu 8+",
		}
		rankByCode["1"] = rank{
			code:      "1",
			title:     "Lecturer <of Major>",
			nextCode:  "2",
			autoSkill: skill.ID_NONE,
			require:   "Edu 8+",
			test:      "automatic",
		}
		rankByCode["2"] = rank{
			code:      "2",
			title:     "Instructor <of Major>",
			nextCode:  "3",
			autoSkill: skill.ID_NONE,
		}
		rankByCode["3"] = rank{
			code:      "3",
			title:     "Assistant Professor <of Major>",
			nextCode:  "4",
			autoSkill: skill.ID_NONE,
		}
		rankByCode["4"] = rank{
			code:           "4",
			title:          "Assosiate Professor <of Major>",
			nextCode:       "5",
			autoSkill:      skill.ID_NONE,
			tenureRequired: true,
		}
		rankByCode["5"] = rank{
			code:           "5",
			title:          "Professor <of Major>",
			nextCode:       "6",
			autoSkill:      skill.ID_NONE,
			tenureRequired: true,
		}
		rankByCode["6"] = rank{
			code:           "6",
			title:          "Distinguished Professor <of Major>",
			autoSkill:      skill.ID_NONE,
			tenureRequired: true,
		}
		// for _, _ := range rankByCode {
		// 	//rankByCode[k].promotion = []string{""}
		// }

	}
	return rankByCode
}

type career struct {
	rank          string
	title         string
	rankTitleMap  map[string]rank
	toBeginTest   []beginTest
	controlChars  []string
	name          string
	descr         string
	ID            int
	forbidden     bool
	haveRetry     bool
	waiverAllowed bool
	musteredOut   bool
	tenure        bool
}

type beginTest struct {
	rank    string
	require string
}

type careerHistory struct {
	career map[int]*career
}

func newCareer(i int) *career {
	c := career{}
	c.ID = i
	c.name = careerName(i)
	c.haveRetry = true
	switch i {
	case Craftsman:
		c.toBeginTest = append(c.toBeginTest, beginTest{"", "automatic"})
	case Scholar:
		c.toBeginTest = append(c.toBeginTest, beginTest{"", "automatic"})
		c.haveRetry = false
		c.waiverAllowed = true
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
	Waiver() counter.Counter
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
			if ch.career[check].forbidden {
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
