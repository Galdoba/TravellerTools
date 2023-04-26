package orbit

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/generation/star"
)

func TestGeneration(t *testing.T) {
	dice := dice.New().SetSeed("Troj 2223 aldo")
	pair, err := star.NewPair("G2 V", "M9 VI")
	if err != nil {
		t.Errorf(err.Error())
	}
	spoMap := StarPlanetOrbits(dice, pair)
	fmt.Println(pair.Class())
	for i := 0; i < 1000000; i++ {
		if val, ok := spoMap[i]; ok {
			fmt.Println(i, val)
		}
	}
}
