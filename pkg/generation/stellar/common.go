package stellar

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

type star struct {
	starType        string // буква определяющая спектр OBAFGKM
	spectralDecimal int    //число определяющее близость к типу 0123456789
	sizeClass       string //римское число определяющее размер Ia Ib II III IV V VI D (D - определяет белого карлика) BD (определяет коричнегого карлика)
	star            string
	temperature     int
	mass            float64
	luminocity      float64
	innerLimit      float64
	habitableLow    float64
	habitableHigh   float64
	snowLine        float64
	outerLimit      float64
	orbitingOn      int
	maxHighOrbit    int
	habitableOrbit  int
}

func NewStar(sType string, decimal int, size string) (*star, error) {
	s := star{}
	err := (error)(nil)
	switch sType {
	default:
		err = fmt.Errorf("Star Type was not adressed")
	case "O", "B", "A", "F", "G", "K", "M":
		s.starType = sType
	case "BD":
		//err = fmt.Errorf("'BD' is a special case of star Type for a White Dwarf")
		s.starType = "LTY"
		err = fmt.Errorf("'BD' is a special case of star Type for a White Dwarf (need specifics to draw data)")
		return &s, err
	case "T", "L", "Y":
		//err = fmt.Errorf("Special case of star Type for a Brown Dwarf")
		s.starType = sType
	}
	switch size {
	default:
		if s.starType != "L" && s.starType != "T" && s.starType != "Y" {
			err = fmt.Errorf("Size not implemented '%v'", size)
		}
	case "Ia", "Ib", "II", "III", "IV", "V", "VI", "D":
		s.sizeClass = size
	}
	switch decimal {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
		s.spectralDecimal = decimal
	}
	if (s.starType == "A" || s.starType == "F") && s.sizeClass == "VI" {
		s.sizeClass = "V"
	}
	s.star = s.starType + fmt.Sprintf("%v", s.spectralDecimal)
	if s.sizeClass != "" {
		s.star += " " + s.sizeClass
	}
	if s.sizeClass == "D" {
		s.star = s.sizeClass + fmt.Sprintf("%v", s.spectralDecimal)
	}

	tv := getTableValues(s.star)
	s.temperature = tv.temperature
	s.mass = tv.mass
	s.luminocity = tv.luminocity
	s.innerLimit = tv.innerLimit
	s.habitableLow = tv.habitableLow
	s.habitableHigh = tv.habitableHigh
	s.snowLine = tv.snowLine
	s.outerLimit = tv.outerLimit
	return &s, err
}

func fixedStarCode(spec string, dec int, siz string) string {
	code := ""
	switch spec {
	case "O", "B", "A":
		if siz == "VI" {
			siz = "V"
		}
	case "F", "G", "K", "M":
		if siz == "Ia" || siz == "Ib" {
			siz = "II"
		}
	case "L", "T", "Y":
		return spec + fmt.Sprintf("%v", dec)
	}
	if spec == "F" && siz == "VI" {
		siz = "V"
	}
	if spec == "BD" {
		return "BD"
	}
	code = spec + fmt.Sprintf("%v", dec)

	if siz != "" {
		code += " " + siz
	}
	if siz == "D" {
		code = siz + fmt.Sprintf("%v", dec)
	}
	return code
}

func brownDwarfSpectral(dice *dice.Dicepool) string {
	switch dice.Sroll("1d4") {
	case 1, 2:
		return "L"
	case 3:
		return "T"
	case 4:
		return "Y"
	}
	return "?"
}

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
	//Temperature(K)	Mass(SM)	Luminocity(SU)	InnerLimit(AU)	HabitableLow(AU)	HabitableHigh(AU)	SnowLine(AU)	OuterLimit(AU)
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
	starMap["O0 Ia"] = tabledata{"O0 Ia", 50000, 160.0, 34100000, 32.0, 5547.54, 6400, -999, 6400}
	starMap["O1 Ia"] = tabledata{"O1 Ia", 47600, 149.3, 2250000, 29.86, 1425, 1950, -999, 5972}
	starMap["O2 Ia"] = tabledata{"O2 Ia", 45200, 148.6, 2140000, 29.72, 1389.73, 1901.74, -999, 5944}
	starMap["O3 Ia"] = tabledata{"O3 Ia", 42800, 148.0, 1850000, 29.60, 1292.14, 1768.19, -999, 5920}
	starMap["O4 Ia"] = tabledata{"O4 Ia", 40400, 147.3, 1740000, 29.46, 1253.14, 1714.82, -999, 5892}
	starMap["O5 Ia"] = tabledata{"O5 Ia", 38000, 142.0, 1480000, 28.40, 1155.72, 1581.52, -999, 5680}
	starMap["O6 Ia"] = tabledata{"O6 Ia", 35400, 120.1, 1360000, 24.02, 1107.88, 1516.05, -999, 4804}
	starMap["O7 Ia"] = tabledata{"O7 Ia", 32800, 100.9, 1120000, 20.18, 1005.39, 1375.79, -999, 4036}
	starMap["O8 Ia"] = tabledata{"O8 Ia", 30200, 81.7, 913000, 16.34, 907.73, 1242.16, -999, 3268}
	starMap["O9 Ia"] = tabledata{"O9 Ia", 27600, 63.1, 731000, 12.62, 812.24, 1111.48, -999, 2524}
	starMap["B0 Ia"] = tabledata{"B0 Ia", 25000, 44.7, 573000, 8.94, 719.12, 984.06, -999, 1788}
	starMap["B1 Ia"] = tabledata{"B1 Ia", 23790, 40.0, 507000, 8.00, 676.44, 925.65, -999, 1600}
	starMap["B2 Ia"] = tabledata{"B2 Ia", 22580, 35.2, 446000, 7.04, 634.44, 868.18, -999, 1408}
	starMap["B3 Ia"] = tabledata{"B3 Ia", 21370, 30.5, 428000, 6.1, 621.51, 850.48, -999, 1220}
	starMap["B4 Ia"] = tabledata{"B4 Ia", 20160, 26.2, 338000, 5.24, 552.31, 755.79, -999, 1048}
	starMap["B5 Ia"] = tabledata{"B5 Ia", 18950, 21.9, 291000, 4.38, 512.47, 701.28, -999, 876}
	starMap["B6 Ia"] = tabledata{"B6 Ia", 17140, 20.2, 229000, 4.04, 454.61, 622.10, -999, 808}
	starMap["B7 Ia"] = tabledata{"B7 Ia", 15330, 18.6, 193000, 3.72, 417.35, 571.11, -999, 744}
	starMap["B8 Ia"] = tabledata{"B8 Ia", 13520, 16.9, 160000, 3.38, 380, 520, -999, 676}
	starMap["B9 Ia"] = tabledata{"B9 Ia", 11710, 15.3, 131000, 3.06, 343.84, 470.52, -999, 612}
	starMap["A0 Ia"] = tabledata{"A0 Ia", 9900, 13.7, 107000, 2.74, 310.75, 425.24, -999, 548}
	starMap["A1 Ia"] = tabledata{"A1 Ia", 9707, 13.1, 114000, 2.62, 320.76, 438.93, -999, 524}
	starMap["A2 Ia"] = tabledata{"A2 Ia", 9513, 12.5, 121000, 2.50, 330.46, 452.21, -999, 500}
	starMap["A3 Ia"] = tabledata{"A3 Ia", 9320, 12.0, 129000, 2.40, 341.21, 466.92, -999, 480}
	starMap["A4 Ia"] = tabledata{"A4 Ia", 9127, 11.4, 138000, 2.28, 352.91, 456, -999, 456}
	starMap["A5 Ia"] = tabledata{"A5 Ia", 8933, 10.8, 161000, 2.16, 381.19, 432, -999, 432}
	starMap["A6 Ia"] = tabledata{"A6 Ia", 8740, 10.8, 189000, 2.16, 413, 432, -999, 432}
	starMap["A7 Ia"] = tabledata{"A7 Ia", 8547, 10.8, 222000, 2.16, -999, -999, -999, 432}
	starMap["A8 Ia"] = tabledata{"A8 Ia", 8353, 10.8, 238000, 2.16, -999, -999, -999, 432}
	starMap["A9 Ia"] = tabledata{"A9 Ia", 8160, 10.8, 233000, 2.16, -999, -999, -999, 432}
	starMap["F0 Ia"] = tabledata{"F0 Ia", 7967, 10.8, 228000, 2.16, -999, -999, -999, 432}
	starMap["F1 Ia"] = tabledata{"F1 Ia", 7773, 10.3, 224000, 2.06, -999, -999, -999, 412}
	starMap["F2 Ia"] = tabledata{"F2 Ia", 7580, 9.9, 202000, 1.98, -999, -999, -999, 396}
	starMap["F3 Ia"] = tabledata{"F3 Ia", 7387, 9.4, 199000, 1.88, -999, -999, -999, 376}
	starMap["F4 Ia"] = tabledata{"F4 Ia", 7193, 9.0, 180000, 1.80, -999, -999, -999, 360}
	starMap["F5 Ia"] = tabledata{"F5 Ia", 7000, 8.6, 163000, 1.72, -999, -999, -999, 344}
	starMap["F6 Ia"] = tabledata{"F6 Ia", 6760, 8.6, 149000, 1.72, -999, -999, -999, 344}
	starMap["F7 Ia"] = tabledata{"F7 Ia", 6520, 8.5, 137000, 1.70, -999, -999, -999, 340}
	starMap["F8 Ia"] = tabledata{"F8 Ia", 6280, 8.5, 130000, 1.70, -999, -999, -999, 340}
	starMap["F9 Ia"] = tabledata{"F9 Ia", 6011, 8.5, 127000, 1.70, 338.55, 340, -999, 340}
	starMap["G0 Ia"] = tabledata{"G0 Ia", 5743, 6.3, 124000, 1.26, -999, -999, -999, 252}
	starMap["G1 Ia"] = tabledata{"G1 Ia", 5474, 6.6, 132000, 1.32, -999, -999, -999, 264}
	starMap["G2 Ia"] = tabledata{"G2 Ia", 5206, 6.9, 131000, 1.38, -999, -999, -999, 276}
	starMap["G3 Ia"] = tabledata{"G3 Ia", 4937, 7.2, 147000, 1.44, -999, -999, -999, 288}
	starMap["G4 Ia"] = tabledata{"G4 Ia", 4669, 7.5, 170000, 1.50, -999, -999, -999, 300}
	starMap["G5 Ia"] = tabledata{"G5 Ia", 4400, 7.8, 186000, 1.56, -999, -999, -999, 312}
	starMap["G6 Ia"] = tabledata{"G6 Ia", 4343, 7.9, 195000, 1.58, -999, -999, -999, 316}
	starMap["G7 Ia"] = tabledata{"G7 Ia", 4286, 8.0, 205000, 1.60, -999, -999, -999, 320}
	starMap["G8 Ia"] = tabledata{"G8 Ia", 4229, 8.1, 196000, 1.62, -999, -999, -999, 324}
	starMap["G9 Ia"] = tabledata{"G9 Ia", 4171, 8.1, 207000, 1.62, -999, -999, -999, 324}
	starMap["K0 Ia"] = tabledata{"K0 Ia", 4114, 8.2, 219000, 1.64, -999, -999, -999, 328}
	starMap["K1 Ia"] = tabledata{"K1 Ia", 4057, 8.3, 212000, 1.66, -999, -999, -999, 332}
	starMap["K2 Ia"] = tabledata{"K2 Ia", 4000, 8.8, 225000, 1.76, -999, -999, -999, 352}
	starMap["K3 Ia"] = tabledata{"K3 Ia", 3900, 9.4, 253000, 1.88, -999, -999, -999, 376}
	starMap["K4 Ia"] = tabledata{"K4 Ia", 3800, 9.9, 262000, 1.98, -999, -999, -999, 396}
	starMap["K5 Ia"] = tabledata{"K5 Ia", 3700, 10.4, 301000, 2.08, -999, -999, -999, 416}
	starMap["K6 Ia"] = tabledata{"K6 Ia", 3717, 10.6, 294000, 2.12, -999, -999, -999, 424}
	starMap["K7 Ia"] = tabledata{"K7 Ia", 3733, 10.7, 262000, 2.14, -999, -999, -999, 428}
	starMap["K8 Ia"] = tabledata{"K8 Ia", 3750, 10.9, 256000, 2.18, -999, -999, -999, 436}
	starMap["K9 Ia"] = tabledata{"K9 Ia", 3725, 11.1, 265000, 2.22, -999, -999, -999, 444}
	starMap["M0 Ia"] = tabledata{"M0 Ia", 3700, 13.3, 274000, 2.66, 497.28, 532, -999, 532}
	starMap["M1 Ia"] = tabledata{"M1 Ia", 3510, 12.7, 337000, 2.54, -999, -999, -999, 508}
	starMap["M2 Ia"] = tabledata{"M2 Ia", 3320, 12.1, 481000, 2.42, -999, -999, -999, 484}
	starMap["M3 Ia"] = tabledata{"M3 Ia", 3130, 11.6, 733000, 2.32, -999, -999, -999, 464}
	starMap["M4 Ia"] = tabledata{"M4 Ia", 2940, 11.0, 1110000, 2.20, -999, -999, -999, 440}
	starMap["M5 Ia"] = tabledata{"M5 Ia", 2750, 10.5, 2020000, 2.10, -999, -999, -999, 420}
	starMap["M6 Ia"] = tabledata{"M6 Ia", 2560, 10.3, 4170000, 2.06, -999, -999, -999, 412}
	starMap["M7 Ia"] = tabledata{"M7 Ia", 2370, 10.2, 9180000, 2.04, -999, -999, -999, 408}
	starMap["M8 Ia"] = tabledata{"M8 Ia", 2180, 10.1, 27000000, 2.02, -999, -999, -999, 404}
	starMap["M9 Ia"] = tabledata{"M9 Ia", 1990, 10.0, 103000000, 2.00, -999, -999, -999, 400}
	starMap["O0 Ib"] = tabledata{"O0 Ib", 50000, 140.0, 2150000, 28.00, 1392.97, 1906.47, -999, 5600}
	starMap["O1 Ib"] = tabledata{"O1 Ib", 47600, 149.3, 2250000, 29.86, 1425, 1950, -999, 5972}
	starMap["O2 Ib"] = tabledata{"O2 Ib", 45200, 137.9, 1620000, 27.58, 1209.15, 1654.63, -999, 5516}
	starMap["O3 Ib"] = tabledata{"O3 Ib", 42800, 136.9, 1400000, 27.38, 1124.06, 1538.18, -999, 5476}
	starMap["O4 Ib"] = tabledata{"O4 Ib", 40400, 135.8, 1200000, 27.16, 1040.67, 1424.08, -999, 5432}
	starMap["O5 Ib"] = tabledata{"O5 Ib", 38000, 125.6, 1030000, 25.12, 964.14, 1319.36, -999, 5024}
	starMap["O6 Ib"] = tabledata{"O6 Ib", 35400, 103.5, 781000, 20.70, 839.55, 1148.86, -999, 4140}
	starMap["O7 Ib"] = tabledata{"O7 Ib", 32800, 86.7, 588000, 17.34, 728.47, 996.86, -999, 3468}
	starMap["O8 Ib"] = tabledata{"O8 Ib", 30200, 69.9, 437000, 13.98, 628.01, 859.38, -999, 2796}
	starMap["O9 Ib"] = tabledata{"O9 Ib", 27600, 54.5, 319000, 10.90, 536.56, 734.24, -999, 2180}
	starMap["B0 Ib"] = tabledata{"B0 Ib", 25000, 39.2, 228000, 7.84, 453.62, 620.74, -999, 1568}
	starMap["B1 Ib"] = tabledata{"B1 Ib", 23790, 34.8, 184000, 6.96, 407.50, 557.64, -999, 1392}
	starMap["B2 Ib"] = tabledata{"B2 Ib", 22580, 30.4, 162000, 6.08, 382.37, 523.24, -999, 1216}
	starMap["B3 Ib"] = tabledata{"B3 Ib", 21370, 25.9, 129000, 5.18, 341.21, 466.92, -999, 1036}
	starMap["B4 Ib"] = tabledata{"B4 Ib", 20160, 22.3, 112000, 4.46, 317.93, 435.06, -999, 892}
	starMap["B5 Ib"] = tabledata{"B5 Ib", 18950, 18.7, 88000, 3.74, 281.82, 385.64, -999, 748}
	starMap["B6 Ib"] = tabledata{"B6 Ib", 17140, 17.2, 63100, 3.44, 238.64, 326.56, -999, 688}
	starMap["B7 Ib"] = tabledata{"B7 Ib", 15330, 15.8, 44300, 3.16, 199.95, 273.62, -999, 632}
	starMap["B8 Ib"] = tabledata{"B8 Ib", 13520, 14.3, 33400, 2.86, 173.62, 237.58, -999, 572}
	starMap["B9 Ib"] = tabledata{"B9 Ib", 11710, 12.9, 22700, 2.58, 143.13, 195.86, -999, 516}
	starMap["A0 Ib"] = tabledata{"A0 Ib", 9900, 11.5, 15400, 2.30, 117.89, 161.33, -999, 460}
	starMap["A1 Ib"] = tabledata{"A1 Ib", 9707, 11.0, 15000, 2.20, 116.35, 159.22, -999, 440}
	starMap["A2 Ib"] = tabledata{"A2 Ib", 9513, 10.5, 13300, 2.10, 109.56, 149.92, -999, 420}
	starMap["A3 Ib"] = tabledata{"A3 Ib", 9320, 10.0, 12900, 2.00, 107.90, 147.65, -999, 400}
	starMap["A4 Ib"] = tabledata{"A4 Ib", 9127, 9.5, 11400, 1.90, 101.43, 138.80, -999, 380}
	starMap["A5 Ib"] = tabledata{"A5 Ib", 8933, 9.0, 11100, 1.8, 100.09, 136.96, -999, 360}
	starMap["A6 Ib"] = tabledata{"A6 Ib", 8710, 9.0, 10900, 1.8, 99.18, 135.72, -999, 360}
	starMap["A7 Ib"] = tabledata{"A7 Ib", 8547, 9.0, 10600, 1.8, 97.81, 133.84, -999, 360}
	starMap["A8 Ib"] = tabledata{"A8 Ib", 8353, 9.0, 10400, 1.8, 96.88, 132.57, -999, 360}
	starMap["A9 Ib"] = tabledata{"A9 Ib", 8160, 9.0, 10200, 1.8, 95.95, 131.29, -999, 360}
	starMap["F0 Ib"] = tabledata{"F0 Ib", 7967, 8.9, 9960, 1.78, 94.81, 129.74, -999, 356}
	starMap["F1 Ib"] = tabledata{"F1 Ib", 7773, 8.6, 9790, 1.72, 94.00, 128.63, -999, 344}
	starMap["F2 Ib"] = tabledata{"F2 Ib", 7580, 8.2, 8800, 1.64, 89.12, 121.95, -999, 328}
	starMap["F3 Ib"] = tabledata{"F3 Ib", 7387, 7.9, 8700, 1.58, 88.61, 121.26, -999, 316}
	starMap["F4 Ib"] = tabledata{"F4 Ib", 7193, 7.5, 7860, 1.50, 84.22, 115.25, -999, 300}
	starMap["F5 Ib"] = tabledata{"F5 Ib", 7000, 7.1, 7820, 1.42, 84.01, 114.96, -999, 284}
	starMap["F6 Ib"] = tabledata{"F6 Ib", 6760, 7.1, 7820, 1.42, 84.01, 114.96, -999, 284}
	starMap["F7 Ib"] = tabledata{"F7 Ib", 6520, 7.1, 7180, 1.42, 80.50, 110.16, -999, 284}
	starMap["F8 Ib"] = tabledata{"F8 Ib", 6280, 7.1, 7290, 1.42, 81.11, 111, -999, 284}
	starMap["F9 Ib"] = tabledata{"F9 Ib", 6011, 7.0, 6840, 1.40, 78.57, 107.52, -999, 280}
	starMap["G0 Ib"] = tabledata{"G0 Ib", 5743, 2.5, 7150, 0.50, 80.33, 100, -999, 100}
	starMap["G1 Ib"] = tabledata{"G1 Ib", 5474, 2.6, 7620, 0.52, 82.93, 104, -999, 104}
	starMap["G2 Ib"] = tabledata{"G2 Ib", 5206, 2.8, 8290, 0.56, 86.50, 112, -999, 112}
	starMap["G3 Ib"] = tabledata{"G3 Ib", 4937, 2.9, 8460, 0.58, 87.38, 116, -999, 116}
	starMap["G4 Ib"] = tabledata{"G4 Ib", 4669, 3.0, 9770, 0.60, 93.90, 120, -999, 120}
	starMap["G5 Ib"] = tabledata{"G5 Ib", 4400, 3.2, 11800, 0.64, 103.20, 128, -999, 128}
	starMap["G6 Ib"] = tabledata{"G6 Ib", 4343, 3.3, 12300, 0.66, 105.36, 132, -999, 132}
	starMap["G7 Ib"] = tabledata{"G7 Ib", 4286, 3.5, 12900, 0.70, 107.90, 140, -999, 140}
	starMap["G8 Ib"] = tabledata{"G8 Ib", 4229, 3.6, 14900, 0.72, 115.96, 144, -999, 144}
	starMap["G9 Ib"] = tabledata{"G9 Ib", 4171, 3.8, 15700, 0.76, 119.03, 152, -999, 152}
	starMap["K0 Ib"] = tabledata{"K0 Ib", 4114, 3.9, 16600, 0.78, 122.40, 156, -999, 156}
	starMap["K1 Ib"] = tabledata{"K1 Ib", 4057, 4.1, 17600, 0.82, 126.03, 164, -999, 164}
	starMap["K2 Ib"] = tabledata{"K2 Ib", 4000, 4.3, 20600, 0.86, 136.35, 172, -999, 172}
	starMap["K3 Ib"] = tabledata{"K3 Ib", 3900, 4.6, 25300, 0.92, 151.11, 184, -999, 184}
	starMap["K4 Ib"] = tabledata{"K4 Ib", 3800, 4.8, 31500, 0.96, 168.61, 192, -999, 192}
	starMap["K5 Ib"] = tabledata{"K5 Ib", 3700, 5.0, 39600, 1.00, 189.05, 200, -999, 200}
	starMap["K6 Ib"] = tabledata{"K6 Ib", 3717, 5.3, 42400, 1.06, 195.62, 212, -999, 212}
	starMap["K7 Ib"] = tabledata{"K7 Ib", 3733, 5.7, 45500, 1.14, 202.64, 228, -999, 228}
	starMap["K8 Ib"] = tabledata{"K8 Ib", 3750, 6.0, 44400, 1.20, 200.18, 240, -999, 240}
	starMap["K9 Ib"] = tabledata{"K9 Ib", 3725, 6.3, 50400, 1.26, 213.27, 252, -999, 252}
	starMap["M0 Ib"] = tabledata{"M0 Ib", 3700, 10.7, 57300, 2.14, 227.41, 311.19, -999, 428}
	starMap["M1 Ib"] = tabledata{"M1 Ib", 3510, 10.2, 77300, 2.04, 264.13, 361.44, -999, 408}
	starMap["M2 Ib"] = tabledata{"M2 Ib", 3320, 9.8, 110000, 1.96, 315.08, 392, -999, 392}
	starMap["M3 Ib"] = tabledata{"M3 Ib", 3130, 9.3, 168000, 1.86, -999, -999, -999, 372}
	starMap["M4 Ib"] = tabledata{"M4 Ib", 2940, 8.9, 277000, 1.78, -999, -999, -999, 356}
	starMap["M5 Ib"] = tabledata{"M5 Ib", 2750, 8.4, 507000, 1.68, -999, -999, -999, 336}
	starMap["M6 Ib"] = tabledata{"M6 Ib", 2560, 8.3, 955000, 1.66, -999, -999, -999, 332}
	starMap["M7 Ib"] = tabledata{"M7 Ib", 2370, 8.2, 2100000, 1.64, -999, -999, -999, 328}
	starMap["M8 Ib"] = tabledata{"M8 Ib", 2180, 8.1, 5150000, 1.62, -999, -999, -999, 324}
	starMap["M9 Ib"] = tabledata{"M9 Ib", 1990, 8.0, 17900000, 1.60, -999, -999, -999, 320}
	starMap["O0 II"] = tabledata{"O0 II", 50000, 130.0, 2150000, 26.00, 1392.97, 1906.17, -999, 5200}
	starMap["O1 II"] = tabledata{"O1 II", 47800, 128.6, 1730000, 25.72, 1249.53, 1709.89, -999, 5144}
	starMap["O2 II"] = tabledata{"O2 II", 45600, 127.2, 1520000, 25.44, 1171.24, 1602.75, -999, 5088}
	starMap["O3 II"] = tabledata{"O3 II", 43400, 125.8, 1210000, 25.16, 1045, 1430, -999, 5032}
	starMap["O4 II"] = tabledata{"O4 II", 41200, 124.4, 960000, 24.88, 930.81, 1273.73, 4898.98, 4976}
	starMap["O5 II"] = tabledata{"O5 II", 39000, 109.2, 759000, 21.84, 827.65, 1132.57, 4356.03, 4368}
	starMap["O6 II"] = tabledata{"O6 II", 36800, 86.9, 654000, 17.38, 768.27, 1051.31, -999, 3476}
	starMap["O7 II"] = tabledata{"O7 II", 34600, 72.5, 510000, 14.50, 678.44, 928.39, -999, 2900}
	starMap["O8 II"] = tabledata{"O8 II", 32400, 58.2, 360000, 11.64, 570, 780, -999, 2328}
	starMap["O9 II"] = tabledata{"O9 II", 30200, 45.9, 276000, 9.18, 499.09, 682.96, -999, 1836}
	starMap["B0 II"] = tabledata{"B0 II", 28000, 33.8, 191000, 6.76, 415.18, 568.15, -999, 1352}
	starMap["B1 II"] = tabledata{"B1 II", 26190, 29.7, 134000, 5.94, 347.76, 475.88, -999, 1188}
	starMap["B2 II"] = tabledata{"B2 II", 24380, 25.5, 93600, 5.10, 290.49, 397.72, -999, 1020}
	starMap["B3 II"] = tabledata{"B3 II", 22570, 21.4, 64500, 4.28, 241.27, 330.16, -999, 856}
	starMap["B4 II"] = tabledata{"B4 II", 20760, 18.4, 43700, 3.68, 198.59, 271.76, -999, 736}
	starMap["B5 II"] = tabledata{"B5 II", 18950, 15.5, 29100, 3.10, 162.06, 221.76, -999, 620}
	starMap["B6 II"] = tabledata{"B6 II", 17140, 14.2, 17400, 2.84, 125.31, 171.48, -999, 568}
	starMap["B7 II"] = tabledata{"B7 II", 15330, 12.9, 11100, 2.58, 100.09, 136.96, -999, 516}
	starMap["B8 II"] = tabledata{"B8 II", 13520, 11.7, 6990, 2.34, 79.43, 108.69, 418.03, 468}
	starMap["B9 II"] = tabledata{"B9 II", 11710, 10.5, 4320, 2.10, 62.44, 85.45, 328.63, 420}
	starMap["A0 II"] = tabledata{"A0 II", 9900, 9.4, 2680, 1.88, 49.18, 67.30, 258.84, 376}
	starMap["A1 II"] = tabledata{"A1 II", 9650, 8.9, 2350, 1.78, 46.05, 63.02, 242.38, 356}
	starMap["A2 II"] = tabledata{"A2 II", 9400, 8.5, 2070, 1.70, 43.22, 59.15, 227.49, 340}
	starMap["A3 II"] = tabledata{"A3 II", 9150, 8.1, 1660, 1.62, 38.71, 52.97, 203.72, 324}
	starMap["A4 II"] = tabledata{"A4 II", 8900, 7.7, 1460, 1.54, 36.30, 49.67, 191.05, 308}
	starMap["A5 II"] = tabledata{"A5 II", 8650, 7.2, 1290, 1.44, 34.12, 46.69, 179.58, 288}
	starMap["A6 II"] = tabledata{"A6 II", 8400, 7.2, 1140, 1.44, 32.08, 43.89, 168.82, 288}
	starMap["A7 II"] = tabledata{"A7 II", 8150, 7.2, 1110, 1.44, 31.65, 43.31, 166.58, 288}
	starMap["A8 II"] = tabledata{"A8 II", 7900, 7.2, 990, 1.44, 29.89, 40.90, 157.32, 288}
	starMap["A9 II"] = tabledata{"A9 II", 7650, 7.1, 970, 1.42, 29.59, 40.49, 155.72, 284}
	starMap["F0 II"] = tabledata{"F0 II", 7400, 7.1, 870, 1.42, 28.02, 38.34, 147.48, 284}
	starMap["F1 II"] = tabledata{"F1 II", 7260, 6.8, 865, 1.36, 27.94, 38.23, 147.05, 272}
	starMap["F2 II"] = tabledata{"F2 II", 7120, 6.5, 860, 1.30, 27.86, 38.12, 146.63, 260}
	starMap["F3 II"] = tabledata{"F3 II", 6980, 6.3, 782, 1.26, 26.57, 36.35, 139.82, 252}
	starMap["F4 II"] = tabledata{"F4 II", 6840, 6.0, 781, 1.20, 26.55, 36.33, 139.73, 240}
	starMap["F5 II"] = tabledata{"F5 II", 6700, 5.7, 783, 1.14, 26.58, 36.38, 139.91, 228}
	starMap["F6 II"] = tabledata{"F6 II", 6560, 5.7, 786, 1.14, 26.63, 36.45, 140.18, 228}
	starMap["F7 II"] = tabledata{"F7 II", 6420, 5.6, 791, 1.12, 26.72, 36.56, 140.62, 224}
	starMap["F8 II"] = tabledata{"F8 II", 6280, 5.6, 729, 1.12, 25.65, 35.10, 135.00, 224}
	starMap["F9 II"] = tabledata{"F9 II", 6140, 5.6, 739, 1.12, 25.83, 35.34, 135.92, 224}
	starMap["G0 II"] = tabledata{"G0 II", 5743, 2.1, 784, 0.42, 26.60, 36.40, -999, 84}
	starMap["G1 II"] = tabledata{"G1 II", 5474, 2.2, 835, 0.44, 27.45, 37.57, -999, 88}
	starMap["G2 II"] = tabledata{"G2 II", 5206, 2.3, 909, 0.46, 28.64, 39.19, -999, 92}
	starMap["G3 II"] = tabledata{"G3 II", 4937, 2.4, 1110, 0.48, 31.65, 43.31, -999, 96}
	starMap["G4 II"] = tabledata{"G4 II", 4669, 2.5, 1290, 0.50, 34.12, 46.69, -999, 100}
	starMap["G5 II"] = tabledata{"G5 II", 4400, 2.6, 1550, 0.52, 37.40, 51.18, -999, 104}
	starMap["G6 II"] = tabledata{"G6 II", 4343, 2.7, 1620, 0.54, 38.24, 52.32, -999, 108}
	starMap["G7 II"] = tabledata{"G7 II", 4286, 2.8, 1700, 0.56, 39.17, 53.60, -999, 112}
	starMap["G8 II"] = tabledata{"G8 II", 4229, 2.9, 1960, 0.58, 42.06, 57.55, -999, 116}
	starMap["G9 II"] = tabledata{"G9 II", 4171, 3.0, 2070, 0.60, 43.22, 59.15, -999, 120}
	starMap["K0 II"] = tabledata{"K0 II", 4114, 3.1, 2190, 0.62, 44.46, 60.84, -999, 124}
	starMap["K1 II"] = tabledata{"K1 II", 4057, 3.3, 2320, 0.66, 45.76, 62.62, -999, 132}
	starMap["K2 II"] = tabledata{"K2 II", 4000, 3.4, 2470, 0.68, 47.21, 64.61, -999, 136}
	starMap["K3 II"] = tabledata{"K3 II", 3900, 3.6, 2770, 0.72, 50.00, 68.42, -999, 144}
	starMap["K4 II"] = tabledata{"K4 II", 3800, 3.8, 3150, 0.76, 53.32, 72.96, -999, 152}
	starMap["K5 II"] = tabledata{"K5 II", 3700, 3.9, 3620, 0.78, 57.16, 78.22, -999, 156}
	starMap["K6 II"] = tabledata{"K6 II", 3717, 4.2, 3530, 0.84, 56.44, 77.24, -999, 168}
	starMap["K7 II"] = tabledata{"K7 II", 3733, 4.4, 3450, 0.88, 55.80, 76.36, -999, 176}
	starMap["K8 II"] = tabledata{"K8 II", 3750, 4.6, 3690, 0.92, 57.71, 79.61, -999, 184}
	starMap["K9 II"] = tabledata{"K9 II", 3725, 4.9, 3830, 0.98, 58.79, 80.45, -999, 196}
	starMap["M0 II"] = tabledata{"M0 II", 3700, 8.2, 3960, 1.64, 59.78, 81.81, 314.64, 328}
	starMap["M1 II"] = tabledata{"M1 II", 3510, 7.8, 5860, 1.56, 72.72, 99.52, -999, 312}
	starMap["M2 II"] = tabledata{"M2 II", 3320, 7.4, 8360, 1.48, 86.86, 118.86, -999, 296}
	starMap["M3 II"] = tabledata{"M3 II", 3130, 7.1, 14400, 1.42, 114, 156, -999, 284}
	starMap["M4 II"] = tabledata{"M4 II", 2940, 6.7, 23310, 1.34, 145.04, 198.45, -999, 268}
	starMap["M5 II"] = tabledata{"M5 II", 2750, 6.4, 46200, 1.28, 204.19, 256, -999, 256}
	starMap["M6 II"] = tabledata{"M6 II", 2560, 6.3, 95500, 1.26, -999, -999, -999, 252}
	starMap["M7 II"] = tabledata{"M7 II", 2370, 6.2, 253000, 1.24, -999, -999, -999, 248}
	starMap["M8 II"] = tabledata{"M8 II", 2180, 6.1, 744000, 1.22, -999, -999, -999, 244}
	starMap["M9 II"] = tabledata{"M9 II", 1990, 6.0, 2830000, 1.20, -999, -999, -999, 240}
	starMap["O0 III"] = tabledata{"O0 III", 50000, 120.0, 2150000, 24.00, 1392.97, 1906.17, -999, 4800}
	starMap["O1 III"] = tabledata{"O1 III", 47800, 118.2, 1580000, 23.64, 1194.13, 1634.07, -999, 4728}
	starMap["O2 III"] = tabledata{"O2 III", 45600, 116.5, 1260000, 23.30, 1066.37, 159.25, -999, 4660}
	starMap["O3 III"] = tabledata{"O3 III", 43400, 114.7, 917000, 22.94, 909.72, 1244.88, -999, 4588}
	starMap["O4 III"] = tabledata{"O4 III", 41200, 112.9, 728000, 22.58, 810.57, 1109.20, -999, 4516}
	starMap["O5 III"] = tabledata{"O5 III", 39000, 92.8, 525000, 18.56, 688.34, 941.94, -999, 3712}
	starMap["O6 III"] = tabledata{"O6 III", 36800, 70.2, 376000, 14.04, 582.53, 797.14, -999, 2808}
	starMap["O7 III"] = tabledata{"O7 III", 34600, 58.4, 294000, 11.68, 515.11, 704.88, -999, 2336}
	starMap["O8 III"] = tabledata{"O8 III", 32400, 46.5, 207000, 9.30, 432.22, 591.46, -999, 1860}
	starMap["O9 III"] = tabledata{"O9 III", 30200, 37.3, 159000, 7.46, 378.81, 518.37, -999, 1492}
	starMap["B0 III"] = tabledata{"B0 III", 28000, 28.4, 109000, 5.68, 313.64, 429.20, -999, 1136}
	starMap["B1 III"] = tabledata{"B1 III", 26190, 24.5, 53400, 4.90, 219.53, 300.41, -999, 980}
	starMap["B2 III"] = tabledata{"B2 III", 24380, 20.6, 28300, 4.12, 159.81, 218.69, -999, 824}
	starMap["B3 III"] = tabledata{"B3 III", 22570, 16.8, 13500, 3.36, 110.38, 151.05, 580.95, 672}
	starMap["B4 III"] = tabledata{"B4 III", 20760, 14.5, 6930, 2.90, 79.08, 108.22, 416.23, 580}
	starMap["B5 III"] = tabledata{"B5 III", 18950, 12.3, 3190, 2.46, 53.66, 73.42, 282.40, 492}
	starMap["B6 III"] = tabledata{"B6 III", 17140, 11.2, 1740, 2.24, 39.63, 54.23, 208.57, 448}
	starMap["B7 III"] = tabledata{"B7 III", 15330, 10.1, 1010, 2.02, 30.19, 41.31, 158.90, 404}
	starMap["B8 III"] = tabledata{"B8 III", 13520, 9.0, 530, 1.80, 21.87, 29.93, 115.11, 360}
	starMap["B9 III"] = tabledata{"B9 III", 11710, 8.1, 299, 1.62, 16.43, 22.48, 86.46, 324}
	starMap["A0 III"] = tabledata{"A0 III", 9900, 7.2, 154.0, 1.44, 11.79, 16.13, 62.05, 288}
	starMap["A1 III"] = tabledata{"A1 III", 9650, 6.9, 124.0, 1.38, 10.58, 14.48, 55.68, 276}
	starMap["A2 III"] = tabledata{"A2 III", 9400, 6.5, 99.0, 1.30, 9.45, 12.93, 49.75, 260}
	starMap["A3 III"] = tabledata{"A3 III", 9150, 6.2, 87.1, 1.24, 8.87, 12.13, 46.66, 248}
	starMap["A4 III"] = tabledata{"A4 III", 8900, 5.8, 70.0, 1.16, 7.95, 10.88, 41.83, 232}
	starMap["A5 III"] = tabledata{"A5 III", 8650, 5.4, 56.4, 1.08, 7.13, 9.76, 37.55, 216}
	starMap["A6 III"] = tabledata{"A6 III", 8400, 5.4, 45.5, 1.08, 6.41, 8.77, 33.73, 216}
	starMap["A7 III"] = tabledata{"A7 III", 8150, 5.4, 36.8, 1.08, 5.76, 7.89, 30.33, 216}
	starMap["A8 III"] = tabledata{"A8 III", 7900, 5.4, 32.8, 1.08, 5.44, 7.45, 28.64, 216}
	starMap["A9 III"] = tabledata{"A9 III", 7650, 5.3, 26.7, 1.06, 4.91, 6.72, 25.84, 212}
	starMap["F0 III"] = tabledata{"F0 III", 7400, 5.3, 21.9, 1.06, 4.45, 6.08, 23.40, 212}
	starMap["F1 III"] = tabledata{"F1 III", 7260, 5.1, 21.7, 1.02, 4.43, 6.06, 23.29, 204}
	starMap["F2 III"] = tabledata{"F2 III", 7120, 4.9, 19.7, 0.98, 4.22, 5.77, 22.19, 196}
	starMap["F3 III"] = tabledata{"F3 III", 6980, 4.7, 19.6, 0.94, 4.21, 5.76, 22.14, 188}
	starMap["F4 III"] = tabledata{"F4 III", 6840, 4.5, 19.6, 0.90, 4.21, 5.76, 22.14, 180}
	starMap["F5 III"] = tabledata{"F5 III", 6700, 4.3, 21.6, 0.86, 4.42, 6.04, 23.24, 172}
	starMap["F6 III"] = tabledata{"F6 III", 6560, 4.2, 23.7, 0.84, 4.62, 6.33, 24.34, 168}
	starMap["F7 III"] = tabledata{"F7 III", 6420, 4.2, 26.2, 0.84, 4.86, 6.65, 25.60, 168}
	starMap["F8 III"] = tabledata{"F8 III", 6280, 4.1, 29.0, 0.82, 5.12, 7.00, 26.93, 164}
	starMap["F9 III"] = tabledata{"F9 III", 6140, 4.1, 32.2, 0.82, 5.39, 7.38, 28.37, 164}
	starMap["G0 III"] = tabledata{"G0 III", 5743, 1.8, 37.5, 0.36, 5.82, 7.96, 30.62, 72}
	starMap["G1 III"] = tabledata{"G1 III", 5474, 1.8, 43.8, 0.36, 6.29, 8.60, 33.09, 72}
	starMap["G2 III"] = tabledata{"G2 III", 5206, 1.9, 52.3, 0.38, 6.87, 9.40, 36.16, 76}
	starMap["G3 III"] = tabledata{"G3 III", 4937, 1.9, 70.3, 0.38, 7.97, 10.90, 41.92, 76}
	starMap["G4 III"] = tabledata{"G4 III", 4669, 2.0, 89.1, 0.40, 8.97, 12.27, 47.20, 80}
	starMap["G5 III"] = tabledata{"G5 III", 4400, 2.0, 118.0, 0.40, 10.32, 14.12, 54.31, 80}
	starMap["G6 III"] = tabledata{"G6 III", 4343, 2.1, 123.0, 0.42, 10.54, 14.42, 55.45, 84}
	starMap["G7 III"] = tabledata{"G7 III", 4286, 2.2, 129.0, 0.44, 10.79, 14.77, 56.79, 88}
	starMap["G8 III"] = tabledata{"G8 III", 4229, 2.2, 124.0, 0.44, 10.57, 14.48, 55.68, 88}
	starMap["G9 III"] = tabledata{"G9 III", 4171, 2.3, 131.0, 0.46, 10.87, 14.88, 57.23, 92}
	starMap["K0 III"] = tabledata{"K0 III", 4114, 2.3, 138.0, 0.46, 11.16, 15.27, 58.74, 92}
	starMap["K1 III"] = tabledata{"K1 III", 4057, 2.4, 161.0, 0.48, 12.05, 16.50, 63.44, 96}
	starMap["K2 III"] = tabledata{"K2 III", 4000, 2.5, 206.0, 0.50, 13.64, 18.66, 71.76, 100}
	starMap["K3 III"] = tabledata{"K3 III", 3900, 2.6, 253.0, 0.52, 15.11, 20.68, 79.53, 104}
	starMap["K4 III"] = tabledata{"K4 III", 3800, 2.7, 345.0, 0.54, 17.65, 24.15, 92.87, 108}
	starMap["K5 III"] = tabledata{"K5 III", 3700, 2.8, 435.0, 0.56, 19.81, 27.11, 104.28, 112}
	starMap["K6 III"] = tabledata{"K6 III", 3717, 3.0, 465.0, 0.60, 20.49, 28.03, 107.82, 120}
	starMap["K7 III"] = tabledata{"K7 III", 3733, 3.1, 499.0, 0.62, 21.22, 29.04, 111.69, 124}
	starMap["K8 III"] = tabledata{"K8 III", 3750, 3.3, 534.0, 0.66, 21.95, 30.04, 115.54, 132}
	starMap["K9 III"] = tabledata{"K9 III", 3725, 3.4, 606.0, 0.68, 23.39, 32.00, 123.09, 136}
	starMap["M0 III"] = tabledata{"M0 III", 3700, 5.6, 689, 1.12, 24.94, 34.12, 131.24, 224}
	starMap["M1 III"] = tabledata{"M1 III", 3510, 5.3, 929, 1.06, 28.96, 39.62, 152.40, 212}
	starMap["M2 III"] = tabledata{"M2 III", 3320, 5.1, 1210, 1.02, 33.05, 45.22, 173.93, 204}
	starMap["M3 III"] = tabledata{"M3 III", 3130, 4.8, 1840, 0.96, 40.75, 55.76, -999, 192}
	starMap["M4 III"] = tabledata{"M4 III", 2940, 4.6, 2770, 0.92, 50.00, 68.42, -999, 184}
	starMap["M5 III"] = tabledata{"M5 III", 2750, 4.3, 5070, 0.86, 67.64, 92.57, -999, 172}
	starMap["M6 III"] = tabledata{"M6 III", 2560, 4.2, 9550, 0.84, 92.84, 127.04, -999, 168}
	starMap["M7 III"] = tabledata{"M7 III", 2370, 4.2, 21000, 0.84, 137.67, 168, -999, 168}
	starMap["M8 III"] = tabledata{"M8 III", 2180, 4.1, 51500, 0.82, -999, -999, -999, 164}
	starMap["M9 III"] = tabledata{"M9 III", 1990, 4.1, 179000, 0.82, -999, -999, -999, 164}
	starMap["O0 IV"] = tabledata{"O0 IV", 50000, 110.0, 1360000, 22.00, 1107.88, 1516.05, -999, 4400}
	starMap["O1 IV"] = tabledata{"O1 IV", 47800, 107.9, 1090000, 21.58, 991.83, 1357.24, -999, 4316}
	starMap["O2 IV"] = tabledata{"O2 IV", 45600, 105.7, 872000, 21.14, 887.12, 1213.95, -999, 4228}
	starMap["O3 IV"] = tabledata{"O3 IV", 43400, 103.6, 696000, 20.72, 792.55, 1084.55, -999, 4144}
	starMap["O4 IV"] = tabledata{"O4 IV", 41200, 101.5, 552000, 20.30, 705.82, 965.86, 3714.84, 4060}
	starMap["O5 IV"] = tabledata{"O5 IV", 39000, 76.4, 437000, 15.28, 628.01, 859.38, -999, 3056}
	starMap["O6 IV"] = tabledata{"O6 IV", 36800, 53.6, 313000, 10.72, 531.49, 727.30, -999, 2144}
	starMap["O7 IV"] = tabledata{"O7 IV", 34600, 44.2, 223000, 8.84, 448.62, 613.90, -999, 1768}
	starMap["O8 IV"] = tabledata{"O8 IV", 32400, 34.7, 157000, 6.94, 376.42, 515.10, -999, 1388}
	starMap["O9 IV"] = tabledata{"O9 IV", 30200, 28.6, 110000, 5.72, 315.08, 431.16, -999, 1144}
	starMap["B0 IV"] = tabledata{"B0 IV", 28000, 28.4, 109000, 5.68, 313.64, 429.20, -999, 1136}
	starMap["B1 IV"] = tabledata{"B1 IV", 26190, 24.5, 53400, 4.90, 219.53, 300.41, -999, 980}
	starMap["B2 IV"] = tabledata{"B2 IV", 24380, 20.6, 28300, 4.12, 159.81, 218.69, -999, 824}
	starMap["B3 IV"] = tabledata{"B3 IV", 22570, 16.8, 13500, 3.36, 110.38, 151.05, 580.95, 672}
	starMap["B4 IV"] = tabledata{"B4 IV", 20760, 14.5, 6930, 2.90, 79.08, 108.22, 416.23, 580}
	starMap["B5 IV"] = tabledata{"B5 IV", 18950, 12.3, 3190, 2.46, 53.66, 73.42, 282.40, 492}
	starMap["B6 IV"] = tabledata{"B6 IV", 17140, 11.2, 1740, 2.24, 39.63, 54.23, 208.57, 448}
	starMap["B7 IV"] = tabledata{"B7 IV", 15330, 10.1, 1010, 2.02, 30.19, 41.31, 158.90, 404}
	starMap["B8 IV"] = tabledata{"B8 IV", 13520, 9.0, 530, 1.80, 21.87, 29.93, 115.11, 360}
	starMap["B9 IV"] = tabledata{"B9 IV", 11710, 8.1, 299, 1.62, 16.43, 22.48, 86.46, 324}
	starMap["A0 IV"] = tabledata{"A0 IV", 9900, 5.1, 88.8, 1.02, 8.95, 12.25, 47.12, 204}
	starMap["A1 IV"] = tabledata{"A1 IV", 9650, 4.8, 71.1, 0.96, 8.01, 10.96, 42.16, 192}
	starMap["A2 IV"] = tabledata{"A2 IV", 9400, 4.5, 57.0, 0.90, 7.17, 9.81, 37.75, 180}
	starMap["A3 IV"] = tabledata{"A3 IV", 9150, 4.3, 41.7, 0.86, 6.13, 8.39, 32.29, 172}
	starMap["A4 IV"] = tabledata{"A4 IV", 8900, 4.0, 33.5, 0.80, 5.50, 7.52, 28.94, 160}
	starMap["A5 IV"] = tabledata{"A5 IV", 8650, 3.7, 27.0, 0.74, 4.94, 6.75, 25.98, 148}
	starMap["A6 IV"] = tabledata{"A6 IV", 8400, 3.6, 21.8, 0.72, 4.44, 6.07, 23.35, 144}
	starMap["A7 IV"] = tabledata{"A7 IV", 8150, 3.6, 19.3, 0.72, 4.17, 5.71, 21.97, 144}
	starMap["A8 IV"] = tabledata{"A8 IV", 7900, 3.6, 15.7, 0.72, 3.76, 5.15, 19.81, 144}
	starMap["A9 IV"] = tabledata{"A9 IV", 7650, 3.5, 14.0, 0.70, 3.55, 4.86, 18.71, 140}
	starMap["F0 IV"] = tabledata{"F0 IV", 7400, 3.4, 11.5, 0.68, 3.22, 4.41, 16.96, 136}
	starMap["F1 IV"] = tabledata{"F1 IV", 7260, 3.3, 12.5, 0.66, 3.36, 4.60, 17.68, 132}
	starMap["F2 IV"] = tabledata{"F2 IV", 7120, 3.2, 14.9, 0.64, 3.67, 5.02, 19.30, 128}
	starMap["F3 IV"] = tabledata{"F3 IV", 6980, 3.1, 16.3, 0.62, 3.84, 5.25, 20.19, 124}
	starMap["F4 IV"] = tabledata{"F4 IV", 6840, 3.0, 19.6, 0.60, 4.21, 5.76, 22.14, 120}
	starMap["F5 IV"] = tabledata{"F5 IV", 6700, 2.8, 21.6, 0.56, 4.42, 6.04, 23.24, 112}
	starMap["F6 IV"] = tabledata{"F6 IV", 6560, 2.8, 16.4, 0.56, 3.85, 5.26, 20.25, 112}
	starMap["F7 IV"] = tabledata{"F7 IV", 6420, 2.7, 12.5, 0.54, 3.36, 4.60, 17.68, 108}
	starMap["F8 IV"] = tabledata{"F8 IV", 6280, 2.7, 10.5, 0.54, 3.08, 4.21, 16.20, 108}
	starMap["F9 IV"] = tabledata{"F9 IV", 6140, 2.6, 8.1, 0.52, 2.70, 3.70, 14.23, 104}
	starMap["G0 IV"] = tabledata{"G0 IV", 6000, 1.4, 6.25, 0.28, 2.38, 3.25, 12.50, 56}
	starMap["G1 IV"] = tabledata{"G1 IV", 5890, 1.4, 6.35, 0.28, 2.39, 3.28, 12.60, 56}
	starMap["G2 IV"] = tabledata{"G2 IV", 5780, 1.4, 6.48, 0.28, 2.42, 3.31, 12.73, 56}
	starMap["G3 IV"] = tabledata{"G3 IV", 5670, 1.5, 6.04, 0.30, 2.33, 3.19, 12.29, 60}
	starMap["G4 IV"] = tabledata{"G4 IV", 5560, 1.5, 6.20, 0.30, 2.37, 3.24, 12.45, 60}
	starMap["G5 IV"] = tabledata{"G5 IV", 5450, 1.5, 6.38, 0.30, 2.40, 3.28, 12.63, 60}
	starMap["G6 IV"] = tabledata{"G6 IV", 5340, 1.5, 6.59, 0.30, 2.44, 3.34, 12.84, 60}
	starMap["G7 IV"] = tabledata{"G7 IV", 5230, 1.5, 6.84, 0.30, 2.48, 3.40, 13.08, 60}
	starMap["G8 IV"] = tabledata{"G8 IV", 5120, 1.5, 6.50, 0.30, 2.42, 3.31, 12.75, 60}
	starMap["G9 IV"] = tabledata{"G9 IV", 5010, 1.6, 6.80, 0.32, 2.48, 3.39, 13.04, 64}
	starMap["K0 IV"] = tabledata{"K0 IV", 4900, 1.6, 7.16, 0.32, 2.54, 3.48, 13.38, 64}
	starMap["K1 IV"] = tabledata{"K1 IV", 4760, 1.6, 7.71, 0.32, 2.64, 3.61, 13.88, 64}
	starMap["K2 IV"] = tabledata{"K2 IV", 4620, 1.6, 8.38, 0.32, 2.75, 3.76, 14.47, 64}
	starMap["K3 IV"] = tabledata{"K3 IV", 4480, 1.7, 9.22, 0.34, 2.88, 3.95, 15.18, 68}
	starMap["K4 IV"] = tabledata{"K4 IV", 4340, 1.7, 10.30, 0.34, 3.05, 4.17, 16.05, 68}
	starMap["K5 IV"] = tabledata{"K5 IV", 4200, 1.8, 10.60, 0.36, 3.09, 4.23, 16.28, 72}
	starMap["K6 IV"] = tabledata{"K6 IV", 4060, 1.8, 12.20, 0.36, 3.32, 4.54, 17.46, 72}
	starMap["K7 IV"] = tabledata{"K7 IV", 3920, 1.9, 14.20, 0.38, 3.58, 4.90, 18.84, 76}
	starMap["K8 IV"] = tabledata{"K8 IV", 3780, 1.9, 17.00, 0.38, 3.92, 5.36, 20.62, 76}
	starMap["K9 IV"] = tabledata{"K9 IV", 3640, 2.0, 20.70, 0.40, 4.32, 5.91, 22.75, 80}
	starMap["M0 IV"] = tabledata{"M0 IV", 3500, 3.1, 26.0, 0.62, 4.84, 6.63, 25.50, 124}
	starMap["M1 IV"] = tabledata{"M1 IV", 3333, 2.9, 35.5, 0.58, 5.66, 7.75, 29.79, 116}
	starMap["M2 IV"] = tabledata{"M2 IV", 3167, 2.7, 50.9, 0.54, 6.78, 9.27, 35.67, 108}
	starMap["M3 IV"] = tabledata{"M3 IV", 3000, 2.6, 77.6, 0.52, 8.37, 11.45, 44.05, 104}
	starMap["M4 IV"] = tabledata{"M4 IV", 2833, 2.4, 127.0, 0.48, 10.71, 14.65, 56.35, 96}
	starMap["M5 IV"] = tabledata{"M5 IV", 2667, 2.3, 207.0, 0.46, 13.67, 18.70, 71.94, 92}
	starMap["M6 IV"] = tabledata{"M6 IV", 2500, 2.2, 410.0, 0.44, 19.24, 26.32, -999, 88}
	starMap["M7 IV"] = tabledata{"M7 IV", 2333, 2.1, 926.0, 0.42, 28.91, 39.56, -999, 84}
	starMap["M8 IV"] = tabledata{"M8 IV", 2167, 2.1, 2440.0, 0.42, 46.93, 64.22, -999, 84}
	starMap["M9 IV"] = tabledata{"M9 IV", 2000, 2.1, 7910.0, 0.42, -999, -999, -999, 84}
	starMap["G0 VI"] = tabledata{"G0 VI", 6000, 0.9, 0.520, 0.18, 0.69, 0.94, 3.61, 36}
	starMap["G1 VI"] = tabledata{"G1 VI", 5890, 0.9, 0.440, 0.18, 0.63, 0.86, 3.32, 36}
	starMap["G2 VI"] = tabledata{"G2 VI", 5780, 0.9, 0.373, 0.18, 0.58, 0.79, 3.05, 36}
	starMap["G3 VI"] = tabledata{"G3 VI", 5670, 0.8, 0.348, 0.16, 0.56, 0.77, 2.95, 32}
	starMap["G4 VI"] = tabledata{"G4 VI", 5560, 0.8, 0.297, 0.16, 0.52, 0.71, 2.72, 32}
	starMap["G5 VI"] = tabledata{"G5 VI", 5450, 0.8, 0.254, 0.16, 0.48, 0.66, 2.52, 32}
	starMap["G6 VI"] = tabledata{"G6 VI", 5340, 0.8, 0.218, 0.16, 0.44, 0.61, 2.33, 32}
	starMap["G7 VI"] = tabledata{"G7 VI", 5230, 0.7, 0.206, 0.14, 0.43, 0.59, 2.27, 28}
	starMap["G8 VI"] = tabledata{"G8 VI", 5120, 0.7, 0.179, 0.14, 0.40, 0.55, 2.12, 28}
	starMap["G9 VI"] = tabledata{"G9 VI", 5010, 0.7, 0.171, 0.14, 0.39, 0.54, 2.08, 28}
	starMap["K0 VI"] = tabledata{"K0 VI", 4900, 0.7, 0.1500, 0.14, 0.38, 0.50, 1.94, 28}
	starMap["K1 VI"] = tabledata{"K1 VI", 4760, 0.6, 0.1340, 0.12, 0.35, 0.48, 1.83, 24}
	starMap["K2 VI"] = tabledata{"K2 VI", 4620, 0.6, 0.1210, 0.12, 0.33, 0.45, 1.74, 24}
	starMap["K3 VI"] = tabledata{"K3 VI", 4480, 0.6, 0.1010, 0.12, 0.30, 0.41, 1.59, 24}
	starMap["K4 VI"] = tabledata{"K4 VI", 4340, 0.5, 0.0936, 0.10, 0.29, 0.40, 1.53, 20}
	starMap["K5 VI"] = tabledata{"K5 VI", 4200, 0.5, 0.0880, 0.10, 0.28, 0.39, 1.48, 20}
	starMap["K6 VI"] = tabledata{"K6 VI", 4060, 0.5, 0.0767, 0.10, 0.26, 0.36, 1.38, 20}
	starMap["K7 VI"] = tabledata{"K7 VI", 3920, 0.4, 0.0680, 0.08, 0.25, 0.34, 1.30, 16}
	starMap["K8 VI"] = tabledata{"K8 VI", 3780, 0.3, 0.0562, 0.06, 0.23, 0.31, 1.19, 16}
	starMap["K9 VI"] = tabledata{"K9 VI", 3640, 0.3, 0.0433, 0.06, 0.20, 0.27, 1.04, 16}
	starMap["M0 VI"] = tabledata{"M0 VI", 3500, 0.20, 0.03760, 0.04, 0.184, 0.252, 0.970, 8}
	starMap["M1 VI"] = tabledata{"M1 VI", 3333, 0.20, 0.01860, 0.04, 0.130, 0.177, 0.682, 8}
	starMap["M2 VI"] = tabledata{"M2 VI", 3167, 0.20, 0.00885, 0.04, 0.089, 0.122, 0.470, 8}
	starMap["M3 VI"] = tabledata{"M3 VI", 3000, 0.20, 0.00490, 0.04, 0.067, 0.091, 0.350, 8}
	starMap["M4 VI"] = tabledata{"M4 VI", 2833, 0.10, 0.00266, 0.02, 0.049, 0.067, 0.258, 4}
	starMap["M5 VI"] = tabledata{"M5 VI", 2667, 0.10, 0.00172, 0.02, 0.039, 0.054, 0.207, 4}
	starMap["M6 VI"] = tabledata{"M6 VI", 2500, 0.10, 0.00163, 0.02, 0.038, 0.053, 0.202, 4}
	starMap["M7 VI"] = tabledata{"M7 VI", 2333, 0.10, 0.00194, 0.02, 0.042, 0.057, 0.220, 4}
	starMap["M8 VI"] = tabledata{"M8 VI", 2167, 0.10, 0.00244, 0.02, 0.047, 0.064, 0.247, 4}
	starMap["M9 VI"] = tabledata{"M9 VI", 2000, 0.10, 0.00415, 0.002, 0.061, 0.084, 0.322, 4}
	starMap["D0"] = tabledata{"D0", 100000, 1.1, 6.91, 0.22, 2.49, 3.41, 13.14, 44}
	starMap["D1"] = tabledata{"D1", 50400, 0.9, 0.265, 0.18, 0.49, 0.67, 2.57, 36}
	starMap["D2"] = tabledata{"D2", 25200, 0.8, 0.0255, 0.16, 0.16, 0.21, 0.798, 32}
	starMap["D3"] = tabledata{"D3", 16800, 0.7, 0.006, 0.14, -999, -999, 0.387, 28}
	starMap["D4"] = tabledata{"D4", 12600, 0.6, 0.0018, 0.12, -999, -999, 0.212, 24}
	starMap["D5"] = tabledata{"D5", 10080, 0.5, 0.000693, 0.10, -999, -999, 0.131, 20}
	starMap["D6"] = tabledata{"D6", 8400, 0.4, 0.000315, 0.08, -999, -999, 0.088, 16}
	starMap["D7"] = tabledata{"D7", 7200, 0.3, 0.000180, 0.06, -999, -999, 0.067, 12}
	starMap["D8"] = tabledata{"D8", 6300, 0.2, 0.000105, 0.04, -999, -999, 0.051, 8}
	starMap["D9"] = tabledata{"D9", 5600, 0.1, 0.0000673, 0.02, -999, -999, 0.041, 4}
	starMap["L0"] = tabledata{"L0", 2200, 0.08, 0.005, 0.016, 0.067, 0.092, 0.354, 3.2}
	starMap["L1"] = tabledata{"L1", 2100, 0.08, 0.004, 0.016, 0.060, 0.082, 0.316, 3.2}
	starMap["L2"] = tabledata{"L2", 2000, 0.07, 0.003, 0.014, 0.052, 0.071, 0.274, 2.8}
	starMap["L3"] = tabledata{"L3", 1900, 0.07, 0.001, 0.014, 0.030, 0.041, 0.158, 2.8}
	starMap["L4"] = tabledata{"L4", 1800, 0.07, 0.0007, 0.014, 0.025, 0.034, 0.132, 2.8}
	starMap["L5"] = tabledata{"L5", 1700, 0.06, 0.0005, 0.012, 0.021, 0.029, 0.112, 2.4}
	starMap["L6"] = tabledata{"L6", 1600, 0.06, 0.0001, 0.012, 0.012, 0.013, 0.050, 2.4}
	starMap["L7"] = tabledata{"L7", 1450, 0.05, 0.00007, 0.010, 0.010, 0.011, 0.418, 2.0}
	starMap["L8"] = tabledata{"L8", 1425, 0.05, 0.00005, 0.010, -999, -999, 0.035, 2.0}
	starMap["L9"] = tabledata{"L9", 1410, 0.05, 0.00001, 0.010, -999, -999, 0.016, 2.0}
	starMap["T0"] = tabledata{"T0", 1400, 0.05, 0.0000060, 0.010, -999, -999, 0.012, 2.0}
	starMap["T1"] = tabledata{"T1", 1350, 0.04, 0.0000060, 0.008, -999, -999, 0.012, 1.6}
	starMap["T2"] = tabledata{"T2", 1300, 0.04, 0.0000055, 0.008, -999, -999, 0.012, 1.6}
	starMap["T3"] = tabledata{"T3", 1200, 0.04, 0.0000050, 0.008, -999, -999, 0.011, 1.6}
	starMap["T4"] = tabledata{"T4", 1100, 0.04, 0.0000040, 0.008, -999, -999, 0.010, 1.6}
	starMap["T5"] = tabledata{"T5", 1000, 0.03, 0.0000040, 0.006, -999, -999, 0.010, 1.2}
	starMap["T6"] = tabledata{"T6", 900, 0.03, 0.0000035, 0.006, -999, -999, 0.009, 1.2}
	starMap["T7"] = tabledata{"T7", 800, 0.03, 0.0000030, 0.006, -999, -999, 0.009, 1.2}
	starMap["T8"] = tabledata{"T8", 750, 0.03, 0.0000020, 0.006, -999, -999, 0.007, 1.2}
	starMap["T9"] = tabledata{"T9", 700, 0.03, 0.0000010, 0.006, -999, -999, 0.006, 1.2}
	starMap["Y0"] = tabledata{"Y0", 448, 0.03, 0.0000006, 0.006, -999, -999, 0.006, 1.2}
	starMap["Y1"] = tabledata{"Y1", 433, 0.03, 0.0000006, 0.006, -999, -999, 0.006, 1.2}
	starMap["Y2"] = tabledata{"Y2", 418, 0.03, 0.0000005, 0.006, -999, -999, 0.006, 1.2}
	starMap["Y3"] = tabledata{"Y3", 403, 0.03, 0.0000003, 0.006, -999, -999, 0.006, 1.2}
	starMap["Y4"] = tabledata{"Y4", 388, 0.03, 0.0000001, 0.006, -999, -999, 0.006, 1.2}
	starMap["Y5"] = tabledata{"Y5", 373, 0.03, 0.00000007, 0.006, -999, -999, 0.006, 1.2}
	starMap["Y6"] = tabledata{"Y6", 358, 0.03, 0.00000005, 0.006, -999, -999, 0.006, 1.2}
	starMap["Y7"] = tabledata{"Y7", 343, 0.03, 0.00000003, 0.006, -999, -999, 0.006, 1.2}
	starMap["Y8"] = tabledata{"Y8", 328, 0.03, 0.00000001, 0.006, -999, -999, 0.006, 1.2}
	starMap["Y9"] = tabledata{"Y9", 298, 0.03, 0.000000007, 0.006, -999, -999, 0.006, 1.2}
	return starMap[star]
}
