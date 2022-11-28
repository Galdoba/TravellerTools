package calendar

import (
	"fmt"
)

const (
	Holiday = iota
	Wonday
	Tuday
	Trirday
	Forday
	Fiday
	Sixday
	Senday
	NEXT_DAY     = "NEXT_DAY"
	NEXT_WEEK    = "NEXT_WEEK"
	NEXT_MONTH   = "NEXT_MONTH"
	NEXT_YEAR    = "NEXT_YEAR"
	PERIOD_NONE  = 0
	PERIOD_DAY   = 1
	PERIOD_WEEK  = 7
	PERIOD_MONTH = 30
	PERIOD_YEAR  = 365
)

/*

int64  : -9223372036854775808 to 922 3372036 854 77 58 07
*/
type Date struct {
	val  uint64
	year int
	day  int
}

func New(code uint64) Date {
	t := Date{
		val: code,
	}
	t.evaluate()
	return t
}

func SetDate(day, year int) Date {
	return Date{
		val:  uint64(day + year*365),
		day:  day,
		year: year,
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (t *Date) evaluate() {
	if t.val == 0 {
		t.val = 1
	}
	t.day = 0
	t.year = 0
	val := t.val
	for val > 3650000 {
		t.year += 10000
		val -= 3650000
	}
	for val > 365000 {
		t.year += 1000
		val -= 365000
	}
	for val > 36500 {
		t.year += 100
		val -= 36500
	}
	for val > 3650 {
		t.year += 10
		val -= 3650
	}
	for val > 365 {
		t.year += 1
		val -= 365
	}
	t.day = int(val)

}

func (t *Date) String() string {
	return fmt.Sprintf("%v-%v", formatUnits(t.day, 3), formatUnits(t.year, 3))

}

func formatUnits(i int, unit int) string {
	s := ""
	s = fmt.Sprintf("%v", i)
	for len(s) < unit {
		s = "0" + s
	}
	return s
}

func (t *Date) Next(code string) {
	switch code {
	case NEXT_DAY:
		t.val++
		t.evaluate()
	case NEXT_WEEK:
		for {
			t.val++
			t.evaluate()
			if isWeekStart(t.day) {
				return
			}
		}
	case NEXT_MONTH:
		for {
			t.val++
			t.evaluate()
			if isMonthStart(t.day) {
				return
			}
		}
	case NEXT_YEAR:
		for {
			t.val++
			t.evaluate()
			if t.day == 1 {
				return
			}
		}
	}
}

func isWeekStart(i int) bool {
	if i == 1 {
		return true
	}
	for w := 1; w < 53; w++ {
		if 2+w*7 == i {
			return true
		}
	}
	return false
}

func isMonthStart(i int) bool {
	for _, w := range []int{1, 30, 58, 93, 121, 149, 184, 212, 240, 275, 303, 330} {
		if w == i {
			return true
		}
	}
	return false
}
