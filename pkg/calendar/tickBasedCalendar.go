package calendar

import (
	"time"
)

type RWSCalendar struct {
	startTime      time.Time
	timeMultiplier int
	gameTick       int64
}

func NewSyncCalendar(y, m, d, mult int) *RWSCalendar {
	cal := RWSCalendar{}
	var tm time.Time
	cal.startTime = tm.AddDate(y, m, d)
	cal.timeMultiplier = mult
	cal.Sync()
	return &cal
}

func (sc *RWSCalendar) Sync() {
	sc.gameTick = int64(time.Since(sc.startTime).Seconds()) * int64(sc.timeMultiplier)
}

func (sc *RWSCalendar) GameTick() int64 {
	return sc.gameTick
}

func (sc *RWSCalendar) String() string {
	date := newDate(int(sc.gameTick / 86400))
	clok := newClock(int(sc.gameTick % 86400))
	return date.String() + " " + clok.String()
}

func ImperialTimeStamp(tick int64) string {
	date := newDate(int(tick / 86400))
	clok := newClock(int(tick % 86400))
	return date.String() + " " + clok.String()
}

// type imperialDate struct {
// 	date int64
// }

// func (c *RWSCalendar) Date() *date {
// 	//date := imperialDate{c.gameTick / 86400}
// 	date2 := newDate(c.gameTick / 86400)
// 	return date2
// }

// func (id *imperialDate) String() string {
// 	yr := id.date / 365
// 	dy := id.date % 365
// 	yrStr := formatUnits(int(yr), 3)
// 	dyStr := formatUnits(int(dy), 3)
// 	return fmt.Sprintf("%v-%v", dyStr, yrStr)
// }

// type clock struct
