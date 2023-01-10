package stellar

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestStellar(t *testing.T) {
	for i := 480; i < 1000; i++ {
		dp := dice.New().SetSeed(i)
		stellar := GenerateStellar(dp)
		stars := Parse(stellar)
		fmt.Printf("stellar: '%v' [%v]\n", stellar, strings.Join(stars, "|"))
		if stellar != strings.Join(stars, " ") {
			t.Errorf("not merging := %v", fmt.Sprintf("stellar: '%v' [%v]\n", stellar, strings.Join(stars, " ")))
		}
	}

}

func TestHabitableOrbits(t *testing.T) {
	for i, code := range listAllStars() {
		hz := HabitableOrbitByCode(code)
		fmt.Printf("%v	%v: %v\n", i, code, hz)
	}
}
