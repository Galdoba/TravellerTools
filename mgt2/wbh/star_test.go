package wbh

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestStarType(t *testing.T) {

	for i := 1; i < 40; i++ {
		dice := dice.New().SetSeed(i)
		ss, err := NewStarSystem(dice, GenerationMethodUnusual, TypeVariantTraditional)
		if err != nil {
			//t.Errorf("roll %v: primary: %v", i, ss.primary)
		}
		//fmt.Println("i =", i)
		//for _, code := range []string{"Aa", "Ab", "Ba", "Bb", "Ca", "Cb", "Da", "Db"} {
		//if v, ok := ss.star[code]; ok {
		//		fmt.Println(code, v)
		//}
		//}
		// if len(ss.Star) != 1 {
		// 	continue
		// }
		fmt.Printf("[%v]	%v %v\n", i, ss.String(), ss.Star["Aa"].MAO)
		for k, v := range ss.Star {
			fmt.Println(k, v)
		}
		//l := 0
		//for _, desig := range []string{"Aa", "Ab", "Ba", "Bb", "Ca", "Cb", "Da", "Db"} {
		//if st, ok := ss.star[desig]; ok {
		//fmt.Println(i, l, desig, st)
		//	l++
		//}
		//}
		// if ss.primary.class != classBD && ss.primary.class != classD {
		// 	continue
		// }
		// st := ss.primary
		// switch ss.primary.class {
		// case classIa, classIb, classII, classIII, classIV, classV, classVI, classBD, classD:
		// 	if st.age < 0 {
		// 		t.Errorf("star has negative age %v", st)
		// 	}
		// 	if st.sttype == "" {
		// 		t.Errorf("star type not defined %v", st)
		// 	}
		// 	if st.subtype == "dwarf" {
		// 		t.Errorf("star subtype not defined %v", st)
		// 	}
		// 	if st.diameter <= 0 {
		// 		t.Errorf("star diameter not defined %v", st)
		// 	}
		// }

		// short := shortStarDescription(ss.primary)
		// switch short {
		// case nebula, protostar, neutronStar, pulsar, blackHole, starcluster, anomaly:
		// 	//fmt.Println(short + "                                   ")
		// case "K5 IV", "K6 IV", "K7 IV", "K8 IV", "K9 IV":
		// 	t.Errorf("invalid class %v", short)
		// case "M0 IV", "M1 IV", "M2 IV", "M3 IV", "M4 IV", "M5 IV", "M6 IV", "M7 IV", "M8 IV", "M9 IV":
		// 	t.Errorf("invalid class %v", short)
		// case "O0 IV", "O1 IV", "O2 IV", "O3 IV", "O4 IV", "O5 IV", "O6 IV", "O7 IV", "O8 IV", "O9 IV":
		// 	t.Errorf("invalid class %v", short)
		// case "A0 VI", "A1 VI", "A2 VI", "A3 VI", "A4 VI", "A5 VI", "A6 VI", "A7 VI", "A8 VI", "A9 VI":
		// 	t.Errorf("invalid class %v", short)
		// case "F0 VI", "F1 VI", "F2 VI", "F3 VI", "F4 VI", "F5 VI", "F6 VI", "F7 VI", "F8 VI", "F9 VI":
		// 	t.Errorf("invalid class %v", short)
		// }

	}

	// for k, v := range primMap {
	// 	fmt.Println(k, v)
	// }
}
