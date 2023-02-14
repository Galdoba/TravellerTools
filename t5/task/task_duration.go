package task

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/calendar"
)

const (
	Ignored = iota
	Absolute
	Variable
	Randomized
)

type taskDurationData struct {
	durationType int
	units        int64
}

var About10Minutes = int64(-21)
var AboutAnHour = int64(-22)
var AllDay = int64(-23)
var AboutAWeek = int64(-24)
var AboutAMonth = int64(-25)
var AboutAYear = int64(-26)

func (tdd *taskDurationData) describe() string {
	s := "["
	switch tdd.durationType {
	default:
		s += fmt.Sprintf("error: unknown time duration type %v", tdd.durationType)
	case Ignored:
		return ""
	case Absolute:
		s += fmt.Sprintf("Absolute: %v", calendar.TicksToText(tdd.units))
	case Variable:
		switch tdd.units {
		default:
			s += fmt.Sprintf("error: unknown time unit %v", tdd.units)
		case About10Minutes:
			s += "About 10 Minutes"
		case AboutAnHour:
			s += "About An Hour"
		case AllDay:
			s += "All Day"
		case AboutAWeek:
			s += "About A Week"
		case AboutAMonth:
			s += "About A Month"
		case AboutAYear:
			s += "About A Year"
		}
	case Randomized:
		s += "Randomized: more or less than " + calendar.TicksToText(tdd.units)

	}
	return s + "]"
}
