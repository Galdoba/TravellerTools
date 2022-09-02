package systemgeneration

/*

O0 V 50,000 100  1,240,000 20    1057.88 – 1447.62 N/A     4000
O1 V 47,800 97.5 994,000   19.5  947.15  – 1296.09 N/A     3900
O2 V 45,600 95   795,000   19    847.05  – 1159.12 N/A     3800
O3 V 43,400 92.5 634,000   18.5  756.43  – 1035.11 N/A     3700
O4 V 41,200 90   504,000   18    674.43  – 922.91  3549.65 3600
O5 V 39,000 60   398,000   12    599.33  – 820.13  N/A     2400
O6 V 36,800 37   260,000   7.4   153.19  – 662.87  N/A     1480
O7 V 34,600 30   154,000   6     372.81  – 510.16  N/A     1200
O8 V 32,400 23   99,100    4.6   299.06  - 409.24  N/A     920
O9 V 30,200 20   57,600    4     228     – 312     N/A     800
B0 V 28,000 17.5 36,200    3.5   180.75  – 247.34  N/A     700
B1 V 26,190 14.2 19,400    2.84  132.32  – 181.07  N/A     568
B2 V 24,380 10.9 9,360     2.18  89.97   – 125.77  N/A     436
B3 V 22,570 7.6  4,890     1.52  66.43   – 90.91   N/A     304
B4 V 20,760 6.7  2,290     1.34  45.46   – 62.21   239.27  268
B5 V 18,950 5.9  1,160     1.18  32.36   – 44.28   170.29  236
B6 V 17,140 5.2  692       1.04  24.99   – 34.20   131.53  208
B7 V 15,330 4.5  404       0.90  19.09   – 26.13   100.50  180
B8 V 13,520 3.8  211       0.76  13.80   – 18.88   76.63   152
B9 V 11,710 3.4  119       0.68  10.36   – 14.18   54.54   136
A0 V 9900   2.9  67.4      0.58  7.80    – 10.67   41.05   116
A1 V 9650   2.7  49.2      0.54  6.66    – 9.12    35.07   108
A2 V 9400   2.5  39.4      0.50  5.96    – 8.16    31.38   100
A3 V 9150   2.4  28.9      0.48  5.11    – 6.99    26.88   96
A4 V 8900   2.1  23.2      0.42  4.58    – 6.26    24.08   84
A5 V 8650   1.9  17.0      0.38  3.92    – 5.36    20.62   76
A6 V 8400   1.8  15.1      0.36  3.69    – 5.05    19.43   72
A7 V 8150   1.8  12.2      0.36  3.32    – 4.54    17.46   72
A8 V 7900   1.8  10.9      0.36  3.14    – 4.30    16.50   72
A9 V 7650   1.7  8.85      0.34  2.83    – 3.87    14.87   68
F0 V 7400   1.6  7.94      0.32  2.68    – 3.66    14.09   64
F1 V 7260   1.6  6.56      0.32  2.43    – 1.64    12.81   64
F2 V 7120   1.5  5.95      0.30  2.32    – 3.17    12.20   60
F3 V 6980   1.5  4.94      0.30  2.11    – 2.89    11.11   60
F4 V 6840   1.4  4.50      0.28  2.02    – 2.76    10.61   56
F5 V 6700   1.4  3.75      0.28  1.84    – 2.52    9.68    56
F6 V 6560   1.3  3.13      0.26  1.68    – 2.30    8.85    52
F7 V 6420   1.3  2.62      0.26  1.54    – 2.10    8.09    52
F8 V 6280   1.2  2.41      0.24  1.47    – 2.02    7.76    48
F9 V 6140   1.1  2.03      0.22  1.35    – 1.85    7.12    44
G0 V 6000   1.1  1.72      0.22  1.25    – 1.70    6.56    44
G1 V 5890   1.0  1.46      0.20  1.15    – 1.57    6.04    40
G2 V 5780   1.0  1.00      0.20  0.95    – 1.30    5.00    40
G3 V 5670   1.0  1.00      0.20  0.95    – 1.30    5.00    40
G4 V 5560   0.9  0.98      0.18  0.94    – 1.29    4.95    36
G5 V 5450   0.9  0.84      0.18  0.87    – 1.19    4.58    36
G6 V 5340   0.9  0.79      0.18  0.84    – 1.16    4.44    36
G7 V 5230   0.9  0.68      0.18  0.78    – 1.07    4.12    36
G8 V 5120   0.8  0.65      0.16  0.77    – 1.05    4.03    32
G9 V 5010   0.8  0.57      0.16  0.72    – 0.98    3.77    32
K0 V 4900   0.8  0.54      0.16  0.70    – 0.96    3.67    32
K1 V 4760   0.8  0.44      0.16  0.63    – 0.86    3.32    32
K2 V 4620   0.7  0.40      0.14  0.60    – 0.82    3.16    28
K3 V 4480   0.7  0.34      0.14  0.55    – 0.76    2.92    28
K4 V 4340   0.7  0.31      0.14  0.53    – 0.72    2.78    28
K5 V 4200   0.7  0.27      0.14  0.49    – 0.68    2.60    28
K6 V 4060   0.6  0.21      0.12  0.44    – 0.60    2.29    24
K7 V 3920   0.6  0.19      0.12  0.41    – 0.57    2.18    24
K8 V 3780   0.6  0.16      0.12  0.38    – 0.52    2.00    24
K9 V 3640   0.5  0.14      0.10  0.36    – 0.49    1.87    20
M0 V 3500   0.5  0.125     0.100 0.336   – 0.460   1.768   20
M1 V 3333   0.5  0.0618    0.100 0.236   – 0.323   1.243   20
M2 V 3167   0.4  0.0321    0.080 0.170   – 0.233   0.896   16
M3 V 3000   0.3  0.0178    0.060 0.127   – 0.173   0.667   12
M4 V 2833   0.3  0.0106    0.060 0.098   – 0.134   0.515   12
M5 V 2667   0.2  0.00624   0.040 0.075   – 0.103   0.395   8
M6 V 2500   0.2  0.00450   0.040 0.0637  – 0.0872  0.335   8
M7 V 2333   0.1  0.00369   0.020 0.0577  – 0.0790  0.960   4
M8 V 2167   0.1  0.00353   0.020 0.0564  – 0.0772  0.297   4
M9 V 2000   0.1  0.00315   0.020 0.0533  – 0.0730  0.281   4
*/
type tabledata struct {
	star          string
	temperature   int
	mass          float64
	luminocity    float64
	innerLimit    float64
	habitableLow  float64
	habitableHigh float64
	snowLine      float64
	outerLimit    float64
}

func getTableValues(star string) tabledata {
	starMap := make(map[string]tabledata)
	starMap["O0 V"] = tabledata{"O0 V", 50000, 100, 1240000, 20, 1057.88, 1447.62, -999, 4000}
	starMap["O1 V"] = tabledata{"O1 V", 47800, 97.5, 994000, 19.5, 947.15, 1296.09, -999, 3900}
	starMap["O2 V"] = tabledata{"O2 V", 45600, 95, 795000, 19, 847.05, 1159.12, -999, 3800}
	starMap["O3 V"] = tabledata{"O3 V", 43400, 92.5, 634000, 18.5, 756.43, 1035.11, -999, 3700}
	starMap["O4 V"] = tabledata{"O4 V", 41200, 90, 504000, 18, 674.43, 922.91, 3549.65, 3600}
	starMap["O5 V"] = tabledata{"O5 V", 39000, 60, 398000, 12, 599.33, 820.13, -999, 2400}
	starMap["O6 V"] = tabledata{"O6 V", 36800, 37, 260000, 7.4, 153.19, 662.87, -999, 1480}
	starMap["O7 V"] = tabledata{"O7 V", 34600, 30, 154000, 6, 372.81, 510.16, -999, 1200}
	starMap["O8 V"] = tabledata{"O8 V", 32400, 23, 99100, 4.6, 299.06, 409.24, -999, 920}
	starMap["O9 V"] = tabledata{"O9 V", 30200, 20, 57600, 4, 228, 312, -999, 800}
	starMap["B0 V"] = tabledata{"B0 V", 28000, 17.5, 36200, 3.5, 180.75, 247.34, -999, 700}
	starMap["B1 V"] = tabledata{"B1 V", 26190, 14.2, 19400, 2.84, 132.32, 181.07, -999, 568}
	starMap["B2 V"] = tabledata{"B2 V", 24380, 10.9, 9360, 2.18, 89.97, 125.77, -999, 436}
	starMap["B3 V"] = tabledata{"B3 V", 22570, 7.6, 4890, 1.52, 66.43, 90.91, -999, 304}
	starMap["B4 V"] = tabledata{"B4 V", 20760, 6.7, 2290, 1.34, 45.46, 62.21, 239.27, 268}
	starMap["B5 V"] = tabledata{"B5 V", 18950, 5.9, 1160, 1.18, 32.36, 44.28, 170.29, 236}
	starMap["B6 V"] = tabledata{"B6 V", 17140, 5.2, 692, 1.04, 24.99, 34.20, 131.53, 208}
	starMap["B7 V"] = tabledata{"B7 V", 15330, 4.5, 404, 0.90, 19.09, 26.13, 100.50, 180}
	starMap["B8 V"] = tabledata{"B8 V", 13520, 3.8, 211, 0.76, 13.80, 18.88, 76.63, 152}
	starMap["B9 V"] = tabledata{"B9 V", 11710, 3.4, 119, 0.68, 10.36, 14.18, 54.54, 136}
	starMap["A0 V"] = tabledata{"A0 V", 9900, 2.9, 67.4, 0.58, 7.80, 10.67, 41.05, 116}
	starMap["A1 V"] = tabledata{"A1 V", 9650, 2.7, 49.2, 0.54, 6.66, 9.12, 35.07, 108}
	starMap["A2 V"] = tabledata{"A2 V", 9400, 2.5, 39.4, 0.50, 5.96, 8.16, 31.38, 100}
	starMap["A3 V"] = tabledata{"A3 V", 9150, 2.4, 28.9, 0.48, 5.11, 6.99, 26.88, 96}
	starMap["A4 V"] = tabledata{"A4 V", 8900, 2.1, 23.2, 0.42, 4.58, 6.26, 24.08, 84}
	starMap["A5 V"] = tabledata{"A5 V", 8650, 1.9, 17.0, 0.38, 3.92, 5.36, 20.62, 76}
	starMap["A6 V"] = tabledata{"A6 V", 8400, 1.8, 15.1, 0.36, 3.69, 5.05, 19.43, 72}
	starMap["A7 V"] = tabledata{"A7 V", 8150, 1.8, 12.2, 0.36, 3.32, 4.54, 17.46, 72}
	starMap["A8 V"] = tabledata{"A8 V", 7900, 1.8, 10.9, 0.36, 3.14, 4.30, 16.50, 72}
	starMap["A9 V"] = tabledata{"A9 V", 7650, 1.7, 8.85, 0.34, 2.83, 3.87, 14.87, 68}
	starMap["F0 V"] = tabledata{"F0 V", 7400, 1.6, 7.94, 0.32, 2.68, 3.66, 14.09, 64}
	starMap["F1 V"] = tabledata{"F1 V", 7260, 1.6, 6.56, 0.32, 2.43, 1.64, 12.81, 64}
	starMap["F2 V"] = tabledata{"F2 V", 7120, 1.5, 5.95, 0.30, 2.32, 3.17, 12.20, 60}
	starMap["F3 V"] = tabledata{"F3 V", 6980, 1.5, 4.94, 0.30, 2.11, 2.89, 11.11, 60}
	starMap["F4 V"] = tabledata{"F4 V", 6840, 1.4, 4.50, 0.28, 2.02, 2.76, 10.61, 56}
	starMap["F5 V"] = tabledata{"F5 V", 6700, 1.4, 3.75, 0.28, 1.84, 2.52, 9.68, 56}
	starMap["F6 V"] = tabledata{"F6 V", 6560, 1.3, 3.13, 0.26, 1.68, 2.30, 8.85, 52}
	starMap["F7 V"] = tabledata{"F7 V", 6420, 1.3, 2.62, 0.26, 1.54, 2.10, 8.09, 52}
	starMap["F8 V"] = tabledata{"F8 V", 6280, 1.2, 2.41, 0.24, 1.47, 2.02, 7.76, 48}
	starMap["F9 V"] = tabledata{"F9 V", 6140, 1.1, 2.03, 0.22, 1.35, 1.85, 7.12, 44}
	starMap["G0 V"] = tabledata{"G0 V", 6000, 1.1, 1.72, 0.22, 1.25, 1.70, 6.56, 44}
	starMap["G1 V"] = tabledata{"G1 V", 5890, 1.0, 1.46, 0.20, 1.15, 1.57, 6.04, 40}
	starMap["G2 V"] = tabledata{"G2 V", 5780, 1.0, 1.00, 0.20, 0.95, 1.30, 5.00, 40}
	starMap["G3 V"] = tabledata{"G3 V", 5670, 1.0, 1.00, 0.20, 0.95, 1.30, 5.00, 40}
	starMap["G4 V"] = tabledata{"G4 V", 5560, 0.9, 0.98, 0.18, 0.94, 1.29, 4.95, 36}
	starMap["G5 V"] = tabledata{"G5 V", 5450, 0.9, 0.84, 0.18, 0.87, 1.19, 4.58, 36}
	starMap["G6 V"] = tabledata{"G6 V", 5340, 0.9, 0.79, 0.18, 0.84, 1.16, 4.44, 36}
	starMap["G7 V"] = tabledata{"G7 V", 5230, 0.9, 0.68, 0.18, 0.78, 1.07, 4.12, 36}
	starMap["G8 V"] = tabledata{"G8 V", 5120, 0.8, 0.65, 0.16, 0.77, 1.05, 4.03, 32}
	starMap["G9 V"] = tabledata{"G9 V", 5010, 0.8, 0.57, 0.16, 0.72, 0.98, 3.77, 32}
	starMap["K0 V"] = tabledata{"K0 V", 4900, 0.8, 0.54, 0.16, 0.70, 0.96, 3.67, 32}
	starMap["K1 V"] = tabledata{"K1 V", 4760, 0.8, 0.44, 0.16, 0.63, 0.86, 3.32, 32}
	starMap["K2 V"] = tabledata{"K2 V", 4620, 0.7, 0.40, 0.14, 0.60, 0.82, 3.16, 28}
	starMap["K3 V"] = tabledata{"K3 V", 4480, 0.7, 0.34, 0.14, 0.55, 0.76, 2.92, 28}
	starMap["K4 V"] = tabledata{"K4 V", 4340, 0.7, 0.31, 0.14, 0.53, 0.72, 2.78, 28}
	starMap["K5 V"] = tabledata{"K5 V", 4200, 0.7, 0.27, 0.14, 0.49, 0.68, 2.60, 28}
	starMap["K6 V"] = tabledata{"K6 V", 4060, 0.6, 0.21, 0.12, 0.44, 0.60, 2.29, 24}
	starMap["K7 V"] = tabledata{"K7 V", 3920, 0.6, 0.19, 0.12, 0.41, 0.57, 2.18, 24}
	starMap["K8 V"] = tabledata{"K8 V", 3780, 0.6, 0.16, 0.12, 0.38, 0.52, 2.00, 24}
	starMap["K9 V"] = tabledata{"K9 V", 3640, 0.5, 0.14, 0.10, 0.36, 0.49, 1.87, 20}
	starMap["M0 V"] = tabledata{"M0 V", 3500, 0.5, 0.125, 0.100, 0.336, 0.460, 1.768, 20}
	starMap["M1 V"] = tabledata{"M1 V", 3333, 0.5, 0.0618, 0.100, 0.236, 0.323, 1.243, 20}
	starMap["M2 V"] = tabledata{"M2 V", 3167, 0.4, 0.0321, 0.080, 0.170, 0.233, 0.896, 16}
	starMap["M3 V"] = tabledata{"M3 V", 3000, 0.3, 0.0178, 0.060, 0.127, 0.173, 0.667, 12}
	starMap["M4 V"] = tabledata{"M4 V", 2833, 0.3, 0.0106, 0.060, 0.098, 0.134, 0.515, 12}
	starMap["M5 V"] = tabledata{"M5 V", 2667, 0.2, 0.00624, 0.040, 0.075, 0.103, 0.395, 8}
	starMap["M6 V"] = tabledata{"M6 V", 2500, 0.2, 0.00450, 0.040, 0.0637, 0.0872, 0.335, 8}
	starMap["M7 V"] = tabledata{"M7 V", 2333, 0.1, 0.00369, 0.020, 0.0577, 0.0790, 0.960, 4}
	starMap["M8 V"] = tabledata{"M8 V", 2167, 0.1, 0.00353, 0.020, 0.0564, 0.0772, 0.297, 4}
	starMap["M9 V"] = tabledata{"M9 V", 2000, 0.1, 0.00315, 0.020, 0.0533, 0.0730, 0.281, 4}
	return starMap[star]
}

func (s *star) LoadValues() error {
	code := s.Code()
	td := getTableValues(code)
	s.temperature = td.temperature
	s.mass = td.mass
	s.luminocity = td.luminocity
	s.innerLimit = td.innerLimit
	s.habitableLow = td.habitableLow
	s.habitableHigh = td.habitableHigh
	s.snowLine = td.snowLine
	s.outerLimit = td.outerLimit
	return nil
}