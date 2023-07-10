package orbitns

import "github.com/Galdoba/TravellerTools/pkg/dice"

func D10(dice *dice.Dicepool) float64 {
	d := dice.Sroll("1d10")
	return float64(d)
}

func tableBaseDistance(i int) float64 {
	switch i {
	default:
		return -1.0
	case 0:
		return 0.0
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

func tableDifference(i int) float64 {
	switch i {
	default:
		return -1
	case 0:
		return 0.4
	case 1:
		return 0.3
	case 2:
		return 0.3
	case 3:
		return 0.6
	case 4:
		return 1.2
	case 5:
		return 2.4
	case 6:
		return 4.8
	case 7:
		return 10.0
	case 8:
		return 20.0
	case 9:
		return 37.0
	case 10:
		return 77.0
	case 11:
		return 154.0
	case 12:
		return 307.0
	case 13:
		return 615.0
	case 14:
		return 1270.0
	case 15:
		return 2400.0
	case 16:
		return 4900.0
	case 17:
		return 9700.0
	case 18:
		return 20000.0
	case 19:
		return 39200.0
	case 20:
		return -1
	}
}
