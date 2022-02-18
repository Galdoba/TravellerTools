package worldprofile

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"github.com/Galdoba/TravellerTools/pkg/survey"
)

func TestNums(t *testing.T) {
	return
	tlMap := make(map[string]int)
	last := "???????-?"
	for i := 0; i < 5; i++ {
		mwUWP := NewMain(last + fmt.Sprintf("_Very  long Seed sdhdhjdfgString%v", i))
		last = mwUWP
		_, err := uwp.FromString(mwUWP)
		if err != nil {
			t.Errorf("%v failed: %v\ninput string: %v", last, err.Error(), mwUWP)
		}
		ssd, err := survey.NewSecondSurvey(
			survey.Instruction(survey.MW_UWP, mwUWP),
			survey.Instruction(survey.MW_Name, last+fmt.Sprintf("_Very  long Seed sdhdhjdfgString%v", i)),
		)
		fmt.Println(ssd)
		secondaryTypes := []int{Hospitable, Planetoid, IceWorld, RadWorld, Inferno, BigWorld, Worldlet, InnerWorld, StormWorld, SGG, LGG, IG, PlanetaryRings, AsteroidBelt}

		seUwp := NewSecondary(ssd, secondaryTypes[i%len(secondaryTypes)], "Ay")
		fmt.Println("Secondary Worldlet =", seUwp)
		suwp, err := uwp.FromString(seUwp)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(suwp.Describe("All"))
		tlMap[last]++
	}
	for k, v := range tlMap {
		if v > 1 {
			fmt.Println(k, v)
		}
	}

}

func TestByMask(t *testing.T) {
	testMasks := []string{"???0???-?"}
	for i, tm := range testMasks {
		uwp, err := ByMask(tm, tm)
		if err != nil {
			t.Errorf("%v testMaks = %v: %v", i, tm, err)
		} else {
			fmt.Printf("success (%v => %v)", tm, uwp)
		}
	}
}

//A:79 B:206 C:328 D:358 E:226 X:99 - mgt2
// A :  6%
// B : 16% 22
// C : 25% 47
// D : 27% 74
// E : 17% 91
// X :  7%
//A:6 B:9 C:11 D:4 E:5 X:1 - T5
// A : 16%
// B : 25% 41
// C : 30% 71
// D : 11% 82
// E : 14% 96
// X :  3%
