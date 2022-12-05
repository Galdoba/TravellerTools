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
type date struct {
	val  uint64
	year int
	day  int
}

type Date interface {
	String() string
	AsInts() (int, int)
	IsPast(Date) bool
	Advance(string)
	Day() int
	Year() int
	Global() uint64
}

func (d *date) Global() uint64 {
	return d.val
}

func (d *date) Day() int {
	return d.day
}

func (d *date) Year() int {
	return d.year
}

func (d *date) AsInts() (int, int) {
	return d.day, d.year
}

func (d *date) IsPast(check Date) bool {
	valD := (d.year * 365) + d.day
	valCheck := (check.Year() * 365) + check.Day()
	if valCheck < valD {
		return true
	}
	return false
}

func After(d Date, period int) Date {
	newDate := d
	for i := 0; i < period; i++ {
		newDate.Advance(NEXT_DAY)
	}
	return newDate
}

func TimeBetween(d1, d2 *date) (int, int) {
	valD1 := (d1.year * 365) + d1.day
	valD2 := (d2.year * 365) + d2.day
	diff := valD1 - valD2
	if diff < 0 {
		diff = diff * -1
	}
	dif := New(uint64(diff))
	return dif.day, dif.year
}

func IsEqual(d1, d2 *date) bool {
	return d1.day == d2.day && d1.year == d2.year
}

func New(code uint64) *date {
	t := date{
		val: code,
	}
	t.evaluate()
	return &t
}

func SetDate(day, year int) *date {
	d := date{
		val:  uint64(day + year*365),
		day:  day,
		year: year,
	}
	d.evaluate()
	return &d
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (t *date) evaluate() {
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

func (t *date) String() string {
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

func (t *date) Advance(code string) {
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
