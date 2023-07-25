package orbitns

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func D10(dice *dice.Dicepool) float64 {
	d := dice.Sroll("1d10")
	return float64(d)
}

func decodeUINT(i uint) (int, int, int64, error) {
	str := fmt.Sprintf("%v", i)
	for len(str) < 10 {
		str += "0"
	}
	data := strings.Split(str, "")
	switch data[0] {
	default:
		return 0, 0, 0, fmt.Errorf("position key incorect (0)")
	case "4":
		if data[1] != "0" {
			return 0, 0, 0, fmt.Errorf("position key incorect (1)")
		}
	case "2", "3":
	}
	switch data[2] {
	default:
		return 0, 0, 0, fmt.Errorf("full orbit# data incorect (2)")
	case "2":
		if data[3] != "0" {
			return 0, 0, 0, fmt.Errorf("full orbit# data incorect (3)")
		}
	case "0", "1":
	}
	prim, _ := strconv.Atoi(strings.Join(data[:2], ""))
	full, _ := strconv.Atoi(strings.Join(data[2:4], ""))
	dist, _ := strconv.Atoi(strings.Join(data[4:], ""))
	return prim - 20, full, int64(dist), nil
}

func encodeUINT(f float64, i ...int) (uint, error) {
	if f < 0 && f >= 20 {
		return 0, fmt.Errorf("incorect orbit#")
	}
	rp := uint(2000000000)
	for _, relative := range i {
		if relative > 0 && relative <= 20 {
			rp += uint(relative * 100000000)
		}
		break
	}
	fl := uint(f * 1000000)
	rp += fl
	return rp, nil
}

func encodeFL2INT(f float64) int {
	if f < 0 && f >= 20 {
		return -1
	}
	return int(f * 1000)
}
func decodeINT2FL(i int) float64 {
	fl := float64(i)
	return fl / 1000.0
}

func decodeINT2FRACK(i int) (int, int) {
	return i / 1000, i % 1000
}

//возвращает 1/1000 (milliAU) as
func tableBaseDistance(i int) int {
	switch i {
	default:
		return -1
	case 0:
		return 0
	case 1:
		return 400
	case 2:
		return 700
	case 3:
		return 1000
	case 4:
		return 1600
	case 5:
		return 2800
	case 6:
		return 5200
	case 7:
		return 10000
	case 8:
		return 20000
	case 9:
		return 40000
	case 10:
		return 77000
	case 11:
		return 154000
	case 12:
		return 308000
	case 13:
		return 615000
	case 14:
		return 1230000
	case 15:
		return 2500000
	case 16:
		return 4900000
	case 17:
		return 9800000
	case 18:
		return 19500000
	case 19:
		return 39500000
	case 20:
		return 78700000
	}
}

func tableDifference(i int) int {
	switch i {
	default:
		return -1
	case 0:
		return 400
	case 1:
		return 300
	case 2:
		return 300
	case 3:
		return 600
	case 4:
		return 1200
	case 5:
		return 2400
	case 6:
		return 4800
	case 7:
		return 10000
	case 8:
		return 20000
	case 9:
		return 37000
	case 10:
		return 77000
	case 11:
		return 154000
	case 12:
		return 307000
	case 13:
		return 615000
	case 14:
		return 1270000
	case 15:
		return 2400000
	case 16:
		return 4900000
	case 17:
		return 9700000
	case 18:
		return 20000000
	case 19:
		return 39200000
	case 20:
		return -1
	}
}
