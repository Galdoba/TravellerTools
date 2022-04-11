package astrogation

import (
	"github.com/Galdoba/convert"
	"github.com/Galdoba/utils"
)

//JumpDriveLimit - Возвращает номер последней орбиты
//которая находится в гравитационной тени звезды
func JumpDriveLimitStar(star string) (orbit int) {
	data := []byte(star)
	if len(data) == 2 {
		return -1
	}
	//fmt.Println("Error:", data, star)
	dec := convert.StoI(string(data[1]))
	if dec < 5 {
		dec = 0
	} else {
		dec = 5
	}
	seq := string(data[0]) + convert.ItoS(dec)
	size := string(data[2:])
	//fmt.Println(size, seq)
	col := 0
	switch size {
	default:
		col = 7
	case "Ia":
		col = 0
	case "Ib":
		col = 1
	case "II":
		col = 2
	case "III":
		col = 3
	case "IV":
		col = 4
		if seq == "K5" || seq == "M5" || seq == "M0" {
			col = 5
		}
	case "V":
		col = 5
	case "VI":
		col = 6
		if seq == "F0" || seq == "A5" || seq == "A0" {
			col = 5
		}
	}
	tableMap := make(map[string][]int)
	tableMap["A0"] = []int{10, 9, 7, 6, 5, 5, -99, -1}
	tableMap["A5"] = []int{10, 9, 7, 5, 4, 4, -99, -1}
	tableMap["F0"] = []int{11, 9, 7, 5, 4, 3, -99, -1}
	tableMap["F5"] = []int{11, 9, 7, 5, 4, 3, 3, -1}
	tableMap["G0"] = []int{11, 10, 8, 6, 4, 2, 2, -1}
	tableMap["G5"] = []int{12, 10, 8, 7, 4, 2, 1, -1}
	tableMap["K0"] = []int{12, 11, 9, 7, 5, 2, 0, -1}
	tableMap["K5"] = []int{13, 12, 10, 9, -99, 1, 0, -1}
	tableMap["M0"] = []int{14, 13, 11, 9, -99, 1, 0, -1}
	tableMap["M5"] = []int{15, 14, 13, 11, -99, 0, -1, -1}
	return tableMap[seq][col]
}

func OrbitToAU(orbit int) float64 {
	switch orbit {
	default:
		return -1.0
	case 0:
		return 0.2
	case 1:
		return 0.4
	case 2:
		return 0.7
	case 3:
		return 1.0
	case 4:
		return 1.6
	case 5:
		return 2.8
	case 6:
		return 5.2
	case 7:
		return 10.0
	case 8:
		return 20.0
	case 9:
		return 40.0
	case 10:
		return 77.0
	case 11:
		return 154.0
	case 12:
		return 308.0
	case 13:
		return 615.0
	case 14:
		return 1230.0
	case 15:
		return 2500.0
	case 16:
		return 4900.0
	case 17:
		return 9800.0
	case 18:
		return 19500.0
	case 19:
		return 39500.0
	case 20:
		return 78700.0
	}
}

func AUnitsToMegameters(au float64) float64 {
	if au < 0 {
		return 0
	}
	return utils.RoundFloat64(au*149597.9, 3)
}

const (
	AU2Megameters        = 149597.9
	SolDiametrMegameters = 13927.7
)

//au = 149597900
