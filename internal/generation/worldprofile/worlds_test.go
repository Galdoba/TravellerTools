package worldprofile

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

func TestNums(t *testing.T) {
	tlMap := make(map[string]int)
	last := "???????-?"
	for i := 0; i < 10000; i++ {
		mwUWP := NewMain(last + fmt.Sprintf("_Very  long Seed sdhdhjdfgString%v", i))
		last = mwUWP
		_, err := uwp.FromString(mwUWP)
		if err != nil {
			t.Errorf("%v failed: %v\ninput string: %v", last, err.Error(), mwUWP)
		}

		tlMap[last]++
	}
	for k, v := range tlMap {
		if v > 1 {
			fmt.Println(k, v)
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
