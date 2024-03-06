package career

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	Android            = "0"
	CorporateAgent     = "1"
	CorporateExec      = "2"
	Colonist           = "3"
	CommersialSpacer   = "4"
	Marine             = "5"
	Marshal            = "6"
	MilitarySpacer     = "7"
	Physician          = "8"
	Ranger             = "9"
	Rogue              = "A"
	Roughneck          = "B"
	Scientist          = "C"
	SurveyScout        = "D"
	Technitian         = "E"
	CommisionPassed    = true
	CommisionNotPassed = false
)

type Career struct {
	Name            string
	code            string
	CommissionState bool
	Rank            int
	TermsCompleted  int
}

func New(code string) (*Career, error) {
	cr := Career{}
	switch code {
	default:
		return nil, fmt.Errorf("can't create career: unknown code '%v'", code)
	case Android:
		return nil, fmt.Errorf("Android career is not implemented")
	}
	return &cr, nil
}

type CareerStats struct {
	Name          string              `json:"Career"`
	Qualification string              `json:"Qualification"`
	Survival      string              `json:"Survival"`
	Commision     string              `json:"Commision"`
	Advance       string              `json:"Advance"`
	ReEnlist      int                 `json:"Re-Enlist"`
	Ranks         []Rank              `json:"Career Ranks"`
	Branch        string              `json:"Career Branch,omitempty"`
	SkillTable    map[string][]string `json:"Skill Tables,omitempty"`
	MusterOut     []string            `json:"Muster Out Benefits"`
}

func init() {
	sklTab := make(map[string][]string)
	sklTab["1"] = []string{"1", "2", "3", "4", "5", "6"}
	sklTab["2"] = []string{"1", "2", "3", "4", "5", "6"}
	cr := CareerStats{
		Name:          "name",
		Qualification: "STR 6+",
		Survival:      "DEX 4+",
		Commision:     "END 5+",
		Advance:       "END 6+",
		ReEnlist:      5,
		Ranks: []Rank{
			Rank{
				Value:             0,
				CommisionRequired: false,
				Position:          "initial",
				AutoSkill:         "skill 1",
			},
			Rank{
				Value:             1,
				CommisionRequired: false,
				Position:          "NCO 1",
				AutoSkill:         "",
			},
			Rank{
				Value:             1,
				CommisionRequired: true,
				Position:          "leitenent",
				AutoSkill:         "pilot 1",
			},
		},
		MusterOut: []string{
			"one",
			"two",
			"three",
			"four",
			"five",
			"six",
			"seven",
		},
		Branch:     "branch",
		SkillTable: sklTab,
	}
	bt, err := json.MarshalIndent(&cr, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	f, err := os.OpenFile(`/home/galdoba/TabletopGames/HOSTILE/careers/template.json`, os.O_CREATE|os.O_WRONLY, 0770)
	f.Write(bt)
}

type Rank struct {
	Value             int    `json:"Rank"`
	CommisionRequired bool   `json:"Commision Required"`
	Position          string `json:"Position"`
	AutoSkill         string `json:"Automatic Skill,omitempty"`
}

func GetCareer(name string) CareerStats {
	switch name {
	default:
		return CareerStats{}

	}
}
