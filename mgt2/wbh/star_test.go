package wbh

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestStarType(t *testing.T) {

	for i := 1; i < 2; i++ {
		dice := dice.New().SetSeed(i)
		ss, err := NewStarSystem(dice, GenerationMethodUnusual, TypeVariantTraditional, GenerationMethodExpanded)
		if err != nil {
			fmt.Println("func reterned error:", err.Error())
		}
		fmt.Printf("\n[%v]	%v\n", i, ss.String())
		for _, desig := range []string{"Aa", "Ab", "Ba", "Bb", "Ca", "Cb", "Da", "Db"} {
			if st, ok := ss.Star[desig]; ok {

				fmt.Println(desig, st)

			}
		}
	}
}
