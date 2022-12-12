package stellar

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/generation/table"
)

const (
	realistic   = "Realistic"
	luminocity1 = "Luminocity1"
	luminocity2 = "Luminocity2"
	mult1       = "mult1"
	mult2       = "mult2"
	mult3       = "mult3"
)

func generateStellar(dice *dice.Dicepool) string {
	diceRolls := []int{}
	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			diceRolls = append(diceRolls, dice.Sroll("1d100"))
		case 1:
			diceRolls = append(diceRolls, dice.Sroll("1d100"))
		case 2:
			diceRolls = append(diceRolls, dice.Sroll("1d100"))
		}
	}
	stars := []string{generateStar(dice)}
	mult := 0
	switch {
	case containsAny(stars[0], "A", "B", "O"):
		mult, _ = strconv.Atoi(chart(mult1).Result(diceRolls[0]))
	case containsAny(stars[0], "K", "F", "G"):
		mult, _ = strconv.Atoi(chart(mult2).Result(diceRolls[1]))
	case containsAny(stars[0], "M", "D"):
		mult, _ = strconv.Atoi(chart(mult3).Result(diceRolls[2]))
	}
	for len(stars) < mult {
		star := generateStar(dice)
		stars = append(stars, star)
	}
	//fmt.Println(stars)
	return agreggateStellar(stars)
}

func Parse(stellar string) []string {
	try := 1
	stars := []string{}
	for len(stellar) > 0 {
		if try > 10 {
			return stars
		}
		for _, l := range listAllStars() {
			//fmt.Printf("stellar: '%v' | stars: '%v' | l: '%v' | try: %v \n", stellar, stars, l, try)
			if strings.HasPrefix(stellar, l+" ") {
				stars = append(stars, l)
				stellar = strings.TrimPrefix(stellar, stars[len(stars)-1])
				stellar = strings.TrimPrefix(stellar, " ")
				stellar = strings.TrimSuffix(stellar, " ")
				//fmt.Printf("Redused stellar: '%v'\n", stellar)
			}
			if stellar == l {
				stars = append(stars, l)
				return stars
			}

		}
		try++
	}
	return stars
}

func agreggateStellar(stars []string) string {
	list := listAllStars()
	stellar := ""
	for _, l := range list {
		for _, s := range stars {

			if l == s {
				stellar += l + " "
			}
		}
	}
	stellar = strings.TrimSuffix(stellar, " ")
	//stellar = strings.TrimPrefix(stellar, " ")
	if stellar == "" {
		fmt.Println(stars)
		panic(1)
	}
	return strings.TrimSuffix(stellar, " ")
}

func generateStar(dp *dice.Dicepool) string {
	diceRolls := []int{}
	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			diceRolls = append(diceRolls, dp.Sroll("1d100"))
		case 1:
			diceRolls = append(diceRolls, dp.Sroll("1d100"))
		case 2:
			diceRolls = append(diceRolls, dp.Sroll("1d10"))
		case 3:
			diceRolls = append(diceRolls, dp.Sroll("1d10")-1)
		}
	}
	sType := chart(realistic).Result(diceRolls[0])
	sClass := chart(luminocity1).Result(diceRolls[1])
	if sClass == "c3" {
		sClass = chart(luminocity2).Result(diceRolls[2])
	}
	sNum := fmt.Sprintf("%v", diceRolls[3])
	if sClass == "D" {
		sType = ""
		return sClass + sNum
	}
	switch sType {
	case "O", "B", "A", "F":
		if sClass == "VI" {
			sClass = "V"
		}
	}
	if sType+sNum+" "+sClass == "M2 " {
		fmt.Println("++++++++", sType+sNum+" "+sClass)
		fmt.Println(diceRolls)
		panic(2)
	}
	return sType + sNum + " " + sClass
}

func containsAny(str string, substr ...string) bool {
	for _, sub := range substr {
		if strings.Contains(str, sub) {
			return true
		}
	}
	return false
}

func chart(s string) *table.DiceResultChart {
	switch s {
	default:
		return table.DiceChart()
	case realistic:
		return table.DiceChart(
			table.Row("01-80", "M"),
			table.Row("81-88", "K"),
			table.Row("89-94", "G"),
			table.Row("95-97", "F"),
			table.Row("98", "A"),
			table.Row("99", "B"),
			table.Row("100", "O"),
		)
	case luminocity1:
		return table.DiceChart(
			table.Row("01-90", "V"),
			table.Row("91-94", "IV"),
			table.Row("95-96", "D"),
			table.Row("97-99", "III"),
			table.Row("100", "c3"),
		)
	case luminocity2:
		return table.DiceChart(
			table.Row("1-4", "II"),
			table.Row("5-6", "VI"),
			table.Row("7-8", "Ia"),
			table.Row("9-10", "Ib"),
		)
	case mult1:
		return table.DiceChart(
			table.Row("01-10", "1"),
			table.Row("11-90", "2"),
			table.Row("91-98", "3"),
			table.Row("99", "4"),
			table.Row("100", "5"),
		)
	case mult2:
		return table.DiceChart(
			table.Row("01-45", "1"),
			table.Row("46-99", "2"),
			table.Row("100", "3"),
		)
	case mult3:
		return table.DiceChart(
			table.Row("01-69", "1"),
			table.Row("70-98", "2"),
			table.Row("99-100", "3"),
		)
	}
}

func listAllStars() []string {
	list := []string{
		"O0 V",
		"O1 V",
		"O2 V",
		"O3 V",
		"O4 V",
		"O5 V",
		"O6 V",
		"O7 V",
		"O8 V",
		"O9 V",
		"B0 V",
		"B1 V",
		"B2 V",
		"B3 V",
		"B4 V",
		"B5 V",
		"B6 V",
		"B7 V",
		"B8 V",
		"B9 V",
		"A0 V",
		"A1 V",
		"A2 V",
		"A3 V",
		"A4 V",
		"A5 V",
		"A6 V",
		"A7 V",
		"A8 V",
		"A9 V",
		"F0 V",
		"F1 V",
		"F2 V",
		"F3 V",
		"F4 V",
		"F5 V",
		"F6 V",
		"F7 V",
		"F8 V",
		"F9 V",
		"G0 V",
		"G1 V",
		"G2 V",
		"G3 V",
		"G4 V",
		"G5 V",
		"G6 V",
		"G7 V",
		"G8 V",
		"G9 V",
		"K0 V",
		"K1 V",
		"K2 V",
		"K3 V",
		"K4 V",
		"K5 V",
		"K6 V",
		"K7 V",
		"K8 V",
		"K9 V",
		"M0 V",
		"M1 V",
		"M2 V",
		"M3 V",
		"M4 V",
		"M5 V",
		"M6 V",
		"M7 V",
		"M8 V",
		"M9 V",
		"O0 Ia",
		"O1 Ia",
		"O2 Ia",
		"O3 Ia",
		"O4 Ia",
		"O5 Ia",
		"O6 Ia",
		"O7 Ia",
		"O8 Ia",
		"O9 Ia",
		"B0 Ia",
		"B1 Ia",
		"B2 Ia",
		"B3 Ia",
		"B4 Ia",
		"B5 Ia",
		"B6 Ia",
		"B7 Ia",
		"B8 Ia",
		"B9 Ia",
		"A0 Ia",
		"A1 Ia",
		"A2 Ia",
		"A3 Ia",
		"A4 Ia",
		"A5 Ia",
		"A6 Ia",
		"A7 Ia",
		"A8 Ia",
		"A9 Ia",
		"F0 Ia",
		"F1 Ia",
		"F2 Ia",
		"F3 Ia",
		"F4 Ia",
		"F5 Ia",
		"F6 Ia",
		"F7 Ia",
		"F8 Ia",
		"F9 Ia",
		"G0 Ia",
		"G1 Ia",
		"G2 Ia",
		"G3 Ia",
		"G4 Ia",
		"G5 Ia",
		"G6 Ia",
		"G7 Ia",
		"G8 Ia",
		"G9 Ia",
		"K0 Ia",
		"K1 Ia",
		"K2 Ia",
		"K3 Ia",
		"K4 Ia",
		"K5 Ia",
		"K6 Ia",
		"K7 Ia",
		"K8 Ia",
		"K9 Ia",
		"M0 Ia",
		"M1 Ia",
		"M2 Ia",
		"M3 Ia",
		"M4 Ia",
		"M5 Ia",
		"M6 Ia",
		"M7 Ia",
		"M8 Ia",
		"M9 Ia",
		"O0 Ib",
		"O1 Ib",
		"O2 Ib",
		"O3 Ib",
		"O4 Ib",
		"O5 Ib",
		"O6 Ib",
		"O7 Ib",
		"O8 Ib",
		"O9 Ib",
		"B0 Ib",
		"B1 Ib",
		"B2 Ib",
		"B3 Ib",
		"B4 Ib",
		"B5 Ib",
		"B6 Ib",
		"B7 Ib",
		"B8 Ib",
		"B9 Ib",
		"A0 Ib",
		"A1 Ib",
		"A2 Ib",
		"A3 Ib",
		"A4 Ib",
		"A5 Ib",
		"A6 Ib",
		"A7 Ib",
		"A8 Ib",
		"A9 Ib",
		"F0 Ib",
		"F1 Ib",
		"F2 Ib",
		"F3 Ib",
		"F4 Ib",
		"F5 Ib",
		"F6 Ib",
		"F7 Ib",
		"F8 Ib",
		"F9 Ib",
		"G0 Ib",
		"G1 Ib",
		"G2 Ib",
		"G3 Ib",
		"G4 Ib",
		"G5 Ib",
		"G6 Ib",
		"G7 Ib",
		"G8 Ib",
		"G9 Ib",
		"K0 Ib",
		"K1 Ib",
		"K2 Ib",
		"K3 Ib",
		"K4 Ib",
		"K5 Ib",
		"K6 Ib",
		"K7 Ib",
		"K8 Ib",
		"K9 Ib",
		"M0 Ib",
		"M1 Ib",
		"M2 Ib",
		"M3 Ib",
		"M4 Ib",
		"M5 Ib",
		"M6 Ib",
		"M7 Ib",
		"M8 Ib",
		"M9 Ib",
		"O0 II",
		"O1 II",
		"O2 II",
		"O3 II",
		"O4 II",
		"O5 II",
		"O6 II",
		"O7 II",
		"O8 II",
		"O9 II",
		"B0 II",
		"B1 II",
		"B2 II",
		"B3 II",
		"B4 II",
		"B5 II",
		"B6 II",
		"B7 II",
		"B8 II",
		"B9 II",
		"A0 II",
		"A1 II",
		"A2 II",
		"A3 II",
		"A4 II",
		"A5 II",
		"A6 II",
		"A7 II",
		"A8 II",
		"A9 II",
		"F0 II",
		"F1 II",
		"F2 II",
		"F3 II",
		"F4 II",
		"F5 II",
		"F6 II",
		"F7 II",
		"F8 II",
		"F9 II",
		"G0 II",
		"G1 II",
		"G2 II",
		"G3 II",
		"G4 II",
		"G5 II",
		"G6 II",
		"G7 II",
		"G8 II",
		"G9 II",
		"K0 II",
		"K1 II",
		"K2 II",
		"K3 II",
		"K4 II",
		"K5 II",
		"K6 II",
		"K7 II",
		"K8 II",
		"K9 II",
		"M0 II",
		"M1 II",
		"M2 II",
		"M3 II",
		"M4 II",
		"M5 II",
		"M6 II",
		"M7 II",
		"M8 II",
		"M9 II",
		"O0 III",
		"O1 III",
		"O2 III",
		"O3 III",
		"O4 III",
		"O5 III",
		"O6 III",
		"O7 III",
		"O8 III",
		"O9 III",
		"B0 III",
		"B1 III",
		"B2 III",
		"B3 III",
		"B4 III",
		"B5 III",
		"B6 III",
		"B7 III",
		"B8 III",
		"B9 III",
		"A0 III",
		"A1 III",
		"A2 III",
		"A3 III",
		"A4 III",
		"A5 III",
		"A6 III",
		"A7 III",
		"A8 III",
		"A9 III",
		"F0 III",
		"F1 III",
		"F2 III",
		"F3 III",
		"F4 III",
		"F5 III",
		"F6 III",
		"F7 III",
		"F8 III",
		"F9 III",
		"G0 III",
		"G1 III",
		"G2 III",
		"G3 III",
		"G4 III",
		"G5 III",
		"G6 III",
		"G7 III",
		"G8 III",
		"G9 III",
		"K0 III",
		"K1 III",
		"K2 III",
		"K3 III",
		"K4 III",
		"K5 III",
		"K6 III",
		"K7 III",
		"K8 III",
		"K9 III",
		"M0 III",
		"M1 III",
		"M2 III",
		"M3 III",
		"M4 III",
		"M5 III",
		"M6 III",
		"M7 III",
		"M8 III",
		"M9 III",
		"O0 IV",
		"O1 IV",
		"O2 IV",
		"O3 IV",
		"O4 IV",
		"O5 IV",
		"O6 IV",
		"O7 IV",
		"O8 IV",
		"O9 IV",
		"B0 IV",
		"B1 IV",
		"B2 IV",
		"B3 IV",
		"B4 IV",
		"B5 IV",
		"B6 IV",
		"B7 IV",
		"B8 IV",
		"B9 IV",
		"A0 IV",
		"A1 IV",
		"A2 IV",
		"A3 IV",
		"A4 IV",
		"A5 IV",
		"A6 IV",
		"A7 IV",
		"A8 IV",
		"A9 IV",
		"F0 IV",
		"F1 IV",
		"F2 IV",
		"F3 IV",
		"F4 IV",
		"F5 IV",
		"F6 IV",
		"F7 IV",
		"F8 IV",
		"F9 IV",
		"G0 IV",
		"G1 IV",
		"G2 IV",
		"G3 IV",
		"G4 IV",
		"G5 IV",
		"G6 IV",
		"G7 IV",
		"G8 IV",
		"G9 IV",
		"K0 IV",
		"K1 IV",
		"K2 IV",
		"K3 IV",
		"K4 IV",
		"K5 IV",
		"K6 IV",
		"K7 IV",
		"K8 IV",
		"K9 IV",
		"M0 IV",
		"M1 IV",
		"M2 IV",
		"M3 IV",
		"M4 IV",
		"M5 IV",
		"M6 IV",
		"M7 IV",
		"M8 IV",
		"M9 IV",
		"G0 VI",
		"G1 VI",
		"G2 VI",
		"G3 VI",
		"G4 VI",
		"G5 VI",
		"G6 VI",
		"G7 VI",
		"G8 VI",
		"G9 VI",
		"K0 VI",
		"K1 VI",
		"K2 VI",
		"K3 VI",
		"K4 VI",
		"K5 VI",
		"K6 VI",
		"K7 VI",
		"K8 VI",
		"K9 VI",
		"M0 VI",
		"M1 VI",
		"M2 VI",
		"M3 VI",
		"M4 VI",
		"M5 VI",
		"M6 VI",
		"M7 VI",
		"M8 VI",
		"M9 VI",
		"D0",
		"D1",
		"D2",
		"D3",
		"D4",
		"D5",
		"D6",
		"D7",
		"D8",
		"D9",
		"L0",
		"L1",
		"L2",
		"L3",
		"L4",
		"L5",
		"L6",
		"L7",
		"L8",
		"L9",
		"T0",
		"T1",
		"T2",
		"T3",
		"T4",
		"T5",
		"T6",
		"T7",
		"T8",
		"T9",
		"Y0",
		"Y1",
		"Y2",
		"Y3",
		"Y4",
		"Y5",
		"Y6",
		"Y7",
		"Y8",
		"Y9",
		"D0",
		"D1",
		"D2",
		"D3",
		"D4",
		"D5",
		"D6",
		"D7",
		"D8",
		"D9",
		"L0",
		"L1",
		"L2",
		"L3",
		"L4",
		"L5",
		"L6",
		"L7",
		"L8",
		"L9",
	}

	return list
}
