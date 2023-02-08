package cascad

import (
	"fmt"
	"testing"
	"time"

	"github.com/Galdoba/TravellerTools/pkg/calendar"
	"github.com/Galdoba/utils"
)

func TestCascad(t *testing.T) {
	pawn := NewPawn(25, 25, 0)
	cal := calendar.NewSyncCalendar(2020, 0, 0, 360)

	//	fmt.Println(nowSeed())
	//seedTime := 60*tm.Minute() + tm.Second()
	for {
		utils.ClearScreen()
		fmt.Println(cal)
		cal.Sync()
		seed := cal.GameTick()

		fmt.Printf("SEED: %v \n", seed)
		pawn.DoSomething(seed)
		fmt.Printf("= %v           \n", pawn)
		time.Sleep(time.Second)
		if pawn.Money > 20 {
			break
		}
	}
}

func nowSeed() int64 {
	// var tm time.Time
	// tm = tm.AddDate(2022, 1, 6)
	// fmt.Println(tm.Unix()) // 1970-01-01 00:00:00 +0000 UTC
	// seed := time.Since(tm)
	// return int64(seed.Seconds())
	cal := calendar.NewSyncCalendar(2000, 0, 0, 24)
	fmt.Println(cal.String())
	return cal.GameTick()
}
