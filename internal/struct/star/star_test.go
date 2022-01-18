package star

import (
	"fmt"
	"strings"
	"testing"
)

func TestMassAndLuma(t *testing.T) {
	tot := 0
	cinv := 0
	for _, code := range allCodes() {
		tot++
		l := baseStellarLuminocity(code)
		m := baseStellarMass(code)
		switch {
		case !CodeValid(code):
			cinv++
			continue
		case l == 0:
			t.Errorf("code '%v': luma is not expected to be %v", code, l)
		case m == 0:
			t.Errorf("code '%v': mass is not expected to be %v", code, m)
		}
	}
	fmt.Printf("Total Tests: %v\nInput Code Invalid: %v\n \n", tot, cinv)
}

func TestStar(t *testing.T) {
	codes := allCodes()
	codes = append(codes, "Fus", "")
	categ := []int{Category_Primary, Category_PrimaryCompanion, Category_Close, Category_CloseCompanion, Category_Near, Category_NearCompanion, Category_Far, Category_FarCompanion, 11, UNDEFINED}
	tst := 0
	vld := 0
	ice := 0
	sce := 0
	for _, code := range codes {
		for _, cat := range categ {
			tst++
			_, err := New("Kimeria", code, cat)
			switch {
			default:
				t.Errorf("unknown error: test %v (%v,%v)", tst, code, categ)
			case err == nil:
				vld++
				continue
			case strings.Contains(err.Error(), "input code invalid"):
				ice++
				continue
			case strings.Contains(err.Error(), "star category invalid"):
				sce++
				continue
			}
		}
	}
	fmt.Printf("Total tests: %v\nValid: %v\nCodeError: %v\nCategoryError %v\n \n", tst, vld, ice, sce)
}

func TestStellarEncode(t *testing.T) {
	spectrals := []string{"O", "B", "A", "F", "G", "K", "M", "BD", "", "f"}
	decimals := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "", "f"}
	sizes := []string{"Ia", "Ib", "II", "III", "IV", "V", "VI", "D", "", "f"}
	tot := 0
	err := 0
	for _, spec := range spectrals {
		for _, sz := range sizes {
			for _, dec := range decimals {
				tot++
				code := EncodeStellar(spec, dec, sz)
				if code == "error" {
					err++
					continue
					t.Errorf("encoding failed with input '%v' '%v' '%v'\n", spec, dec, sz)
				}
			}
		}
	}
	fmt.Printf("Total Tests: %v\nErrors: %v\n \n", tot, err)
}

func TestStellarDecode(t *testing.T) {
	codes := allCodes()
	codes = append(codes, " ")
	codes = append(codes, "sdgf")
	codes = append(codes, "12345678")
	codes = append(codes, "G2 V ")
	tot := 0
	errNum := 0
	correct := 0
	for _, code := range codes {
		tot++
		_, _, _, err := DecodeStellar(code)
		switch {
		default:
			t.Errorf("unhandled err: %v", err.Error())
		case err == nil:
			correct++
		case strings.Contains(err.Error(), "code lenght incorrect"):
			errNum++
		case strings.Contains(err.Error(), "code input incorrect"):
			errNum++
		}
	}
	fmt.Printf("Total Tests: %v\nValid: %v\nErrors: %v\n \n", tot, correct, errNum)
}

func allCodes() []string {
	spectrals := []string{"O", "B", "A", "F", "G", "K", "M", "BD"}
	decimals := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	sizes := []string{"Ia", "Ib", "II", "III", "IV", "V", "VI", "D"}
	allCodes := []string{}
	for _, sz := range sizes {
		for _, spec := range spectrals {

			for _, dec := range decimals {
				allCodes = appendUnique(allCodes, EncodeStellar(spec, dec, sz))
			}
		}
	}
	return allCodes
}

func appendUnique(sl []string, s string) []string {
	for _, v := range sl {
		if v == s {
			return sl
		}
	}
	sl = append(sl, s)
	return sl
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
				code := EncodeStellar(sp, d, si)
				if codeErrorExpected(code) {
					continue
				}
				if !CodeValid(code) {
					t.Errorf("code '%v' is invalid!", code)
				}
			}
		}
	}
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

func TestParse(t *testing.T) {
	//переделать
	// for i, tl := range parcingLines() {
	// stCodes, err := ParseStellar(tl)
	// switch {
	// default:
	// t.Errorf("Test %v: input %v - err: %v", i, tl, err)
	// case err == nil:
	// fmt.Println("Tst ==", tl)
	// comp := rollSystemComposition("Test Name", len(stCodes))
	// for i, code := range stCodes {
	// st, _ := New("Test Name", code, comp[i])
	// fmt.Println(st)
	// }
	// }
	// }
}

func parcingLines() []string {
	return []string{
		//"F1 V M9 II M7 V",
		//"G2   V ",
		//"G2 V",
		"G3 V G8 V K1 V M3 VI M3 V G8 VI B2 Ia DK",
		//"M0 V M0 V",
	}
}
