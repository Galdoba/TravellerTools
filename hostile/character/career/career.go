package career

import (
	"fmt"
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
	Name          string `json:"Career"`
	Qualification string `json:"Qualification"`
	Survival      string `json:"Survival"`
	Commision     string `json:"Commision"`
	Advance       string `json:"Advance"`
	ReEnlist      int    `json:"Re-Enlist"`
}
