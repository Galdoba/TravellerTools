package components

import "fmt"

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

func contains(sl []int, e int) bool {
	for _, val := range sl {
		if val == e {
			return true
		}
	}
	return false
}
