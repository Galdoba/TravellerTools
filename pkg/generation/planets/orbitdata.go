package planets

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

/*
Планетарная орбита кодируется в  хексов
1 - пара вокруг которой вращается планета
2 - номер орбиты от пары
3 - орбита спутника вокруг планеты
------------
Хексы 2 и 3 так же определяют тип мира (Планета\Спутник)
*/

func GeneratePlanetOrbit(dice *dice.Dicepool) ehex.Ehex {
	fl1 := dice.Flux()
	fl2 := dice.Flux() / 3
	vars := []int{12, 10, 8, 6, 4, 2, 0, 1, 3, 5, 7, 9, 11}[fl1+fl2+6]
	return ehex.New().Set(vars)
}

const (
	IsPlanet = 0
	IsClose  = 1
	IsFar    = 2
)

func GenerateSateliteOrbit(dice *dice.Dicepool, satMode int) ehex.Ehex {
	fl1 := dice.Flux()
	fl2 := dice.Flux() / 3
	isSat := []int{2, 2, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0}[fl1+fl2+6]
	switch satMode {
	case 0, 1, 2:
		isSat = satMode
	}
	switch isSat {
	default:
		return ehex.New().Set("*")
	case 1:
		fl1 = dice.Flux()
		fl2 = dice.Flux() / 3
		clSat := []string{"A", "B", "C", "D", "E", "F", "G", "H", "1", "J", "K", "L", "M"}[fl1+fl2+6]
		return ehex.New().Set(clSat)
	case 2:
		fl1 = dice.Flux()
		fl2 = dice.Flux() / 3
		frSat := []string{"N", "0", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}[fl1+fl2+6]
		return ehex.New().Set(frSat)
	}
}
