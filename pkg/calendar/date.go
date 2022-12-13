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

type cal struct {
	era   string
	date  *date
	clock *clock
}

type Calendar interface {
	Clock
	Date
}

func NewImperial(name string, day, second int) *cal {
	c := cal{}
	c.era = name
	c.date = newDate(day)
	c.clock = newClock(second)
	return &c
}

func (c *cal) Hours() int {
	return c.clock.Hours()
}

func (c *cal) Minutes() int {
	return c.clock.Minutes()
}

func (c *cal) Seconds() int {
	return c.clock.Seconds()
}

func (c *cal) Day() int {
	return c.date.Day()
}

func (c *cal) Year() int {
	return c.date.Year()
}

func (c *cal) String() string {
	return fmt.Sprintf("%v %v %v", c.date.String(), c.era, c.clock.String())
}

type Clock interface {
	String() string
	Seconds() int
	Minutes() int
	Hours() int
}

type Date interface {
	String() string
	Day() int
	Year() int
}

type clock struct {
	val int
}

func newClock(val ...int) *clock {
	d := clock{}
	for _, v := range val {
		d.val += v
	}
	if d.val < 0 {
		d.val = d.val * -1
	}
	d.val = d.val % 86400
	return &d
}

func timeSegment(d *clock) int {
	return int(d.val % 86400)
}

func (d *clock) Hours() int {
	hh := timeSegment(d) / 3600 % 24
	return hh
}

func (d *clock) Minutes() int {
	mm := (timeSegment(d) / 60) % 60
	return mm
}

func (d *clock) Seconds() int {
	ss := timeSegment(d) % 60
	return ss
}

func (d *clock) Val() int {
	return d.val
}

func (cl *clock) String() string {
	return fmt.Sprintf("%v:%v:%v", formatUnits(cl.Hours(), 2), formatUnits(cl.Minutes(), 2), formatUnits(cl.Seconds(), 2))
}

/*

int64  : -9223372036854775808 to 922 3372036 854 77 58 07
*/
type date struct {
	day int
}

func newDate(d int) *date {
	if d < 0 {
		d = d * -1
	}
	dt := date{day: d}
	return &dt
}

func (d *date) Day() int {
	return d.day
}

func (d *date) Year() int {
	return d.day / 365
}

func (t *date) String() string {
	return fmt.Sprintf("%v-%v", formatUnits((t.day%365)+1, 3), formatUnits(t.day/365, 3))

}

func formatUnits(i int, unit int) string {
	s := ""
	s = fmt.Sprintf("%v", i)
	for len(s) < unit {
		s = "0" + s
	}
	return s
}

// func (t *date) Advance(code string) {
// 	switch code {
// 	case NEXT_DAY:
// 		t.val++
// 		t.evaluate()
// 	case NEXT_WEEK:
// 		for {
// 			t.val++
// 			t.evaluate()
// 			if isWeekStart(t.day) {
// 				return
// 			}
// 		}
// 	case NEXT_MONTH:
// 		for {
// 			t.val++
// 			t.evaluate()
// 			if isMonthStart(t.day) {
// 				return
// 			}
// 		}
// 	case NEXT_YEAR:
// 		for {
// 			t.val++
// 			t.evaluate()
// 			if t.day == 1 {
// 				return
// 			}
// 		}
// 	}
// }

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
