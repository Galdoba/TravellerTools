package helper

func SliceStrContains(slice []string, element string) bool {
	for _, val := range slice {
		if element == val {
			return true
		}
	}
	return false
}

func SliceIntContains(slice []int, element int) bool {
	for _, val := range slice {
		if element == val {
			return true
		}
	}
	return false
}

func AppendIfUniqueInt(sl []int, newElem int) []int {
	for _, elem := range sl {
		if elem == newElem {
			return sl
		}
	}
	sl = append(sl, newElem)
	return sl
}
