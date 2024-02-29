package career

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	EnlistedByQualific = "Q"
	EnlistetByDraft    = "D"
	///////
	Android          = "0"
	CorporateAgent   = "1"
	CorporateExec    = "2"
	Colonist         = "3"
	CommersialSpacer = "4"
	Marine           = "5"
	Marshal          = "6"
	MilitarySpacer   = "7"
	Physician        = "8"
	Ranger           = "9"
	Rogue            = "A"
	Roughneck        = "B"
	Scientist        = "C"
	SurveyScout      = "D"
	Technitian       = "E"
	///////
	PersDevel    = "1"
	Service      = "2"
	Specialist   = "3"
	AdvancedEdu  = "4"
	ComMilSpacer = "5"
	Scout        = "6"
	PhysSkill    = "7"
	ScienceSkill = "8"
	TechSkill    = "9"
	////////
	SurvivalFailed = "S"
	MishapEvent    = "M"
	////////
	CommisionPassed = "Y"
	CommisionFailed = "N"
	////////
	AdvancementGained = "+"
	AdvancementMissed = "-"
	AdvancementHalted = "="
	////////
	ReenlistFailed = "N"
	ReenlistPassed = "Y"
	ReenlistForsed = "F"
)

type CareerPath struct {
	careersFinished int
	draftUsed       bool
}

type term struct {
	careerCode      string
	skillTable      string
	skillRoll       string
	commisionRoll   string
	advancementRoll string
	reenlistRoll    string
}

// 125101
func Injury(dice *dice.Dicepool) string {
	code := "I"
	r := dice.Sroll("1d6")
	code += fmt.Sprintf("%v", r)
	switch dice.Sroll("1d6") {
	case 1:
		c := dice.Sroll("1d3") - 1
		for i := 0; i < 3; i++ {
			if i == c {
				r2 := dice.Sroll("1d6")

			}
		}

	case 2:
	case 3:
	case 4:
	case 5:
	case 6:

	}
	return code
}
