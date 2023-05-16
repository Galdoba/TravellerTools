package education

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/t5/pawn/skill"
)

func Test_hash(t *testing.T) {
	for i := BasicSchoolED5; i <= MarineSchool; i++ {
		fmt.Println("==============")
		list, err := listMajorMinorSkillID(i)
		if err != nil {
			t.Errorf(err.Error())
		} else {
			fmt.Println("institution", i)
			for _, id := range list {
				fmt.Println(skill.NameByID(id))
			}

		}
	}

}
