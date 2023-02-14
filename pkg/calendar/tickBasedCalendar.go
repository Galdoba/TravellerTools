package calendar

import (
	"fmt"
	"strings"
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
var Tick = int64(1)
var Minute = int64(10 * Tick)
var Hour = int64(60 * Minute)
var Day = int64(24 * Hour)
var Week = int64(7 * Day)
var Month = int64(30 * Day)
var Year = int64(365 * Day)
var Decade = int64(10 * Year)
var Century = int64(100 * Year)
var Millenium = int64(1000 * Year)

func TicksToText(ticks int64) string {
	text := ""
	if ticks < 0 {
		return ""
	}
	if ticks == 0 {
		return "Instant"
	}
	if ticks >= Year {
		x := ticks / Year
		text += fmt.Sprintf(" %v Year", x)
		if (x-1)%10 != 0 {
			text += "s"
		}
		ticks = ticks - (x * Year)
	}
	if ticks >= Day {
		x := ticks / Day
		text += fmt.Sprintf(" %v Day", x)
		if (x-1)%10 != 0 {
			text += "s"
		}
		ticks = ticks - (x * Day)
	}
	if ticks >= Hour {
		x := ticks / Hour
		text += fmt.Sprintf(" %v Hour", x)
		if (x-1)%10 != 0 {
			text += "s"
		}
		ticks = ticks - (x * Hour)
	}
	if ticks >= Minute {
		x := ticks / Minute
		text += fmt.Sprintf(" %v Minute", x)
		if (x-1)%10 != 0 {
			text += "s"
		}
		ticks = ticks - (x * Minute)
	}
	if ticks >= Tick {
		x := ticks / Tick
		text += fmt.Sprintf(" %v Second", x*6)
		if ((x*6)-1)%10 != 0 {
			text += "s"
		}
		ticks = ticks - (x * Tick)
	}

	return strings.TrimSpace(text)
}
