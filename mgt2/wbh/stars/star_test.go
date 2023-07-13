package stars

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestStarType(t *testing.T) {
	primMap := make(map[string]int)
	for i := 0; i < 1000; i++ {
		dice := dice.New().SetSeed(i)
		ss, err := NewStarSystem(dice, GenerationMethodUnusual, TypeVariantTraditional)
		if err != nil {
			t.Errorf("roll %v: primary: %v", i, ss.primary)
		}
		if ss.primary.class == classV {
			//continue
		}
		// if !strings.HasPrefix(ss.primary.class, "Class ") || !strings.HasPrefix(ss.primary.sttype, "Type ") {
		// 	t.Errorf("roll %v: primary: %v", i, ss.primary)
		// }
		if strings.Contains(ss.primary.class, "BD") && ss.primary.sttype != "" {
			t.Errorf("roll %v: primary: %v", i, ss.primary)
		}
		short := shortStarDescription(ss.primary)
		switch short {
		case nebula, protostar, neutronStar, pulsar, blackHole, starcluster, anomaly:
			//fmt.Println(short + "                                   ")
		case "K5 IV", "K6 IV", "K7 IV", "K8 IV", "K9 IV":
			t.Errorf("invalid class %v", short)
		case "M0 IV", "M1 IV", "M2 IV", "M3 IV", "M4 IV", "M5 IV", "M6 IV", "M7 IV", "M8 IV", "M9 IV":
			t.Errorf("invalid class %v", short)
		case "O0 IV", "O1 IV", "O2 IV", "O3 IV", "O4 IV", "O5 IV", "O6 IV", "O7 IV", "O8 IV", "O9 IV":
			t.Errorf("invalid class %v", short)
		case "A0 VI", "A1 VI", "A2 VI", "A3 VI", "A4 VI", "A5 VI", "A6 VI", "A7 VI", "A8 VI", "A9 VI":
			t.Errorf("invalid class %v", short)
		case "F0 VI", "F1 VI", "F2 VI", "F3 VI", "F4 VI", "F5 VI", "F6 VI", "F7 VI", "F8 VI", "F9 VI":
			t.Errorf("invalid class %v", short)
		}
		primMap[short]++

		fmt.Printf("%v %v\n", i, ss.primary)

	}

	// for k, v := range primMap {
	// 	fmt.Println(k, v)
	// }
}
