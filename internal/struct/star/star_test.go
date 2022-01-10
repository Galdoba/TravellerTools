package star

import (
	"fmt"
	"testing"
)

func TestInterpolation(t *testing.T) {
	// c := 0
	// tags := strs()
	// pairs := pairs()
	// for ind := range pairs {
	// 	if ind == 0 {
	// 		continue
	// 	}
	// 	div := 5.0
	// 	dif := 5
	// 	if ind == len(pairs)-1 && strings.Contains(tags[c], "M") {
	// 		div = 4.0
	// 	}
	// 	for i := 0; i < dif; i++ {
	// 		res := interpalation(pairs[ind-1], pairs[ind], float64(i), div)
	// 		if res != -999.9 {
	// 			marker := ""
	// 			//t.Errorf("Interpolation ERROR: \nhave %v ", res)
	// 			res = utils.RoundFloat64(res, 2)
	// 			if strings.Contains(tags[c], "0") || strings.Contains(tags[c], "5") {
	// 				marker = " //check " + fmt.Sprintf("%v", pairs[ind-1])
	// 			}
	// 			fmt.Printf("lumaMap[%v VI] = %v%v\r", tags[c], res, marker)
	// 		}
	// 		c++
	// 	}
	// }
	for _, code := range allCodes() {
		l := baseStellarLuminocity(code)
		m := baseStellarMass(code)
		if l == 0 || m == 0 {
			t.Errorf("code '%v': luma/mass is not expected to be %v/%v", code, l, m)
		}

	}
	fmt.Println("////////////////")
}

func TestStar(t *testing.T) {
	st, err := New("Test star", "", Category_Primary)
	fmt.Println(st)
	if err != nil {
		t.Errorf("error encountered: %v", err.Error())
	}

}

func TestStellarEncode(t *testing.T) {
	spectrals := []string{"O", "B", "A", "F", "G", "K", "M", "BD"}
	decimals := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	sizes := []string{"Ia", "Ib", "II", "III", "IV", "V", "VI", "D"}
	for _, spec := range spectrals {
		for _, sz := range sizes {
			for _, dec := range decimals {
				code := encodeStellar(spec, dec, sz)
				if code == "error" {
					t.Errorf("encoding failed with input '%v' '%v' '%v'\n", spec, dec, sz)
				}
				//fmt.Printf("encoded stellar = '%v' with input '%v' '%v' '%v'\n", code, spec, dec, sz)
			}
		}
	}
}

func allCodes() []string {
	spectrals := []string{"O", "B", "A", "F", "G", "K", "M", "BD"}
	decimals := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	sizes := []string{"Ia", "Ib", "II", "III", "IV", "V", "VI", "D"}
	allCodes := []string{}
	for _, spec := range spectrals {
		for _, sz := range sizes {
			for _, dec := range decimals {
				allCodes = append(allCodes, encodeStellar(spec, dec, sz))

			}
		}
	}
	return allCodes
}

func pairs() []float64 {
	//return []float64{50, 25, 16, 13, 12, 10, 10, 12, 13, 16, 16, 20, 25} Ib
	//return []float64{30, 20, 14, 11, 10, 8.1, 8.1, 10, 11, 14, 14, 16, 18} // II
	//return []float64{25, 15, 12, 9, 8, 5, 2.5, 3.2, 4, 5, 6.3, 7.4, 9.2} // III
	//return []float64{20, 10, 6, 4, 2.5, 2, 1.75, 2, 2.3} // IV
	//return []float64{18, 6.5, 3.2, 2.1, 1.7, 1.3, 1.04, 0.94, 0.825, 0.570, 0.489, 0.331, 0.215} // V
	//return []float64{0.8, 0.6, 0.528, 0.43, 0.33, 0.154, 0.104, 0.058} // VI
	//LUMA
	//return []float64{27.36, 21.25, 18.09, 16.87, 15.72, 15.03, 16.09, 17.27, 17.65, 18.09, 18.49, 18.95, 19.38} // Ia
	//return []float64{22.80, 14.70, 11.07, 10.40, 9.27, 8.45, 8.84, 9.49, 10.40, 11.95, 14.65, 17.27, 18.49} // Ib
	//return []float64{20.31, 11.68, 6.85, 5.40, 4.95, 4.75, 4.86, 5.22, 5.46, 7.04, 8.24, 11.05, 11.28} // II
	//return []float64{18.09, 9.05, 4.09, 3.08, 2.70, 2.56, 2.66, 2.94, 3.12, 4.23, 4.66, 6.91, 7.20} // III
	//return []float64{16.87, 6.69, 3.53, 2.47, 2.09, 1.86, 1.60, 1.49, 1.47} // IV
	//return []float64{15.38, 6.12, 3.08, 2.00, 1.69, 1.37, 1.05, 0.90, 0.81, 0.53, 0.45, 0.29, 0.18} // V
	return []float64{0.99, 0.75, 0.66, 0.58, 0.40, 0.32, 0.21, 0.09} // VI

}

func strs() []string {
	return []string{
		//"B0",
		//"B1",
		//"B2",
		//"B3",
		//"B4",
		//"B5",
		//"B6",
		//"B7",
		//"B8",
		//"B9",
		//"A0",
		//"A1",
		//"A2",
		//"A3",
		//"A4",
		//"A5",
		//"A6",
		//"A7",
		//"A8",
		//"A9",
		//"F0",
		//"F1",
		//"F2",
		//"F3",
		//"F4",
		"F5",
		"F6",
		"F7",
		"F8",
		"F9",
		"G0",
		"G1",
		"G2",
		"G3",
		"G4",
		"G5",
		"G6",
		"G7",
		"G8",
		"G9",
		"K0",
		"K1",
		"K2",
		"K3",
		"K4",
		"K5",
		"K6",
		"K7",
		"K8",
		"K9",
		"M0",
		"M1",
		"M2",
		"M3",
		"M4",
		"M5",
		"M6",
		"M7",
		"M8",
		"M9",
	}
}

func TestCode(t *testing.T) {
	spectr := []string{"B", "A", "F", "G", "K", "M"}
	dec := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	siz := []string{"Ia", "Ib", "II", "III", "IV", "V", "VI", "D"}
	for _, sp := range spectr {
		for _, d := range dec {
			for _, si := range siz {
				code := sp + d + " " + si
				if si == "D" {
					code = sp + si
				}
				if codeErrorExpected(code) {
					continue
				}
				if !codeValid(code) {
					t.Errorf("code '%v' is invalid!", code)
				}
			}
		}
	}
}

func codeValid(code string) bool {
	checkmap := make(map[string]bool)
	checkmap["B0 Ia"] = true
	checkmap["B1 Ia"] = true
	checkmap["B2 Ia"] = true
	checkmap["B3 Ia"] = true
	checkmap["B4 Ia"] = true
	checkmap["B5 Ia"] = true
	checkmap["B6 Ia"] = true
	checkmap["B7 Ia"] = true
	checkmap["B8 Ia"] = true
	checkmap["B9 Ia"] = true
	checkmap["A0 Ia"] = true
	checkmap["A1 Ia"] = true
	checkmap["A2 Ia"] = true
	checkmap["A3 Ia"] = true
	checkmap["A4 Ia"] = true
	checkmap["A5 Ia"] = true
	checkmap["A6 Ia"] = true
	checkmap["A7 Ia"] = true
	checkmap["A8 Ia"] = true
	checkmap["A9 Ia"] = true
	checkmap["F0 Ia"] = true
	checkmap["F1 Ia"] = true
	checkmap["F2 Ia"] = true
	checkmap["F3 Ia"] = true
	checkmap["F4 Ia"] = true
	checkmap["F5 Ia"] = true
	checkmap["F6 Ia"] = true
	checkmap["F7 Ia"] = true
	checkmap["F8 Ia"] = true
	checkmap["F9 Ia"] = true
	checkmap["G0 Ia"] = true
	checkmap["G1 Ia"] = true
	checkmap["G2 Ia"] = true
	checkmap["G3 Ia"] = true
	checkmap["G4 Ia"] = true
	checkmap["G5 Ia"] = true
	checkmap["G6 Ia"] = true
	checkmap["G7 Ia"] = true
	checkmap["G8 Ia"] = true
	checkmap["G9 Ia"] = true
	checkmap["K0 Ia"] = true
	checkmap["K1 Ia"] = true
	checkmap["K2 Ia"] = true
	checkmap["K3 Ia"] = true
	checkmap["K4 Ia"] = true
	checkmap["K5 Ia"] = true
	checkmap["K6 Ia"] = true
	checkmap["K7 Ia"] = true
	checkmap["K8 Ia"] = true
	checkmap["K9 Ia"] = true
	checkmap["M0 Ia"] = true
	checkmap["M1 Ia"] = true
	checkmap["M2 Ia"] = true
	checkmap["M3 Ia"] = true
	checkmap["M4 Ia"] = true
	checkmap["M5 Ia"] = true
	checkmap["M6 Ia"] = true
	checkmap["M7 Ia"] = true
	checkmap["M8 Ia"] = true
	checkmap["M9 Ia"] = true
	checkmap["B0 Ib"] = true
	checkmap["B1 Ib"] = true
	checkmap["B2 Ib"] = true
	checkmap["B3 Ib"] = true
	checkmap["B4 Ib"] = true
	checkmap["B5 Ib"] = true
	checkmap["B6 Ib"] = true
	checkmap["B7 Ib"] = true
	checkmap["B8 Ib"] = true
	checkmap["B9 Ib"] = true
	checkmap["A0 Ib"] = true
	checkmap["A1 Ib"] = true
	checkmap["A2 Ib"] = true
	checkmap["A3 Ib"] = true
	checkmap["A4 Ib"] = true
	checkmap["A5 Ib"] = true
	checkmap["A6 Ib"] = true
	checkmap["A7 Ib"] = true
	checkmap["A8 Ib"] = true
	checkmap["A9 Ib"] = true
	checkmap["F0 Ib"] = true
	checkmap["F1 Ib"] = true
	checkmap["F2 Ib"] = true
	checkmap["F3 Ib"] = true
	checkmap["F4 Ib"] = true
	checkmap["F5 Ib"] = true
	checkmap["F6 Ib"] = true
	checkmap["F7 Ib"] = true
	checkmap["F8 Ib"] = true
	checkmap["F9 Ib"] = true
	checkmap["G0 Ib"] = true
	checkmap["G1 Ib"] = true
	checkmap["G2 Ib"] = true
	checkmap["G3 Ib"] = true
	checkmap["G4 Ib"] = true
	checkmap["G5 Ib"] = true
	checkmap["G6 Ib"] = true
	checkmap["G7 Ib"] = true
	checkmap["G8 Ib"] = true
	checkmap["G9 Ib"] = true
	checkmap["K0 Ib"] = true
	checkmap["K1 Ib"] = true
	checkmap["K2 Ib"] = true
	checkmap["K3 Ib"] = true
	checkmap["K4 Ib"] = true
	checkmap["K5 Ib"] = true
	checkmap["K6 Ib"] = true
	checkmap["K7 Ib"] = true
	checkmap["K8 Ib"] = true
	checkmap["K9 Ib"] = true
	checkmap["M0 Ib"] = true
	checkmap["M1 Ib"] = true
	checkmap["M2 Ib"] = true
	checkmap["M3 Ib"] = true
	checkmap["M4 Ib"] = true
	checkmap["M5 Ib"] = true
	checkmap["M6 Ib"] = true
	checkmap["M7 Ib"] = true
	checkmap["M8 Ib"] = true
	checkmap["M9 Ib"] = true
	checkmap["B0 II"] = true
	checkmap["B1 II"] = true
	checkmap["B2 II"] = true
	checkmap["B3 II"] = true
	checkmap["B4 II"] = true
	checkmap["B5 II"] = true
	checkmap["B6 II"] = true
	checkmap["B7 II"] = true
	checkmap["B8 II"] = true
	checkmap["B9 II"] = true
	checkmap["A0 II"] = true
	checkmap["A1 II"] = true
	checkmap["A2 II"] = true
	checkmap["A3 II"] = true
	checkmap["A4 II"] = true
	checkmap["A5 II"] = true
	checkmap["A6 II"] = true
	checkmap["A7 II"] = true
	checkmap["A8 II"] = true
	checkmap["A9 II"] = true
	checkmap["F0 II"] = true
	checkmap["F1 II"] = true
	checkmap["F2 II"] = true
	checkmap["F3 II"] = true
	checkmap["F4 II"] = true
	checkmap["F5 II"] = true
	checkmap["F6 II"] = true
	checkmap["F7 II"] = true
	checkmap["F8 II"] = true
	checkmap["F9 II"] = true
	checkmap["G0 II"] = true
	checkmap["G1 II"] = true
	checkmap["G2 II"] = true
	checkmap["G3 II"] = true
	checkmap["G4 II"] = true
	checkmap["G5 II"] = true
	checkmap["G6 II"] = true
	checkmap["G7 II"] = true
	checkmap["G8 II"] = true
	checkmap["G9 II"] = true
	checkmap["K0 II"] = true
	checkmap["K1 II"] = true
	checkmap["K2 II"] = true
	checkmap["K3 II"] = true
	checkmap["K4 II"] = true
	checkmap["K5 II"] = true
	checkmap["K6 II"] = true
	checkmap["K7 II"] = true
	checkmap["K8 II"] = true
	checkmap["K9 II"] = true
	checkmap["M0 II"] = true
	checkmap["M1 II"] = true
	checkmap["M2 II"] = true
	checkmap["M3 II"] = true
	checkmap["M4 II"] = true
	checkmap["M5 II"] = true
	checkmap["M6 II"] = true
	checkmap["M7 II"] = true
	checkmap["M8 II"] = true
	checkmap["M9 II"] = true
	checkmap["B0 III"] = true
	checkmap["B1 III"] = true
	checkmap["B2 III"] = true
	checkmap["B3 III"] = true
	checkmap["B4 III"] = true
	checkmap["B5 III"] = true
	checkmap["B6 III"] = true
	checkmap["B7 III"] = true
	checkmap["B8 III"] = true
	checkmap["B9 III"] = true
	checkmap["A0 III"] = true
	checkmap["A1 III"] = true
	checkmap["A2 III"] = true
	checkmap["A3 III"] = true
	checkmap["A4 III"] = true
	checkmap["A5 III"] = true
	checkmap["A6 III"] = true
	checkmap["A7 III"] = true
	checkmap["A8 III"] = true
	checkmap["A9 III"] = true
	checkmap["F0 III"] = true
	checkmap["F1 III"] = true
	checkmap["F2 III"] = true
	checkmap["F3 III"] = true
	checkmap["F4 III"] = true
	checkmap["F5 III"] = true
	checkmap["F6 III"] = true
	checkmap["F7 III"] = true
	checkmap["F8 III"] = true
	checkmap["F9 III"] = true
	checkmap["G0 III"] = true
	checkmap["G1 III"] = true
	checkmap["G2 III"] = true
	checkmap["G3 III"] = true
	checkmap["G4 III"] = true
	checkmap["G5 III"] = true
	checkmap["G6 III"] = true
	checkmap["G7 III"] = true
	checkmap["G8 III"] = true
	checkmap["G9 III"] = true
	checkmap["K0 III"] = true
	checkmap["K1 III"] = true
	checkmap["K2 III"] = true
	checkmap["K3 III"] = true
	checkmap["K4 III"] = true
	checkmap["K5 III"] = true
	checkmap["K6 III"] = true
	checkmap["K7 III"] = true
	checkmap["K8 III"] = true
	checkmap["K9 III"] = true
	checkmap["M0 III"] = true
	checkmap["M1 III"] = true
	checkmap["M2 III"] = true
	checkmap["M3 III"] = true
	checkmap["M4 III"] = true
	checkmap["M5 III"] = true
	checkmap["M6 III"] = true
	checkmap["M7 III"] = true
	checkmap["M8 III"] = true
	checkmap["M9 III"] = true
	checkmap["B0 IV"] = true
	checkmap["B1 IV"] = true
	checkmap["B2 IV"] = true
	checkmap["B3 IV"] = true
	checkmap["B4 IV"] = true
	checkmap["B5 IV"] = true
	checkmap["B6 IV"] = true
	checkmap["B7 IV"] = true
	checkmap["B8 IV"] = true
	checkmap["B9 IV"] = true
	checkmap["A0 IV"] = true
	checkmap["A1 IV"] = true
	checkmap["A2 IV"] = true
	checkmap["A3 IV"] = true
	checkmap["A4 IV"] = true
	checkmap["A5 IV"] = true
	checkmap["A6 IV"] = true
	checkmap["A7 IV"] = true
	checkmap["A8 IV"] = true
	checkmap["A9 IV"] = true
	checkmap["F0 IV"] = true
	checkmap["F1 IV"] = true
	checkmap["F2 IV"] = true
	checkmap["F3 IV"] = true
	checkmap["F4 IV"] = true
	checkmap["F5 IV"] = true
	checkmap["F6 IV"] = true
	checkmap["F7 IV"] = true
	checkmap["F8 IV"] = true
	checkmap["F9 IV"] = true
	checkmap["G0 IV"] = true
	checkmap["G1 IV"] = true
	checkmap["G2 IV"] = true
	checkmap["G3 IV"] = true
	checkmap["G4 IV"] = true
	checkmap["G5 IV"] = true
	checkmap["G6 IV"] = true
	checkmap["G7 IV"] = true
	checkmap["G8 IV"] = true
	checkmap["G9 IV"] = true
	checkmap["K0 IV"] = true
	checkmap["K1 IV"] = true
	checkmap["K2 IV"] = true
	checkmap["K3 IV"] = true
	checkmap["K4 IV"] = true
	checkmap["B0 V"] = true
	checkmap["B1 V"] = true
	checkmap["B2 V"] = true
	checkmap["B3 V"] = true
	checkmap["B4 V"] = true
	checkmap["B5 V"] = true
	checkmap["B6 V"] = true
	checkmap["B7 V"] = true
	checkmap["B8 V"] = true
	checkmap["B9 V"] = true
	checkmap["A0 V"] = true
	checkmap["A1 V"] = true
	checkmap["A2 V"] = true
	checkmap["A3 V"] = true
	checkmap["A4 V"] = true
	checkmap["A5 V"] = true
	checkmap["A6 V"] = true
	checkmap["A7 V"] = true
	checkmap["A8 V"] = true
	checkmap["A9 V"] = true
	checkmap["F0 V"] = true
	checkmap["F1 V"] = true
	checkmap["F2 V"] = true
	checkmap["F3 V"] = true
	checkmap["F4 V"] = true
	checkmap["F5 V"] = true
	checkmap["F6 V"] = true
	checkmap["F7 V"] = true
	checkmap["F8 V"] = true
	checkmap["F9 V"] = true
	checkmap["G0 V"] = true
	checkmap["G1 V"] = true
	checkmap["G2 V"] = true
	checkmap["G3 V"] = true
	checkmap["G4 V"] = true
	checkmap["G5 V"] = true
	checkmap["G6 V"] = true
	checkmap["G7 V"] = true
	checkmap["G8 V"] = true
	checkmap["G9 V"] = true
	checkmap["K0 V"] = true
	checkmap["K1 V"] = true
	checkmap["K2 V"] = true
	checkmap["K3 V"] = true
	checkmap["K4 V"] = true
	checkmap["K5 V"] = true
	checkmap["K6 V"] = true
	checkmap["K7 V"] = true
	checkmap["K8 V"] = true
	checkmap["K9 V"] = true
	checkmap["M0 V"] = true
	checkmap["M1 V"] = true
	checkmap["M2 V"] = true
	checkmap["M3 V"] = true
	checkmap["M4 V"] = true
	checkmap["M5 V"] = true
	checkmap["M6 V"] = true
	checkmap["M7 V"] = true
	checkmap["M8 V"] = true
	checkmap["M9 V"] = true
	checkmap["F5 VI"] = true
	checkmap["F6 VI"] = true
	checkmap["F7 VI"] = true
	checkmap["F8 VI"] = true
	checkmap["F9 VI"] = true
	checkmap["G0 VI"] = true
	checkmap["G1 VI"] = true
	checkmap["G2 VI"] = true
	checkmap["G3 VI"] = true
	checkmap["G4 VI"] = true
	checkmap["G5 VI"] = true
	checkmap["G6 VI"] = true
	checkmap["G7 VI"] = true
	checkmap["G8 VI"] = true
	checkmap["G9 VI"] = true
	checkmap["K0 VI"] = true
	checkmap["K1 VI"] = true
	checkmap["K2 VI"] = true
	checkmap["K3 VI"] = true
	checkmap["K4 VI"] = true
	checkmap["K5 VI"] = true
	checkmap["K6 VI"] = true
	checkmap["K7 VI"] = true
	checkmap["K8 VI"] = true
	checkmap["K9 VI"] = true
	checkmap["M0 VI"] = true
	checkmap["M1 VI"] = true
	checkmap["M2 VI"] = true
	checkmap["M3 VI"] = true
	checkmap["M4 VI"] = true
	checkmap["M5 VI"] = true
	checkmap["M6 VI"] = true
	checkmap["M7 VI"] = true
	checkmap["M8 VI"] = true
	checkmap["M9 VI"] = true
	checkmap["BD"] = true
	checkmap["AD"] = true
	checkmap["FD"] = true
	checkmap["GD"] = true
	checkmap["KD"] = true
	checkmap["MD"] = true
	return checkmap[code]
}

func codeErrorExpected(code string) bool {
	errMap := make(map[string]bool)
	errMap["B0 VI"] = true
	errMap["B1 VI"] = true
	errMap["B2 VI"] = true
	errMap["B3 VI"] = true
	errMap["B4 VI"] = true
	errMap["B5 VI"] = true
	errMap["B6 VI"] = true
	errMap["B7 VI"] = true
	errMap["B8 VI"] = true
	errMap["B9 VI"] = true
	errMap["A0 VI"] = true
	errMap["A1 VI"] = true
	errMap["A2 VI"] = true
	errMap["A3 VI"] = true
	errMap["A4 VI"] = true
	errMap["A5 VI"] = true
	errMap["A6 VI"] = true
	errMap["A7 VI"] = true
	errMap["A8 VI"] = true
	errMap["A9 VI"] = true
	errMap["F0 VI"] = true
	errMap["F1 VI"] = true
	errMap["F2 VI"] = true
	errMap["F3 VI"] = true
	errMap["F4 VI"] = true
	//errMap["K1 IV"] = true
	//errMap["K2 IV"] = true
	//errMap["K3 IV"] = true
	//errMap["K4 IV"] = true
	errMap["K5 IV"] = true
	errMap["K6 IV"] = true
	errMap["K7 IV"] = true
	errMap["K8 IV"] = true
	errMap["K9 IV"] = true
	errMap["M0 IV"] = true
	errMap["M1 IV"] = true
	errMap["M2 IV"] = true
	errMap["M3 IV"] = true
	errMap["M4 IV"] = true
	errMap["M5 IV"] = true
	errMap["M6 IV"] = true
	errMap["M7 IV"] = true
	errMap["M8 IV"] = true
	errMap["M9 IV"] = true
	return errMap[code]
}
