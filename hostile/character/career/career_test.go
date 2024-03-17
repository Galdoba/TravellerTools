package career

import (
	"fmt"
	"testing"
)

func TestCareer(t *testing.T) {
	GetCareer("name")
}

func TestLoadCareer(t *testing.T) {
	fmt.Println("Test Load")
	for _, name := range []string{
		Colonist,
		CorporateAgent,
		CorporateExec,
		CommersialSpacer,
		Marine,
		Marshal,
		MilitarySpacer,
		Physician,
		Ranger,
		Rogue,
		Roughneck,
		Scientist,
		SurveyScout,
		Technician,
	} {
		cs, err := LoadCareerStats(name)
		if err != nil {

			t.Errorf("err %v - %v\n%v\n", err, name, cs)
		}
	}
}
