package probability

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/utils"
)

const (
	RESULT_IS_SAME = iota
	RESULT_IS_SAME_OR_MORE
	RESULT_IS_SAME_OR_LESS
	RESULT_IS_MORE
	RESULT_IS_LESS
)

func Calculate(operation int, code string, tn int) (float64, error) {
	dp := dice.New()
	dp.Roll(code)
	die, edge := dp.DrawData()
	if die > 10 {
		return 0, fmt.Errorf("dice number must be [0-10]")
	}
	dm := dp.ModTotal()
	var totalCombinations int64
	var validCombinations int64
	totalCombinations = 1
	validCombinations = 0
	resultArray := []int{}
	for i := 0; i < die; i++ {
		totalCombinations = totalCombinations * int64(edge)
		resultArray = append(resultArray, 1)
	}
	if totalCombinations > 1000000000 {
		return 0, fmt.Errorf("total combinations is above 1,000,000,000 (%v)", totalCombinations)
	}
	for arrSum(resultArray) != 0 {
		switch operation {
		case RESULT_IS_SAME:
			if arrSum(resultArray)+dm == tn {
				validCombinations++
			}
		case RESULT_IS_SAME_OR_MORE:
			if arrSum(resultArray)+dm >= tn {
				validCombinations++
			}
		case RESULT_IS_SAME_OR_LESS:
			if arrSum(resultArray)+dm <= tn {
				validCombinations++
			}
		case RESULT_IS_MORE:
			if arrSum(resultArray)+dm > tn {
				validCombinations++
			}
		case RESULT_IS_LESS:
			if arrSum(resultArray)+dm < tn {
				validCombinations++
			}
		}
		resultArray = nextResult(resultArray, edge)
	}
	res := float64(validCombinations) / float64(totalCombinations)
	res = utils.RoundFloat64(res, 6)
	return res, nil
}

func arrSum(sl []int) int {
	s := 0
	for _, v := range sl {
		s += v
	}
	return s
}

func nextResult(resSl []int, edges int) []int {
	last := len(resSl) - 1
	for {
		if resSl[last] < edges {
			resSl[last]++
			return resSl
		}
		if resSl[last] == edges {
			last--
			resSl[last+1] = 1
		}
		if last == -1 {
			break
		}
	}
	return []int{0}
}
