package weapon

func isAnyOf(source int, check ...int) bool {
	for _, c := range check {
		if c == source {
			return true
		}
	}
	return false
}
