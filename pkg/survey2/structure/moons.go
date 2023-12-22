package structure

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

//MoonsQuantity - return Moons Quantity AND Ring Allowance
func MoonsQuantity(parentSize string, dice *dice.Dicepool, dmPerDice int) (int, bool) {
	diceNum := 0
	totalDM := 0
	dm := 0
	switch parentSize {
	case "1", "2":
		diceNum = 1
		dm = -5
	case "3", "4", "5", "6", "7", "8", "9":
		diceNum = 2
		dm = -8
	case "A", "B", "C", "D", "E", "F":
		diceNum = 2
		dm = -6
	case "GS":
		diceNum = 3
		dm = -7
	case "GM", "GL":
		diceNum = 4
		dm = -6
	}
	diceCode := fmt.Sprintf("%vd6")
	totalDM = (diceNum * dmPerDice) + dm
	rSize := dice.Sroll(diceCode) + totalDM
	rings := false
	if rSize == 0 {
		rings = true
		rSize, _ = MoonsQuantity(parentSize, dice, dmPerDice)
	}
	return rSize, rings
	//moonSize := ""
	panic("TODO: разобраться результат 0: перебрасываем количество или нет ")
	/*
			ПРИМЕР:
		if 0 {
			parentSize = parentSize + parentSize

		}
	*/
	// rSize := dice.Sroll(diceCode) + totalDM
	// if rSize < 0 {
	// 	fmt.Println("no moons")
	// }
	// if rSize == 0 {
	// 	fmt.Println("Rings Present")
	// }
	return 0 //rSize
}
