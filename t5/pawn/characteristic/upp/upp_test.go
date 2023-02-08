package upp

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/t5/pawn/characteristic"
)

func corectProfiles() []string {
	gp := []string{}
	for _, c1 := range []string{"S"} {
		for _, c2 := range []string{"D", "A", "G"} {
			for _, c3 := range []string{"E", "S", "V"} {
				for _, c4 := range []string{"I"} {
					for _, c5 := range []string{"E", "T", "I"} {
						for _, c6 := range []string{"S", "C", "K"} {
							gp = append(gp, c1+c2+c3+c4+c5+c6)
						}
					}
				}
			}
		}
	}
	return gp
}

func TestUPP(t *testing.T) {
	pass := 0
	fail := 0
	total := 0
	for i, genetics := range corectProfiles() {
		total++
		t.Logf("test %d\n", i+1)
		up, err := newProfile(genetics)
		if err != nil {
			fail++
			t.Errorf("newProfile(%v) returned error: %v", genetics, err.Error())
			continue
		}
		t.Logf("%v\n", up)
		pass++
	}
	t.Logf("Result: %v/%v/%v (PASS/FAIL/TOTAL)\n", pass, fail, total)
	human, _ := newProfile(characteristic.GP_Human)
	fmt.Println(human)
}
