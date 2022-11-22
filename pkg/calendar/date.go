package calendar

import "fmt"

const (
	Holiday = iota
	Wonday
	Tuday
	Trirday
	Forday
	Fiday
	Sixday
	Senday
	Minute  = 60
	Hour    = 3600
	Day     = 86400
	Year    = 31536000
	bigBang = 1000000000000000
)

/*

int64  : -9223372036854775808 to 922 3372036 854 77 58 07
*/
type Time struct {
	val uint64
}

func Encode(code uint64) *Time {
	return &Time{code}
}

func (t *Time) Second() uint64 {
	return t.val % 60
}
func (t *Time) Minute() uint64 {
	return (t.val / Minute) % 60
}
func (t *Time) Hour() uint64 {
	return (t.val / Hour) % 24
}
func (t *Time) Day() uint64 {
	return (t.val / Day) % 365
}
func (t *Time) Year() uint64 {
	return t.val / Year
}

func (t *Time) String() string {
	return fmt.Sprintf("%v-%v %v:%v:%v", formatInt(t.Day(), 3), formatInt(t.Year(), 3), formatInt(t.Hour(), 2), formatInt(t.Minute(), 2), formatInt(t.Second(), 2))
}

func formatInt(i uint64, power int) string {
	s := fmt.Sprintf("%v", i)
	for len(s) < power {
		s = "0" + s
	}
	return s
}
