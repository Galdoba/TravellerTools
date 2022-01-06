package star

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
	decimal    int
	spectral   string
	mass       float64
	luminocity float64
	orbit      int
	category   int
}

func New(name string) Star {
	s := Star{}
	s.name = name
	return s
}

func interpalation(a, b float64, dif, div float64) float64 {
	res := a - ((a - b) * dif / div)

	return res
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
	massMap["BD"] = 0.26
	massMap["AD"] = 0.36
	massMap["FD"] = 0.42
	massMap["GD"] = 0.63
	massMap["KD"] = 0.83
	massMap["MD"] = 1.11
	return massMap[class]
}
