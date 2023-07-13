package stars

import (
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func massOf(st star, dice *dice.Dicepool) float64 {
	averageMass := averageMassMap(st)
	flux := dice.Flux()
	variance := (averageMass / 100) * (4 * flux)
	switch st.class {
	case classIa, classIb, classII, classIII:
		flux = (dice.Flux() * 10) + dice.Flux()
		variance = (averageMass / 100) * flux
	}
	return float64(averageMass+variance) / 1000000
}

func temperatureOf(st star, dice *dice.Dicepool) int {
	temp := averageTempMap(st)
	flux := dice.Flux()
	variance := (temp / 200) * flux
	return temp + variance
}

func massAbbreviation(st star) string {
	abb := "[UNKNOWN]"

	switch st.sttype {
	case typeO, typeB, typeA, typeF, typeG, typeK, typeM:
		abb = strings.TrimPrefix(st.sttype, "Type ") + st.subtype + strings.TrimPrefix(st.class, "Class")
	}

	return abb
}

func averageMassMap(st star) int {
	massMap := make(map[string]int)
	//KBIYTT
	massMap["O0 Ia"] = 200000000
	massMap["O5 Ia"] = 80000000
	massMap["B0 Ia"] = 60000000
	massMap["B5 Ia"] = 30000000
	massMap["A0 Ia"] = 20000000
	massMap["A5 Ia"] = 15000000
	massMap["F0 Ia"] = 13000000
	massMap["F5 Ia"] = 12000000
	massMap["G0 Ia"] = 12000000
	massMap["G5 Ia"] = 13000000
	massMap["K0 Ia"] = 14000000
	massMap["K5 Ia"] = 18000000
	massMap["M0 Ia"] = 20000000
	massMap["M5 Ia"] = 25000000
	massMap["M9 Ia"] = 30000000

	massMap["O1 Ia"] = massMap["O5 Ia"] + ((massMap["O0 Ia"] - massMap["O5 Ia"]) / 5 * 4)
	massMap["O2 Ia"] = massMap["O5 Ia"] + ((massMap["O0 Ia"] - massMap["O5 Ia"]) / 5 * 3)
	massMap["O3 Ia"] = massMap["O5 Ia"] + ((massMap["O0 Ia"] - massMap["O5 Ia"]) / 5 * 2)
	massMap["O4 Ia"] = massMap["O5 Ia"] + ((massMap["O0 Ia"] - massMap["O5 Ia"]) / 5 * 1)
	massMap["B1 Ia"] = massMap["B5 Ia"] + ((massMap["B0 Ia"] - massMap["B5 Ia"]) / 5 * 4)
	massMap["B2 Ia"] = massMap["B5 Ia"] + ((massMap["B0 Ia"] - massMap["B5 Ia"]) / 5 * 3)
	massMap["B3 Ia"] = massMap["B5 Ia"] + ((massMap["B0 Ia"] - massMap["B5 Ia"]) / 5 * 2)
	massMap["B4 Ia"] = massMap["B5 Ia"] + ((massMap["B0 Ia"] - massMap["B5 Ia"]) / 5 * 1)
	massMap["A1 Ia"] = massMap["A5 Ia"] + ((massMap["A0 Ia"] - massMap["A5 Ia"]) / 5 * 4)
	massMap["A2 Ia"] = massMap["A5 Ia"] + ((massMap["A0 Ia"] - massMap["A5 Ia"]) / 5 * 3)
	massMap["A3 Ia"] = massMap["A5 Ia"] + ((massMap["A0 Ia"] - massMap["A5 Ia"]) / 5 * 2)
	massMap["A4 Ia"] = massMap["A5 Ia"] + ((massMap["A0 Ia"] - massMap["A5 Ia"]) / 5 * 1)
	massMap["F1 Ia"] = massMap["F5 Ia"] + ((massMap["F0 Ia"] - massMap["F5 Ia"]) / 5 * 4)
	massMap["F2 Ia"] = massMap["F5 Ia"] + ((massMap["F0 Ia"] - massMap["F5 Ia"]) / 5 * 3)
	massMap["F3 Ia"] = massMap["F5 Ia"] + ((massMap["F0 Ia"] - massMap["F5 Ia"]) / 5 * 2)
	massMap["F4 Ia"] = massMap["F5 Ia"] + ((massMap["F0 Ia"] - massMap["F5 Ia"]) / 5 * 1)
	massMap["G1 Ia"] = massMap["G5 Ia"] + ((massMap["G0 Ia"] - massMap["G5 Ia"]) / 5 * 4)
	massMap["G2 Ia"] = massMap["G5 Ia"] + ((massMap["G0 Ia"] - massMap["G5 Ia"]) / 5 * 3)
	massMap["G3 Ia"] = massMap["G5 Ia"] + ((massMap["G0 Ia"] - massMap["G5 Ia"]) / 5 * 2)
	massMap["G4 Ia"] = massMap["G5 Ia"] + ((massMap["G0 Ia"] - massMap["G5 Ia"]) / 5 * 1)
	massMap["K1 Ia"] = massMap["K5 Ia"] + ((massMap["K0 Ia"] - massMap["K5 Ia"]) / 5 * 4)
	massMap["K2 Ia"] = massMap["K5 Ia"] + ((massMap["K0 Ia"] - massMap["K5 Ia"]) / 5 * 3)
	massMap["K3 Ia"] = massMap["K5 Ia"] + ((massMap["K0 Ia"] - massMap["K5 Ia"]) / 5 * 2)
	massMap["K4 Ia"] = massMap["K5 Ia"] + ((massMap["K0 Ia"] - massMap["K5 Ia"]) / 5 * 1)
	massMap["M1 Ia"] = massMap["M5 Ia"] + ((massMap["M0 Ia"] - massMap["M5 Ia"]) / 5 * 4)
	massMap["M2 Ia"] = massMap["M5 Ia"] + ((massMap["M0 Ia"] - massMap["M5 Ia"]) / 5 * 3)
	massMap["M3 Ia"] = massMap["M5 Ia"] + ((massMap["M0 Ia"] - massMap["M5 Ia"]) / 5 * 2)
	massMap["M4 Ia"] = massMap["M5 Ia"] + ((massMap["M0 Ia"] - massMap["M5 Ia"]) / 5 * 1)

	massMap["O6 Ia"] = massMap["B0 Ia"] + ((massMap["O5 Ia"] - massMap["B0 Ia"]) / 5 * 4)
	massMap["O7 Ia"] = massMap["B0 Ia"] + ((massMap["O5 Ia"] - massMap["B0 Ia"]) / 5 * 3)
	massMap["O8 Ia"] = massMap["B0 Ia"] + ((massMap["O5 Ia"] - massMap["B0 Ia"]) / 5 * 2)
	massMap["O9 Ia"] = massMap["B0 Ia"] + ((massMap["O5 Ia"] - massMap["B0 Ia"]) / 5 * 1)
	massMap["B6 Ia"] = massMap["A0 Ia"] + ((massMap["B5 Ia"] - massMap["A0 Ia"]) / 5 * 4)
	massMap["B7 Ia"] = massMap["A0 Ia"] + ((massMap["B5 Ia"] - massMap["A0 Ia"]) / 5 * 3)
	massMap["B8 Ia"] = massMap["A0 Ia"] + ((massMap["B5 Ia"] - massMap["A0 Ia"]) / 5 * 2)
	massMap["B9 Ia"] = massMap["A0 Ia"] + ((massMap["B5 Ia"] - massMap["A0 Ia"]) / 5 * 1)
	massMap["A6 Ia"] = massMap["F0 Ia"] + ((massMap["A5 Ia"] - massMap["F0 Ia"]) / 5 * 4)
	massMap["A7 Ia"] = massMap["F0 Ia"] + ((massMap["A5 Ia"] - massMap["F0 Ia"]) / 5 * 3)
	massMap["A8 Ia"] = massMap["F0 Ia"] + ((massMap["A5 Ia"] - massMap["F0 Ia"]) / 5 * 2)
	massMap["A9 Ia"] = massMap["F0 Ia"] + ((massMap["A5 Ia"] - massMap["F0 Ia"]) / 5 * 1)
	massMap["F6 Ia"] = massMap["G0 Ia"] + ((massMap["F5 Ia"] - massMap["G0 Ia"]) / 5 * 4)
	massMap["F7 Ia"] = massMap["G0 Ia"] + ((massMap["F5 Ia"] - massMap["G0 Ia"]) / 5 * 3)
	massMap["F8 Ia"] = massMap["G0 Ia"] + ((massMap["F5 Ia"] - massMap["G0 Ia"]) / 5 * 2)
	massMap["F9 Ia"] = massMap["G0 Ia"] + ((massMap["F5 Ia"] - massMap["G0 Ia"]) / 5 * 1)
	massMap["G6 Ia"] = massMap["K0 Ia"] + ((massMap["G5 Ia"] - massMap["K0 Ia"]) / 5 * 4)
	massMap["G7 Ia"] = massMap["K0 Ia"] + ((massMap["G5 Ia"] - massMap["K0 Ia"]) / 5 * 3)
	massMap["G8 Ia"] = massMap["K0 Ia"] + ((massMap["G5 Ia"] - massMap["K0 Ia"]) / 5 * 2)
	massMap["G9 Ia"] = massMap["K0 Ia"] + ((massMap["G5 Ia"] - massMap["K0 Ia"]) / 5 * 1)
	massMap["K6 Ia"] = massMap["M0 Ia"] + ((massMap["K5 Ia"] - massMap["M0 Ia"]) / 5 * 4)
	massMap["K7 Ia"] = massMap["M0 Ia"] + ((massMap["K5 Ia"] - massMap["M0 Ia"]) / 5 * 3)
	massMap["K8 Ia"] = massMap["M0 Ia"] + ((massMap["K5 Ia"] - massMap["M0 Ia"]) / 5 * 2)
	massMap["K9 Ia"] = massMap["M0 Ia"] + ((massMap["K5 Ia"] - massMap["M0 Ia"]) / 5 * 1)
	massMap["M6 Ia"] = massMap["M9 Ia"] + ((massMap["M5 Ia"] - massMap["M9 Ia"]) / 4 * 3)
	massMap["M7 Ia"] = massMap["M9 Ia"] + ((massMap["M5 Ia"] - massMap["M9 Ia"]) / 4 * 2)
	massMap["M8 Ia"] = massMap["M9 Ia"] + ((massMap["M5 Ia"] - massMap["M9 Ia"]) / 4 * 1)

	////////////////////Ib

	massMap["O0 Ib"] = 150000000
	massMap["O5 Ib"] = 60000000
	massMap["B0 Ib"] = 40000000
	massMap["B5 Ib"] = 25000000
	massMap["A0 Ib"] = 15000000
	massMap["A5 Ib"] = 13000000
	massMap["F0 Ib"] = 12000000
	massMap["F5 Ib"] = 10000000
	massMap["G0 Ib"] = 10000000
	massMap["G5 Ib"] = 11000000
	massMap["K0 Ib"] = 12000000
	massMap["K5 Ib"] = 13000000
	massMap["M0 Ib"] = 15000000
	massMap["M5 Ib"] = 20000000
	massMap["M9 Ib"] = 25000000

	massMap["O1 Ib"] = massMap["O5 Ib"] + ((massMap["O0 Ib"] - massMap["O5 Ib"]) / 5 * 4)
	massMap["O2 Ib"] = massMap["O5 Ib"] + ((massMap["O0 Ib"] - massMap["O5 Ib"]) / 5 * 3)
	massMap["O3 Ib"] = massMap["O5 Ib"] + ((massMap["O0 Ib"] - massMap["O5 Ib"]) / 5 * 2)
	massMap["O4 Ib"] = massMap["O5 Ib"] + ((massMap["O0 Ib"] - massMap["O5 Ib"]) / 5 * 1)
	massMap["B1 Ib"] = massMap["B5 Ib"] + ((massMap["B0 Ib"] - massMap["B5 Ib"]) / 5 * 4)
	massMap["B2 Ib"] = massMap["B5 Ib"] + ((massMap["B0 Ib"] - massMap["B5 Ib"]) / 5 * 3)
	massMap["B3 Ib"] = massMap["B5 Ib"] + ((massMap["B0 Ib"] - massMap["B5 Ib"]) / 5 * 2)
	massMap["B4 Ib"] = massMap["B5 Ib"] + ((massMap["B0 Ib"] - massMap["B5 Ib"]) / 5 * 1)
	massMap["A1 Ib"] = massMap["A5 Ib"] + ((massMap["A0 Ib"] - massMap["A5 Ib"]) / 5 * 4)
	massMap["A2 Ib"] = massMap["A5 Ib"] + ((massMap["A0 Ib"] - massMap["A5 Ib"]) / 5 * 3)
	massMap["A3 Ib"] = massMap["A5 Ib"] + ((massMap["A0 Ib"] - massMap["A5 Ib"]) / 5 * 2)
	massMap["A4 Ib"] = massMap["A5 Ib"] + ((massMap["A0 Ib"] - massMap["A5 Ib"]) / 5 * 1)
	massMap["F1 Ib"] = massMap["F5 Ib"] + ((massMap["F0 Ib"] - massMap["F5 Ib"]) / 5 * 4)
	massMap["F2 Ib"] = massMap["F5 Ib"] + ((massMap["F0 Ib"] - massMap["F5 Ib"]) / 5 * 3)
	massMap["F3 Ib"] = massMap["F5 Ib"] + ((massMap["F0 Ib"] - massMap["F5 Ib"]) / 5 * 2)
	massMap["F4 Ib"] = massMap["F5 Ib"] + ((massMap["F0 Ib"] - massMap["F5 Ib"]) / 5 * 1)
	massMap["G1 Ib"] = massMap["G5 Ib"] + ((massMap["G0 Ib"] - massMap["G5 Ib"]) / 5 * 4)
	massMap["G2 Ib"] = massMap["G5 Ib"] + ((massMap["G0 Ib"] - massMap["G5 Ib"]) / 5 * 3)
	massMap["G3 Ib"] = massMap["G5 Ib"] + ((massMap["G0 Ib"] - massMap["G5 Ib"]) / 5 * 2)
	massMap["G4 Ib"] = massMap["G5 Ib"] + ((massMap["G0 Ib"] - massMap["G5 Ib"]) / 5 * 1)
	massMap["K1 Ib"] = massMap["K5 Ib"] + ((massMap["K0 Ib"] - massMap["K5 Ib"]) / 5 * 4)
	massMap["K2 Ib"] = massMap["K5 Ib"] + ((massMap["K0 Ib"] - massMap["K5 Ib"]) / 5 * 3)
	massMap["K3 Ib"] = massMap["K5 Ib"] + ((massMap["K0 Ib"] - massMap["K5 Ib"]) / 5 * 2)
	massMap["K4 Ib"] = massMap["K5 Ib"] + ((massMap["K0 Ib"] - massMap["K5 Ib"]) / 5 * 1)
	massMap["M1 Ib"] = massMap["M5 Ib"] + ((massMap["M0 Ib"] - massMap["M5 Ib"]) / 5 * 4)
	massMap["M2 Ib"] = massMap["M5 Ib"] + ((massMap["M0 Ib"] - massMap["M5 Ib"]) / 5 * 3)
	massMap["M3 Ib"] = massMap["M5 Ib"] + ((massMap["M0 Ib"] - massMap["M5 Ib"]) / 5 * 2)
	massMap["M4 Ib"] = massMap["M5 Ib"] + ((massMap["M0 Ib"] - massMap["M5 Ib"]) / 5 * 1)

	massMap["O6 Ib"] = massMap["B0 Ib"] + ((massMap["O5 Ib"] - massMap["B0 Ib"]) / 5 * 4)
	massMap["O7 Ib"] = massMap["B0 Ib"] + ((massMap["O5 Ib"] - massMap["B0 Ib"]) / 5 * 3)
	massMap["O8 Ib"] = massMap["B0 Ib"] + ((massMap["O5 Ib"] - massMap["B0 Ib"]) / 5 * 2)
	massMap["O9 Ib"] = massMap["B0 Ib"] + ((massMap["O5 Ib"] - massMap["B0 Ib"]) / 5 * 1)
	massMap["B6 Ib"] = massMap["A0 Ib"] + ((massMap["B5 Ib"] - massMap["A0 Ib"]) / 5 * 4)
	massMap["B7 Ib"] = massMap["A0 Ib"] + ((massMap["B5 Ib"] - massMap["A0 Ib"]) / 5 * 3)
	massMap["B8 Ib"] = massMap["A0 Ib"] + ((massMap["B5 Ib"] - massMap["A0 Ib"]) / 5 * 2)
	massMap["B9 Ib"] = massMap["A0 Ib"] + ((massMap["B5 Ib"] - massMap["A0 Ib"]) / 5 * 1)
	massMap["A6 Ib"] = massMap["F0 Ib"] + ((massMap["A5 Ib"] - massMap["F0 Ib"]) / 5 * 4)
	massMap["A7 Ib"] = massMap["F0 Ib"] + ((massMap["A5 Ib"] - massMap["F0 Ib"]) / 5 * 3)
	massMap["A8 Ib"] = massMap["F0 Ib"] + ((massMap["A5 Ib"] - massMap["F0 Ib"]) / 5 * 2)
	massMap["A9 Ib"] = massMap["F0 Ib"] + ((massMap["A5 Ib"] - massMap["F0 Ib"]) / 5 * 1)
	massMap["F6 Ib"] = massMap["G0 Ib"] + ((massMap["F5 Ib"] - massMap["G0 Ib"]) / 5 * 4)
	massMap["F7 Ib"] = massMap["G0 Ib"] + ((massMap["F5 Ib"] - massMap["G0 Ib"]) / 5 * 3)
	massMap["F8 Ib"] = massMap["G0 Ib"] + ((massMap["F5 Ib"] - massMap["G0 Ib"]) / 5 * 2)
	massMap["F9 Ib"] = massMap["G0 Ib"] + ((massMap["F5 Ib"] - massMap["G0 Ib"]) / 5 * 1)
	massMap["G6 Ib"] = massMap["K0 Ib"] + ((massMap["G5 Ib"] - massMap["K0 Ib"]) / 5 * 4)
	massMap["G7 Ib"] = massMap["K0 Ib"] + ((massMap["G5 Ib"] - massMap["K0 Ib"]) / 5 * 3)
	massMap["G8 Ib"] = massMap["K0 Ib"] + ((massMap["G5 Ib"] - massMap["K0 Ib"]) / 5 * 2)
	massMap["G9 Ib"] = massMap["K0 Ib"] + ((massMap["G5 Ib"] - massMap["K0 Ib"]) / 5 * 1)
	massMap["K6 Ib"] = massMap["M0 Ib"] + ((massMap["K5 Ib"] - massMap["M0 Ib"]) / 5 * 4)
	massMap["K7 Ib"] = massMap["M0 Ib"] + ((massMap["K5 Ib"] - massMap["M0 Ib"]) / 5 * 3)
	massMap["K8 Ib"] = massMap["M0 Ib"] + ((massMap["K5 Ib"] - massMap["M0 Ib"]) / 5 * 2)
	massMap["K9 Ib"] = massMap["M0 Ib"] + ((massMap["K5 Ib"] - massMap["M0 Ib"]) / 5 * 1)
	massMap["M6 Ib"] = massMap["M9 Ib"] + ((massMap["M5 Ib"] - massMap["M9 Ib"]) / 4 * 3)
	massMap["M7 Ib"] = massMap["M9 Ib"] + ((massMap["M5 Ib"] - massMap["M9 Ib"]) / 4 * 2)
	massMap["M8 Ib"] = massMap["M9 Ib"] + ((massMap["M5 Ib"] - massMap["M9 Ib"]) / 4 * 1)

	////////////////////II

	massMap["O0 II"] = 130000000
	massMap["O5 II"] = 40000000
	massMap["B0 II"] = 30000000
	massMap["B5 II"] = 20000000
	massMap["A0 II"] = 14000000
	massMap["A5 II"] = 11000000
	massMap["F0 II"] = 10000000
	massMap["F5 II"] = 8000000
	massMap["G0 II"] = 8000000
	massMap["G5 II"] = 10000000
	massMap["K0 II"] = 10000000
	massMap["K5 II"] = 12000000
	massMap["M0 II"] = 14000000
	massMap["M5 II"] = 16000000
	massMap["M9 II"] = 18000000

	massMap["O1 II"] = massMap["O5 II"] + ((massMap["O0 II"] - massMap["O5 II"]) / 5 * 4)
	massMap["O2 II"] = massMap["O5 II"] + ((massMap["O0 II"] - massMap["O5 II"]) / 5 * 3)
	massMap["O3 II"] = massMap["O5 II"] + ((massMap["O0 II"] - massMap["O5 II"]) / 5 * 2)
	massMap["O4 II"] = massMap["O5 II"] + ((massMap["O0 II"] - massMap["O5 II"]) / 5 * 1)
	massMap["B1 II"] = massMap["B5 II"] + ((massMap["B0 II"] - massMap["B5 II"]) / 5 * 4)
	massMap["B2 II"] = massMap["B5 II"] + ((massMap["B0 II"] - massMap["B5 II"]) / 5 * 3)
	massMap["B3 II"] = massMap["B5 II"] + ((massMap["B0 II"] - massMap["B5 II"]) / 5 * 2)
	massMap["B4 II"] = massMap["B5 II"] + ((massMap["B0 II"] - massMap["B5 II"]) / 5 * 1)
	massMap["A1 II"] = massMap["A5 II"] + ((massMap["A0 II"] - massMap["A5 II"]) / 5 * 4)
	massMap["A2 II"] = massMap["A5 II"] + ((massMap["A0 II"] - massMap["A5 II"]) / 5 * 3)
	massMap["A3 II"] = massMap["A5 II"] + ((massMap["A0 II"] - massMap["A5 II"]) / 5 * 2)
	massMap["A4 II"] = massMap["A5 II"] + ((massMap["A0 II"] - massMap["A5 II"]) / 5 * 1)
	massMap["F1 II"] = massMap["F5 II"] + ((massMap["F0 II"] - massMap["F5 II"]) / 5 * 4)
	massMap["F2 II"] = massMap["F5 II"] + ((massMap["F0 II"] - massMap["F5 II"]) / 5 * 3)
	massMap["F3 II"] = massMap["F5 II"] + ((massMap["F0 II"] - massMap["F5 II"]) / 5 * 2)
	massMap["F4 II"] = massMap["F5 II"] + ((massMap["F0 II"] - massMap["F5 II"]) / 5 * 1)
	massMap["G1 II"] = massMap["G5 II"] + ((massMap["G0 II"] - massMap["G5 II"]) / 5 * 4)
	massMap["G2 II"] = massMap["G5 II"] + ((massMap["G0 II"] - massMap["G5 II"]) / 5 * 3)
	massMap["G3 II"] = massMap["G5 II"] + ((massMap["G0 II"] - massMap["G5 II"]) / 5 * 2)
	massMap["G4 II"] = massMap["G5 II"] + ((massMap["G0 II"] - massMap["G5 II"]) / 5 * 1)
	massMap["K1 II"] = massMap["K5 II"] + ((massMap["K0 II"] - massMap["K5 II"]) / 5 * 4)
	massMap["K2 II"] = massMap["K5 II"] + ((massMap["K0 II"] - massMap["K5 II"]) / 5 * 3)
	massMap["K3 II"] = massMap["K5 II"] + ((massMap["K0 II"] - massMap["K5 II"]) / 5 * 2)
	massMap["K4 II"] = massMap["K5 II"] + ((massMap["K0 II"] - massMap["K5 II"]) / 5 * 1)
	massMap["M1 II"] = massMap["M5 II"] + ((massMap["M0 II"] - massMap["M5 II"]) / 5 * 4)
	massMap["M2 II"] = massMap["M5 II"] + ((massMap["M0 II"] - massMap["M5 II"]) / 5 * 3)
	massMap["M3 II"] = massMap["M5 II"] + ((massMap["M0 II"] - massMap["M5 II"]) / 5 * 2)
	massMap["M4 II"] = massMap["M5 II"] + ((massMap["M0 II"] - massMap["M5 II"]) / 5 * 1)

	massMap["O6 II"] = massMap["B0 II"] + ((massMap["O5 II"] - massMap["B0 II"]) / 5 * 4)
	massMap["O7 II"] = massMap["B0 II"] + ((massMap["O5 II"] - massMap["B0 II"]) / 5 * 3)
	massMap["O8 II"] = massMap["B0 II"] + ((massMap["O5 II"] - massMap["B0 II"]) / 5 * 2)
	massMap["O9 II"] = massMap["B0 II"] + ((massMap["O5 II"] - massMap["B0 II"]) / 5 * 1)
	massMap["B6 II"] = massMap["A0 II"] + ((massMap["B5 II"] - massMap["A0 II"]) / 5 * 4)
	massMap["B7 II"] = massMap["A0 II"] + ((massMap["B5 II"] - massMap["A0 II"]) / 5 * 3)
	massMap["B8 II"] = massMap["A0 II"] + ((massMap["B5 II"] - massMap["A0 II"]) / 5 * 2)
	massMap["B9 II"] = massMap["A0 II"] + ((massMap["B5 II"] - massMap["A0 II"]) / 5 * 1)
	massMap["A6 II"] = massMap["F0 II"] + ((massMap["A5 II"] - massMap["F0 II"]) / 5 * 4)
	massMap["A7 II"] = massMap["F0 II"] + ((massMap["A5 II"] - massMap["F0 II"]) / 5 * 3)
	massMap["A8 II"] = massMap["F0 II"] + ((massMap["A5 II"] - massMap["F0 II"]) / 5 * 2)
	massMap["A9 II"] = massMap["F0 II"] + ((massMap["A5 II"] - massMap["F0 II"]) / 5 * 1)
	massMap["F6 II"] = massMap["G0 II"] + ((massMap["F5 II"] - massMap["G0 II"]) / 5 * 4)
	massMap["F7 II"] = massMap["G0 II"] + ((massMap["F5 II"] - massMap["G0 II"]) / 5 * 3)
	massMap["F8 II"] = massMap["G0 II"] + ((massMap["F5 II"] - massMap["G0 II"]) / 5 * 2)
	massMap["F9 II"] = massMap["G0 II"] + ((massMap["F5 II"] - massMap["G0 II"]) / 5 * 1)
	massMap["G6 II"] = massMap["K0 II"] + ((massMap["G5 II"] - massMap["K0 II"]) / 5 * 4)
	massMap["G7 II"] = massMap["K0 II"] + ((massMap["G5 II"] - massMap["K0 II"]) / 5 * 3)
	massMap["G8 II"] = massMap["K0 II"] + ((massMap["G5 II"] - massMap["K0 II"]) / 5 * 2)
	massMap["G9 II"] = massMap["K0 II"] + ((massMap["G5 II"] - massMap["K0 II"]) / 5 * 1)
	massMap["K6 II"] = massMap["M0 II"] + ((massMap["K5 II"] - massMap["M0 II"]) / 5 * 4)
	massMap["K7 II"] = massMap["M0 II"] + ((massMap["K5 II"] - massMap["M0 II"]) / 5 * 3)
	massMap["K8 II"] = massMap["M0 II"] + ((massMap["K5 II"] - massMap["M0 II"]) / 5 * 2)
	massMap["K9 II"] = massMap["M0 II"] + ((massMap["K5 II"] - massMap["M0 II"]) / 5 * 1)
	massMap["M6 II"] = massMap["M9 II"] + ((massMap["M5 II"] - massMap["M9 II"]) / 4 * 3)
	massMap["M7 II"] = massMap["M9 II"] + ((massMap["M5 II"] - massMap["M9 II"]) / 4 * 2)
	massMap["M8 II"] = massMap["M9 II"] + ((massMap["M5 II"] - massMap["M9 II"]) / 4 * 1)

	////////////////////III

	massMap["O0 III"] = 110000000
	massMap["O5 III"] = 30000000
	massMap["B0 III"] = 20000000
	massMap["B5 III"] = 10000000
	massMap["A0 III"] = 8000000
	massMap["A5 III"] = 6000000
	massMap["F0 III"] = 4000000
	massMap["F5 III"] = 3000000
	massMap["G0 III"] = 2500000
	massMap["G5 III"] = 2400000
	massMap["K0 III"] = 1100000
	massMap["K5 III"] = 1500000
	massMap["M0 III"] = 1800000
	massMap["M5 III"] = 2400000
	massMap["M9 III"] = 8000000

	massMap["O1 III"] = massMap["O5 III"] + ((massMap["O0 III"] - massMap["O5 III"]) / 5 * 4)
	massMap["O2 III"] = massMap["O5 III"] + ((massMap["O0 III"] - massMap["O5 III"]) / 5 * 3)
	massMap["O3 III"] = massMap["O5 III"] + ((massMap["O0 III"] - massMap["O5 III"]) / 5 * 2)
	massMap["O4 III"] = massMap["O5 III"] + ((massMap["O0 III"] - massMap["O5 III"]) / 5 * 1)
	massMap["B1 III"] = massMap["B5 III"] + ((massMap["B0 III"] - massMap["B5 III"]) / 5 * 4)
	massMap["B2 III"] = massMap["B5 III"] + ((massMap["B0 III"] - massMap["B5 III"]) / 5 * 3)
	massMap["B3 III"] = massMap["B5 III"] + ((massMap["B0 III"] - massMap["B5 III"]) / 5 * 2)
	massMap["B4 III"] = massMap["B5 III"] + ((massMap["B0 III"] - massMap["B5 III"]) / 5 * 1)
	massMap["A1 III"] = massMap["A5 III"] + ((massMap["A0 III"] - massMap["A5 III"]) / 5 * 4)
	massMap["A2 III"] = massMap["A5 III"] + ((massMap["A0 III"] - massMap["A5 III"]) / 5 * 3)
	massMap["A3 III"] = massMap["A5 III"] + ((massMap["A0 III"] - massMap["A5 III"]) / 5 * 2)
	massMap["A4 III"] = massMap["A5 III"] + ((massMap["A0 III"] - massMap["A5 III"]) / 5 * 1)
	massMap["F1 III"] = massMap["F5 III"] + ((massMap["F0 III"] - massMap["F5 III"]) / 5 * 4)
	massMap["F2 III"] = massMap["F5 III"] + ((massMap["F0 III"] - massMap["F5 III"]) / 5 * 3)
	massMap["F3 III"] = massMap["F5 III"] + ((massMap["F0 III"] - massMap["F5 III"]) / 5 * 2)
	massMap["F4 III"] = massMap["F5 III"] + ((massMap["F0 III"] - massMap["F5 III"]) / 5 * 1)
	massMap["G1 III"] = massMap["G5 III"] + ((massMap["G0 III"] - massMap["G5 III"]) / 5 * 4)
	massMap["G2 III"] = massMap["G5 III"] + ((massMap["G0 III"] - massMap["G5 III"]) / 5 * 3)
	massMap["G3 III"] = massMap["G5 III"] + ((massMap["G0 III"] - massMap["G5 III"]) / 5 * 2)
	massMap["G4 III"] = massMap["G5 III"] + ((massMap["G0 III"] - massMap["G5 III"]) / 5 * 1)
	massMap["K1 III"] = massMap["K5 III"] + ((massMap["K0 III"] - massMap["K5 III"]) / 5 * 4)
	massMap["K2 III"] = massMap["K5 III"] + ((massMap["K0 III"] - massMap["K5 III"]) / 5 * 3)
	massMap["K3 III"] = massMap["K5 III"] + ((massMap["K0 III"] - massMap["K5 III"]) / 5 * 2)
	massMap["K4 III"] = massMap["K5 III"] + ((massMap["K0 III"] - massMap["K5 III"]) / 5 * 1)
	massMap["M1 III"] = massMap["M5 III"] + ((massMap["M0 III"] - massMap["M5 III"]) / 5 * 4)
	massMap["M2 III"] = massMap["M5 III"] + ((massMap["M0 III"] - massMap["M5 III"]) / 5 * 3)
	massMap["M3 III"] = massMap["M5 III"] + ((massMap["M0 III"] - massMap["M5 III"]) / 5 * 2)
	massMap["M4 III"] = massMap["M5 III"] + ((massMap["M0 III"] - massMap["M5 III"]) / 5 * 1)

	massMap["O6 III"] = massMap["B0 III"] + ((massMap["O5 III"] - massMap["B0 III"]) / 5 * 4)
	massMap["O7 III"] = massMap["B0 III"] + ((massMap["O5 III"] - massMap["B0 III"]) / 5 * 3)
	massMap["O8 III"] = massMap["B0 III"] + ((massMap["O5 III"] - massMap["B0 III"]) / 5 * 2)
	massMap["O9 III"] = massMap["B0 III"] + ((massMap["O5 III"] - massMap["B0 III"]) / 5 * 1)
	massMap["B6 III"] = massMap["A0 III"] + ((massMap["B5 III"] - massMap["A0 III"]) / 5 * 4)
	massMap["B7 III"] = massMap["A0 III"] + ((massMap["B5 III"] - massMap["A0 III"]) / 5 * 3)
	massMap["B8 III"] = massMap["A0 III"] + ((massMap["B5 III"] - massMap["A0 III"]) / 5 * 2)
	massMap["B9 III"] = massMap["A0 III"] + ((massMap["B5 III"] - massMap["A0 III"]) / 5 * 1)
	massMap["A6 III"] = massMap["F0 III"] + ((massMap["A5 III"] - massMap["F0 III"]) / 5 * 4)
	massMap["A7 III"] = massMap["F0 III"] + ((massMap["A5 III"] - massMap["F0 III"]) / 5 * 3)
	massMap["A8 III"] = massMap["F0 III"] + ((massMap["A5 III"] - massMap["F0 III"]) / 5 * 2)
	massMap["A9 III"] = massMap["F0 III"] + ((massMap["A5 III"] - massMap["F0 III"]) / 5 * 1)
	massMap["F6 III"] = massMap["G0 III"] + ((massMap["F5 III"] - massMap["G0 III"]) / 5 * 4)
	massMap["F7 III"] = massMap["G0 III"] + ((massMap["F5 III"] - massMap["G0 III"]) / 5 * 3)
	massMap["F8 III"] = massMap["G0 III"] + ((massMap["F5 III"] - massMap["G0 III"]) / 5 * 2)
	massMap["F9 III"] = massMap["G0 III"] + ((massMap["F5 III"] - massMap["G0 III"]) / 5 * 1)
	massMap["G6 III"] = massMap["K0 III"] + ((massMap["G5 III"] - massMap["K0 III"]) / 5 * 4)
	massMap["G7 III"] = massMap["K0 III"] + ((massMap["G5 III"] - massMap["K0 III"]) / 5 * 3)
	massMap["G8 III"] = massMap["K0 III"] + ((massMap["G5 III"] - massMap["K0 III"]) / 5 * 2)
	massMap["G9 III"] = massMap["K0 III"] + ((massMap["G5 III"] - massMap["K0 III"]) / 5 * 1)
	massMap["K6 III"] = massMap["M0 III"] + ((massMap["K5 III"] - massMap["M0 III"]) / 5 * 4)
	massMap["K7 III"] = massMap["M0 III"] + ((massMap["K5 III"] - massMap["M0 III"]) / 5 * 3)
	massMap["K8 III"] = massMap["M0 III"] + ((massMap["K5 III"] - massMap["M0 III"]) / 5 * 2)
	massMap["K9 III"] = massMap["M0 III"] + ((massMap["K5 III"] - massMap["M0 III"]) / 5 * 1)
	massMap["M6 III"] = massMap["M9 III"] + ((massMap["M5 III"] - massMap["M9 III"]) / 4 * 3)
	massMap["M7 III"] = massMap["M9 III"] + ((massMap["M5 III"] - massMap["M9 III"]) / 4 * 2)
	massMap["M8 III"] = massMap["M9 III"] + ((massMap["M5 III"] - massMap["M9 III"]) / 4 * 1)

	////////////////////V

	massMap["O0 V"] = 90000000
	massMap["O5 V"] = 60000000
	massMap["B0 V"] = 18000000
	massMap["B5 V"] = 5000000
	massMap["A0 V"] = 2200000
	massMap["A5 V"] = 1800000
	massMap["F0 V"] = 1500000
	massMap["F5 V"] = 1300000
	massMap["G0 V"] = 1100000
	massMap["G5 V"] = 900000
	massMap["K0 V"] = 800000
	massMap["K5 V"] = 700000
	massMap["M0 V"] = 500000
	massMap["M5 V"] = 160000
	massMap["M9 V"] = 80000

	massMap["O1 V"] = massMap["O5 V"] + ((massMap["O0 V"] - massMap["O5 V"]) / 5 * 4)
	massMap["O2 V"] = massMap["O5 V"] + ((massMap["O0 V"] - massMap["O5 V"]) / 5 * 3)
	massMap["O3 V"] = massMap["O5 V"] + ((massMap["O0 V"] - massMap["O5 V"]) / 5 * 2)
	massMap["O4 V"] = massMap["O5 V"] + ((massMap["O0 V"] - massMap["O5 V"]) / 5 * 1)
	massMap["B1 V"] = massMap["B5 V"] + ((massMap["B0 V"] - massMap["B5 V"]) / 5 * 4)
	massMap["B2 V"] = massMap["B5 V"] + ((massMap["B0 V"] - massMap["B5 V"]) / 5 * 3)
	massMap["B3 V"] = massMap["B5 V"] + ((massMap["B0 V"] - massMap["B5 V"]) / 5 * 2)
	massMap["B4 V"] = massMap["B5 V"] + ((massMap["B0 V"] - massMap["B5 V"]) / 5 * 1)
	massMap["A1 V"] = massMap["A5 V"] + ((massMap["A0 V"] - massMap["A5 V"]) / 5 * 4)
	massMap["A2 V"] = massMap["A5 V"] + ((massMap["A0 V"] - massMap["A5 V"]) / 5 * 3)
	massMap["A3 V"] = massMap["A5 V"] + ((massMap["A0 V"] - massMap["A5 V"]) / 5 * 2)
	massMap["A4 V"] = massMap["A5 V"] + ((massMap["A0 V"] - massMap["A5 V"]) / 5 * 1)
	massMap["F1 V"] = massMap["F5 V"] + ((massMap["F0 V"] - massMap["F5 V"]) / 5 * 4)
	massMap["F2 V"] = massMap["F5 V"] + ((massMap["F0 V"] - massMap["F5 V"]) / 5 * 3)
	massMap["F3 V"] = massMap["F5 V"] + ((massMap["F0 V"] - massMap["F5 V"]) / 5 * 2)
	massMap["F4 V"] = massMap["F5 V"] + ((massMap["F0 V"] - massMap["F5 V"]) / 5 * 1)
	massMap["G1 V"] = massMap["G5 V"] + ((massMap["G0 V"] - massMap["G5 V"]) / 5 * 4)
	massMap["G2 V"] = massMap["G5 V"] + ((massMap["G0 V"] - massMap["G5 V"]) / 5 * 3)
	massMap["G3 V"] = massMap["G5 V"] + ((massMap["G0 V"] - massMap["G5 V"]) / 5 * 2)
	massMap["G4 V"] = massMap["G5 V"] + ((massMap["G0 V"] - massMap["G5 V"]) / 5 * 1)
	massMap["K1 V"] = massMap["K5 V"] + ((massMap["K0 V"] - massMap["K5 V"]) / 5 * 4)
	massMap["K2 V"] = massMap["K5 V"] + ((massMap["K0 V"] - massMap["K5 V"]) / 5 * 3)
	massMap["K3 V"] = massMap["K5 V"] + ((massMap["K0 V"] - massMap["K5 V"]) / 5 * 2)
	massMap["K4 V"] = massMap["K5 V"] + ((massMap["K0 V"] - massMap["K5 V"]) / 5 * 1)
	massMap["M1 V"] = massMap["M5 V"] + ((massMap["M0 V"] - massMap["M5 V"]) / 5 * 4)
	massMap["M2 V"] = massMap["M5 V"] + ((massMap["M0 V"] - massMap["M5 V"]) / 5 * 3)
	massMap["M3 V"] = massMap["M5 V"] + ((massMap["M0 V"] - massMap["M5 V"]) / 5 * 2)
	massMap["M4 V"] = massMap["M5 V"] + ((massMap["M0 V"] - massMap["M5 V"]) / 5 * 1)

	massMap["O6 V"] = massMap["B0 V"] + ((massMap["O5 V"] - massMap["B0 V"]) / 5 * 4)
	massMap["O7 V"] = massMap["B0 V"] + ((massMap["O5 V"] - massMap["B0 V"]) / 5 * 3)
	massMap["O8 V"] = massMap["B0 V"] + ((massMap["O5 V"] - massMap["B0 V"]) / 5 * 2)
	massMap["O9 V"] = massMap["B0 V"] + ((massMap["O5 V"] - massMap["B0 V"]) / 5 * 1)
	massMap["B6 V"] = massMap["A0 V"] + ((massMap["B5 V"] - massMap["A0 V"]) / 5 * 4)
	massMap["B7 V"] = massMap["A0 V"] + ((massMap["B5 V"] - massMap["A0 V"]) / 5 * 3)
	massMap["B8 V"] = massMap["A0 V"] + ((massMap["B5 V"] - massMap["A0 V"]) / 5 * 2)
	massMap["B9 V"] = massMap["A0 V"] + ((massMap["B5 V"] - massMap["A0 V"]) / 5 * 1)
	massMap["A6 V"] = massMap["F0 V"] + ((massMap["A5 V"] - massMap["F0 V"]) / 5 * 4)
	massMap["A7 V"] = massMap["F0 V"] + ((massMap["A5 V"] - massMap["F0 V"]) / 5 * 3)
	massMap["A8 V"] = massMap["F0 V"] + ((massMap["A5 V"] - massMap["F0 V"]) / 5 * 2)
	massMap["A9 V"] = massMap["F0 V"] + ((massMap["A5 V"] - massMap["F0 V"]) / 5 * 1)
	massMap["F6 V"] = massMap["G0 V"] + ((massMap["F5 V"] - massMap["G0 V"]) / 5 * 4)
	massMap["F7 V"] = massMap["G0 V"] + ((massMap["F5 V"] - massMap["G0 V"]) / 5 * 3)
	massMap["F8 V"] = massMap["G0 V"] + ((massMap["F5 V"] - massMap["G0 V"]) / 5 * 2)
	massMap["F9 V"] = massMap["G0 V"] + ((massMap["F5 V"] - massMap["G0 V"]) / 5 * 1)
	massMap["G6 V"] = massMap["K0 V"] + ((massMap["G5 V"] - massMap["K0 V"]) / 5 * 4)
	massMap["G7 V"] = massMap["K0 V"] + ((massMap["G5 V"] - massMap["K0 V"]) / 5 * 3)
	massMap["G8 V"] = massMap["K0 V"] + ((massMap["G5 V"] - massMap["K0 V"]) / 5 * 2)
	massMap["G9 V"] = massMap["K0 V"] + ((massMap["G5 V"] - massMap["K0 V"]) / 5 * 1)
	massMap["K6 V"] = massMap["M0 V"] + ((massMap["K5 V"] - massMap["M0 V"]) / 5 * 4)
	massMap["K7 V"] = massMap["M0 V"] + ((massMap["K5 V"] - massMap["M0 V"]) / 5 * 3)
	massMap["K8 V"] = massMap["M0 V"] + ((massMap["K5 V"] - massMap["M0 V"]) / 5 * 2)
	massMap["K9 V"] = massMap["M0 V"] + ((massMap["K5 V"] - massMap["M0 V"]) / 5 * 1)
	massMap["M6 V"] = massMap["M9 V"] + ((massMap["M5 V"] - massMap["M9 V"]) / 4 * 3)
	massMap["M7 V"] = massMap["M9 V"] + ((massMap["M5 V"] - massMap["M9 V"]) / 4 * 2)
	massMap["M8 V"] = massMap["M9 V"] + ((massMap["M5 V"] - massMap["M9 V"]) / 4 * 1)

	////////////////////IV

	massMap["B0 IV"] = 20000000
	massMap["B5 IV"] = 10000000
	massMap["A0 IV"] = 4000000
	massMap["A5 IV"] = 2300000
	massMap["F0 IV"] = 2000000
	massMap["F5 IV"] = 1500000
	massMap["G0 IV"] = 1700000
	massMap["G5 IV"] = 1200000
	massMap["K0 IV"] = 1500000
	massMap["K4 IV"] = 1600000

	massMap["B1 IV"] = massMap["B5 IV"] + ((massMap["B0 IV"] - massMap["B5 IV"]) / 5 * 4)
	massMap["B2 IV"] = massMap["B5 IV"] + ((massMap["B0 IV"] - massMap["B5 IV"]) / 5 * 3)
	massMap["B3 IV"] = massMap["B5 IV"] + ((massMap["B0 IV"] - massMap["B5 IV"]) / 5 * 2)
	massMap["B4 IV"] = massMap["B5 IV"] + ((massMap["B0 IV"] - massMap["B5 IV"]) / 5 * 1)
	massMap["A1 IV"] = massMap["A5 IV"] + ((massMap["A0 IV"] - massMap["A5 IV"]) / 5 * 4)
	massMap["A2 IV"] = massMap["A5 IV"] + ((massMap["A0 IV"] - massMap["A5 IV"]) / 5 * 3)
	massMap["A3 IV"] = massMap["A5 IV"] + ((massMap["A0 IV"] - massMap["A5 IV"]) / 5 * 2)
	massMap["A4 IV"] = massMap["A5 IV"] + ((massMap["A0 IV"] - massMap["A5 IV"]) / 5 * 1)
	massMap["F1 IV"] = massMap["F5 IV"] + ((massMap["F0 IV"] - massMap["F5 IV"]) / 5 * 4)
	massMap["F2 IV"] = massMap["F5 IV"] + ((massMap["F0 IV"] - massMap["F5 IV"]) / 5 * 3)
	massMap["F3 IV"] = massMap["F5 IV"] + ((massMap["F0 IV"] - massMap["F5 IV"]) / 5 * 2)
	massMap["F4 IV"] = massMap["F5 IV"] + ((massMap["F0 IV"] - massMap["F5 IV"]) / 5 * 1)
	massMap["G1 IV"] = massMap["G5 IV"] + ((massMap["G0 IV"] - massMap["G5 IV"]) / 5 * 4)
	massMap["G2 IV"] = massMap["G5 IV"] + ((massMap["G0 IV"] - massMap["G5 IV"]) / 5 * 3)
	massMap["G3 IV"] = massMap["G5 IV"] + ((massMap["G0 IV"] - massMap["G5 IV"]) / 5 * 2)
	massMap["G4 IV"] = massMap["G5 IV"] + ((massMap["G0 IV"] - massMap["G5 IV"]) / 5 * 1)
	massMap["K1 IV"] = massMap["K5 IV"] + ((massMap["K0 IV"] - massMap["K5 IV"]) / 4 * 3)
	massMap["K2 IV"] = massMap["K5 IV"] + ((massMap["K0 IV"] - massMap["K5 IV"]) / 4 * 2)
	massMap["K3 IV"] = massMap["K5 IV"] + ((massMap["K0 IV"] - massMap["K5 IV"]) / 4 * 1)

	massMap["B6 IV"] = massMap["A0 IV"] + ((massMap["B5 IV"] - massMap["A0 IV"]) / 5 * 4)
	massMap["B7 IV"] = massMap["A0 IV"] + ((massMap["B5 IV"] - massMap["A0 IV"]) / 5 * 3)
	massMap["B8 IV"] = massMap["A0 IV"] + ((massMap["B5 IV"] - massMap["A0 IV"]) / 5 * 2)
	massMap["B9 IV"] = massMap["A0 IV"] + ((massMap["B5 IV"] - massMap["A0 IV"]) / 5 * 1)
	massMap["A6 IV"] = massMap["F0 IV"] + ((massMap["A5 IV"] - massMap["F0 IV"]) / 5 * 4)
	massMap["A7 IV"] = massMap["F0 IV"] + ((massMap["A5 IV"] - massMap["F0 IV"]) / 5 * 3)
	massMap["A8 IV"] = massMap["F0 IV"] + ((massMap["A5 IV"] - massMap["F0 IV"]) / 5 * 2)
	massMap["A9 IV"] = massMap["F0 IV"] + ((massMap["A5 IV"] - massMap["F0 IV"]) / 5 * 1)
	massMap["F6 IV"] = massMap["G0 IV"] + ((massMap["F5 IV"] - massMap["G0 IV"]) / 5 * 4)
	massMap["F7 IV"] = massMap["G0 IV"] + ((massMap["F5 IV"] - massMap["G0 IV"]) / 5 * 3)
	massMap["F8 IV"] = massMap["G0 IV"] + ((massMap["F5 IV"] - massMap["G0 IV"]) / 5 * 2)
	massMap["F9 IV"] = massMap["G0 IV"] + ((massMap["F5 IV"] - massMap["G0 IV"]) / 5 * 1)
	massMap["G6 IV"] = massMap["K0 IV"] + ((massMap["G5 IV"] - massMap["K0 IV"]) / 5 * 4)
	massMap["G7 IV"] = massMap["K0 IV"] + ((massMap["G5 IV"] - massMap["K0 IV"]) / 5 * 3)
	massMap["G8 IV"] = massMap["K0 IV"] + ((massMap["G5 IV"] - massMap["K0 IV"]) / 5 * 2)
	massMap["G9 IV"] = massMap["K0 IV"] + ((massMap["G5 IV"] - massMap["K0 IV"]) / 5 * 1)

	////////////////////V

	massMap["O0 VI"] = 2000000
	massMap["O5 VI"] = 1500000
	massMap["B0 VI"] = 500000
	massMap["B5 VI"] = 400000
	massMap["B9 VI"] = 350000
	massMap["G0 VI"] = 800000
	massMap["G5 VI"] = 700000
	massMap["K0 VI"] = 600000
	massMap["K5 VI"] = 500000
	massMap["M0 VI"] = 400000
	massMap["M5 VI"] = 120000
	massMap["M9 VI"] = 75000

	massMap["O1 VI"] = massMap["O5 VI"] + ((massMap["O0 VI"] - massMap["O5 VI"]) / 5 * 4)
	massMap["O2 VI"] = massMap["O5 VI"] + ((massMap["O0 VI"] - massMap["O5 VI"]) / 5 * 3)
	massMap["O3 VI"] = massMap["O5 VI"] + ((massMap["O0 VI"] - massMap["O5 VI"]) / 5 * 2)
	massMap["O4 VI"] = massMap["O5 VI"] + ((massMap["O0 VI"] - massMap["O5 VI"]) / 5 * 1)
	massMap["B1 VI"] = massMap["B5 VI"] + ((massMap["B0 VI"] - massMap["B5 VI"]) / 5 * 4)
	massMap["B2 VI"] = massMap["B5 VI"] + ((massMap["B0 VI"] - massMap["B5 VI"]) / 5 * 3)
	massMap["B3 VI"] = massMap["B5 VI"] + ((massMap["B0 VI"] - massMap["B5 VI"]) / 5 * 2)
	massMap["B4 VI"] = massMap["B5 VI"] + ((massMap["B0 VI"] - massMap["B5 VI"]) / 5 * 1)
	massMap["G1 VI"] = massMap["G5 VI"] + ((massMap["G0 VI"] - massMap["G5 VI"]) / 5 * 4)
	massMap["G2 VI"] = massMap["G5 VI"] + ((massMap["G0 VI"] - massMap["G5 VI"]) / 5 * 3)
	massMap["G3 VI"] = massMap["G5 VI"] + ((massMap["G0 VI"] - massMap["G5 VI"]) / 5 * 2)
	massMap["G4 VI"] = massMap["G5 VI"] + ((massMap["G0 VI"] - massMap["G5 VI"]) / 5 * 1)
	massMap["K1 VI"] = massMap["K5 VI"] + ((massMap["K0 VI"] - massMap["K5 VI"]) / 5 * 4)
	massMap["K2 VI"] = massMap["K5 VI"] + ((massMap["K0 VI"] - massMap["K5 VI"]) / 5 * 3)
	massMap["K3 VI"] = massMap["K5 VI"] + ((massMap["K0 VI"] - massMap["K5 VI"]) / 5 * 2)
	massMap["K4 VI"] = massMap["K5 VI"] + ((massMap["K0 VI"] - massMap["K5 VI"]) / 5 * 1)
	massMap["M1 VI"] = massMap["M5 VI"] + ((massMap["M0 VI"] - massMap["M5 VI"]) / 5 * 4)
	massMap["M2 VI"] = massMap["M5 VI"] + ((massMap["M0 VI"] - massMap["M5 VI"]) / 5 * 3)
	massMap["M3 VI"] = massMap["M5 VI"] + ((massMap["M0 VI"] - massMap["M5 VI"]) / 5 * 2)
	massMap["M4 VI"] = massMap["M5 VI"] + ((massMap["M0 VI"] - massMap["M5 VI"]) / 5 * 1)

	massMap["O6 VI"] = massMap["B0 VI"] + ((massMap["O5 VI"] - massMap["B0 VI"]) / 5 * 4)
	massMap["O7 VI"] = massMap["B0 VI"] + ((massMap["O5 VI"] - massMap["B0 VI"]) / 5 * 3)
	massMap["O8 VI"] = massMap["B0 VI"] + ((massMap["O5 VI"] - massMap["B0 VI"]) / 5 * 2)
	massMap["O9 VI"] = massMap["B0 VI"] + ((massMap["O5 VI"] - massMap["B0 VI"]) / 5 * 1)
	massMap["B6 VI"] = massMap["B9 VI"] + ((massMap["B5 VI"] - massMap["B9 VI"]) / 4 * 3)
	massMap["B7 VI"] = massMap["B9 VI"] + ((massMap["B5 VI"] - massMap["B9 VI"]) / 4 * 2)
	massMap["B8 VI"] = massMap["B9 VI"] + ((massMap["B5 VI"] - massMap["B9 VI"]) / 4 * 1)
	massMap["G6 VI"] = massMap["K0 VI"] + ((massMap["G5 VI"] - massMap["K0 VI"]) / 5 * 4)
	massMap["G7 VI"] = massMap["K0 VI"] + ((massMap["G5 VI"] - massMap["K0 VI"]) / 5 * 3)
	massMap["G8 VI"] = massMap["K0 VI"] + ((massMap["G5 VI"] - massMap["K0 VI"]) / 5 * 2)
	massMap["G9 VI"] = massMap["K0 VI"] + ((massMap["G5 VI"] - massMap["K0 VI"]) / 5 * 1)
	massMap["K6 VI"] = massMap["M0 VI"] + ((massMap["K5 VI"] - massMap["M0 VI"]) / 5 * 4)
	massMap["K7 VI"] = massMap["M0 VI"] + ((massMap["K5 VI"] - massMap["M0 VI"]) / 5 * 3)
	massMap["K8 VI"] = massMap["M0 VI"] + ((massMap["K5 VI"] - massMap["M0 VI"]) / 5 * 2)
	massMap["K9 VI"] = massMap["M0 VI"] + ((massMap["K5 VI"] - massMap["M0 VI"]) / 5 * 1)
	massMap["M6 VI"] = massMap["M9 VI"] + ((massMap["M5 VI"] - massMap["M9 VI"]) / 4 * 3)
	massMap["M7 VI"] = massMap["M9 VI"] + ((massMap["M5 VI"] - massMap["M9 VI"]) / 4 * 2)
	massMap["M8 VI"] = massMap["M9 VI"] + ((massMap["M5 VI"] - massMap["M9 VI"]) / 4 * 1)

	return massMap[shortStarDescription(st)]
}

func averageTempMap(st star) int {
	massMap := make(map[string]int)
	//KBIYTT
	massMap["O0"] = 50000
	massMap["O5"] = 40000
	massMap["B0"] = 30000
	massMap["B5"] = 15000
	massMap["A0"] = 10000
	massMap["A5"] = 8000
	massMap["F0"] = 7500
	massMap["F5"] = 6500
	massMap["G0"] = 6000
	massMap["G5"] = 5600
	massMap["K0"] = 5200
	massMap["K5"] = 4400
	massMap["M0"] = 3700
	massMap["M5"] = 3000
	massMap["M9"] = 2400

	massMap["O1"] = massMap["O5"] + ((massMap["O0"] - massMap["O5"]) / 5 * 4)
	massMap["O2"] = massMap["O5"] + ((massMap["O0"] - massMap["O5"]) / 5 * 3)
	massMap["O3"] = massMap["O5"] + ((massMap["O0"] - massMap["O5"]) / 5 * 2)
	massMap["O4"] = massMap["O5"] + ((massMap["O0"] - massMap["O5"]) / 5 * 1)
	massMap["B1"] = massMap["B5"] + ((massMap["B0"] - massMap["B5"]) / 5 * 4)
	massMap["B2"] = massMap["B5"] + ((massMap["B0"] - massMap["B5"]) / 5 * 3)
	massMap["B3"] = massMap["B5"] + ((massMap["B0"] - massMap["B5"]) / 5 * 2)
	massMap["B4"] = massMap["B5"] + ((massMap["B0"] - massMap["B5"]) / 5 * 1)
	massMap["A1"] = massMap["A5"] + ((massMap["A0"] - massMap["A5"]) / 5 * 4)
	massMap["A2"] = massMap["A5"] + ((massMap["A0"] - massMap["A5"]) / 5 * 3)
	massMap["A3"] = massMap["A5"] + ((massMap["A0"] - massMap["A5"]) / 5 * 2)
	massMap["A4"] = massMap["A5"] + ((massMap["A0"] - massMap["A5"]) / 5 * 1)
	massMap["F1"] = massMap["F5"] + ((massMap["F0"] - massMap["F5"]) / 5 * 4)
	massMap["F2"] = massMap["F5"] + ((massMap["F0"] - massMap["F5"]) / 5 * 3)
	massMap["F3"] = massMap["F5"] + ((massMap["F0"] - massMap["F5"]) / 5 * 2)
	massMap["F4"] = massMap["F5"] + ((massMap["F0"] - massMap["F5"]) / 5 * 1)
	massMap["G1"] = massMap["G5"] + ((massMap["G0"] - massMap["G5"]) / 5 * 4)
	massMap["G2"] = massMap["G5"] + ((massMap["G0"] - massMap["G5"]) / 5 * 3)
	massMap["G3"] = massMap["G5"] + ((massMap["G0"] - massMap["G5"]) / 5 * 2)
	massMap["G4"] = massMap["G5"] + ((massMap["G0"] - massMap["G5"]) / 5 * 1)
	massMap["K1"] = massMap["K5"] + ((massMap["K0"] - massMap["K5"]) / 5 * 4)
	massMap["K2"] = massMap["K5"] + ((massMap["K0"] - massMap["K5"]) / 5 * 3)
	massMap["K3"] = massMap["K5"] + ((massMap["K0"] - massMap["K5"]) / 5 * 2)
	massMap["K4"] = massMap["K5"] + ((massMap["K0"] - massMap["K5"]) / 5 * 1)
	massMap["M1"] = massMap["M5"] + ((massMap["M0"] - massMap["M5"]) / 5 * 4)
	massMap["M2"] = massMap["M5"] + ((massMap["M0"] - massMap["M5"]) / 5 * 3)
	massMap["M3"] = massMap["M5"] + ((massMap["M0"] - massMap["M5"]) / 5 * 2)
	massMap["M4"] = massMap["M5"] + ((massMap["M0"] - massMap["M5"]) / 5 * 1)

	massMap["O6"] = massMap["B0"] + ((massMap["O5"] - massMap["B0"]) / 5 * 4)
	massMap["O7"] = massMap["B0"] + ((massMap["O5"] - massMap["B0"]) / 5 * 3)
	massMap["O8"] = massMap["B0"] + ((massMap["O5"] - massMap["B0"]) / 5 * 2)
	massMap["O9"] = massMap["B0"] + ((massMap["O5"] - massMap["B0"]) / 5 * 1)
	massMap["B6"] = massMap["A0"] + ((massMap["B5"] - massMap["A0"]) / 5 * 4)
	massMap["B7"] = massMap["A0"] + ((massMap["B5"] - massMap["A0"]) / 5 * 3)
	massMap["B8"] = massMap["A0"] + ((massMap["B5"] - massMap["A0"]) / 5 * 2)
	massMap["B9"] = massMap["A0"] + ((massMap["B5"] - massMap["A0"]) / 5 * 1)
	massMap["A6"] = massMap["F0"] + ((massMap["A5"] - massMap["F0"]) / 5 * 4)
	massMap["A7"] = massMap["F0"] + ((massMap["A5"] - massMap["F0"]) / 5 * 3)
	massMap["A8"] = massMap["F0"] + ((massMap["A5"] - massMap["F0"]) / 5 * 2)
	massMap["A9"] = massMap["F0"] + ((massMap["A5"] - massMap["F0"]) / 5 * 1)
	massMap["F6"] = massMap["G0"] + ((massMap["F5"] - massMap["G0"]) / 5 * 4)
	massMap["F7"] = massMap["G0"] + ((massMap["F5"] - massMap["G0"]) / 5 * 3)
	massMap["F8"] = massMap["G0"] + ((massMap["F5"] - massMap["G0"]) / 5 * 2)
	massMap["F9"] = massMap["G0"] + ((massMap["F5"] - massMap["G0"]) / 5 * 1)
	massMap["G6"] = massMap["K0"] + ((massMap["G5"] - massMap["K0"]) / 5 * 4)
	massMap["G7"] = massMap["K0"] + ((massMap["G5"] - massMap["K0"]) / 5 * 3)
	massMap["G8"] = massMap["K0"] + ((massMap["G5"] - massMap["K0"]) / 5 * 2)
	massMap["G9"] = massMap["K0"] + ((massMap["G5"] - massMap["K0"]) / 5 * 1)
	massMap["K6"] = massMap["M0"] + ((massMap["K5"] - massMap["M0"]) / 5 * 4)
	massMap["K7"] = massMap["M0"] + ((massMap["K5"] - massMap["M0"]) / 5 * 3)
	massMap["K8"] = massMap["M0"] + ((massMap["K5"] - massMap["M0"]) / 5 * 2)
	massMap["K9"] = massMap["M0"] + ((massMap["K5"] - massMap["M0"]) / 5 * 1)
	massMap["M6"] = massMap["M9"] + ((massMap["M5"] - massMap["M9"]) / 4 * 3)
	massMap["M7"] = massMap["M9"] + ((massMap["M5"] - massMap["M9"]) / 4 * 2)
	massMap["M8"] = massMap["M9"] + ((massMap["M5"] - massMap["M9"]) / 4 * 1)
	short := shortStarDescription(st)
	temp := -1
	for k, v := range massMap {
		if !strings.HasPrefix(short, k) {
			continue
		}
		return v
	}
	return temp
}
