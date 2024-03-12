package career

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Galdoba/TravellerTools/hostile/character/skill"
)

const (
	Android                     = "0"
	CorporateAgent              = "1"
	CorporateExec               = "2"
	Colonist                    = "3"
	CommersialSpacer            = "4"
	Marine                      = "5"
	Marshal                     = "6"
	MilitarySpacer              = "7"
	Physician                   = "8"
	Ranger                      = "9"
	Rogue                       = "A"
	Roughneck                   = "B"
	Scientist                   = "C"
	SurveyScout                 = "D"
	Technitian                  = "E"
	CommisionPassed             = true
	CommisionNotPassed          = false
	Benefit_1STR                = "+1 STR"
	Benefit_1DEX                = "+1 DEX"
	Benefit_1END                = "+1 END"
	Benefit_1INT                = "+1 INT"
	Benefit_1EDU                = "+1 EDU"
	Benefit_1SOC                = "+1 SOC"
	Benefit_EliteTicket         = "Elite Ticket"
	Benefit_StarEnvoyClubMember = "Star Envoy Club Member"
	Benefit_TicketStandard      = "Standard Ticket"
	Benefit_TicketElite         = "Elite Ticket"
	Benefit_Weapon              = "Weapon"
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
	Name             string              `json:"Career"`
	Qualification    string              `json:"Qualification,omitempty"`
	QualificationReq string              `json:"Qualification Requirement,omitempty"`
	Survival         string              `json:"Survival"`
	Commision        string              `json:"Commision"`
	Advance          string              `json:"Advance"`
	ReEnlist         int                 `json:"Re-Enlist"`
	Ranks            []Rank              `json:"Career Ranks"`
	Branch           string              `json:"Career Branch,omitempty"`
	SkillTable       map[string][]string `json:"Skill Tables,omitempty"`
	MusterOut        []string            `json:"Muster Out Benefits"`
}

func init() {
	sklTab := make(map[string][]string)
	sklTab["Personal Development"] = []string{"______", "______", "______", "______", "______", "______"}
	sklTab["Service Skills"] = []string{"______", "______", "______", "______", "______", "______"}
	sklTab["Specialist Skills"] = []string{"______", "______", "______", "______", "______", "______"}
	sklTab["Advanced Education Skills"] = []string{"______", "______", "______", "______", "______", "______"}
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
				Position:          "-----------",
				AutoSkill:         "___________",
			},
			Rank{
				Value:             1,
				CommisionRequired: false,
				Position:          "-----------",
				AutoSkill:         "___________",
			},
			Rank{
				Value:             1,
				CommisionRequired: true,
				Position:          "leitenent",
				AutoSkill:         skill.SkillStr(skill.Pilot) + " 1",
			},
		},
		MusterOut: []string{
			"______",
			"______",
			"______",
			"______",
			"______",
			"______",
			"______",
		},
		Branch:     "branch",
		SkillTable: sklTab,
	}
	bt, err := json.MarshalIndent(&cr, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	usrHome, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	sep := string(filepath.Separator)
	// path := usrHome + sep + `TabletopGames` + sep + `HOSTILE` + sep + `careers` + sep + `template.json`
	// fmt.Println(path)

	f, err := os.OpenFile(usrHome+sep+`TabletopGames`+sep+`HOSTILE`+sep+`careers`+sep+`template.json`, os.O_CREATE|os.O_WRONLY, 0770)
	f.Write(bt)
	defer f.Close()
	//CORPORATE_AGENT
	caSkills := make(map[string][]string)
	caSkills["Personal Development"] = []string{"+1 STR", "+1 DEX", text(skill.Bribery), "+1 INT", "+1 EDU", "+1 SOC"}
	caSkills["Service Skills"] = []string{text(skill.Gun_Combat), text(skill.Vacc_Suit), text(skill.Vechicle), text(skill.Streetwise), text(skill.Gun_Combat), text(skill.Recon)}
	caSkills["Specialist Skills"] = []string{text(skill.Forgery), text(skill.Investigate), text(skill.Computer), text(skill.Carousing), text(skill.Comms), text(skill.Jack_of_All_Trades)}
	caSkills["Advanced Education Skills"] = []string{text(skill.Medical), text(skill.Security), text(skill.Administration), text(skill.Computer), text(skill.Leader), text(skill.Tactics)}
	corpAgent := CareerStats{
		Name:          "Corporate Agent",
		Qualification: "INT 6+",
		Survival:      "DEX 5+",
		Commision:     "SOC 5+",
		Advance:       "INT 7+",
		ReEnlist:      5,
		Ranks: []Rank{
			Rank{
				Value:             0,
				CommisionRequired: false,
				Position:          "Agent",
				AutoSkill:         "",
			},
			Rank{
				Value:             1,
				CommisionRequired: true,
				Position:          "Senior Agent",
				AutoSkill:         "",
			},
			Rank{
				Value:             2,
				CommisionRequired: true,
				Position:          "Supervisor",
				AutoSkill:         "",
			},
			Rank{
				Value:             3,
				CommisionRequired: true,
				Position:          "Assistant Project Leader",
				AutoSkill:         "",
			},
			Rank{
				Value:             4,
				CommisionRequired: true,
				Position:          "Project Leader",
				AutoSkill:         "",
			},
			Rank{
				Value:             5,
				CommisionRequired: true,
				Position:          "Assistant Division Chief",
				AutoSkill:         "",
			},
			Rank{
				Value:             6,
				CommisionRequired: true,
				Position:          "Division Chief",
				AutoSkill:         "",
			},
		},
		MusterOut: []string{
			Benefit_Weapon,
			Benefit_TicketStandard,
			Benefit_Weapon,
			Benefit_1INT,
			Benefit_1EDU,
			Benefit_1SOC,
			Benefit_1SOC,
		},
		SkillTable: caSkills,
	}
	bt, err = json.MarshalIndent(&corpAgent, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	f, err = os.OpenFile(usrHome+sep+`TabletopGames`+sep+`HOSTILE`+sep+`careers`+sep+`corporate_agent.json`, os.O_CREATE|os.O_WRONLY, 0770)
	f.Write(bt)
	defer f.Close()

	//CORPORATE_EXEC
	ceSkills := make(map[string][]string)
	ceSkills["Personal Development"] = []string{text(skill.Streetwise), "+1 INT", "+1 EDU", "+1 SOC", text(skill.Carousing), text(skill.Bribery)}
	ceSkills["Service Skills"] = []string{text(skill.Gambling), text(skill.Administration), text(skill.Carousing), text(skill.Leader), text(skill.Bribery), text(skill.Forgery)}
	ceSkills["Specialist Skills"] = []string{text(skill.Broker), text(skill.Liason), text(skill.Vechicle), text(skill.Broker), text(skill.Computer), text(skill.Leader)}
	ceSkills["Advanced Education Skills"] = []string{text(skill.Medical), text(skill.Computer), text(skill.Administration), text(skill.Liason), text(skill.Leader), text(skill.Jack_of_All_Trades)}
	corpExec := CareerStats{
		Name:             "Corporate Exec",
		QualificationReq: "SOC 10+",
		Survival:         "Basic 3+",
		Commision:        "EDU 5+",
		Advance:          "INT 10+",
		ReEnlist:         4,
		Ranks: []Rank{
			Rank{
				Value:             1,
				CommisionRequired: false,
				Position:          "Vice President",
				AutoSkill:         text(skill.Broker) + " 1",
			},
			Rank{
				Value:             2,
				CommisionRequired: true,
				Position:          "Senior Vice President",
				AutoSkill:         "",
			},
			Rank{
				Value:             3,
				CommisionRequired: true,
				Position:          "Executive Senior Vice President",
				AutoSkill:         "",
			},
			Rank{
				Value:             4,
				CommisionRequired: true,
				Position:          "Senior Executive Vice President",
				AutoSkill:         "",
			},
			Rank{
				Value:             5,
				CommisionRequired: true,
				Position:          "Director",
				AutoSkill:         "",
			},
			Rank{
				Value:             6,
				CommisionRequired: true,
				Position:          "President",
				AutoSkill:         "",
			},
		},
		MusterOut: []string{
			Benefit_EliteTicket,
			Benefit_1EDU,
			Benefit_1INT,
			Benefit_EliteTicket,
			Benefit_StarEnvoyClubMember,
			Benefit_1SOC,
			Benefit_1SOC,
		},
		SkillTable: ceSkills,
	}
	bt, err = json.MarshalIndent(&corpExec, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	f, err = os.OpenFile(usrHome+sep+`TabletopGames`+sep+`HOSTILE`+sep+`careers`+sep+`corporate_exec.json`, os.O_CREATE|os.O_WRONLY, 0770)
	f.Write(bt)
	defer f.Close()

	//COLONIST
	coSkills := make(map[string][]string)
	coSkills["Personal Development"] = []string{"+1 STR", "+1 DEX", "+1 END", "+1 INT", text(skill.Brawling), text(skill.Gun_Combat)}
	coSkills["Service Skills"] = []string{text(skill.Mechanical), text(skill.Comms), text(skill.Agriculture), text(skill.Electronics), text(skill.Survival), text(skill.Vechicle)}
	coSkills["Specialist Skills"] = []string{text(skill.Loader), text(skill.Carousing), text(skill.Jack_of_All_Trades), text(skill.Engineering), text(skill.Agriculture), text(skill.Vechicle)}
	coSkills["Advanced Education Skills"] = []string{text(skill.Medical), text(skill.Agriculture), text(skill.Jack_of_All_Trades), text(skill.Liason), text(skill.Administration), text(skill.Leader)}
	colonist := CareerStats{
		Name:          "Colonist",
		Qualification: "END 5+",
		Survival:      "END 6+",
		Commision:     "INT 7+",
		Advance:       "EDU 6+",
		ReEnlist:      5,
		Ranks: []Rank{
			Rank{
				Value:             0,
				CommisionRequired: false,
				Position:          "Colonist",
				AutoSkill:         "",
			},
			Rank{
				Value:             1,
				CommisionRequired: false,
				Position:          "Vice President",
				AutoSkill:         text(skill.Broker) + " 1",
			},
			Rank{
				Value:             2,
				CommisionRequired: true,
				Position:          "Senior Vice President",
				AutoSkill:         "",
			},
			Rank{
				Value:             3,
				CommisionRequired: true,
				Position:          "Executive Senior Vice President",
				AutoSkill:         "",
			},
			Rank{
				Value:             4,
				CommisionRequired: true,
				Position:          "Senior Executive Vice President",
				AutoSkill:         "",
			},
			Rank{
				Value:             5,
				CommisionRequired: true,
				Position:          "Director",
				AutoSkill:         "",
			},
			Rank{
				Value:             6,
				CommisionRequired: true,
				Position:          "President",
				AutoSkill:         "",
			},
		},
		MusterOut: []string{
			Benefit_EliteTicket,
			Benefit_1EDU,
			Benefit_1INT,
			Benefit_EliteTicket,
			Benefit_StarEnvoyClubMember,
			Benefit_1SOC,
			Benefit_1SOC,
		},
		SkillTable: coSkills,
	}
	bt, err = json.MarshalIndent(&colonist, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	f, err = os.OpenFile(usrHome+sep+`TabletopGames`+sep+`HOSTILE`+sep+`careers`+sep+`corporate_agent.json`, os.O_CREATE|os.O_WRONLY, 0770)
	f.Write(bt)
	defer f.Close()

}

func text(i int) string {
	return skill.SkillStr(i)
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
