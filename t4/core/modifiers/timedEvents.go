package modifiers

import "github.com/Galdoba/TravellerTools/pkg/calendar"

type EventModifier struct {
	StartDate     calendar.Date
	EndDate       calendar.Date
	ReductionStep int
	Name          string
	Description   string
	MaxEffect     int
}

/*
Funcs:
NewEventModifier(name string, effect) EventModifier

EffectOf(em EventModifier, currentDate calendar.Date) int
EventModifier Methods:
SetNewEffect(newMaxEffect int)
*/

func NewEventModifier(name string, effect int) EventModifier {
	return EventModifier{
		Name:      name,
		MaxEffect: effect,
	}
}

func EffectOf(em EventModifier, currentDate calendar.Date) int {
	eff := em.MaxEffect
	if em.EndDate == nil {
		return 0
	}
	return eff
}
