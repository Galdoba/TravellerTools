package calendar

import (
	"fmt"
	"strconv"
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
	Minute = 60
	Hour   = 3600
	Day    = 86400
	Year   = 31536000
)

type Time struct {
	Second int
	Minute int
	Hour   int
	Day    int
	Year   int
	era    string
	//WeekDay int - By demand
	//timeInt int64
}

func (t *Time) String() string {
	yS := strconv.Itoa(t.Year)
	for len(yS) < 3 {
		yS = "0" + yS
	}
	dS := strconv.Itoa(t.Day)
	for len(dS) < 3 {
		dS = "0" + dS
	}
	hS := strconv.Itoa(t.Hour)
	for len(hS) < 2 {
		hS = "0" + hS
	}
	mS := strconv.Itoa(t.Minute)
	for len(mS) < 2 {
		mS = "0" + mS
	}
	sS := strconv.Itoa(t.Second)
	for len(sS) < 2 {
		sS = "0" + sS
	}
	return fmt.Sprintf("%v-%v %v %v:%v:%v", dS, yS, t.era, hS, mS, sS)
}

func decodeTime(t int64) (Time, error) {
	switch {
	case t > 0:
		return decodeTimeIC(t)
	case t < 0:
		return decodeTimeBI(t)
	case t == 0:
		return Time{0, 0, 0, 1, 0, "IC"}, nil
	}
	return Time{}, fmt.Errorf("unexpected decoding error")
}

func decodeTimeIC(t int64) (Time, error) {
	if t < 0 {
		return Time{}, fmt.Errorf("timecode negative")
	}
	era := "IC"
	s := int(t % Minute)
	m := int(t/Minute) % 60
	h := int(t/Hour) % 24
	d := (int(t/Day) % 365) + 1
	y := int(t / Year)
	for d > 365 {
		d = d - 365
		y++
	}
	return Time{s, int(m), int(h), d, y, era}, nil
}

func decodeTimeBI(t int64) (Time, error) {
	if t >= 0 {
		return Time{}, fmt.Errorf("timecode positive")
	}
	//t = t * -1
	era := "BI"
	nt := t + Year
	yr := 1
	for nt < 0 {
		nt = nt + Year
		yr++
	}
	t2, err := decodeTimeIC(nt)
	if err != nil {
		return Time{}, fmt.Errorf("decode BI: %v", err.Error())
	}
	return Time{t2.Second, t2.Minute, t2.Hour, t2.Day, yr, era}, nil
}

func FromInt64(i int64) (*Time, error) {
	return &Time{}, fmt.Errorf("Not Implemented")
}

/*
-2 = 365-001 BI 23:59:58
-1 = 365-001 BI 23:59:59
 0 = 001-000 IC 00:00:00
 1 = 001-000 IC 00:00:01
 2 = 001-000 IC 00:00:02


time as Int64:
   yyyyyyy|ddd|hh|mm|ss
   1234567|890|12|34|56
         2 147 48 36 47
9223372036 854 77 58 07
         0 000 00 00 00
*/
