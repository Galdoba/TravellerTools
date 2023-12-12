package hex

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestHex(t *testing.T) {
	for q := 0; q < 5; q++ {
		for r := 0; r < 5; r++ {
			for s := 0; s < 5; s++ {
				coord_key := fmt.Sprintf("%v:%v:%v", q, r, s)
				dice := dice.New().SetSeed(coord_key)
				hx := New(coord_key, DENSITY_STANDARD)
				err1 := hx.RollStarSystemPresence(dice)
				if err1 != nil {
					t.Error(err1)
				}
				//			err2 := hx.RollCentralObject(dice)
				//			if err2 != nil {
				//				t.Error(err2)
				//			}
				fmt.Printf("%v - %v\n", coord_key, hx)
			}
		}
	}
}
