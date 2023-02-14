package upp

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestUPP(t *testing.T) {
	return
	pass := 0
	fail := 0
	total := 0
	dice := dice.New()
	max := len(corectGenMaps()) * len(corectProfiles())
	for _, genetics := range corectProfiles() {
		for _, genMap := range corectGenMaps() {
			fmt.Printf("Test: genetics '%v' map '%v' || %v/%v    \r", genetics, genMap, total, max)
			total++
			t.Logf("test %d\n", total)
			gd, err := GeneDataManual(genetics, genMap)
			if err != nil {
				t.Errorf("GeneDataManual error: %v", err.Error())
				fail++
			}
			up := NewUniversalPersonalityProfile(dice, gd)
			t.Logf("up: %v\n", up.String())
			pass++
		}
	}
	t.Logf("Result: %v/%v/%v (PASS/FAIL/TOTAL)\n", pass, fail, total)

	human := NewUniversalPersonalityProfile(dice, GeneDataHuman())
	fmt.Println(human)
}

func TestHumanUPP(t *testing.T) {
	dice := dice.New()
	for i := 0; i < 20; i++ {
		human := NewUniversalPersonalityProfile(dice, GeneDataHuman())
		fmt.Println(human)
	}
}
