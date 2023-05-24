package education

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/classifications"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/pkg/struct/world"
	"github.com/Galdoba/TravellerTools/t5/pawn"
)

func Test_hash(t *testing.T) {
	dice := dice.New()
	for i := 0; i < 10; i++ {
		w, _ := world.NewWorld(
			world.KnownData(world.IsMainworld, world.FLAG_TRUE),
			world.KnownData(profile.KEY_SIZE, "7"),
			world.KnownData(profile.KEY_HYDR, "7"),
		)
		w.GenerateFull(dice)
		tc := classifications.Evaluate(w)
		fmt.Println(w)
		pawn, _ := pawn.New(dice, 1, tc)
		//fmt.Println(pawn)
		opt := CurrentOptions(pawn)
		institutionID := pawn.ChooseOne(opt)
		//fmt.Println(opt)
		out := Attend(pawn, institutionID)
		//gainedMajor, gainedMinor, yearsPassed, waiversUsed, degreeGained, newEducationVal, skillsGained := Outcome(out)
		pawn.InjectEducationOutcome(Outcome(out))
		fmt.Println(out)
		fmt.Println(pawn)
		for _, ev := range out.events {
			fmt.Println(ev)
		}
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
