package pocketempire

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/calendar"
)

func TestGame(t *testing.T) {
	eventPl := eventPlanner{}
	crrnt := calendar.New(616)
	gm := Game{}
	eventPl.current = crrnt
	gm.CurrentDate = crrnt
	event1 := "Foo"
	event2 := "Bar"
	eventPl.plannedEvents = make(map[uint64]event)
	eventPl.plannedEvents[calendar.New(620).Global()] = event{event1}
	eventPl.plannedEvents[calendar.New(630).Global()] = event{event2}
	fmt.Printf("Start on %v\n", gm.CurrentDate.String())
	for !gm.CurrentDate.IsPast(calendar.New(636)) {
		fmt.Printf("Today is %v\n", gm.CurrentDate.String())
		gm.CurrentDate.Advance(calendar.NEXT_DAY)
		if event, ok := eventPl.plannedEvents[gm.CurrentDate.Global()]; ok == true {
			fmt.Println("fire event: ", event)
		}
	}
}
