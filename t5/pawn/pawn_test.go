package pawn

import (
	"fmt"
	"testing"
	"time"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/pkg/struct/world"
	"github.com/Galdoba/TravellerTools/t5/pawn/characteristic"
	"github.com/Galdoba/TravellerTools/t5/pawn/education"
)

func TestPawn(t *testing.T) {
	for i := 0; i < 1; i++ {
		dice := dice.New()
		w := &world.World{}
		switch dice.Sroll("2d6") {
		default:
			w, _ = world.NewWorld(
				world.KnownData(world.IsMainworld, world.FLAG_TRUE),
				world.KnownData(profile.KEY_SIZE, "7"),
				world.KnownData(profile.KEY_HYDR, "7"),
			)
			w.GenerateFull(dice)
		case 12:
			fmt.Println("DEEP SPACE CHARACTER")
			w = world.DeepSpace()
		}

		//gt := genetics.NewTemplate("SDEIES", "222222")
		//fmt.Println("==============")
		chr2, err := New(dice, control_Random, w.ListTC())
		if err != nil {
			t.Errorf(err.Error())
		}
		//dice := dice.New()
		//fmt.Println("==============")
		//genome := genetics.NewGeneData("SDEIES", "222222")
		//chr2.InjectGenetics(genome)
		//if err := chr2.RollCharacteristics(dice); err != nil {
		//t.Errorf(err.Error())
		//}

		fmt.Println("==============")
		fmt.Println(chr2)

		fmt.Println(chr2.CheckCharacteristic(CheckAverage, CHAR_TRAINING))
		for i := CHAR_STRENGHT; i < 18; i++ {
			fmt.Println(characteristic.FromProfile(chr2.profile, i))
		}
		for _, ev := range chr2.generationEvents {
			fmt.Println(ev)
		}
		chr2.StartEducationProgram(education.BasicSchoolTrainingCourse)
		for _, ev := range chr2.generationEvents {
			fmt.Println(ev)
		}
		chr2.StartEducationProgram(education.Mentor)
		for _, ev := range chr2.generationEvents {
			fmt.Println(ev)
		}

		chr2.StartEducationProgram(education.BasicSchoolApprentice)
		fmt.Println("----------")
		for _, ev := range chr2.generationEvents {

			fmt.Println(ev)
		}
		time.Sleep(time.Millisecond * 100)
	}

}
