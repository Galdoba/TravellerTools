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
	Benefit_2INT                = "+2 INT"
	Benefit_1EDU                = "+1 EDU"
	Benefit_2EDU                = "+2 EDU"
	Benefit_1SOC                = "+1 SOC"
	Benefit_StarEnvoyClubMember = "Star Envoy Club Member"
	Benefit_TicketStandard      = "Standard Ticket"
	Benefit_TicketElite         = "Elite Ticket"
	Benefit_Weapon              = "Weapon"
	Benefit_Award               = "Award"
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
			Benefit_TicketElite,
			Benefit_1EDU,
			Benefit_1INT,
			Benefit_TicketElite,
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
				CommisionRequired: true,
				Position:          "Team Leader",
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
				Position:          "Department Chief",
				AutoSkill:         "",
			},
			Rank{
				Value:             4,
				CommisionRequired: true,
				Position:          "Assistant Operations Manager",
				AutoSkill:         "",
			},
			Rank{
				Value:             5,
				CommisionRequired: true,
				Position:          "Operations Manager",
				AutoSkill:         "",
			},
			Rank{
				Value:             6,
				CommisionRequired: true,
				Position:          "Colonial Administrator",
				AutoSkill:         "",
			},
		},
		MusterOut: []string{
			Benefit_TicketStandard,
			Benefit_1INT,
			Benefit_Weapon,
			Benefit_TicketStandard,
			Benefit_TicketStandard,
			Benefit_TicketElite,
			Benefit_1SOC,
		},
		SkillTable: coSkills,
	}
	bt, err = json.MarshalIndent(&colonist, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	f, err = os.OpenFile(usrHome+sep+`TabletopGames`+sep+`HOSTILE`+sep+`careers`+sep+`colonist.json`, os.O_CREATE|os.O_WRONLY, 0770)
	f.Write(bt)
	defer f.Close()

	//COMMERCIAL SPACER
	csSkills := make(map[string][]string)
	csSkills["Personal Development"] = []string{"+1 STR", "+1 DEX", "+1 END", "+1 STR", text(skill.Brawling), text(skill.Vacc_Suit)}
	csSkills["Service Skills"] = []string{text(skill.Comms), text(skill.Bribery), text(skill.Gun_Combat), text(skill.Loader), text(skill.Broker), text(skill.Vechicle)}
	csSkills["Specialist Skills"] = []string{text(skill.Vacc_Suit), text(skill.Mechanical), text(skill.Loader), text(skill.Electronics), text(skill.Steward), text(skill.Navigation)}
	csSkills["Advanced Education Skills"] = []string{text(skill.Medical), text(skill.Computer), text(skill.Broker), text(skill.Pilot), text(skill.Engineering), text(skill.Navigation)}
	commercial_spacer := CareerStats{
		Name:          "Commercial Spacer",
		Qualification: "INT 4+",
		Survival:      "INT 5+",
		Commision:     "INT 5+",
		Advance:       "EDU 8+",
		ReEnlist:      4,
		Ranks: []Rank{
			Rank{
				Value:             0,
				CommisionRequired: false,
				Position:          "Crewman",
				AutoSkill:         text(skill.Vacc_Suit) + " 1",
			},
			Rank{
				Value:             1,
				CommisionRequired: true,
				Position:          "Fourth Officer",
				AutoSkill:         "",
			},
			Rank{
				Value:             2,
				CommisionRequired: true,
				Position:          "Third Officer",
				AutoSkill:         "",
			},
			Rank{
				Value:             3,
				CommisionRequired: true,
				Position:          "Second Officer",
				AutoSkill:         "",
			},
			Rank{
				Value:             4,
				CommisionRequired: true,
				Position:          "First Officer",
				AutoSkill:         text(skill.Pilot) + " 1",
			},
			Rank{
				Value:             5,
				CommisionRequired: true,
				Position:          "Captain",
				AutoSkill:         "",
			},
			Rank{
				Value:             6,
				CommisionRequired: true,
				Position:          "Senior Captain",
				AutoSkill:         "",
			},
		},
		MusterOut: []string{
			Benefit_TicketStandard,
			Benefit_1EDU,
			Benefit_Weapon,
			Benefit_TicketElite,
			Benefit_1INT,
			Benefit_TicketElite,
			Benefit_StarEnvoyClubMember,
		},
		SkillTable: csSkills,
	}
	bt, err = json.MarshalIndent(&commercial_spacer, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	f, err = os.OpenFile(usrHome+sep+`TabletopGames`+sep+`HOSTILE`+sep+`careers`+sep+`commercial_spacer.json`, os.O_CREATE|os.O_WRONLY, 0770)
	f.Write(bt)
	defer f.Close()

	//MARINE
	maSkills := make(map[string][]string)
	maSkills["Personal Development"] = []string{"+1 STR", "+1 DEX", "+1 END", text(skill.Gambling), text(skill.Brawling), text(skill.Blade_Combat)}
	maSkills["Service Skills"] = []string{text(skill.Ground_Vechicle), text(skill.Comms), text(skill.Gun_Combat), text(skill.Survival), text(skill.Gun_Combat), text(skill.Vacc_Suit)}
	maSkills["Specialist Skills"] = []string{text(skill.Vechicle), text(skill.Mechanical), text(skill.Electronics), text(skill.Demolitions), text(skill.Recon), text(skill.Heavy_Weapons)}
	maSkills["Advanced Education Skills"] = []string{text(skill.Medical), text(skill.Security), text(skill.Tactics), text(skill.Computer), text(skill.Leader), text(skill.Administration)}
	marine := CareerStats{
		Name:          "Marine",
		Qualification: "INT 4+",
		Survival:      "END 6+",
		Commision:     "EDU 9+",
		Advance:       "EDU 6+",
		ReEnlist:      6,
		Ranks: []Rank{
			Rank{
				Value:             0,
				CommisionRequired: false,
				Position:          "Private",
				AutoSkill:         text(skill.Gun_Combat) + " 1",
			},
			Rank{
				Value:             1,
				CommisionRequired: false,
				Position:          "Lance Copral",
				AutoSkill:         text(skill.Vacc_Suit) + " 1",
			},
			Rank{
				Value:             2,
				CommisionRequired: false,
				Position:          "Copral",
				AutoSkill:         "",
			},
			Rank{
				Value:             3,
				CommisionRequired: false,
				Position:          "Sergeant",
				AutoSkill:         "",
			},
			Rank{
				Value:             4,
				CommisionRequired: false,
				Position:          "Staff Sergeant",
				AutoSkill:         "",
			},
			Rank{
				Value:             5,
				CommisionRequired: false,
				Position:          "Gunnery Sergeant",
				AutoSkill:         "",
			},
			Rank{
				Value:             6,
				CommisionRequired: false,
				Position:          "Master Sergeant",
				AutoSkill:         "",
			},
			Rank{
				Value:             1,
				CommisionRequired: true,
				Position:          "Lieutenant",
				AutoSkill:         text(skill.Vacc_Suit) + " 1",
			},
			Rank{
				Value:             2,
				CommisionRequired: true,
				Position:          "Captain",
				AutoSkill:         "",
			},
			Rank{
				Value:             3,
				CommisionRequired: true,
				Position:          "Major",
				AutoSkill:         "",
			},
			Rank{
				Value:             4,
				CommisionRequired: true,
				Position:          "Lieut. Colonel",
				AutoSkill:         "",
			},
			Rank{
				Value:             5,
				CommisionRequired: true,
				Position:          "Colonel",
				AutoSkill:         "",
			},
			Rank{
				Value:             6,
				CommisionRequired: true,
				Position:          "General",
				AutoSkill:         "",
			},
		},
		MusterOut: []string{
			Benefit_TicketStandard,
			Benefit_Award,
			Benefit_1EDU,
			Benefit_Weapon,
			Benefit_StarEnvoyClubMember,
			Benefit_TicketElite,
			Benefit_2EDU,
		},
		SkillTable: maSkills,
	}
	bt, err = json.MarshalIndent(&marine, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	f, err = os.OpenFile(usrHome+sep+`TabletopGames`+sep+`HOSTILE`+sep+`careers`+sep+`marine.json`, os.O_CREATE|os.O_WRONLY, 0770)
	f.Write(bt)
	defer f.Close()

	//MARSHAL
	marshSkills := make(map[string][]string)
	marshSkills["Personal Development"] = []string{"+1 STR", "+1 DEX", "+1 END", "+1 INT", "+1 EDU", "+1 SOC"}
	marshSkills["Service Skills"] = []string{text(skill.Streetwise), text(skill.Brawling), text(skill.Brawling), text(skill.Investigate), text(skill.Recon), text(skill.Gun_Combat)}
	marshSkills["Specialist Skills"] = []string{text(skill.Gun_Combat), text(skill.Comms), text(skill.Ground_Vechicle), text(skill.Security), text(skill.Computer), text(skill.Medical)}
	marshSkills["Advanced Education Skills"] = []string{text(skill.Medical), text(skill.Tactics), text(skill.Investigate), text(skill.Computer), text(skill.Tactics), text(skill.Administration)}
	marshal := CareerStats{
		Name:          "Marshal",
		Qualification: "INT 7+",
		Survival:      "DEX 6+",
		Commision:     "EDU 8+",
		Advance:       "INT 7+",
		ReEnlist:      5,
		Ranks: []Rank{
			Rank{
				Value:             0,
				CommisionRequired: false,
				Position:          "Deputy",
				AutoSkill:         text(skill.Investigate) + " 1",
			},
			Rank{
				Value:             1,
				CommisionRequired: true,
				Position:          "Senior Deputy",
				AutoSkill:         "",
			},
			Rank{
				Value:             2,
				CommisionRequired: true,
				Position:          "Supervisory Deputy",
				AutoSkill:         "",
			},
			Rank{
				Value:             3,
				CommisionRequired: true,
				Position:          "Assistant Chief",
				AutoSkill:         "",
			},
			Rank{
				Value:             4,
				CommisionRequired: true,
				Position:          "Chief Deputy",
				AutoSkill:         "",
			},
			Rank{
				Value:             5,
				CommisionRequired: true,
				Position:          "Marshal",
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
			Benefit_TicketStandard,
			Benefit_TicketElite,
			Benefit_2INT,
			Benefit_1EDU,
			Benefit_Weapon,
			Benefit_TicketElite,
			Benefit_1SOC,
		},
		SkillTable: marshSkills,
	}
	bt, err = json.MarshalIndent(&marshal, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	f, err = os.OpenFile(usrHome+sep+`TabletopGames`+sep+`HOSTILE`+sep+`careers`+sep+`marshal.json`, os.O_CREATE|os.O_WRONLY, 0770)
	f.Write(bt)
	defer f.Close()

	//MILITARY SPACER
	msSkills := make(map[string][]string)
	msSkills["Personal Development"] = []string{"+1 STR", "+1 DEX", "+1 END", "+1 INT", "+1 EDU", "+1 SOC"}
	msSkills["Service Skills"] = []string{text(skill.Vacc_Suit), text(skill.Computer), text(skill.Loader), text(skill.Gunnery), text(skill.Brawling), text(skill.Gun_Combat)}
	msSkills["Specialist Skills"] = []string{text(skill.Gunnery), text(skill.Mechanical), text(skill.Electronics), text(skill.Engineering), text(skill.Leader), text(skill.Comms)}
	msSkills["Advanced Education Skills"] = []string{text(skill.Medical), text(skill.Navigation), text(skill.Engineering), text(skill.Computer), text(skill.Pilot), text(skill.Administration)}
	mSpacer := CareerStats{
		Name:          "Military Spacer",
		Qualification: "INT 6+",
		Survival:      "INT 5+",
		Commision:     "SOC 7+",
		Advance:       "EDU 6+",
		ReEnlist:      5,
		Ranks: []Rank{
			Rank{
				Value:             0,
				CommisionRequired: false,
				Position:          "Airman",
				AutoSkill:         text(skill.Vacc_Suit) + " 1",
			},
			Rank{
				Value:             1,
				CommisionRequired: false,
				Position:          "Airman First Class",
				AutoSkill:         "",
			},
			Rank{
				Value:             2,
				CommisionRequired: false,
				Position:          "Senior Airman",
				AutoSkill:         "",
			},
			Rank{
				Value:             3,
				CommisionRequired: false,
				Position:          "Sergeant",
				AutoSkill:         "",
			},
			Rank{
				Value:             4,
				CommisionRequired: false,
				Position:          "Staff Sergeant",
				AutoSkill:         "",
			},
			Rank{
				Value:             5,
				CommisionRequired: false,
				Position:          "Technical Sergeant",
				AutoSkill:         "",
			},
			Rank{
				Value:             6,
				CommisionRequired: false,
				Position:          "Master Sergeant",
				AutoSkill:         "",
			},
			Rank{
				Value:             1,
				CommisionRequired: true,
				Position:          "Lieutenant",
				AutoSkill:         text(skill.Vacc_Suit) + " 1",
			},
			Rank{
				Value:             2,
				CommisionRequired: true,
				Position:          "Captain",
				AutoSkill:         "",
			},
			Rank{
				Value:             3,
				CommisionRequired: true,
				Position:          "Major",
				AutoSkill:         "",
			},
			Rank{
				Value:             4,
				CommisionRequired: true,
				Position:          "Lieut. Colonel",
				AutoSkill:         "",
			},
			Rank{
				Value:             5,
				CommisionRequired: true,
				Position:          "Colonel",
				AutoSkill:         Benefit_1SOC,
			},
			Rank{
				Value:             6,
				CommisionRequired: true,
				Position:          "General",
				AutoSkill:         "",
			},
		},
		MusterOut: []string{
			Benefit_Award,
			Benefit_1EDU,
			Benefit_2EDU,
			Benefit_TicketElite,
			Benefit_TicketStandard,
			Benefit_TicketElite,
			Benefit_StarEnvoyClubMember,
		},
		SkillTable: maSkills,
	}
	bt, err = json.MarshalIndent(&mSpacer, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	f, err = os.OpenFile(usrHome+sep+`TabletopGames`+sep+`HOSTILE`+sep+`careers`+sep+`military_spacer.json`, os.O_CREATE|os.O_WRONLY, 0770)
	f.Write(bt)
	defer f.Close()

	//MARSHAL

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
