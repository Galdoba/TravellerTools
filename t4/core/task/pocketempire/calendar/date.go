package calendar

import "fmt"

const (
	NextDAY = iota
	NextWEEK
	NextMONTH
	NextYEAR
)

type Date struct {
	day int
}

func New() *Date {
	dt := &Date{}
	dt.Set(1, 0)
	return dt
}

func (d *Date) String() string {
	y := fmt.Sprintf("%v", d.Year())
	dy := fmt.Sprintf("%v", d.Day()+1)
	for len(y) < 3 {
		y = "0" + y
	}
	for len(dy) < 3 {
		dy = "0" + dy
	}
	return fmt.Sprintf("%v-%v", dy, y)
}

func (d *Date) Day() int {
	return d.day % 365
}

func (d *Date) Year() int {
	return d.day / 365
}

func (d *Date) Set(day, year int) *Date {
	d.day = (year * 365) + day
	return d
}

func (d *Date) Advance(milestone int) error {
	switch milestone {
	case NextDAY:
		d.day++
	case NextWEEK:
		for _, check := range weekStarts() {
			if check > d.Day() {
				d.Set(check, d.Year())
				return nil
			}
		}
		d.Advance(NextYEAR)
	case NextMONTH:
		for _, check := range []int{1, 30, 58, 93, 121, 149, 184, 212, 240, 275, 303, 330} {
			if check > d.Day() {
				d.Set(check, d.Year())
				return nil
			}
		}
		d.Advance(NextYEAR)
	case NextYEAR:
		d.Set(1, d.Year()+1)
	}
	return nil
}

func weekStarts() []int {
	ws := append([]int{}, 1)
	for i := 1; i <= 52; i++ {
		ws = append(ws, (i*7)+1)
	}
	return ws
}

func Match(current, target Date) bool {
	return current.day == target.day
}

func Passed(current, target Date) bool {
	return current.day > target.day
}

func After(current, past Date) Date {
	days := current.day - past.day
	d := days + 1%365
	y := days / 365
	p := New().Set(d, y)
	return *p
}

func Add(dt Date, days, years int) Date {
	dt.day = dt.day + days + (365 * years)
	return dt
}

/*
12345678
 9012345
 6789012

*/
