package components

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func TestFurniture(t *testing.T) {
	fmt.Println("------FURNITURE--------")
	input := []int{FRNTR_STOCKLESS, FRNTR_STOCK_FOLDING, FRNTR_STOCK_FULL, FRNTR_MODULARIZATION, FRNTR_BIPOD_ABSENT, FRNTR_BIPOD_FIXED, FRNTR_BIPOD_DETACHABLE, FRNTR_SUPPORT_MOUNT, WRONG_INSTRUCTION}
	dtStr := []string{}
	for _, s := range input {
		dtStr = append(dtStr, strconv.Itoa(s))
	}
	fmt.Println("calculating all combinations for", dtStr)
	comb := CombinationsTracked(dtStr, 0, false)
	testNum := 0
	errors := 0
	errrMap := make(map[string]int)
	for _, strComb := range comb {
		testNum++
		inp := []int{}
		for _, sInp := range strComb {
			i, _ := strconv.Atoi(sInp)
			inp = append(inp, i)
		}
		_, err := NewFurniture(inp...)

		if err != nil {
			errrMap[err.Error()]++
			errors++
		}
	}
	fmt.Println("Total", testNum, " | errors", errors, "| correct =", testNum-errors)

	errNames := []string{}
	for k, _ := range errrMap {
		errNames = append(errNames, k)
	}
	sort.Strings(errNames)
	for _, name := range errNames {
		switch name {
		default:
			fmt.Println("Error:", name, errrMap[name])
		case "":
			fmt.Println("Correct:", errrMap[name])
		}
	}
	fmt.Println("-----------------------")
}
