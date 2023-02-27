package orbit

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestGeneration(t *testing.T) {
	dice := dice.New()
	// orb := New("G2 V", 0, 4, 1)

	// fmt.Println(orb)

	// return
	orbMap := Generate(dice, "G2 V M6 V")
	for k, v := range orbMap.orb {
		fmt.Println(k, v)

	}
	fmt.Println(len(orbMap.orb))
}
