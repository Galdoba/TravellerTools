package stars

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestStarType(t *testing.T) {
	primMap := make(map[string]int)
	for i := 0; i < 5000000; i++ {
		dice := dice.New().SetSeed(i)
		ss, err := NewStarSystem(dice, GenerationMethodUnusual, TypeVariantRealistic)
		if err != nil {
			t.Errorf("roll %v: primary: %v", i, ss.primary)
		}
		// if !strings.HasPrefix(ss.primary.class, "Class ") || !strings.HasPrefix(ss.primary.sttype, "Type ") {
		// 	t.Errorf("roll %v: primary: %v", i, ss.primary)
		// }
		if strings.Contains(ss.primary.class, "BD") && ss.primary.sttype != "" {
			t.Errorf("roll %v: primary: %v", i, ss.primary)
		}
		primMap[ss.primary.sttype+" "+ss.primary.class]++

	}

	for k, v := range primMap {
		fmt.Println(k, v)
	}
}
