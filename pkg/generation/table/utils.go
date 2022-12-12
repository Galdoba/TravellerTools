package table

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func RollChart() {

}

type DiceResultChart struct {
	rows     []drcRow
	verified bool
}

type drcRow struct {
	codeColumn string
	result     string
}

func DiceChart(rows ...drcRow) *DiceResultChart {
	drc := DiceResultChart{}
	for _, r := range rows {
		drc.rows = append(drc.rows, r)
	}
	drc.verified = veryfy(drc)
	return &drc
}

func (drc *DiceResultChart) Result(i int) string {
	for _, r := range drc.rows {
		arr := parse(r.codeColumn)
		for _, val := range arr {
			if val == i {
				return r.result
			}
		}
	}
	return ""
}

func veryfy(drc DiceResultChart) bool {
	allVals := []int{}
	for _, r := range drc.rows {
		vals := parse(r.codeColumn)
		allVals = append(allVals, vals...)
	}
	sort.Ints(allVals)
	last := 0
	for i, n := range allVals {
		if i == 0 {
			last = n - 1
			continue
		}
		if n != last+1 {
			return false
		}
		last = n
	}
	return true
}

func Row(code string, result string) drcRow {
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
	}
	vals := strings.Split(code, "-")
	if len(vals) == 2 {
		n, e := strconv.Atoi(vals[0])
		if e != nil {
			return []int{}
		}
		m, e := strconv.Atoi(vals[1])
		if e != nil {
			return []int{}
		}
		for n <= m {
			codeArray = append(codeArray, n)
			n++
		}
	}
	return codeArray
}

type Roller interface {
	Roll(string) Roller
	DM(int) Roller
	Sum() int
}

func RollOnChart(chart DiceResultChart, r Roller) ([]string, error) {
	if !chart.verified {
		return []string{}, fmt.Errorf("chart was not verified")
	}

	return []string{}, nil
}

/*
code:
n
n+
n-
n-m

*/
