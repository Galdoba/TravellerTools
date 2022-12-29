package orbit

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestGeneration(t *testing.T) {
	dice := dice.New()

	orbMap := Generate(dice, "G2 V")
	for k, v := range orbMap {
		fmt.Println(k, v)
	}
	fmt.Println(len(orbMap))
}
