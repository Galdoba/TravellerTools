package weapon

import "fmt"

/*
1 Decide general type (pistol, rifle, ect...)
2 choose ammnition/powertype
3 choose receiver
4 choose receiver's mode of operation (breechloader, semi-automatic, ect...)
5 assign barrel lenght
6 assign furniture
7 choose feed device
8 add accessories that come as standard
9 total cost and weight

*/

type Weapon struct {
	rcvr *receiver
	brl  *barrel
}

func contains(sl []int, e int) bool {
	for _, val := range sl {
		if val == e {
			return true
		}
	}
	return false
}

func inErr() error {
	return fmt.Errorf("Input is incorrect")
}

func timesCrossed(aSlice []int, bSlice []int) int {
	//сколько раз встречаются элементы слайса А в слайсе Б?
	met := 0
	for _, bElem := range bSlice {
		for _, aElem := range aSlice {
			if aElem == bElem {
				met++
			}
		}
	}
	return met
}
