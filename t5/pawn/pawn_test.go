package pawn

import (
	"fmt"
	"testing"
	"time"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/struct/world"
)

func TestPawn(t *testing.T) {
	for i := 0; i < 1; i++ {
		dice := dice.New()
		w := &world.World{}
		switch dice.Sroll("2d6") {
		default:
			w, _ = world.NewWorld(
				world.KnownData(world.IsMainworld, world.FLAG_TRUE),
			)
			w.GenerateFull(dice)
		case 12:
			fmt.Println("DEEP SPACE CHARACTER")
			w = world.DeepSpace()
		}
		//fmt.Println(w)
		//gt := genetics.NewTemplate("SDEIES", "222222")
		//fmt.Println("==============")
		chr2 := New(control_Random, w.ListTC())
		//dice := dice.New()
		//fmt.Println("==============")
		//genome := genetics.NewGeneData("SDEIES", "222222")
		//chr2.InjectGenetics(genome)
		//if err := chr2.RollCharacteristics(dice); err != nil {
		//t.Errorf(err.Error())
		//}

		fmt.Println("==============")
		fmt.Println(chr2.profile)
		time.Sleep(time.Millisecond * 100)

	}

}
