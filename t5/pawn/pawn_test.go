package pawn

import (
	"fmt"
	"testing"
	"time"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/struct/world"
	"github.com/Galdoba/TravellerTools/t5/genetics"
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
		fmt.Println(w)
		gt := genetics.NewTemplate("SDEIES", "222222")
		chr, err := New(control_Random, gt, w)
		fmt.Println(chr.chrSet)
		fmt.Println(chr.sklSet)
		fmt.Println(err)
		time.Sleep(time.Millisecond * 100)
	}

}
