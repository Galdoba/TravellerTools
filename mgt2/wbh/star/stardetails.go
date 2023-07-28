package star

import (
	"fmt"
	"math"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func massOf(st Star, dice *dice.Dicepool) float64 {
	mass := -0.1
	switch st.Class {
	case ClassIa, ClassIb, ClassII, ClassIII, ClassIV, ClassV, ClassVI:
		averageMass := averageMassMap(st)
		flux := dice.Flux()
		variance := (averageMass / 100) * (4 * flux)
		switch st.Class {
		case ClassIa, ClassIb, ClassII, ClassIII:
			flux = (dice.Flux() * 10) + dice.Flux()
			variance = (averageMass / 100) * flux
		}
		mass = float64(averageMass+variance) / 1000000
	case ClassBD:
		r1 := float64(dice.Sroll("1d6")) * 0.01
		r2 := float64(dice.Sroll("4d6-1")) * 0.001
		r1Int := int(r1 * 100)
		r1 = float64(r1Int / 100)
		r2Int := int(r2 * 1000)
		r2 = float64(r2Int / 1000)
		mass = r1 + r2
	case ClassD:
		mass = float64((dice.Sroll("2d6")-1)/10) + ((float64(dice.Sroll("1d10"))) / 100)
	}

	return mass
}

func temperatureOf(st Star, dice *dice.Dicepool) int {
	temp := averageTempMap(st)
	flux := dice.Flux()
	variance := (temp / 200) * flux
	t := temp + variance
	if t <= -1 && st.Class == ClassD {
		temp = wdTemp(st.Age)
	}
	return temp + variance
}

func wdTemp(age float64) int {
	ageTempMap := make(map[float64]int)
	ageTempMap[0] = 100000
	ageTempMap[0.1] = 25000
	ageTempMap[0.5] = 10000
	ageTempMap[1] = 8000
	ageTempMap[1.5] = 7000
	ageTempMap[2.5] = 5500
	ageTempMap[5] = 5000
	ageTempMap[10] = 4000
	ageTempMap[13] = 3800
	ageTempMap[13.8] = 3725
	ageLow := -1.0
	ageHigh := 20.0
	for k, _ := range ageTempMap {
		if k >= ageLow && k < age {
			ageLow = k
		}
		if k <= ageHigh && k > age {
			ageHigh = k
		}
	}
	// fmt.Println(ageLow, "|", age, "|", ageHigh)
	tDiff := ageTempMap[ageHigh] - ageTempMap[ageLow]
	aDiff := age / ageHigh
	// fmt.Println(ageTempMap[ageHigh], tDiff, aDiff)
	return ageTempMap[ageLow] + int(float64(tDiff)*aDiff)
}

func diameterOf(st Star, dice *dice.Dicepool) float64 {
	diameter := -1.0
	switch st.Class {
	case ClassIa, ClassIb, ClassII, ClassIII, ClassIV, ClassV, ClassVI:
		diam := averageDiameterMap(st)
		flux := dice.Flux()
		variance := (diam / 100) * (4 * flux)
		diameter = float64(diam+variance) / 1000000
	case ClassBD:
		diam := 60000 + (dice.Flux() * 1000)
		diameter = float64(diam) / 1000000
	case ClassD:
		diameter = 0.017
	}

	return diameter
}

func luminocityOf(st Star) float64 {
	diamRatio := st.Diameter * st.Diameter
	tempRatio := math.Pow(float64(st.Temperature)/5772, 4)
	lum := diamRatio * tempRatio
	lum = float64(int(lum*1000000)) / 1000000
	return lum
}

func ageOf(st Star, dice *dice.Dicepool) float64 {
	age := -0.1
	switch st.Class {

	case ClassD:
		//fullMass := float64(2+dice.Sroll("1d3")) * st.Mass
		age = smallStarAge(dice) + starFinalAge(st.Mass, dice)
	case ClassV, ClassBD, ClassVI, ClassII, ClassIb, ClassIa:

		switch st.Mass <= 0.9 {
		case true:
			age = smallStarAge(dice)
		case false:
			age = largeStarAge(st.Mass, dice)
		}
	case ClassIV:
		age = mainSeqLifespan(st.Mass) + (subGigantLifespan(st.Mass) * d100variance(dice))
	case ClassIII:
		age = mainSeqLifespan(st.Mass) + subGigantLifespan(st.Mass) + (gigantLifespan(st.Mass) * d100variance(dice))
	}
	switch st.Specialcase {
	case Pulsar:
		age = (0.1/float64(2*dice.Sroll("1d10")) + starFinalAge(st.Mass, dice))
	case NeutronStar, BlackHole:
		age = smallStarAge(dice) + starFinalAge(st.Mass, dice)
	case Protostar:
		age = 0.01 / float64(dice.Sroll("2d10"))
	}

	age = age * 1000
	ageInt := int(age)
	age = float64(ageInt) / 1000
	if age > 13.5 {
		age = 13.5
	}
	return age
}

func starFinalAge(mass float64, dice *dice.Dicepool) float64 {
	m := float64(dice.Sroll("1d3")+2) * mass
	return mainSeqLifespan(m) + subGigantLifespan(m) + gigantLifespan(m)
}

func mainSeqLifespan(mass float64) float64 {
	return 10.0 / math.Pow(mass, 2.5)
}

func subGigantLifespan(mass float64) float64 {
	return mainSeqLifespan(mass) / (4 + mass)
}

func gigantLifespan(mass float64) float64 {
	return mainSeqLifespan(mass) / (10.0 * math.Pow(mass, 3))
}

func smallStarAge(dice *dice.Dicepool) float64 {
	return float64((dice.Sroll("1d6")*2)+dice.Sroll("1d3")-2) + (float64(dice.Sroll("1d10")) / 10)
}

func largeStarAge(mass float64, dice *dice.Dicepool) float64 {
	return mainSeqLifespan(mass) * d100variance(dice)
}

func d100variance(dice *dice.Dicepool) float64 {
	return float64(dice.Sroll("1d100")) / 100.0
}

func evaluateBDClassData(mass float64) (string, string) {
	for i, l := range []float64{0.08, 0.076, 0.072, 0.068, 0.064} {
		if mass >= l {
			return TypeL, fmt.Sprintf("%v", i)
		}
	}
	for i, l := range []float64{0.06, 0.058, 0.056, 0.054, 0.052} {
		if mass >= l {
			return TypeL, fmt.Sprintf("%v", i+5)
		}
	}

	for i, l := range []float64{0.05, 0.048, 0.046, 0.044, 0.042} {
		if mass >= l {
			return TypeT, fmt.Sprintf("%v", i)
		}
	}
	for i, l := range []float64{0.04, 0.037, 0.034, 0.031, 0.028} {
		if mass >= l {
			return TypeT, fmt.Sprintf("%v", i+5)
		}
	}
	for i, l := range []float64{0.025, 0.0226, 0.0202, 0.0178, 0.0154} {
		if mass >= l {
			return TypeY, fmt.Sprintf("%v", i)
		}
	}
	return TypeY, "5"

}

func averageDiameterMap(st Star) int {
	diameterMap := make(map[string]int)
	//KBIYTT
	diameterMap["O0 Ia"] = 25000000
	diameterMap["O5 Ia"] = 22000000
	diameterMap["B0 Ia"] = 20000000
	diameterMap["B5 Ia"] = 60000000
	diameterMap["A0 Ia"] = 120000000
	diameterMap["A5 Ia"] = 180000000
	diameterMap["F0 Ia"] = 210000000
	diameterMap["F5 Ia"] = 280000000
	diameterMap["G0 Ia"] = 330000000
	diameterMap["G5 Ia"] = 360000000
	diameterMap["K0 Ia"] = 420000000
	diameterMap["K5 Ia"] = 600000000
	diameterMap["M0 Ia"] = 900000000
	diameterMap["M5 Ia"] = 1200000000
	diameterMap["M9 Ia"] = 1800000000
	diameterMap["O0 Ib"] = 24000000
	diameterMap["O5 Ib"] = 20000000
	diameterMap["B0 Ib"] = 14000000
	diameterMap["B5 Ib"] = 25000000
	diameterMap["A0 Ib"] = 50000000
	diameterMap["A5 Ib"] = 75000000
	diameterMap["F0 Ib"] = 85000000
	diameterMap["F5 Ib"] = 115000000
	diameterMap["G0 Ib"] = 135000000
	diameterMap["G5 Ib"] = 150000000
	diameterMap["K0 Ib"] = 180000000
	diameterMap["K5 Ib"] = 260000000
	diameterMap["M0 Ib"] = 380000000
	diameterMap["M5 Ib"] = 600000000
	diameterMap["M9 Ib"] = 800000000
	diameterMap["O0 II"] = 22000000
	diameterMap["O5 II"] = 18000000
	diameterMap["B0 II"] = 12000000
	diameterMap["B5 II"] = 14000000
	diameterMap["A0 II"] = 30000000
	diameterMap["A5 II"] = 45000000
	diameterMap["F0 II"] = 50000000
	diameterMap["F5 II"] = 66000000
	diameterMap["G0 II"] = 77000000
	diameterMap["G5 II"] = 90000000
	diameterMap["K0 II"] = 110000000
	diameterMap["K5 II"] = 160000000
	diameterMap["M0 II"] = 230000000
	diameterMap["M5 II"] = 350000000
	diameterMap["M9 II"] = 500000000
	diameterMap["O0 III"] = 21000000
	diameterMap["O5 III"] = 15000000
	diameterMap["B0 III"] = 10000000
	diameterMap["B5 III"] = 6000000
	diameterMap["A0 III"] = 5000000
	diameterMap["A5 III"] = 5000000
	diameterMap["F0 III"] = 5000000
	diameterMap["F5 III"] = 5000000
	diameterMap["G0 III"] = 10000000
	diameterMap["G5 III"] = 15000000
	diameterMap["K0 III"] = 20000000
	diameterMap["K5 III"] = 40000000
	diameterMap["M0 III"] = 60000000
	diameterMap["M5 III"] = 100000000
	diameterMap["M9 III"] = 200000000
	diameterMap["O0 V"] = 20000000
	diameterMap["O5 V"] = 12000000
	diameterMap["B0 V"] = 7000000
	diameterMap["B5 V"] = 3500000
	diameterMap["A0 V"] = 2200000
	diameterMap["A5 V"] = 2000000
	diameterMap["F0 V"] = 1700000
	diameterMap["F5 V"] = 1500000
	diameterMap["G0 V"] = 1100000
	diameterMap["G5 V"] = 950000
	diameterMap["K0 V"] = 900000
	diameterMap["K5 V"] = 800000
	diameterMap["M0 V"] = 700000
	diameterMap["M5 V"] = 200000
	diameterMap["M9 V"] = 100000
	diameterMap["B0 IV"] = 8000000
	diameterMap["B5 IV"] = 5000000
	diameterMap["A0 IV"] = 4000000
	diameterMap["A5 IV"] = 3000000
	diameterMap["F0 IV"] = 3000000
	diameterMap["F5 IV"] = 2000000
	diameterMap["G0 IV"] = 3000000
	diameterMap["G5 IV"] = 3000000
	diameterMap["K0 IV"] = 4000000
	diameterMap["K4 IV"] = 5000000
	diameterMap["O0 VI"] = 180000
	diameterMap["O5 VI"] = 180000
	diameterMap["B0 VI"] = 200000
	diameterMap["B5 VI"] = 500000
	diameterMap["B9 VI"] = 800000
	diameterMap["G0 VI"] = 800000
	diameterMap["G5 VI"] = 700000
	diameterMap["K0 VI"] = 600000
	diameterMap["K5 VI"] = 500000
	diameterMap["M0 VI"] = 400000
	diameterMap["M5 VI"] = 100000
	diameterMap["M9 VI"] = 80000

	diameterMap["L0"] = 100000
	diameterMap["L5"] = 80000
	diameterMap["T0"] = 90000
	diameterMap["T5"] = 110000
	diameterMap["Y0"] = 100000
	diameterMap["Y5"] = 100000
	diameterMap, err := extrapolate(diameterMap)
	if err != nil {
		panic(err.Error())
	}
	return diameterMap[ShortStarDescription(st)]
}

func averageMassMap(st Star) int {
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

	massMap["L0"] = 80000
	massMap["L5"] = 60000
	massMap["T0"] = 50000
	massMap["T5"] = 40000
	massMap["Y0"] = 25000
	massMap["Y5"] = 13000
	massMap, err := extrapolate(massMap)
	if err != nil {
		panic(err.Error())
	}
	return massMap[ShortStarDescription(st)]
}

func averageTempMap(st Star) int {
	temperatureMap := make(map[string]int)
	//KBIYTT
	temperatureMap["O0"] = 50000
	temperatureMap["O5"] = 40000
	temperatureMap["B0"] = 30000
	temperatureMap["B5"] = 15000
	temperatureMap["A0"] = 10000
	temperatureMap["A5"] = 8000
	temperatureMap["F0"] = 7500
	temperatureMap["F5"] = 6500
	temperatureMap["G0"] = 6000
	temperatureMap["G5"] = 5600
	temperatureMap["K0"] = 5200
	temperatureMap["K5"] = 4400
	temperatureMap["M0"] = 3700
	temperatureMap["M5"] = 3000
	temperatureMap["M9"] = 2400
	temperatureMap["L0"] = 2400
	temperatureMap["L5"] = 1850
	temperatureMap["T0"] = 1300
	temperatureMap["T5"] = 900
	temperatureMap["Y0"] = 550
	temperatureMap["Y5"] = 300

	temperatureMap["O1"] = temperatureMap["O5"] + ((temperatureMap["O0"] - temperatureMap["O5"]) / 5 * 4)
	temperatureMap["O2"] = temperatureMap["O5"] + ((temperatureMap["O0"] - temperatureMap["O5"]) / 5 * 3)
	temperatureMap["O3"] = temperatureMap["O5"] + ((temperatureMap["O0"] - temperatureMap["O5"]) / 5 * 2)
	temperatureMap["O4"] = temperatureMap["O5"] + ((temperatureMap["O0"] - temperatureMap["O5"]) / 5 * 1)
	temperatureMap["B1"] = temperatureMap["B5"] + ((temperatureMap["B0"] - temperatureMap["B5"]) / 5 * 4)
	temperatureMap["B2"] = temperatureMap["B5"] + ((temperatureMap["B0"] - temperatureMap["B5"]) / 5 * 3)
	temperatureMap["B3"] = temperatureMap["B5"] + ((temperatureMap["B0"] - temperatureMap["B5"]) / 5 * 2)
	temperatureMap["B4"] = temperatureMap["B5"] + ((temperatureMap["B0"] - temperatureMap["B5"]) / 5 * 1)
	temperatureMap["A1"] = temperatureMap["A5"] + ((temperatureMap["A0"] - temperatureMap["A5"]) / 5 * 4)
	temperatureMap["A2"] = temperatureMap["A5"] + ((temperatureMap["A0"] - temperatureMap["A5"]) / 5 * 3)
	temperatureMap["A3"] = temperatureMap["A5"] + ((temperatureMap["A0"] - temperatureMap["A5"]) / 5 * 2)
	temperatureMap["A4"] = temperatureMap["A5"] + ((temperatureMap["A0"] - temperatureMap["A5"]) / 5 * 1)
	temperatureMap["F1"] = temperatureMap["F5"] + ((temperatureMap["F0"] - temperatureMap["F5"]) / 5 * 4)
	temperatureMap["F2"] = temperatureMap["F5"] + ((temperatureMap["F0"] - temperatureMap["F5"]) / 5 * 3)
	temperatureMap["F3"] = temperatureMap["F5"] + ((temperatureMap["F0"] - temperatureMap["F5"]) / 5 * 2)
	temperatureMap["F4"] = temperatureMap["F5"] + ((temperatureMap["F0"] - temperatureMap["F5"]) / 5 * 1)
	temperatureMap["G1"] = temperatureMap["G5"] + ((temperatureMap["G0"] - temperatureMap["G5"]) / 5 * 4)
	temperatureMap["G2"] = temperatureMap["G5"] + ((temperatureMap["G0"] - temperatureMap["G5"]) / 5 * 3)
	temperatureMap["G3"] = temperatureMap["G5"] + ((temperatureMap["G0"] - temperatureMap["G5"]) / 5 * 2)
	temperatureMap["G4"] = temperatureMap["G5"] + ((temperatureMap["G0"] - temperatureMap["G5"]) / 5 * 1)
	temperatureMap["K1"] = temperatureMap["K5"] + ((temperatureMap["K0"] - temperatureMap["K5"]) / 5 * 4)
	temperatureMap["K2"] = temperatureMap["K5"] + ((temperatureMap["K0"] - temperatureMap["K5"]) / 5 * 3)
	temperatureMap["K3"] = temperatureMap["K5"] + ((temperatureMap["K0"] - temperatureMap["K5"]) / 5 * 2)
	temperatureMap["K4"] = temperatureMap["K5"] + ((temperatureMap["K0"] - temperatureMap["K5"]) / 5 * 1)
	temperatureMap["M1"] = temperatureMap["M5"] + ((temperatureMap["M0"] - temperatureMap["M5"]) / 5 * 4)
	temperatureMap["M2"] = temperatureMap["M5"] + ((temperatureMap["M0"] - temperatureMap["M5"]) / 5 * 3)
	temperatureMap["M3"] = temperatureMap["M5"] + ((temperatureMap["M0"] - temperatureMap["M5"]) / 5 * 2)
	temperatureMap["M4"] = temperatureMap["M5"] + ((temperatureMap["M0"] - temperatureMap["M5"]) / 5 * 1)

	temperatureMap["O6"] = temperatureMap["B0"] + ((temperatureMap["O5"] - temperatureMap["B0"]) / 5 * 4)
	temperatureMap["O7"] = temperatureMap["B0"] + ((temperatureMap["O5"] - temperatureMap["B0"]) / 5 * 3)
	temperatureMap["O8"] = temperatureMap["B0"] + ((temperatureMap["O5"] - temperatureMap["B0"]) / 5 * 2)
	temperatureMap["O9"] = temperatureMap["B0"] + ((temperatureMap["O5"] - temperatureMap["B0"]) / 5 * 1)
	temperatureMap["B6"] = temperatureMap["A0"] + ((temperatureMap["B5"] - temperatureMap["A0"]) / 5 * 4)
	temperatureMap["B7"] = temperatureMap["A0"] + ((temperatureMap["B5"] - temperatureMap["A0"]) / 5 * 3)
	temperatureMap["B8"] = temperatureMap["A0"] + ((temperatureMap["B5"] - temperatureMap["A0"]) / 5 * 2)
	temperatureMap["B9"] = temperatureMap["A0"] + ((temperatureMap["B5"] - temperatureMap["A0"]) / 5 * 1)
	temperatureMap["A6"] = temperatureMap["F0"] + ((temperatureMap["A5"] - temperatureMap["F0"]) / 5 * 4)
	temperatureMap["A7"] = temperatureMap["F0"] + ((temperatureMap["A5"] - temperatureMap["F0"]) / 5 * 3)
	temperatureMap["A8"] = temperatureMap["F0"] + ((temperatureMap["A5"] - temperatureMap["F0"]) / 5 * 2)
	temperatureMap["A9"] = temperatureMap["F0"] + ((temperatureMap["A5"] - temperatureMap["F0"]) / 5 * 1)
	temperatureMap["F6"] = temperatureMap["G0"] + ((temperatureMap["F5"] - temperatureMap["G0"]) / 5 * 4)
	temperatureMap["F7"] = temperatureMap["G0"] + ((temperatureMap["F5"] - temperatureMap["G0"]) / 5 * 3)
	temperatureMap["F8"] = temperatureMap["G0"] + ((temperatureMap["F5"] - temperatureMap["G0"]) / 5 * 2)
	temperatureMap["F9"] = temperatureMap["G0"] + ((temperatureMap["F5"] - temperatureMap["G0"]) / 5 * 1)
	temperatureMap["G6"] = temperatureMap["K0"] + ((temperatureMap["G5"] - temperatureMap["K0"]) / 5 * 4)
	temperatureMap["G7"] = temperatureMap["K0"] + ((temperatureMap["G5"] - temperatureMap["K0"]) / 5 * 3)
	temperatureMap["G8"] = temperatureMap["K0"] + ((temperatureMap["G5"] - temperatureMap["K0"]) / 5 * 2)
	temperatureMap["G9"] = temperatureMap["K0"] + ((temperatureMap["G5"] - temperatureMap["K0"]) / 5 * 1)
	temperatureMap["K6"] = temperatureMap["M0"] + ((temperatureMap["K5"] - temperatureMap["M0"]) / 5 * 4)
	temperatureMap["K7"] = temperatureMap["M0"] + ((temperatureMap["K5"] - temperatureMap["M0"]) / 5 * 3)
	temperatureMap["K8"] = temperatureMap["M0"] + ((temperatureMap["K5"] - temperatureMap["M0"]) / 5 * 2)
	temperatureMap["K9"] = temperatureMap["M0"] + ((temperatureMap["K5"] - temperatureMap["M0"]) / 5 * 1)
	temperatureMap["M6"] = temperatureMap["M9"] + ((temperatureMap["M5"] - temperatureMap["M9"]) / 4 * 3)
	temperatureMap["M7"] = temperatureMap["M9"] + ((temperatureMap["M5"] - temperatureMap["M9"]) / 4 * 2)
	temperatureMap["M8"] = temperatureMap["M9"] + ((temperatureMap["M5"] - temperatureMap["M9"]) / 4 * 1)
	short := ShortStarDescription(st)
	temp := -1
	for k, v := range temperatureMap {
		if !strings.HasPrefix(short, k) {
			continue
		}
		return v
	}
	return temp
}

func basicKeys() []string {
	Classes := []string{
		ClassIa, ClassIb, ClassII, ClassIII, ClassV,
	}
	Types := []string{"O0", "O5", "B0", "B5", "A0", "A5", "F0", "F5", "G0", "G5", "K0", "K5", "M0", "M5", "M9"}
	keys := []string{}
	for _, Class := range Classes {
		for _, tpe := range Types {
			keys = append(keys, tpe+strings.TrimPrefix(Class, "Class"))
		}
	}
	for _, cl4 := range []string{"B0", "B5", "A0", "A5", "F0", "F5", "G0", "G5", "K0", "K4"} {
		keys = append(keys, cl4+" IV")
	}
	for _, cl6 := range []string{"O0", "O5", "B0", "B5", "B9", "G0", "G5", "K0", "K5", "M0", "M5", "M9"} {
		keys = append(keys, cl6+" VI")
	}
	return keys
}

func extrapolate(data map[string]int) (map[string]int, error) {
	for _, key := range basicKeys() {
		if data[key] == 0 {
			return nil, fmt.Errorf("key [%v] is not filled", key)
		}
	}
	data["O1 Ia"] = data["O5 Ia"] + ((data["O0 Ia"] - data["O5 Ia"]) / 5 * 4)
	data["O2 Ia"] = data["O5 Ia"] + ((data["O0 Ia"] - data["O5 Ia"]) / 5 * 3)
	data["O3 Ia"] = data["O5 Ia"] + ((data["O0 Ia"] - data["O5 Ia"]) / 5 * 2)
	data["O4 Ia"] = data["O5 Ia"] + ((data["O0 Ia"] - data["O5 Ia"]) / 5 * 1)
	data["B1 Ia"] = data["B5 Ia"] + ((data["B0 Ia"] - data["B5 Ia"]) / 5 * 4)
	data["B2 Ia"] = data["B5 Ia"] + ((data["B0 Ia"] - data["B5 Ia"]) / 5 * 3)
	data["B3 Ia"] = data["B5 Ia"] + ((data["B0 Ia"] - data["B5 Ia"]) / 5 * 2)
	data["B4 Ia"] = data["B5 Ia"] + ((data["B0 Ia"] - data["B5 Ia"]) / 5 * 1)
	data["A1 Ia"] = data["A5 Ia"] + ((data["A0 Ia"] - data["A5 Ia"]) / 5 * 4)
	data["A2 Ia"] = data["A5 Ia"] + ((data["A0 Ia"] - data["A5 Ia"]) / 5 * 3)
	data["A3 Ia"] = data["A5 Ia"] + ((data["A0 Ia"] - data["A5 Ia"]) / 5 * 2)
	data["A4 Ia"] = data["A5 Ia"] + ((data["A0 Ia"] - data["A5 Ia"]) / 5 * 1)
	data["F1 Ia"] = data["F5 Ia"] + ((data["F0 Ia"] - data["F5 Ia"]) / 5 * 4)
	data["F2 Ia"] = data["F5 Ia"] + ((data["F0 Ia"] - data["F5 Ia"]) / 5 * 3)
	data["F3 Ia"] = data["F5 Ia"] + ((data["F0 Ia"] - data["F5 Ia"]) / 5 * 2)
	data["F4 Ia"] = data["F5 Ia"] + ((data["F0 Ia"] - data["F5 Ia"]) / 5 * 1)
	data["G1 Ia"] = data["G5 Ia"] + ((data["G0 Ia"] - data["G5 Ia"]) / 5 * 4)
	data["G2 Ia"] = data["G5 Ia"] + ((data["G0 Ia"] - data["G5 Ia"]) / 5 * 3)
	data["G3 Ia"] = data["G5 Ia"] + ((data["G0 Ia"] - data["G5 Ia"]) / 5 * 2)
	data["G4 Ia"] = data["G5 Ia"] + ((data["G0 Ia"] - data["G5 Ia"]) / 5 * 1)
	data["K1 Ia"] = data["K5 Ia"] + ((data["K0 Ia"] - data["K5 Ia"]) / 5 * 4)
	data["K2 Ia"] = data["K5 Ia"] + ((data["K0 Ia"] - data["K5 Ia"]) / 5 * 3)
	data["K3 Ia"] = data["K5 Ia"] + ((data["K0 Ia"] - data["K5 Ia"]) / 5 * 2)
	data["K4 Ia"] = data["K5 Ia"] + ((data["K0 Ia"] - data["K5 Ia"]) / 5 * 1)
	data["M1 Ia"] = data["M5 Ia"] + ((data["M0 Ia"] - data["M5 Ia"]) / 5 * 4)
	data["M2 Ia"] = data["M5 Ia"] + ((data["M0 Ia"] - data["M5 Ia"]) / 5 * 3)
	data["M3 Ia"] = data["M5 Ia"] + ((data["M0 Ia"] - data["M5 Ia"]) / 5 * 2)
	data["M4 Ia"] = data["M5 Ia"] + ((data["M0 Ia"] - data["M5 Ia"]) / 5 * 1)
	data["O6 Ia"] = data["B0 Ia"] + ((data["O5 Ia"] - data["B0 Ia"]) / 5 * 4)
	data["O7 Ia"] = data["B0 Ia"] + ((data["O5 Ia"] - data["B0 Ia"]) / 5 * 3)
	data["O8 Ia"] = data["B0 Ia"] + ((data["O5 Ia"] - data["B0 Ia"]) / 5 * 2)
	data["O9 Ia"] = data["B0 Ia"] + ((data["O5 Ia"] - data["B0 Ia"]) / 5 * 1)
	data["B6 Ia"] = data["A0 Ia"] + ((data["B5 Ia"] - data["A0 Ia"]) / 5 * 4)
	data["B7 Ia"] = data["A0 Ia"] + ((data["B5 Ia"] - data["A0 Ia"]) / 5 * 3)
	data["B8 Ia"] = data["A0 Ia"] + ((data["B5 Ia"] - data["A0 Ia"]) / 5 * 2)
	data["B9 Ia"] = data["A0 Ia"] + ((data["B5 Ia"] - data["A0 Ia"]) / 5 * 1)
	data["A6 Ia"] = data["F0 Ia"] + ((data["A5 Ia"] - data["F0 Ia"]) / 5 * 4)
	data["A7 Ia"] = data["F0 Ia"] + ((data["A5 Ia"] - data["F0 Ia"]) / 5 * 3)
	data["A8 Ia"] = data["F0 Ia"] + ((data["A5 Ia"] - data["F0 Ia"]) / 5 * 2)
	data["A9 Ia"] = data["F0 Ia"] + ((data["A5 Ia"] - data["F0 Ia"]) / 5 * 1)
	data["F6 Ia"] = data["G0 Ia"] + ((data["F5 Ia"] - data["G0 Ia"]) / 5 * 4)
	data["F7 Ia"] = data["G0 Ia"] + ((data["F5 Ia"] - data["G0 Ia"]) / 5 * 3)
	data["F8 Ia"] = data["G0 Ia"] + ((data["F5 Ia"] - data["G0 Ia"]) / 5 * 2)
	data["F9 Ia"] = data["G0 Ia"] + ((data["F5 Ia"] - data["G0 Ia"]) / 5 * 1)
	data["G6 Ia"] = data["K0 Ia"] + ((data["G5 Ia"] - data["K0 Ia"]) / 5 * 4)
	data["G7 Ia"] = data["K0 Ia"] + ((data["G5 Ia"] - data["K0 Ia"]) / 5 * 3)
	data["G8 Ia"] = data["K0 Ia"] + ((data["G5 Ia"] - data["K0 Ia"]) / 5 * 2)
	data["G9 Ia"] = data["K0 Ia"] + ((data["G5 Ia"] - data["K0 Ia"]) / 5 * 1)
	data["K6 Ia"] = data["M0 Ia"] + ((data["K5 Ia"] - data["M0 Ia"]) / 5 * 4)
	data["K7 Ia"] = data["M0 Ia"] + ((data["K5 Ia"] - data["M0 Ia"]) / 5 * 3)
	data["K8 Ia"] = data["M0 Ia"] + ((data["K5 Ia"] - data["M0 Ia"]) / 5 * 2)
	data["K9 Ia"] = data["M0 Ia"] + ((data["K5 Ia"] - data["M0 Ia"]) / 5 * 1)
	data["M6 Ia"] = data["M9 Ia"] + ((data["M5 Ia"] - data["M9 Ia"]) / 4 * 3)
	data["M7 Ia"] = data["M9 Ia"] + ((data["M5 Ia"] - data["M9 Ia"]) / 4 * 2)
	data["M8 Ia"] = data["M9 Ia"] + ((data["M5 Ia"] - data["M9 Ia"]) / 4 * 1)
	////////////////////Ib
	data["O1 Ib"] = data["O5 Ib"] + ((data["O0 Ib"] - data["O5 Ib"]) / 5 * 4)
	data["O2 Ib"] = data["O5 Ib"] + ((data["O0 Ib"] - data["O5 Ib"]) / 5 * 3)
	data["O3 Ib"] = data["O5 Ib"] + ((data["O0 Ib"] - data["O5 Ib"]) / 5 * 2)
	data["O4 Ib"] = data["O5 Ib"] + ((data["O0 Ib"] - data["O5 Ib"]) / 5 * 1)
	data["B1 Ib"] = data["B5 Ib"] + ((data["B0 Ib"] - data["B5 Ib"]) / 5 * 4)
	data["B2 Ib"] = data["B5 Ib"] + ((data["B0 Ib"] - data["B5 Ib"]) / 5 * 3)
	data["B3 Ib"] = data["B5 Ib"] + ((data["B0 Ib"] - data["B5 Ib"]) / 5 * 2)
	data["B4 Ib"] = data["B5 Ib"] + ((data["B0 Ib"] - data["B5 Ib"]) / 5 * 1)
	data["A1 Ib"] = data["A5 Ib"] + ((data["A0 Ib"] - data["A5 Ib"]) / 5 * 4)
	data["A2 Ib"] = data["A5 Ib"] + ((data["A0 Ib"] - data["A5 Ib"]) / 5 * 3)
	data["A3 Ib"] = data["A5 Ib"] + ((data["A0 Ib"] - data["A5 Ib"]) / 5 * 2)
	data["A4 Ib"] = data["A5 Ib"] + ((data["A0 Ib"] - data["A5 Ib"]) / 5 * 1)
	data["F1 Ib"] = data["F5 Ib"] + ((data["F0 Ib"] - data["F5 Ib"]) / 5 * 4)
	data["F2 Ib"] = data["F5 Ib"] + ((data["F0 Ib"] - data["F5 Ib"]) / 5 * 3)
	data["F3 Ib"] = data["F5 Ib"] + ((data["F0 Ib"] - data["F5 Ib"]) / 5 * 2)
	data["F4 Ib"] = data["F5 Ib"] + ((data["F0 Ib"] - data["F5 Ib"]) / 5 * 1)
	data["G1 Ib"] = data["G5 Ib"] + ((data["G0 Ib"] - data["G5 Ib"]) / 5 * 4)
	data["G2 Ib"] = data["G5 Ib"] + ((data["G0 Ib"] - data["G5 Ib"]) / 5 * 3)
	data["G3 Ib"] = data["G5 Ib"] + ((data["G0 Ib"] - data["G5 Ib"]) / 5 * 2)
	data["G4 Ib"] = data["G5 Ib"] + ((data["G0 Ib"] - data["G5 Ib"]) / 5 * 1)
	data["K1 Ib"] = data["K5 Ib"] + ((data["K0 Ib"] - data["K5 Ib"]) / 5 * 4)
	data["K2 Ib"] = data["K5 Ib"] + ((data["K0 Ib"] - data["K5 Ib"]) / 5 * 3)
	data["K3 Ib"] = data["K5 Ib"] + ((data["K0 Ib"] - data["K5 Ib"]) / 5 * 2)
	data["K4 Ib"] = data["K5 Ib"] + ((data["K0 Ib"] - data["K5 Ib"]) / 5 * 1)
	data["M1 Ib"] = data["M5 Ib"] + ((data["M0 Ib"] - data["M5 Ib"]) / 5 * 4)
	data["M2 Ib"] = data["M5 Ib"] + ((data["M0 Ib"] - data["M5 Ib"]) / 5 * 3)
	data["M3 Ib"] = data["M5 Ib"] + ((data["M0 Ib"] - data["M5 Ib"]) / 5 * 2)
	data["M4 Ib"] = data["M5 Ib"] + ((data["M0 Ib"] - data["M5 Ib"]) / 5 * 1)
	data["O6 Ib"] = data["B0 Ib"] + ((data["O5 Ib"] - data["B0 Ib"]) / 5 * 4)
	data["O7 Ib"] = data["B0 Ib"] + ((data["O5 Ib"] - data["B0 Ib"]) / 5 * 3)
	data["O8 Ib"] = data["B0 Ib"] + ((data["O5 Ib"] - data["B0 Ib"]) / 5 * 2)
	data["O9 Ib"] = data["B0 Ib"] + ((data["O5 Ib"] - data["B0 Ib"]) / 5 * 1)
	data["B6 Ib"] = data["A0 Ib"] + ((data["B5 Ib"] - data["A0 Ib"]) / 5 * 4)
	data["B7 Ib"] = data["A0 Ib"] + ((data["B5 Ib"] - data["A0 Ib"]) / 5 * 3)
	data["B8 Ib"] = data["A0 Ib"] + ((data["B5 Ib"] - data["A0 Ib"]) / 5 * 2)
	data["B9 Ib"] = data["A0 Ib"] + ((data["B5 Ib"] - data["A0 Ib"]) / 5 * 1)
	data["A6 Ib"] = data["F0 Ib"] + ((data["A5 Ib"] - data["F0 Ib"]) / 5 * 4)
	data["A7 Ib"] = data["F0 Ib"] + ((data["A5 Ib"] - data["F0 Ib"]) / 5 * 3)
	data["A8 Ib"] = data["F0 Ib"] + ((data["A5 Ib"] - data["F0 Ib"]) / 5 * 2)
	data["A9 Ib"] = data["F0 Ib"] + ((data["A5 Ib"] - data["F0 Ib"]) / 5 * 1)
	data["F6 Ib"] = data["G0 Ib"] + ((data["F5 Ib"] - data["G0 Ib"]) / 5 * 4)
	data["F7 Ib"] = data["G0 Ib"] + ((data["F5 Ib"] - data["G0 Ib"]) / 5 * 3)
	data["F8 Ib"] = data["G0 Ib"] + ((data["F5 Ib"] - data["G0 Ib"]) / 5 * 2)
	data["F9 Ib"] = data["G0 Ib"] + ((data["F5 Ib"] - data["G0 Ib"]) / 5 * 1)
	data["G6 Ib"] = data["K0 Ib"] + ((data["G5 Ib"] - data["K0 Ib"]) / 5 * 4)
	data["G7 Ib"] = data["K0 Ib"] + ((data["G5 Ib"] - data["K0 Ib"]) / 5 * 3)
	data["G8 Ib"] = data["K0 Ib"] + ((data["G5 Ib"] - data["K0 Ib"]) / 5 * 2)
	data["G9 Ib"] = data["K0 Ib"] + ((data["G5 Ib"] - data["K0 Ib"]) / 5 * 1)
	data["K6 Ib"] = data["M0 Ib"] + ((data["K5 Ib"] - data["M0 Ib"]) / 5 * 4)
	data["K7 Ib"] = data["M0 Ib"] + ((data["K5 Ib"] - data["M0 Ib"]) / 5 * 3)
	data["K8 Ib"] = data["M0 Ib"] + ((data["K5 Ib"] - data["M0 Ib"]) / 5 * 2)
	data["K9 Ib"] = data["M0 Ib"] + ((data["K5 Ib"] - data["M0 Ib"]) / 5 * 1)
	data["M6 Ib"] = data["M9 Ib"] + ((data["M5 Ib"] - data["M9 Ib"]) / 4 * 3)
	data["M7 Ib"] = data["M9 Ib"] + ((data["M5 Ib"] - data["M9 Ib"]) / 4 * 2)
	data["M8 Ib"] = data["M9 Ib"] + ((data["M5 Ib"] - data["M9 Ib"]) / 4 * 1)
	////////////////////II
	data["O1 II"] = data["O5 II"] + ((data["O0 II"] - data["O5 II"]) / 5 * 4)
	data["O2 II"] = data["O5 II"] + ((data["O0 II"] - data["O5 II"]) / 5 * 3)
	data["O3 II"] = data["O5 II"] + ((data["O0 II"] - data["O5 II"]) / 5 * 2)
	data["O4 II"] = data["O5 II"] + ((data["O0 II"] - data["O5 II"]) / 5 * 1)
	data["B1 II"] = data["B5 II"] + ((data["B0 II"] - data["B5 II"]) / 5 * 4)
	data["B2 II"] = data["B5 II"] + ((data["B0 II"] - data["B5 II"]) / 5 * 3)
	data["B3 II"] = data["B5 II"] + ((data["B0 II"] - data["B5 II"]) / 5 * 2)
	data["B4 II"] = data["B5 II"] + ((data["B0 II"] - data["B5 II"]) / 5 * 1)
	data["A1 II"] = data["A5 II"] + ((data["A0 II"] - data["A5 II"]) / 5 * 4)
	data["A2 II"] = data["A5 II"] + ((data["A0 II"] - data["A5 II"]) / 5 * 3)
	data["A3 II"] = data["A5 II"] + ((data["A0 II"] - data["A5 II"]) / 5 * 2)
	data["A4 II"] = data["A5 II"] + ((data["A0 II"] - data["A5 II"]) / 5 * 1)
	data["F1 II"] = data["F5 II"] + ((data["F0 II"] - data["F5 II"]) / 5 * 4)
	data["F2 II"] = data["F5 II"] + ((data["F0 II"] - data["F5 II"]) / 5 * 3)
	data["F3 II"] = data["F5 II"] + ((data["F0 II"] - data["F5 II"]) / 5 * 2)
	data["F4 II"] = data["F5 II"] + ((data["F0 II"] - data["F5 II"]) / 5 * 1)
	data["G1 II"] = data["G5 II"] + ((data["G0 II"] - data["G5 II"]) / 5 * 4)
	data["G2 II"] = data["G5 II"] + ((data["G0 II"] - data["G5 II"]) / 5 * 3)
	data["G3 II"] = data["G5 II"] + ((data["G0 II"] - data["G5 II"]) / 5 * 2)
	data["G4 II"] = data["G5 II"] + ((data["G0 II"] - data["G5 II"]) / 5 * 1)
	data["K1 II"] = data["K5 II"] + ((data["K0 II"] - data["K5 II"]) / 5 * 4)
	data["K2 II"] = data["K5 II"] + ((data["K0 II"] - data["K5 II"]) / 5 * 3)
	data["K3 II"] = data["K5 II"] + ((data["K0 II"] - data["K5 II"]) / 5 * 2)
	data["K4 II"] = data["K5 II"] + ((data["K0 II"] - data["K5 II"]) / 5 * 1)
	data["M1 II"] = data["M5 II"] + ((data["M0 II"] - data["M5 II"]) / 5 * 4)
	data["M2 II"] = data["M5 II"] + ((data["M0 II"] - data["M5 II"]) / 5 * 3)
	data["M3 II"] = data["M5 II"] + ((data["M0 II"] - data["M5 II"]) / 5 * 2)
	data["M4 II"] = data["M5 II"] + ((data["M0 II"] - data["M5 II"]) / 5 * 1)
	data["O6 II"] = data["B0 II"] + ((data["O5 II"] - data["B0 II"]) / 5 * 4)
	data["O7 II"] = data["B0 II"] + ((data["O5 II"] - data["B0 II"]) / 5 * 3)
	data["O8 II"] = data["B0 II"] + ((data["O5 II"] - data["B0 II"]) / 5 * 2)
	data["O9 II"] = data["B0 II"] + ((data["O5 II"] - data["B0 II"]) / 5 * 1)
	data["B6 II"] = data["A0 II"] + ((data["B5 II"] - data["A0 II"]) / 5 * 4)
	data["B7 II"] = data["A0 II"] + ((data["B5 II"] - data["A0 II"]) / 5 * 3)
	data["B8 II"] = data["A0 II"] + ((data["B5 II"] - data["A0 II"]) / 5 * 2)
	data["B9 II"] = data["A0 II"] + ((data["B5 II"] - data["A0 II"]) / 5 * 1)
	data["A6 II"] = data["F0 II"] + ((data["A5 II"] - data["F0 II"]) / 5 * 4)
	data["A7 II"] = data["F0 II"] + ((data["A5 II"] - data["F0 II"]) / 5 * 3)
	data["A8 II"] = data["F0 II"] + ((data["A5 II"] - data["F0 II"]) / 5 * 2)
	data["A9 II"] = data["F0 II"] + ((data["A5 II"] - data["F0 II"]) / 5 * 1)
	data["F6 II"] = data["G0 II"] + ((data["F5 II"] - data["G0 II"]) / 5 * 4)
	data["F7 II"] = data["G0 II"] + ((data["F5 II"] - data["G0 II"]) / 5 * 3)
	data["F8 II"] = data["G0 II"] + ((data["F5 II"] - data["G0 II"]) / 5 * 2)
	data["F9 II"] = data["G0 II"] + ((data["F5 II"] - data["G0 II"]) / 5 * 1)
	data["G6 II"] = data["K0 II"] + ((data["G5 II"] - data["K0 II"]) / 5 * 4)
	data["G7 II"] = data["K0 II"] + ((data["G5 II"] - data["K0 II"]) / 5 * 3)
	data["G8 II"] = data["K0 II"] + ((data["G5 II"] - data["K0 II"]) / 5 * 2)
	data["G9 II"] = data["K0 II"] + ((data["G5 II"] - data["K0 II"]) / 5 * 1)
	data["K6 II"] = data["M0 II"] + ((data["K5 II"] - data["M0 II"]) / 5 * 4)
	data["K7 II"] = data["M0 II"] + ((data["K5 II"] - data["M0 II"]) / 5 * 3)
	data["K8 II"] = data["M0 II"] + ((data["K5 II"] - data["M0 II"]) / 5 * 2)
	data["K9 II"] = data["M0 II"] + ((data["K5 II"] - data["M0 II"]) / 5 * 1)
	data["M6 II"] = data["M9 II"] + ((data["M5 II"] - data["M9 II"]) / 4 * 3)
	data["M7 II"] = data["M9 II"] + ((data["M5 II"] - data["M9 II"]) / 4 * 2)
	data["M8 II"] = data["M9 II"] + ((data["M5 II"] - data["M9 II"]) / 4 * 1)
	////////////////////III
	data["O1 III"] = data["O5 III"] + ((data["O0 III"] - data["O5 III"]) / 5 * 4)
	data["O2 III"] = data["O5 III"] + ((data["O0 III"] - data["O5 III"]) / 5 * 3)
	data["O3 III"] = data["O5 III"] + ((data["O0 III"] - data["O5 III"]) / 5 * 2)
	data["O4 III"] = data["O5 III"] + ((data["O0 III"] - data["O5 III"]) / 5 * 1)
	data["B1 III"] = data["B5 III"] + ((data["B0 III"] - data["B5 III"]) / 5 * 4)
	data["B2 III"] = data["B5 III"] + ((data["B0 III"] - data["B5 III"]) / 5 * 3)
	data["B3 III"] = data["B5 III"] + ((data["B0 III"] - data["B5 III"]) / 5 * 2)
	data["B4 III"] = data["B5 III"] + ((data["B0 III"] - data["B5 III"]) / 5 * 1)
	data["A1 III"] = data["A5 III"] + ((data["A0 III"] - data["A5 III"]) / 5 * 4)
	data["A2 III"] = data["A5 III"] + ((data["A0 III"] - data["A5 III"]) / 5 * 3)
	data["A3 III"] = data["A5 III"] + ((data["A0 III"] - data["A5 III"]) / 5 * 2)
	data["A4 III"] = data["A5 III"] + ((data["A0 III"] - data["A5 III"]) / 5 * 1)
	data["F1 III"] = data["F5 III"] + ((data["F0 III"] - data["F5 III"]) / 5 * 4)
	data["F2 III"] = data["F5 III"] + ((data["F0 III"] - data["F5 III"]) / 5 * 3)
	data["F3 III"] = data["F5 III"] + ((data["F0 III"] - data["F5 III"]) / 5 * 2)
	data["F4 III"] = data["F5 III"] + ((data["F0 III"] - data["F5 III"]) / 5 * 1)
	data["G1 III"] = data["G5 III"] + ((data["G0 III"] - data["G5 III"]) / 5 * 4)
	data["G2 III"] = data["G5 III"] + ((data["G0 III"] - data["G5 III"]) / 5 * 3)
	data["G3 III"] = data["G5 III"] + ((data["G0 III"] - data["G5 III"]) / 5 * 2)
	data["G4 III"] = data["G5 III"] + ((data["G0 III"] - data["G5 III"]) / 5 * 1)
	data["K1 III"] = data["K5 III"] + ((data["K0 III"] - data["K5 III"]) / 5 * 4)
	data["K2 III"] = data["K5 III"] + ((data["K0 III"] - data["K5 III"]) / 5 * 3)
	data["K3 III"] = data["K5 III"] + ((data["K0 III"] - data["K5 III"]) / 5 * 2)
	data["K4 III"] = data["K5 III"] + ((data["K0 III"] - data["K5 III"]) / 5 * 1)
	data["M1 III"] = data["M5 III"] + ((data["M0 III"] - data["M5 III"]) / 5 * 4)
	data["M2 III"] = data["M5 III"] + ((data["M0 III"] - data["M5 III"]) / 5 * 3)
	data["M3 III"] = data["M5 III"] + ((data["M0 III"] - data["M5 III"]) / 5 * 2)
	data["M4 III"] = data["M5 III"] + ((data["M0 III"] - data["M5 III"]) / 5 * 1)
	data["O6 III"] = data["B0 III"] + ((data["O5 III"] - data["B0 III"]) / 5 * 4)
	data["O7 III"] = data["B0 III"] + ((data["O5 III"] - data["B0 III"]) / 5 * 3)
	data["O8 III"] = data["B0 III"] + ((data["O5 III"] - data["B0 III"]) / 5 * 2)
	data["O9 III"] = data["B0 III"] + ((data["O5 III"] - data["B0 III"]) / 5 * 1)
	data["B6 III"] = data["A0 III"] + ((data["B5 III"] - data["A0 III"]) / 5 * 4)
	data["B7 III"] = data["A0 III"] + ((data["B5 III"] - data["A0 III"]) / 5 * 3)
	data["B8 III"] = data["A0 III"] + ((data["B5 III"] - data["A0 III"]) / 5 * 2)
	data["B9 III"] = data["A0 III"] + ((data["B5 III"] - data["A0 III"]) / 5 * 1)
	data["A6 III"] = data["F0 III"] + ((data["A5 III"] - data["F0 III"]) / 5 * 4)
	data["A7 III"] = data["F0 III"] + ((data["A5 III"] - data["F0 III"]) / 5 * 3)
	data["A8 III"] = data["F0 III"] + ((data["A5 III"] - data["F0 III"]) / 5 * 2)
	data["A9 III"] = data["F0 III"] + ((data["A5 III"] - data["F0 III"]) / 5 * 1)
	data["F6 III"] = data["G0 III"] + ((data["F5 III"] - data["G0 III"]) / 5 * 4)
	data["F7 III"] = data["G0 III"] + ((data["F5 III"] - data["G0 III"]) / 5 * 3)
	data["F8 III"] = data["G0 III"] + ((data["F5 III"] - data["G0 III"]) / 5 * 2)
	data["F9 III"] = data["G0 III"] + ((data["F5 III"] - data["G0 III"]) / 5 * 1)
	data["G6 III"] = data["K0 III"] + ((data["G5 III"] - data["K0 III"]) / 5 * 4)
	data["G7 III"] = data["K0 III"] + ((data["G5 III"] - data["K0 III"]) / 5 * 3)
	data["G8 III"] = data["K0 III"] + ((data["G5 III"] - data["K0 III"]) / 5 * 2)
	data["G9 III"] = data["K0 III"] + ((data["G5 III"] - data["K0 III"]) / 5 * 1)
	data["K6 III"] = data["M0 III"] + ((data["K5 III"] - data["M0 III"]) / 5 * 4)
	data["K7 III"] = data["M0 III"] + ((data["K5 III"] - data["M0 III"]) / 5 * 3)
	data["K8 III"] = data["M0 III"] + ((data["K5 III"] - data["M0 III"]) / 5 * 2)
	data["K9 III"] = data["M0 III"] + ((data["K5 III"] - data["M0 III"]) / 5 * 1)
	data["M6 III"] = data["M9 III"] + ((data["M5 III"] - data["M9 III"]) / 4 * 3)
	data["M7 III"] = data["M9 III"] + ((data["M5 III"] - data["M9 III"]) / 4 * 2)
	data["M8 III"] = data["M9 III"] + ((data["M5 III"] - data["M9 III"]) / 4 * 1)
	////////////////////V
	data["O1 V"] = data["O5 V"] + ((data["O0 V"] - data["O5 V"]) / 5 * 4)
	data["O2 V"] = data["O5 V"] + ((data["O0 V"] - data["O5 V"]) / 5 * 3)
	data["O3 V"] = data["O5 V"] + ((data["O0 V"] - data["O5 V"]) / 5 * 2)
	data["O4 V"] = data["O5 V"] + ((data["O0 V"] - data["O5 V"]) / 5 * 1)
	data["B1 V"] = data["B5 V"] + ((data["B0 V"] - data["B5 V"]) / 5 * 4)
	data["B2 V"] = data["B5 V"] + ((data["B0 V"] - data["B5 V"]) / 5 * 3)
	data["B3 V"] = data["B5 V"] + ((data["B0 V"] - data["B5 V"]) / 5 * 2)
	data["B4 V"] = data["B5 V"] + ((data["B0 V"] - data["B5 V"]) / 5 * 1)
	data["A1 V"] = data["A5 V"] + ((data["A0 V"] - data["A5 V"]) / 5 * 4)
	data["A2 V"] = data["A5 V"] + ((data["A0 V"] - data["A5 V"]) / 5 * 3)
	data["A3 V"] = data["A5 V"] + ((data["A0 V"] - data["A5 V"]) / 5 * 2)
	data["A4 V"] = data["A5 V"] + ((data["A0 V"] - data["A5 V"]) / 5 * 1)
	data["F1 V"] = data["F5 V"] + ((data["F0 V"] - data["F5 V"]) / 5 * 4)
	data["F2 V"] = data["F5 V"] + ((data["F0 V"] - data["F5 V"]) / 5 * 3)
	data["F3 V"] = data["F5 V"] + ((data["F0 V"] - data["F5 V"]) / 5 * 2)
	data["F4 V"] = data["F5 V"] + ((data["F0 V"] - data["F5 V"]) / 5 * 1)
	data["G1 V"] = data["G5 V"] + ((data["G0 V"] - data["G5 V"]) / 5 * 4)
	data["G2 V"] = data["G5 V"] + ((data["G0 V"] - data["G5 V"]) / 5 * 3)
	data["G3 V"] = data["G5 V"] + ((data["G0 V"] - data["G5 V"]) / 5 * 2)
	data["G4 V"] = data["G5 V"] + ((data["G0 V"] - data["G5 V"]) / 5 * 1)
	data["K1 V"] = data["K5 V"] + ((data["K0 V"] - data["K5 V"]) / 5 * 4)
	data["K2 V"] = data["K5 V"] + ((data["K0 V"] - data["K5 V"]) / 5 * 3)
	data["K3 V"] = data["K5 V"] + ((data["K0 V"] - data["K5 V"]) / 5 * 2)
	data["K4 V"] = data["K5 V"] + ((data["K0 V"] - data["K5 V"]) / 5 * 1)
	data["M1 V"] = data["M5 V"] + ((data["M0 V"] - data["M5 V"]) / 5 * 4)
	data["M2 V"] = data["M5 V"] + ((data["M0 V"] - data["M5 V"]) / 5 * 3)
	data["M3 V"] = data["M5 V"] + ((data["M0 V"] - data["M5 V"]) / 5 * 2)
	data["M4 V"] = data["M5 V"] + ((data["M0 V"] - data["M5 V"]) / 5 * 1)
	data["O6 V"] = data["B0 V"] + ((data["O5 V"] - data["B0 V"]) / 5 * 4)
	data["O7 V"] = data["B0 V"] + ((data["O5 V"] - data["B0 V"]) / 5 * 3)
	data["O8 V"] = data["B0 V"] + ((data["O5 V"] - data["B0 V"]) / 5 * 2)
	data["O9 V"] = data["B0 V"] + ((data["O5 V"] - data["B0 V"]) / 5 * 1)
	data["B6 V"] = data["A0 V"] + ((data["B5 V"] - data["A0 V"]) / 5 * 4)
	data["B7 V"] = data["A0 V"] + ((data["B5 V"] - data["A0 V"]) / 5 * 3)
	data["B8 V"] = data["A0 V"] + ((data["B5 V"] - data["A0 V"]) / 5 * 2)
	data["B9 V"] = data["A0 V"] + ((data["B5 V"] - data["A0 V"]) / 5 * 1)
	data["A6 V"] = data["F0 V"] + ((data["A5 V"] - data["F0 V"]) / 5 * 4)
	data["A7 V"] = data["F0 V"] + ((data["A5 V"] - data["F0 V"]) / 5 * 3)
	data["A8 V"] = data["F0 V"] + ((data["A5 V"] - data["F0 V"]) / 5 * 2)
	data["A9 V"] = data["F0 V"] + ((data["A5 V"] - data["F0 V"]) / 5 * 1)
	data["F6 V"] = data["G0 V"] + ((data["F5 V"] - data["G0 V"]) / 5 * 4)
	data["F7 V"] = data["G0 V"] + ((data["F5 V"] - data["G0 V"]) / 5 * 3)
	data["F8 V"] = data["G0 V"] + ((data["F5 V"] - data["G0 V"]) / 5 * 2)
	data["F9 V"] = data["G0 V"] + ((data["F5 V"] - data["G0 V"]) / 5 * 1)
	data["G6 V"] = data["K0 V"] + ((data["G5 V"] - data["K0 V"]) / 5 * 4)
	data["G7 V"] = data["K0 V"] + ((data["G5 V"] - data["K0 V"]) / 5 * 3)
	data["G8 V"] = data["K0 V"] + ((data["G5 V"] - data["K0 V"]) / 5 * 2)
	data["G9 V"] = data["K0 V"] + ((data["G5 V"] - data["K0 V"]) / 5 * 1)
	data["K6 V"] = data["M0 V"] + ((data["K5 V"] - data["M0 V"]) / 5 * 4)
	data["K7 V"] = data["M0 V"] + ((data["K5 V"] - data["M0 V"]) / 5 * 3)
	data["K8 V"] = data["M0 V"] + ((data["K5 V"] - data["M0 V"]) / 5 * 2)
	data["K9 V"] = data["M0 V"] + ((data["K5 V"] - data["M0 V"]) / 5 * 1)
	data["M6 V"] = data["M9 V"] + ((data["M5 V"] - data["M9 V"]) / 4 * 3)
	data["M7 V"] = data["M9 V"] + ((data["M5 V"] - data["M9 V"]) / 4 * 2)
	data["M8 V"] = data["M9 V"] + ((data["M5 V"] - data["M9 V"]) / 4 * 1)
	////////////////////IV
	data["B1 IV"] = data["B5 IV"] + ((data["B0 IV"] - data["B5 IV"]) / 5 * 4)
	data["B2 IV"] = data["B5 IV"] + ((data["B0 IV"] - data["B5 IV"]) / 5 * 3)
	data["B3 IV"] = data["B5 IV"] + ((data["B0 IV"] - data["B5 IV"]) / 5 * 2)
	data["B4 IV"] = data["B5 IV"] + ((data["B0 IV"] - data["B5 IV"]) / 5 * 1)
	data["A1 IV"] = data["A5 IV"] + ((data["A0 IV"] - data["A5 IV"]) / 5 * 4)
	data["A2 IV"] = data["A5 IV"] + ((data["A0 IV"] - data["A5 IV"]) / 5 * 3)
	data["A3 IV"] = data["A5 IV"] + ((data["A0 IV"] - data["A5 IV"]) / 5 * 2)
	data["A4 IV"] = data["A5 IV"] + ((data["A0 IV"] - data["A5 IV"]) / 5 * 1)
	data["F1 IV"] = data["F5 IV"] + ((data["F0 IV"] - data["F5 IV"]) / 5 * 4)
	data["F2 IV"] = data["F5 IV"] + ((data["F0 IV"] - data["F5 IV"]) / 5 * 3)
	data["F3 IV"] = data["F5 IV"] + ((data["F0 IV"] - data["F5 IV"]) / 5 * 2)
	data["F4 IV"] = data["F5 IV"] + ((data["F0 IV"] - data["F5 IV"]) / 5 * 1)
	data["G1 IV"] = data["G5 IV"] + ((data["G0 IV"] - data["G5 IV"]) / 5 * 4)
	data["G2 IV"] = data["G5 IV"] + ((data["G0 IV"] - data["G5 IV"]) / 5 * 3)
	data["G3 IV"] = data["G5 IV"] + ((data["G0 IV"] - data["G5 IV"]) / 5 * 2)
	data["G4 IV"] = data["G5 IV"] + ((data["G0 IV"] - data["G5 IV"]) / 5 * 1)
	data["K1 IV"] = data["K5 IV"] + ((data["K0 IV"] - data["K5 IV"]) / 4 * 3)
	data["K2 IV"] = data["K5 IV"] + ((data["K0 IV"] - data["K5 IV"]) / 4 * 2)
	data["K3 IV"] = data["K5 IV"] + ((data["K0 IV"] - data["K5 IV"]) / 4 * 1)
	data["B6 IV"] = data["A0 IV"] + ((data["B5 IV"] - data["A0 IV"]) / 5 * 4)
	data["B7 IV"] = data["A0 IV"] + ((data["B5 IV"] - data["A0 IV"]) / 5 * 3)
	data["B8 IV"] = data["A0 IV"] + ((data["B5 IV"] - data["A0 IV"]) / 5 * 2)
	data["B9 IV"] = data["A0 IV"] + ((data["B5 IV"] - data["A0 IV"]) / 5 * 1)
	data["A6 IV"] = data["F0 IV"] + ((data["A5 IV"] - data["F0 IV"]) / 5 * 4)
	data["A7 IV"] = data["F0 IV"] + ((data["A5 IV"] - data["F0 IV"]) / 5 * 3)
	data["A8 IV"] = data["F0 IV"] + ((data["A5 IV"] - data["F0 IV"]) / 5 * 2)
	data["A9 IV"] = data["F0 IV"] + ((data["A5 IV"] - data["F0 IV"]) / 5 * 1)
	data["F6 IV"] = data["G0 IV"] + ((data["F5 IV"] - data["G0 IV"]) / 5 * 4)
	data["F7 IV"] = data["G0 IV"] + ((data["F5 IV"] - data["G0 IV"]) / 5 * 3)
	data["F8 IV"] = data["G0 IV"] + ((data["F5 IV"] - data["G0 IV"]) / 5 * 2)
	data["F9 IV"] = data["G0 IV"] + ((data["F5 IV"] - data["G0 IV"]) / 5 * 1)
	data["G6 IV"] = data["K0 IV"] + ((data["G5 IV"] - data["K0 IV"]) / 5 * 4)
	data["G7 IV"] = data["K0 IV"] + ((data["G5 IV"] - data["K0 IV"]) / 5 * 3)
	data["G8 IV"] = data["K0 IV"] + ((data["G5 IV"] - data["K0 IV"]) / 5 * 2)
	data["G9 IV"] = data["K0 IV"] + ((data["G5 IV"] - data["K0 IV"]) / 5 * 1)
	////////////////////VI
	data["O1 VI"] = data["O5 VI"] + ((data["O0 VI"] - data["O5 VI"]) / 5 * 4)
	data["O2 VI"] = data["O5 VI"] + ((data["O0 VI"] - data["O5 VI"]) / 5 * 3)
	data["O3 VI"] = data["O5 VI"] + ((data["O0 VI"] - data["O5 VI"]) / 5 * 2)
	data["O4 VI"] = data["O5 VI"] + ((data["O0 VI"] - data["O5 VI"]) / 5 * 1)
	data["B1 VI"] = data["B5 VI"] + ((data["B0 VI"] - data["B5 VI"]) / 5 * 4)
	data["B2 VI"] = data["B5 VI"] + ((data["B0 VI"] - data["B5 VI"]) / 5 * 3)
	data["B3 VI"] = data["B5 VI"] + ((data["B0 VI"] - data["B5 VI"]) / 5 * 2)
	data["B4 VI"] = data["B5 VI"] + ((data["B0 VI"] - data["B5 VI"]) / 5 * 1)
	data["G1 VI"] = data["G5 VI"] + ((data["G0 VI"] - data["G5 VI"]) / 5 * 4)
	data["G2 VI"] = data["G5 VI"] + ((data["G0 VI"] - data["G5 VI"]) / 5 * 3)
	data["G3 VI"] = data["G5 VI"] + ((data["G0 VI"] - data["G5 VI"]) / 5 * 2)
	data["G4 VI"] = data["G5 VI"] + ((data["G0 VI"] - data["G5 VI"]) / 5 * 1)
	data["K1 VI"] = data["K5 VI"] + ((data["K0 VI"] - data["K5 VI"]) / 5 * 4)
	data["K2 VI"] = data["K5 VI"] + ((data["K0 VI"] - data["K5 VI"]) / 5 * 3)
	data["K3 VI"] = data["K5 VI"] + ((data["K0 VI"] - data["K5 VI"]) / 5 * 2)
	data["K4 VI"] = data["K5 VI"] + ((data["K0 VI"] - data["K5 VI"]) / 5 * 1)
	data["M1 VI"] = data["M5 VI"] + ((data["M0 VI"] - data["M5 VI"]) / 5 * 4)
	data["M2 VI"] = data["M5 VI"] + ((data["M0 VI"] - data["M5 VI"]) / 5 * 3)
	data["M3 VI"] = data["M5 VI"] + ((data["M0 VI"] - data["M5 VI"]) / 5 * 2)
	data["M4 VI"] = data["M5 VI"] + ((data["M0 VI"] - data["M5 VI"]) / 5 * 1)
	data["O6 VI"] = data["B0 VI"] + ((data["O5 VI"] - data["B0 VI"]) / 5 * 4)
	data["O7 VI"] = data["B0 VI"] + ((data["O5 VI"] - data["B0 VI"]) / 5 * 3)
	data["O8 VI"] = data["B0 VI"] + ((data["O5 VI"] - data["B0 VI"]) / 5 * 2)
	data["O9 VI"] = data["B0 VI"] + ((data["O5 VI"] - data["B0 VI"]) / 5 * 1)
	data["B6 VI"] = data["B9 VI"] + ((data["B5 VI"] - data["B9 VI"]) / 4 * 3)
	data["B7 VI"] = data["B9 VI"] + ((data["B5 VI"] - data["B9 VI"]) / 4 * 2)
	data["B8 VI"] = data["B9 VI"] + ((data["B5 VI"] - data["B9 VI"]) / 4 * 1)
	data["G6 VI"] = data["K0 VI"] + ((data["G5 VI"] - data["K0 VI"]) / 5 * 4)
	data["G7 VI"] = data["K0 VI"] + ((data["G5 VI"] - data["K0 VI"]) / 5 * 3)
	data["G8 VI"] = data["K0 VI"] + ((data["G5 VI"] - data["K0 VI"]) / 5 * 2)
	data["G9 VI"] = data["K0 VI"] + ((data["G5 VI"] - data["K0 VI"]) / 5 * 1)
	data["K6 VI"] = data["M0 VI"] + ((data["K5 VI"] - data["M0 VI"]) / 5 * 4)
	data["K7 VI"] = data["M0 VI"] + ((data["K5 VI"] - data["M0 VI"]) / 5 * 3)
	data["K8 VI"] = data["M0 VI"] + ((data["K5 VI"] - data["M0 VI"]) / 5 * 2)
	data["K9 VI"] = data["M0 VI"] + ((data["K5 VI"] - data["M0 VI"]) / 5 * 1)
	data["M6 VI"] = data["M9 VI"] + ((data["M5 VI"] - data["M9 VI"]) / 4 * 3)
	data["M7 VI"] = data["M9 VI"] + ((data["M5 VI"] - data["M9 VI"]) / 4 * 2)
	data["M8 VI"] = data["M9 VI"] + ((data["M5 VI"] - data["M9 VI"]) / 4 * 1)

	data["L1"] = data["L5"] + ((data["L0"] - data["L5"]) / 5 * 4)
	data["L2"] = data["L5"] + ((data["L0"] - data["L5"]) / 5 * 3)
	data["L3"] = data["L5"] + ((data["L0"] - data["L5"]) / 5 * 2)
	data["L4"] = data["L5"] + ((data["L0"] - data["L5"]) / 5 * 1)
	data["L6"] = data["T0"] + ((data["L5"] - data["T0"]) / 5 * 4)
	data["L7"] = data["T0"] + ((data["L5"] - data["T0"]) / 5 * 3)
	data["L8"] = data["T0"] + ((data["L5"] - data["T0"]) / 5 * 2)
	data["L9"] = data["T0"] + ((data["L5"] - data["T0"]) / 5 * 1)
	data["T1"] = data["T5"] + ((data["T0"] - data["T5"]) / 5 * 4)
	data["T2"] = data["T5"] + ((data["T0"] - data["T5"]) / 5 * 3)
	data["T3"] = data["T5"] + ((data["T0"] - data["T5"]) / 5 * 2)
	data["T4"] = data["T5"] + ((data["T0"] - data["T5"]) / 5 * 1)
	data["T6"] = data["Y0"] + ((data["T5"] - data["Y0"]) / 5 * 4)
	data["T7"] = data["Y0"] + ((data["T5"] - data["Y0"]) / 5 * 3)
	data["T8"] = data["Y0"] + ((data["T5"] - data["Y0"]) / 5 * 2)
	data["T9"] = data["Y0"] + ((data["T5"] - data["Y0"]) / 5 * 1)
	data["Y1"] = data["Y5"] + ((data["Y0"] - data["Y5"]) / 5 * 4)
	data["Y2"] = data["Y5"] + ((data["Y0"] - data["Y5"]) / 5 * 3)
	data["Y3"] = data["Y5"] + ((data["Y0"] - data["Y5"]) / 5 * 2)
	data["Y4"] = data["Y5"] + ((data["Y0"] - data["Y5"]) / 5 * 1)

	return data, nil
}

func maoOf(st Star) float64 {
	mao := -0.1
	switch st.Class {
	case ClassIa, ClassIb, ClassII, ClassIII, ClassIV, ClassV, ClassVI:
		averageMAO := averageMAOMap(st)
		mao = float64(averageMAO) / 1000
	case ClassBD:
		mao = 0.005
	case ClassD, BlackHole, NeutronStar, Pulsar:
		mao = 0.001
	}
	return mao
}

func averageMAOMap(st Star) int {
	massMap := make(map[string]int)
	//KBIYTT
	massMap["O0 Ia"] = 630
	massMap["O5 Ia"] = 550
	massMap["B0 Ia"] = 500
	massMap["B5 Ia"] = 1670
	massMap["A0 Ia"] = 3340
	massMap["A5 Ia"] = 4170
	massMap["F0 Ia"] = 4420
	massMap["F5 Ia"] = 5000
	massMap["G0 Ia"] = 5210
	massMap["G5 Ia"] = 5340
	massMap["K0 Ia"] = 5590
	massMap["K5 Ia"] = 6170
	massMap["M0 Ia"] = 6800
	massMap["M5 Ia"] = 7200
	massMap["M9 Ia"] = 7800
	massMap["O0 Ib"] = 600
	massMap["O5 Ib"] = 500
	massMap["B0 Ib"] = 350
	massMap["B5 Ib"] = 630
	massMap["A0 Ib"] = 1400
	massMap["A5 Ib"] = 2170
	massMap["F0 Ib"] = 2500
	massMap["F5 Ib"] = 3250
	massMap["G0 Ib"] = 3590
	massMap["G5 Ib"] = 3840
	massMap["K0 Ib"] = 4170
	massMap["K5 Ib"] = 4840
	massMap["M0 Ib"] = 5420
	massMap["M5 Ib"] = 6170
	massMap["M9 Ib"] = 6590
	massMap["O0 II"] = 550
	massMap["O5 II"] = 450
	massMap["B0 II"] = 300
	massMap["B5 II"] = 350
	massMap["A0 II"] = 750
	massMap["A5 II"] = 1170
	massMap["F0 II"] = 1330
	massMap["F5 II"] = 1870
	massMap["G0 II"] = 2240
	massMap["G5 II"] = 2670
	massMap["K0 II"] = 3170
	massMap["K5 II"] = 4000
	massMap["M0 II"] = 4590
	massMap["M5 II"] = 5300
	massMap["M9 II"] = 5920
	massMap["O0 III"] = 530
	massMap["O5 III"] = 380
	massMap["B0 III"] = 250
	massMap["B5 III"] = 150
	massMap["A0 III"] = 130
	massMap["A5 III"] = 130
	massMap["F0 III"] = 130
	massMap["F5 III"] = 130
	massMap["G0 III"] = 250
	massMap["G5 III"] = 380
	massMap["K0 III"] = 500
	massMap["K5 III"] = 1000
	massMap["M0 III"] = 1680
	massMap["M5 III"] = 3000
	massMap["M9 III"] = 4340
	massMap["O0 V"] = 500
	massMap["O5 V"] = 300
	massMap["B0 V"] = 180
	massMap["B5 V"] = 90
	massMap["A0 V"] = 60
	massMap["A5 V"] = 50
	massMap["F0 V"] = 40
	massMap["F5 V"] = 30
	massMap["G0 V"] = 30
	massMap["G5 V"] = 20
	massMap["K0 V"] = 20
	massMap["K5 V"] = 20
	massMap["M0 V"] = 20
	massMap["M5 V"] = 10
	massMap["M9 V"] = 10
	massMap["B0 IV"] = 200
	massMap["B5 IV"] = 130
	massMap["A0 IV"] = 100
	massMap["A5 IV"] = 70
	massMap["F0 IV"] = 70
	massMap["F5 IV"] = 60
	massMap["G0 IV"] = 70
	massMap["G5 IV"] = 100
	massMap["K0 IV"] = 150
	massMap["K4 IV"] = 200
	massMap["O0 VI"] = 10
	massMap["O5 VI"] = 10
	massMap["B0 VI"] = 10
	massMap["B5 VI"] = 10
	massMap["B9 VI"] = 10
	massMap["G0 VI"] = 20
	massMap["G5 VI"] = 20
	massMap["K0 VI"] = 20
	massMap["K5 VI"] = 10
	massMap["M0 VI"] = 10
	massMap["M5 VI"] = 10
	massMap["M9 VI"] = 10

	massMap["L0"] = 1
	massMap["L5"] = 1
	massMap["T0"] = 1
	massMap["T5"] = 1
	massMap["Y0"] = 1
	massMap["Y5"] = 1
	massMap, err := extrapolate(massMap)
	if err != nil {
		panic(err.Error())
	}
	return massMap[ShortStarDescription(st)]
}
