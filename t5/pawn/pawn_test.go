package pawn

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestPawn(t *testing.T) {
	// for i := 0; i < 1; i++ {
	// 	dice := dice.New()
	// 	w := &world.World{}
	// 	switch dice.Sroll("2d6") {
	// 	default:
	// 		w, _ = world.NewWorld(
	// 			world.KnownData(world.IsMainworld, world.FLAG_TRUE),
	// 		)
	// 		w.GenerateFull(dice)
	// 	case 12:
	// 		fmt.Println("DEEP SPACE CHARACTER")
	// 		w = world.DeepSpace()
	// 	}
	// 	fmt.Println(w)
	// 	gt := genetics.NewTemplate("SDEIES", "222222")
	// 	chr, err := New(control_Random, gt, w)
	// 	fmt.Println(chr.chrSet)
	// 	fmt.Println(chr.sklSet)
	// 	fmt.Println(err)
	// 	time.Sleep(time.Millisecond * 100)
	// }
	fmt.Println("==============")
	chr2 := New2()
	dice := dice.New()
	// if err := chr2.InjectGenetics(genetics.NewTemplate("SDEIES", "332222")); err != nil {
	// 	t.Errorf(err.Error())
	// }
	if err := chr2.RollCharacteristics(dice); err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(chr2)
	fmt.Println(chr2.profile)

}
