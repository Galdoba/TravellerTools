package calendar

import (
	"fmt"
	"testing"
	"time"
)

func TestCalendar(t *testing.T) {
	dt := New()
	dt.Set(1, 1000)
	for dt.day < 1000000 {
		dt.Advance(NextWEEK)
		fmt.Printf("%v\r", dt.String())
		time.Sleep(time.Second)
	}
	fmt.Println("")
}
