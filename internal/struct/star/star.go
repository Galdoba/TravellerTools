package star

import (
	"fmt"
	"strings"
)

const (
	UNDEFINED = iota
	Category_Primary
	Category_PrimaryCompanion
	Category_Close
	Category_CloseCompanion
	Category_Near
	Category_NearCompanion
	Category_Far
	Category_FarCompanion
)

type Star struct {
	name       string
	size       string
	decimal    string
	spectral   string
	mass       float64
	luminocity float64
	orbit      int
	category   int
	code       string
}

func New(name, code string, category int) (Star, error) {
	err := fmt.Errorf("NewRandom func is not implemented")
	s := Star{}
	s.name = strings.TrimSuffix(name+" "+categoryString(category), " ")

	s.category = category
	s.orbit = -2
	//dp := dice.New().SetSeed(name)
	switch s.category {
	default:
		return s, fmt.Errorf("star category undefined")
	case Category_Primary:

	case Category_PrimaryCompanion, Category_Close, Category_CloseCompanion, Category_Near, Category_NearCompanion, Category_Far, Category_FarCompanion:

	}
	s.code = code
	s.luminocity = baseStellarLuminocity(code)
	s.mass = baseStellarMass(code)
	err = s.checkStruct()
	return s, err
}

func categoryString(category int) string {
	switch category {
	case Category_Primary:
		return "Primary"
	case Category_PrimaryCompanion:
		return "Primary Companion"
	case Category_Close:
		return "Close"
	case Category_CloseCompanion:
		return "Close Companion"
	case Category_Near:
		return "Near"
	case Category_NearCompanion:
		return "Near Companion"
	case Category_Far:
		return "Far"
	case Category_FarCompanion:
		return "Far Companion"
	}
	return ""
}

func (s *Star) checkStruct() error {
	switch {
	case s.code == "":
		return fmt.Errorf("code undefined")
	case s.size == "":
		return fmt.Errorf("size undefined")
	case s.spectral == "":
		return fmt.Errorf("spectral undefined")
	case s.mass == 0:
		return fmt.Errorf("mass undefined")
	case s.luminocity == 0:
		return fmt.Errorf("luminocity undefined")
	case s.orbit == -2:
		return fmt.Errorf("orbit undefined")
	case s.category == UNDEFINED:
		return fmt.Errorf("category undefined")
	}
	return nil
}

func interpalation(a, b float64, dif, div float64) float64 {
	res := a - ((a - b) * dif / div)

	return res
}

func encodeStellar(spectral, dec, size string) string {
	switch {
	case spectral == "BD":
		return "BD"
	case size == "D":
		return size + spectral
	case dec == "0" || dec == "1" || dec == "2" || dec == "3" || dec == "4" || dec == "5" || dec == "6" || dec == "7" || dec == "8" || dec == "9":
		return spectral + dec + " " + size
	default:
		return "error"
	}
}

func baseStellarMass(class string) float64 {
	massMap := make(map[string]float64)
	massMap["B0 Ia"] = 60
	massMap["B1 Ia"] = 54
	massMap["B2 Ia"] = 48
	massMap["B3 Ia"] = 42
	massMap["B4 Ia"] = 36
	massMap["B5 Ia"] = 30
	massMap["B6 Ia"] = 27.6
	massMap["B7 Ia"] = 25.2
	massMap["B8 Ia"] = 22.8
	massMap["B9 Ia"] = 20.4
	massMap["A0 Ia"] = 18
	massMap["A1 Ia"] = 17.4
	massMap["A2 Ia"] = 16.8
	massMap["A3 Ia"] = 16.2
	massMap["A4 Ia"] = 15.6
	massMap["A5 Ia"] = 15
	massMap["A6 Ia"] = 14.6
	massMap["A7 Ia"] = 14.2
	massMap["A8 Ia"] = 13.8
	massMap["A9 Ia"] = 13.4
	massMap["F0 Ia"] = 13
	massMap["F1 Ia"] = 12.8
	massMap["F2 Ia"] = 12.6
	massMap["F3 Ia"] = 12.4
	massMap["F4 Ia"] = 12.2
	massMap["F5 Ia"] = 12
	massMap["F6 Ia"] = 12
	massMap["F7 Ia"] = 12
	massMap["F8 Ia"] = 12
	massMap["F9 Ia"] = 12
	massMap["G0 Ia"] = 12
	massMap["G1 Ia"] = 12.2
	massMap["G2 Ia"] = 12.4
	massMap["G3 Ia"] = 12.6
	massMap["G4 Ia"] = 12.8
	massMap["G5 Ia"] = 13
	massMap["G6 Ia"] = 13.2
	massMap["G7 Ia"] = 13.4
	massMap["G8 Ia"] = 13.6
	massMap["G9 Ia"] = 13.8
	massMap["K0 Ia"] = 14
	massMap["K1 Ia"] = 14.8
	massMap["K2 Ia"] = 15.6
	massMap["K3 Ia"] = 16.4
	massMap["K4 Ia"] = 17.2
	massMap["K5 Ia"] = 18
	massMap["K6 Ia"] = 18.4
	massMap["K7 Ia"] = 18.8
	massMap["K8 Ia"] = 19.2
	massMap["K9 Ia"] = 19.6
	massMap["M0 Ia"] = 20
	massMap["M1 Ia"] = 21
	massMap["M2 Ia"] = 22
	massMap["M3 Ia"] = 23
	massMap["M4 Ia"] = 24
	massMap["M5 Ia"] = 25
	massMap["M6 Ia"] = 26.25
	massMap["M7 Ia"] = 27.5
	massMap["M8 Ia"] = 28.75
	massMap["M9 Ia"] = 30
	////////////////
	massMap["B0 Ib"] = 50 //check
	massMap["B1 Ib"] = 45
	massMap["B2 Ib"] = 40
	massMap["B3 Ib"] = 35
	massMap["B4 Ib"] = 30
	massMap["B5 Ib"] = 25 //check
	massMap["B6 Ib"] = 23.2
	massMap["B7 Ib"] = 21.4
	massMap["B8 Ib"] = 19.6
	massMap["B9 Ib"] = 17.8
	massMap["A0 Ib"] = 16 //check
	massMap["A1 Ib"] = 15.4
	massMap["A2 Ib"] = 14.8
	massMap["A3 Ib"] = 14.2
	massMap["A4 Ib"] = 13.6
	massMap["A5 Ib"] = 13 //check
	massMap["A6 Ib"] = 12.8
	massMap["A7 Ib"] = 12.6
	massMap["A8 Ib"] = 12.4
	massMap["A9 Ib"] = 12.2
	massMap["F0 Ib"] = 12 //check
	massMap["F1 Ib"] = 11.6
	massMap["F2 Ib"] = 11.2
	massMap["F3 Ib"] = 10.8
	massMap["F4 Ib"] = 10.4
	massMap["F5 Ib"] = 10 //check
	massMap["F6 Ib"] = 10
	massMap["F7 Ib"] = 10
	massMap["F8 Ib"] = 10
	massMap["F9 Ib"] = 10
	massMap["G0 Ib"] = 10 //check
	massMap["G1 Ib"] = 10.4
	massMap["G2 Ib"] = 10.8
	massMap["G3 Ib"] = 11.2
	massMap["G4 Ib"] = 11.6
	massMap["G5 Ib"] = 12 //check
	massMap["G6 Ib"] = 12.2
	massMap["G7 Ib"] = 12.4
	massMap["G8 Ib"] = 12.6
	massMap["G9 Ib"] = 12.8
	massMap["K0 Ib"] = 13 //check
	massMap["K1 Ib"] = 13.6
	massMap["K2 Ib"] = 14.2
	massMap["K3 Ib"] = 14.8
	massMap["K4 Ib"] = 15.4
	massMap["K5 Ib"] = 16 //check
	massMap["K6 Ib"] = 16
	massMap["K7 Ib"] = 16
	massMap["K8 Ib"] = 16
	massMap["K9 Ib"] = 16
	massMap["M0 Ib"] = 16 //check
	massMap["M1 Ib"] = 16.8
	massMap["M2 Ib"] = 17.6
	massMap["M3 Ib"] = 18.4
	massMap["M4 Ib"] = 19.2
	massMap["M5 Ib"] = 20 //check
	massMap["M6 Ib"] = 21.25
	massMap["M7 Ib"] = 22.5
	massMap["M8 Ib"] = 23.75
	massMap["M9 Ib"] = 25
	////////////////
	massMap["B0 II"] = 30 //check 30
	massMap["B1 II"] = 28
	massMap["B2 II"] = 26
	massMap["B3 II"] = 24
	massMap["B4 II"] = 22
	massMap["B5 II"] = 20 //check 20
	massMap["B6 II"] = 18.8
	massMap["B7 II"] = 17.6
	massMap["B8 II"] = 16.4
	massMap["B9 II"] = 15.2
	massMap["A0 II"] = 14 //check 14
	massMap["A1 II"] = 13.4
	massMap["A2 II"] = 12.8
	massMap["A3 II"] = 12.2
	massMap["A4 II"] = 11.6
	massMap["A5 II"] = 11 //check 11
	massMap["A6 II"] = 10.8
	massMap["A7 II"] = 10.6
	massMap["A8 II"] = 10.4
	massMap["A9 II"] = 10.2
	massMap["F0 II"] = 10 //check 10
	massMap["F1 II"] = 9.62
	massMap["F2 II"] = 9.24
	massMap["F3 II"] = 8.86
	massMap["F4 II"] = 8.48
	massMap["F5 II"] = 8.1 //check 8.1
	massMap["F6 II"] = 8.1
	massMap["F7 II"] = 8.1
	massMap["F8 II"] = 8.1
	massMap["F9 II"] = 8.1
	massMap["G0 II"] = 8.1 //check 8.1
	massMap["G1 II"] = 8.48
	massMap["G2 II"] = 8.86
	massMap["G3 II"] = 9.24
	massMap["G4 II"] = 9.62
	massMap["G5 II"] = 10 //check 10
	massMap["G6 II"] = 10.2
	massMap["G7 II"] = 10.4
	massMap["G8 II"] = 10.6
	massMap["G9 II"] = 10.8
	massMap["K0 II"] = 11 //check 11
	massMap["K1 II"] = 11.6
	massMap["K2 II"] = 12.2
	massMap["K3 II"] = 12.8
	massMap["K4 II"] = 13.4
	massMap["K5 II"] = 14 //check 14
	massMap["K6 II"] = 14
	massMap["K7 II"] = 14
	massMap["K8 II"] = 14
	massMap["K9 II"] = 14
	massMap["M0 II"] = 14 //check 14
	massMap["M1 II"] = 14.4
	massMap["M2 II"] = 14.8
	massMap["M3 II"] = 15.2
	massMap["M4 II"] = 15.6
	massMap["M5 II"] = 16 //check 16
	massMap["M6 II"] = 16.5
	massMap["M7 II"] = 17
	massMap["M8 II"] = 17.5
	massMap["M9 II"] = 18
	////////////////
	massMap["B0 III"] = 25 //check 25
	massMap["B1 III"] = 23
	massMap["B2 III"] = 21
	massMap["B3 III"] = 19
	massMap["B4 III"] = 17
	massMap["B5 III"] = 15 //check 15
	massMap["B6 III"] = 14.4
	massMap["B7 III"] = 13.8
	massMap["B8 III"] = 13.2
	massMap["B9 III"] = 12.6
	massMap["A0 III"] = 12 //check 12
	massMap["A1 III"] = 11.4
	massMap["A2 III"] = 10.8
	massMap["A3 III"] = 10.2
	massMap["A4 III"] = 9.6
	massMap["A5 III"] = 9 //check 9
	massMap["A6 III"] = 8.8
	massMap["A7 III"] = 8.6
	massMap["A8 III"] = 8.4
	massMap["A9 III"] = 8.2
	massMap["F0 III"] = 8 //check 8
	massMap["F1 III"] = 7.4
	massMap["F2 III"] = 6.8
	massMap["F3 III"] = 6.2
	massMap["F4 III"] = 5.6
	massMap["F5 III"] = 5 //check 5
	massMap["F6 III"] = 4.5
	massMap["F7 III"] = 4
	massMap["F8 III"] = 3.5
	massMap["F9 III"] = 3
	massMap["G0 III"] = 2.5 //check 2.5
	massMap["G1 III"] = 2.64
	massMap["G2 III"] = 2.78
	massMap["G3 III"] = 2.92
	massMap["G4 III"] = 3.06
	massMap["G5 III"] = 3.2 //check 3.2
	massMap["G6 III"] = 3.36
	massMap["G7 III"] = 3.52
	massMap["G8 III"] = 3.68
	massMap["G9 III"] = 3.84
	massMap["K0 III"] = 4 //check 4
	massMap["K1 III"] = 4.2
	massMap["K2 III"] = 4.4
	massMap["K3 III"] = 4.6
	massMap["K4 III"] = 4.8
	massMap["K5 III"] = 5 //check 5
	massMap["K6 III"] = 5.26
	massMap["K7 III"] = 5.52
	massMap["K8 III"] = 5.78
	massMap["K9 III"] = 6.04
	massMap["M0 III"] = 6.3 //check 6.3
	massMap["M1 III"] = 6.52
	massMap["M2 III"] = 6.74
	massMap["M3 III"] = 6.96
	massMap["M4 III"] = 7.18
	massMap["M5 III"] = 7.4 //check 7.4
	massMap["M6 III"] = 7.85
	massMap["M7 III"] = 8.3
	massMap["M8 III"] = 8.75
	massMap["M9 III"] = 9.2
	////////////////
	massMap["B0 IV"] = 20 //check 20
	massMap["B1 IV"] = 18
	massMap["B2 IV"] = 16
	massMap["B3 IV"] = 14
	massMap["B4 IV"] = 12
	massMap["B5 IV"] = 10 //check 10
	massMap["B6 IV"] = 9.2
	massMap["B7 IV"] = 8.4
	massMap["B8 IV"] = 7.6
	massMap["B9 IV"] = 6.8
	massMap["A0 IV"] = 6 //check 6
	massMap["A1 IV"] = 5.6
	massMap["A2 IV"] = 5.2
	massMap["A3 IV"] = 4.8
	massMap["A4 IV"] = 4.4
	massMap["A5 IV"] = 4 //check 4
	massMap["A6 IV"] = 3.7
	massMap["A7 IV"] = 3.4
	massMap["A8 IV"] = 3.1
	massMap["A9 IV"] = 2.8
	massMap["F0 IV"] = 2.5 //check 2.5
	massMap["F1 IV"] = 2.4
	massMap["F2 IV"] = 2.3
	massMap["F3 IV"] = 2.2
	massMap["F4 IV"] = 2.1
	massMap["F5 IV"] = 2 //check 2
	massMap["F6 IV"] = 1.95
	massMap["F7 IV"] = 1.9
	massMap["F8 IV"] = 1.85
	massMap["F9 IV"] = 1.8
	massMap["G0 IV"] = 1.75 //check 1.75
	massMap["G1 IV"] = 1.8
	massMap["G2 IV"] = 1.85
	massMap["G3 IV"] = 1.9
	massMap["G4 IV"] = 1.95
	massMap["G5 IV"] = 2 //check 2
	massMap["G6 IV"] = 2.06
	massMap["G7 IV"] = 2.12
	massMap["G8 IV"] = 2.18
	massMap["G9 IV"] = 2.24
	massMap["K0 IV"] = 2.3 //check 2.3
	massMap["K1 IV"] = 2.35
	massMap["K2 IV"] = 2.39
	massMap["K3 IV"] = 2.42
	massMap["K4 IV"] = 2.44
	////////////////
	massMap["B0 V"] = 18 //check 18
	massMap["B1 V"] = 15.7
	massMap["B2 V"] = 13.4
	massMap["B3 V"] = 11.1
	massMap["B4 V"] = 8.8
	massMap["B5 V"] = 6.5 //check 6.5
	massMap["B6 V"] = 5.84
	massMap["B7 V"] = 5.18
	massMap["B8 V"] = 4.52
	massMap["B9 V"] = 3.86
	massMap["A0 V"] = 3.2 //check 3.2
	massMap["A1 V"] = 2.98
	massMap["A2 V"] = 2.76
	massMap["A3 V"] = 2.54
	massMap["A4 V"] = 2.32
	massMap["A5 V"] = 2.1 //check 2.1
	massMap["A6 V"] = 2.02
	massMap["A7 V"] = 1.94
	massMap["A8 V"] = 1.86
	massMap["A9 V"] = 1.78
	massMap["F0 V"] = 1.7 //check 1.7
	massMap["F1 V"] = 1.62
	massMap["F2 V"] = 1.54
	massMap["F3 V"] = 1.46
	massMap["F4 V"] = 1.38
	massMap["F5 V"] = 1.3 //check 1.3
	massMap["F6 V"] = 1.248
	massMap["F7 V"] = 1.196
	massMap["F8 V"] = 1.144
	massMap["F9 V"] = 1.092
	massMap["G0 V"] = 1.04 //check 1.04
	massMap["G1 V"] = 1.02
	massMap["G2 V"] = 1
	massMap["G3 V"] = 0.98
	massMap["G4 V"] = 0.96
	massMap["G5 V"] = 0.94 //check 0.94
	massMap["G6 V"] = 0.917
	massMap["G7 V"] = 0.894
	massMap["G8 V"] = 0.871
	massMap["G9 V"] = 0.848
	massMap["K0 V"] = 0.825 //check 0.825
	massMap["K1 V"] = 0.774
	massMap["K2 V"] = 0.723
	massMap["K3 V"] = 0.672
	massMap["K4 V"] = 0.621
	massMap["K5 V"] = 0.57 //check 0.57
	massMap["K6 V"] = 0.5538
	massMap["K7 V"] = 0.5376
	massMap["K8 V"] = 0.5214
	massMap["K9 V"] = 0.5052
	massMap["M0 V"] = 0.489 //check 0.489
	massMap["M1 V"] = 0.4574
	massMap["M2 V"] = 0.4258
	massMap["M3 V"] = 0.3942
	massMap["M4 V"] = 0.3626
	massMap["M5 V"] = 0.331 //check 0.331
	massMap["M6 V"] = 0.302
	massMap["M7 V"] = 0.273
	massMap["M8 V"] = 0.244
	massMap["M9 V"] = 0.215
	////////////////
	massMap["F5 VI"] = 0.8 //check 0.8
	massMap["F6 VI"] = 0.76
	massMap["F7 VI"] = 0.72
	massMap["F8 VI"] = 0.68
	massMap["F9 VI"] = 0.64
	massMap["G0 VI"] = 0.6 //check 0.6
	massMap["G1 VI"] = 0.5856
	massMap["G2 VI"] = 0.5712
	massMap["G3 VI"] = 0.5568
	massMap["G4 VI"] = 0.5424
	massMap["G5 VI"] = 0.528 //check 0.528
	massMap["G6 VI"] = 0.5084
	massMap["G7 VI"] = 0.4888
	massMap["G8 VI"] = 0.4692
	massMap["G9 VI"] = 0.4496
	massMap["K0 VI"] = 0.43 //check 0.43
	massMap["K1 VI"] = 0.41
	massMap["K2 VI"] = 0.39
	massMap["K3 VI"] = 0.37
	massMap["K4 VI"] = 0.35
	massMap["K5 VI"] = 0.33 //check 0.33
	massMap["K6 VI"] = 0.2948
	massMap["K7 VI"] = 0.2596
	massMap["K8 VI"] = 0.2244
	massMap["K9 VI"] = 0.1892
	massMap["M0 VI"] = 0.154 //check 0.154
	massMap["M1 VI"] = 0.144
	massMap["M2 VI"] = 0.134
	massMap["M3 VI"] = 0.124
	massMap["M4 VI"] = 0.114
	massMap["M5 VI"] = 0.104 //check 0.104
	massMap["M6 VI"] = 0.0925
	massMap["M7 VI"] = 0.081
	massMap["M8 VI"] = 0.0695
	massMap["M9 VI"] = 0.058
	////////////////
	massMap["DO"] = 0.21
	massMap["DB"] = 0.26
	massMap["DA"] = 0.36
	massMap["DF"] = 0.42
	massMap["DG"] = 0.63
	massMap["DK"] = 0.83
	massMap["DM"] = 1.11
	return massMap[class]
}

func baseStellarLuminocity(class string) float64 {
	lumaMap := make(map[string]float64)
	lumaMap["B0 Ia"] = 27.36 //check 27.36
	lumaMap["B1 Ia"] = 26.14
	lumaMap["B2 Ia"] = 24.92
	lumaMap["B3 Ia"] = 23.69
	lumaMap["B4 Ia"] = 22.47
	lumaMap["B5 Ia"] = 21.25 //check 21.25
	lumaMap["B6 Ia"] = 20.62
	lumaMap["B7 Ia"] = 19.99
	lumaMap["B8 Ia"] = 19.35
	lumaMap["B9 Ia"] = 18.72
	lumaMap["A0 Ia"] = 18.09 //check 18.09
	lumaMap["A1 Ia"] = 17.85
	lumaMap["A2 Ia"] = 17.60
	lumaMap["A3 Ia"] = 17.36
	lumaMap["A4 Ia"] = 17.11
	lumaMap["A5 Ia"] = 16.87 //check 16.87
	lumaMap["A6 Ia"] = 16.64
	lumaMap["A7 Ia"] = 16.41
	lumaMap["A8 Ia"] = 16.18
	lumaMap["A9 Ia"] = 15.95
	lumaMap["F0 Ia"] = 15.72 //check 15.72
	lumaMap["F1 Ia"] = 15.58
	lumaMap["F2 Ia"] = 15.44
	lumaMap["F3 Ia"] = 15.31
	lumaMap["F4 Ia"] = 15.17
	lumaMap["F5 Ia"] = 15.03 //check 15.03
	lumaMap["F6 Ia"] = 15.24
	lumaMap["F7 Ia"] = 15.45
	lumaMap["F8 Ia"] = 15.67
	lumaMap["F9 Ia"] = 15.88
	lumaMap["G0 Ia"] = 16.09 //check 16.09
	lumaMap["G1 Ia"] = 16.33
	lumaMap["G2 Ia"] = 16.56
	lumaMap["G3 Ia"] = 16.80
	lumaMap["G4 Ia"] = 17.03
	lumaMap["G5 Ia"] = 17.27 //check 17.27
	lumaMap["G6 Ia"] = 17.35
	lumaMap["G7 Ia"] = 17.42
	lumaMap["G8 Ia"] = 17.50
	lumaMap["G9 Ia"] = 17.57
	lumaMap["K0 Ia"] = 17.65 //check 17.65
	lumaMap["K1 Ia"] = 17.74
	lumaMap["K2 Ia"] = 17.83
	lumaMap["K3 Ia"] = 17.91
	lumaMap["K4 Ia"] = 18.00
	lumaMap["K5 Ia"] = 18.09 //check 18.09
	lumaMap["K6 Ia"] = 18.17
	lumaMap["K7 Ia"] = 18.25
	lumaMap["K8 Ia"] = 18.33
	lumaMap["K9 Ia"] = 18.41
	lumaMap["M0 Ia"] = 18.49 //check 18.49
	lumaMap["M1 Ia"] = 18.58
	lumaMap["M2 Ia"] = 18.67
	lumaMap["M3 Ia"] = 18.77
	lumaMap["M4 Ia"] = 18.86
	lumaMap["M5 Ia"] = 18.95 //check 18.95
	lumaMap["M6 Ia"] = 19.06
	lumaMap["M7 Ia"] = 19.17
	lumaMap["M8 Ia"] = 19.27
	lumaMap["M9 Ia"] = 19.38
	////////////////
	lumaMap["B0 Ib"] = 22.8 //check 22.8
	lumaMap["B1 Ib"] = 21.18
	lumaMap["B2 Ib"] = 19.56
	lumaMap["B3 Ib"] = 17.94
	lumaMap["B4 Ib"] = 16.32
	lumaMap["B5 Ib"] = 14.7 //check 14.70
	lumaMap["B6 Ib"] = 13.97
	lumaMap["B7 Ib"] = 13.25
	lumaMap["B8 Ib"] = 12.52
	lumaMap["B9 Ib"] = 11.8
	lumaMap["A0 Ib"] = 11.07 //check 11.07
	lumaMap["A1 Ib"] = 10.94
	lumaMap["A2 Ib"] = 10.8
	lumaMap["A3 Ib"] = 10.67
	lumaMap["A4 Ib"] = 10.53
	lumaMap["A5 Ib"] = 10.4 //check 10.40
	lumaMap["A6 Ib"] = 10.17
	lumaMap["A7 Ib"] = 9.95
	lumaMap["A8 Ib"] = 9.72
	lumaMap["A9 Ib"] = 9.5
	lumaMap["F0 Ib"] = 9.27 //check 9.27
	lumaMap["F1 Ib"] = 9.11
	lumaMap["F2 Ib"] = 8.94
	lumaMap["F3 Ib"] = 8.78
	lumaMap["F4 Ib"] = 8.61
	lumaMap["F5 Ib"] = 8.45 //check 8.45
	lumaMap["F6 Ib"] = 8.53
	lumaMap["F7 Ib"] = 8.61
	lumaMap["F8 Ib"] = 8.68
	lumaMap["F9 Ib"] = 8.76
	lumaMap["G0 Ib"] = 8.84 //check 8.84
	lumaMap["G1 Ib"] = 8.97
	lumaMap["G2 Ib"] = 9.1
	lumaMap["G3 Ib"] = 9.23
	lumaMap["G4 Ib"] = 9.36
	lumaMap["G5 Ib"] = 9.49 //check 9.49
	lumaMap["G6 Ib"] = 9.67
	lumaMap["G7 Ib"] = 9.85
	lumaMap["G8 Ib"] = 10.04
	lumaMap["G9 Ib"] = 10.22
	lumaMap["K0 Ib"] = 10.4 //check 10.40
	lumaMap["K1 Ib"] = 10.71
	lumaMap["K2 Ib"] = 11.02
	lumaMap["K3 Ib"] = 11.33
	lumaMap["K4 Ib"] = 11.64
	lumaMap["K5 Ib"] = 11.95 //check 11.95
	lumaMap["K6 Ib"] = 12.49
	lumaMap["K7 Ib"] = 13.03
	lumaMap["K8 Ib"] = 13.57
	lumaMap["K9 Ib"] = 14.11
	lumaMap["M0 Ib"] = 14.65 //check 14.65
	lumaMap["M1 Ib"] = 15.17
	lumaMap["M2 Ib"] = 15.7
	lumaMap["M3 Ib"] = 16.22
	lumaMap["M4 Ib"] = 16.75
	lumaMap["M5 Ib"] = 17.27 //check 17.27
	lumaMap["M6 Ib"] = 17.58
	lumaMap["M7 Ib"] = 17.88
	lumaMap["M8 Ib"] = 18.18
	lumaMap["M9 Ib"] = 18.49
	////////////////
	lumaMap["B0 II"] = 20.31 //check 20.31
	lumaMap["B1 II"] = 18.58
	lumaMap["B2 II"] = 16.86
	lumaMap["B3 II"] = 15.13
	lumaMap["B4 II"] = 13.41
	lumaMap["B5 II"] = 11.68 //check 11.68
	lumaMap["B6 II"] = 10.71
	lumaMap["B7 II"] = 9.75
	lumaMap["B8 II"] = 8.78
	lumaMap["B9 II"] = 7.82
	lumaMap["A0 II"] = 6.85 //check 6.85
	lumaMap["A1 II"] = 6.56
	lumaMap["A2 II"] = 6.27
	lumaMap["A3 II"] = 5.98
	lumaMap["A4 II"] = 5.69
	lumaMap["A5 II"] = 5.4 //check 5.4
	lumaMap["A6 II"] = 5.31
	lumaMap["A7 II"] = 5.22
	lumaMap["A8 II"] = 5.13
	lumaMap["A9 II"] = 5.04
	lumaMap["F0 II"] = 4.95 //check 4.95
	lumaMap["F1 II"] = 4.91
	lumaMap["F2 II"] = 4.87
	lumaMap["F3 II"] = 4.83
	lumaMap["F4 II"] = 4.79
	lumaMap["F5 II"] = 4.75 //check 4.75
	lumaMap["F6 II"] = 4.77
	lumaMap["F7 II"] = 4.79
	lumaMap["F8 II"] = 4.82
	lumaMap["F9 II"] = 4.84
	lumaMap["G0 II"] = 4.86 //check 4.86
	lumaMap["G1 II"] = 4.93
	lumaMap["G2 II"] = 5
	lumaMap["G3 II"] = 5.08
	lumaMap["G4 II"] = 5.15
	lumaMap["G5 II"] = 5.22 //check 5.22
	lumaMap["G6 II"] = 5.27
	lumaMap["G7 II"] = 5.32
	lumaMap["G8 II"] = 5.36
	lumaMap["G9 II"] = 5.41
	lumaMap["K0 II"] = 5.46 //check 5.46
	lumaMap["K1 II"] = 5.78
	lumaMap["K2 II"] = 6.09
	lumaMap["K3 II"] = 6.41
	lumaMap["K4 II"] = 6.72
	lumaMap["K5 II"] = 7.04 //check 7.04
	lumaMap["K6 II"] = 7.28
	lumaMap["K7 II"] = 7.52
	lumaMap["K8 II"] = 7.76
	lumaMap["K9 II"] = 8
	lumaMap["M0 II"] = 8.24 //check 8.24
	lumaMap["M1 II"] = 8.8
	lumaMap["M2 II"] = 9.36
	lumaMap["M3 II"] = 9.93
	lumaMap["M4 II"] = 10.49
	lumaMap["M5 II"] = 11.05 //check 11.05
	lumaMap["M6 II"] = 11.11
	lumaMap["M7 II"] = 11.17
	lumaMap["M8 II"] = 11.22
	lumaMap["M9 II"] = 11.28
	////////////////
	lumaMap["B0 III"] = 18.09 //check 18.09
	lumaMap["B1 III"] = 16.28
	lumaMap["B2 III"] = 14.47
	lumaMap["B3 III"] = 12.67
	lumaMap["B4 III"] = 10.86
	lumaMap["B5 III"] = 9.05 //check 9.05
	lumaMap["B6 III"] = 8.06
	lumaMap["B7 III"] = 7.07
	lumaMap["B8 III"] = 6.07
	lumaMap["B9 III"] = 5.08
	lumaMap["A0 III"] = 4.09 //check 4.09
	lumaMap["A1 III"] = 3.89
	lumaMap["A2 III"] = 3.69
	lumaMap["A3 III"] = 3.48
	lumaMap["A4 III"] = 3.28
	lumaMap["A5 III"] = 3.08 //check 3.08
	lumaMap["A6 III"] = 3
	lumaMap["A7 III"] = 2.93
	lumaMap["A8 III"] = 2.85
	lumaMap["A9 III"] = 2.78
	lumaMap["F0 III"] = 2.7 //check 2.7
	lumaMap["F1 III"] = 2.67
	lumaMap["F2 III"] = 2.64
	lumaMap["F3 III"] = 2.62
	lumaMap["F4 III"] = 2.59
	lumaMap["F5 III"] = 2.56 //check 2.56
	lumaMap["F6 III"] = 2.58
	lumaMap["F7 III"] = 2.6
	lumaMap["F8 III"] = 2.62
	lumaMap["F9 III"] = 2.64
	lumaMap["G0 III"] = 2.66 //check 2.66
	lumaMap["G1 III"] = 2.72
	lumaMap["G2 III"] = 2.77
	lumaMap["G3 III"] = 2.83
	lumaMap["G4 III"] = 2.88
	lumaMap["G5 III"] = 2.94 //check 2.94
	lumaMap["G6 III"] = 2.98
	lumaMap["G7 III"] = 3.01
	lumaMap["G8 III"] = 3.05
	lumaMap["G9 III"] = 3.08
	lumaMap["K0 III"] = 3.12 //check 3.12
	lumaMap["K1 III"] = 3.34
	lumaMap["K2 III"] = 3.56
	lumaMap["K3 III"] = 3.79
	lumaMap["K4 III"] = 4.01
	lumaMap["K5 III"] = 4.23 //check 4.23
	lumaMap["K6 III"] = 4.32
	lumaMap["K7 III"] = 4.4
	lumaMap["K8 III"] = 4.49
	lumaMap["K9 III"] = 4.57
	lumaMap["M0 III"] = 4.66 //check 4.66
	lumaMap["M1 III"] = 5.11
	lumaMap["M2 III"] = 5.56
	lumaMap["M3 III"] = 6.01
	lumaMap["M4 III"] = 6.46
	lumaMap["M5 III"] = 6.91 //check 6.91
	lumaMap["M6 III"] = 6.98
	lumaMap["M7 III"] = 7.06
	lumaMap["M8 III"] = 7.13
	lumaMap["M9 III"] = 7.2
	////////////////
	lumaMap["B0 IV"] = 16.87 //check 16.87
	lumaMap["B1 IV"] = 14.83
	lumaMap["B2 IV"] = 12.8
	lumaMap["B3 IV"] = 10.76
	lumaMap["B4 IV"] = 8.73
	lumaMap["B5 IV"] = 6.69 //check 6.69
	lumaMap["B6 IV"] = 6.06
	lumaMap["B7 IV"] = 5.43
	lumaMap["B8 IV"] = 4.79
	lumaMap["B9 IV"] = 4.16
	lumaMap["A0 IV"] = 3.53 //check 3.53
	lumaMap["A1 IV"] = 3.32
	lumaMap["A2 IV"] = 3.11
	lumaMap["A3 IV"] = 2.89
	lumaMap["A4 IV"] = 2.68
	lumaMap["A5 IV"] = 2.47 //check 2.47
	lumaMap["A6 IV"] = 2.39
	lumaMap["A7 IV"] = 2.32
	lumaMap["A8 IV"] = 2.24
	lumaMap["A9 IV"] = 2.17
	lumaMap["F0 IV"] = 2.09 //check 2.09
	lumaMap["F1 IV"] = 2.04
	lumaMap["F2 IV"] = 2
	lumaMap["F3 IV"] = 1.95
	lumaMap["F4 IV"] = 1.91
	lumaMap["F5 IV"] = 1.86 //check 1.86
	lumaMap["F6 IV"] = 1.81
	lumaMap["F7 IV"] = 1.76
	lumaMap["F8 IV"] = 1.7
	lumaMap["F9 IV"] = 1.65
	lumaMap["G0 IV"] = 1.6 //check 1.6
	lumaMap["G1 IV"] = 1.58
	lumaMap["G2 IV"] = 1.56
	lumaMap["G3 IV"] = 1.53
	lumaMap["G4 IV"] = 1.51
	lumaMap["G5 IV"] = 1.49 //check 1.49
	lumaMap["G6 IV"] = 1.49
	lumaMap["G7 IV"] = 1.48
	lumaMap["G8 IV"] = 1.48
	lumaMap["G9 IV"] = 1.47
	lumaMap["K0 IV"] = 1.47
	lumaMap["K1 IV"] = 1.47
	lumaMap["K2 IV"] = 1.46
	lumaMap["K3 IV"] = 1.46
	lumaMap["K4 IV"] = 1.46
	///////////////
	lumaMap["B0 V"] = 15.38 //check 15.38
	lumaMap["B1 V"] = 13.53
	lumaMap["B2 V"] = 11.68
	lumaMap["B3 V"] = 9.82
	lumaMap["B4 V"] = 7.97
	lumaMap["B5 V"] = 6.12 //check 6.12
	lumaMap["B6 V"] = 5.51
	lumaMap["B7 V"] = 4.9
	lumaMap["B8 V"] = 4.3
	lumaMap["B9 V"] = 3.69
	lumaMap["A0 V"] = 3.08 //check 3.08
	lumaMap["A1 V"] = 2.86
	lumaMap["A2 V"] = 2.65
	lumaMap["A3 V"] = 2.43
	lumaMap["A4 V"] = 2.22
	lumaMap["A5 V"] = 2 //check 2
	lumaMap["A6 V"] = 1.94
	lumaMap["A7 V"] = 1.88
	lumaMap["A8 V"] = 1.81
	lumaMap["A9 V"] = 1.75
	lumaMap["F0 V"] = 1.69 //check 1.69
	lumaMap["F1 V"] = 1.63
	lumaMap["F2 V"] = 1.56
	lumaMap["F3 V"] = 1.5
	lumaMap["F4 V"] = 1.43
	lumaMap["F5 V"] = 1.37 //check 1.37
	lumaMap["F6 V"] = 1.31
	lumaMap["F7 V"] = 1.24
	lumaMap["F8 V"] = 1.18
	lumaMap["F9 V"] = 1.11
	lumaMap["G0 V"] = 1.05 //check 1.05
	lumaMap["G1 V"] = 1.02
	lumaMap["G2 V"] = 0.99
	lumaMap["G3 V"] = 0.96
	lumaMap["G4 V"] = 0.93
	lumaMap["G5 V"] = 0.9 //check 0.9
	lumaMap["G6 V"] = 0.88
	lumaMap["G7 V"] = 0.86
	lumaMap["G8 V"] = 0.85
	lumaMap["G9 V"] = 0.83
	lumaMap["K0 V"] = 0.81 //check 0.81
	lumaMap["K1 V"] = 0.75
	lumaMap["K2 V"] = 0.7
	lumaMap["K3 V"] = 0.64
	lumaMap["K4 V"] = 0.59
	lumaMap["K5 V"] = 0.53 //check 0.53
	lumaMap["K6 V"] = 0.51
	lumaMap["K7 V"] = 0.5
	lumaMap["K8 V"] = 0.48
	lumaMap["K9 V"] = 0.47
	lumaMap["M0 V"] = 0.45 //check 0.45
	lumaMap["M1 V"] = 0.42
	lumaMap["M2 V"] = 0.39
	lumaMap["M3 V"] = 0.35
	lumaMap["M4 V"] = 0.32
	lumaMap["M5 V"] = 0.29 //check 0.29
	lumaMap["M6 V"] = 0.26
	lumaMap["M7 V"] = 0.24
	lumaMap["M8 V"] = 0.21
	lumaMap["M9 V"] = 0.18
	////////////////
	lumaMap["F5 VI"] = 0.99 //check 0.99
	lumaMap["F6 VI"] = 0.94
	lumaMap["F7 VI"] = 0.89
	lumaMap["F8 VI"] = 0.85
	lumaMap["F9 VI"] = 0.8
	lumaMap["G0 VI"] = 0.75 //check 0.75
	lumaMap["G1 VI"] = 0.73
	lumaMap["G2 VI"] = 0.71
	lumaMap["G3 VI"] = 0.7
	lumaMap["G4 VI"] = 0.68
	lumaMap["G5 VI"] = 0.66 //check 0.66
	lumaMap["G6 VI"] = 0.64
	lumaMap["G7 VI"] = 0.63
	lumaMap["G8 VI"] = 0.61
	lumaMap["G9 VI"] = 0.6
	lumaMap["K0 VI"] = 0.58 //check 0.58
	lumaMap["K1 VI"] = 0.54
	lumaMap["K2 VI"] = 0.51
	lumaMap["K3 VI"] = 0.47
	lumaMap["K4 VI"] = 0.44
	lumaMap["K5 VI"] = 0.4 //check 0.4
	lumaMap["K6 VI"] = 0.38
	lumaMap["K7 VI"] = 0.37
	lumaMap["K8 VI"] = 0.35
	lumaMap["K9 VI"] = 0.34
	lumaMap["M0 VI"] = 0.32 //check 0.32
	lumaMap["M1 VI"] = 0.3
	lumaMap["M2 VI"] = 0.28
	lumaMap["M3 VI"] = 0.25
	lumaMap["M4 VI"] = 0.23
	lumaMap["M5 VI"] = 0.21 //check 0.21
	lumaMap["M6 VI"] = 0.18
	lumaMap["M7 VI"] = 0.15
	lumaMap["M8 VI"] = 0.12
	lumaMap["M9 VI"] = 0.09
	////////////////
	lumaMap["DB"] = 0.46
	lumaMap["DA"] = 0.27
	lumaMap["DF"] = 0.13
	lumaMap["DG"] = 0.09
	lumaMap["DK"] = 0.08
	lumaMap["DM"] = 0.07
	return lumaMap[class]
}
