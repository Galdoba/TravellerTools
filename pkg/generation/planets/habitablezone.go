package planets

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

/*
Mainworld Habitable Zone:
...
-3= 7
-2= 8
-1= 9
0 = A
1 = B
2 = C
3 = D
...
*/

const (
	HZ_INNER      = "Inner"
	HZ_HOSPITABLE = "Hospitable"
	HZ_OUTER      = "Outer"
)

// Передаем отступ от самой благоприятной орбиты звезды
// 0  - отлично
// -1 - жарко
// 1  - холодно
func SetHZVar(i int) ehex.Ehex {
	hz := ehex.New().Set(i + 10)
	switch i {
	default:
		if i > 10 {
			hz.Encode(HZ_OUTER)
		}
		if i < 10 {
			hz.Encode(HZ_INNER)
		}
	case 9, 10, 11:
		hz.Encode(HZ_HOSPITABLE)
	}
	return hz
}

// вычисление отступа при известных данных звезды и орбиты планеты
func HabitableZoneVar(starHZ int, planetOrbit int) ehex.Ehex {
	hz := ehex.New().Set(10 + planetOrbit - starHZ)
	switch hz.Value() {
	default:
		hz.Encode(HZ_OUTER)
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
		hz.Encode(HZ_INNER)
	case 9, 10, 11:
		hz.Encode(HZ_HOSPITABLE)
	}
	return hz
}

func RollMainworldHabitableZone(dice *dice.Dicepool) ehex.Ehex {
	fl1 := dice.Flux()
	fl2 := dice.Flux() / 3
	vars := []int{8, 9, 9, 9, 10, 10, 10, 10, 10, 11, 11, 11, 12}[fl1+fl2+6]
	return ehex.New().Set(vars)
}
