package education

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/classifications"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/struct/world"
	"github.com/Galdoba/TravellerTools/t5/pawn"
)

func Test_hash(t *testing.T) {
	for i := 0; i < 20; i++ {
		w, _ := world.NewWorld()
		tc := classifications.Evaluate(w)
		pawn, _ := pawn.New(dice.New(), 1, tc)
		//fmt.Println(pawn)
		opt := CurrentOptions(pawn)
		institutionID := pawn.ChooseOne(opt)
		fmt.Println(opt)
		out := Attend(pawn, institutionID)
		//gainedMajor, gainedMinor, yearsPassed, waiversUsed, degreeGained, newEducationVal, skillsGained := Outcome(out)
		pawn.InjectEducationOutcome(Outcome(out))
		fmt.Println(out)
		fmt.Println(pawn)
	}
	// for i := BasicSchoolED5; i <= MarineSchool; i++ {
	// 	fmt.Println("==============")
	// 	list, err := listMajorMinorSkillID(i)
	// 	if err != nil {
	// 		t.Errorf(err.Error())
	// 	} else {
	// 		fmt.Println("institution", i)
	// 		for _, id := range list {
	// 			fmt.Println(skill.NameByID(id))
	// 		}

	// 	}
	// }

}
