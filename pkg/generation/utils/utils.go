package utils

import (
	"sort"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func RollChart() {

}

type DiceResultChart struct {
	rows []drcRow
}

type drcRow struct {
	codeColumn   string
	resultColumn []string
}

func New(rows ...drcRow) DiceResultChart {
	drc := DiceResultChart{}
	for _, r := range rows {
		drc.rows = append(drc.rows, r)
	}
	drc.veryfy()
	return drc
}

func (drc *DiceResultChart) veryfy() error {
	allVals := []int{}
	for _, r := range drc.rows {
		vals := parse(r.codeColumn)
		allVals = append(allVals, vals...)
	}
	sort.Ints(allVals)
	min := -1000
	max := 1000
	valMap := make(map[int]int)
	for _, v := range allVals {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

}

func intPresent(i int,sl []int) bool {
	for i
}

func Row(code string, result ...string) drcRow {
	return drcRow{code, result}
}

func parse(code string) []int {
	codeArray := []int{}
	n, e := strconv.Atoi(code)
	switch {
	case e == nil:
		codeArray = append(codeArray, n)
		return codeArray
	case strings.HasSuffix(code, "-"):
		code = strings.TrimSuffix(code, "-")
		n, e := strconv.Atoi(code)
		if e == nil {
			for i := n; i > -100; i-- {
				codeArray = append(codeArray, n)
			}
			return codeArray
		}
	case strings.HasSuffix(code, "+"):
		code = strings.TrimSuffix(code, "+")
		n, e := strconv.Atoi(code)
		if e == nil {
			for i := n; i < 100; i++ {
				codeArray = append(codeArray, n)
			}
			return codeArray
		}
	default:
		vals := strings.Split(code, " -- ")
		if len(vals) == 2 {
			n, e := strconv.Atoi(code)
			if e != nil {
				return []int{}
			}
			m, e := strconv.Atoi(code)
			if e != nil {
				return []int{}
			}
			for n <= m {
				codeArray = append(codeArray, n)
				n++
			}
		}
	}
	return codeArray
}

func RollOnChart(chart DiceResultChart, dice *dice.Dicepool) []string {

}

/*
code:
n
n+
n-
n-m

*/
