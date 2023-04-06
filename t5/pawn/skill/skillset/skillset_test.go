package skillset

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/t5/pawn/skill"
)

func TestSkillset(t *testing.T) {
	ss := NewSkillSet()
	fmt.Println(ss)
	fmt.Println("================")
	for i := 0; i < 4; i++ {
		errorArr := []error{}

		if err := ss.Increase(skill.ID_Astrogator); err != nil {
			errorArr = append(errorArr, err)
		}
		if err := ss.Increase(skill.ID_Engineer); err != nil {
			errorArr = append(errorArr, err)
		}
		if err := ss.Increase(skill.ID_Blades); err != nil {
			errorArr = append(errorArr, err)
		}
		if err := ss.Increase(skill.ID_Power); err != nil {
			errorArr = append(errorArr, err)
		}

		if len(errorArr) != 0 {
			for _, err := range errorArr {
				t.Errorf("error: %v", err.Error())
			}
		}
	}
	fmt.Println(ss)

}
