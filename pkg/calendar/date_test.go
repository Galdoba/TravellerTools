package calendar

import (
	"fmt"
	"testing"
)

func testVals(i int64) string {
	expected := make(map[int64]string)
	expected[-1] = "365-001 BI 23:59:59"
	expected[-2] = "365-001 BI 23:59:58"
	expected[-60] = "365-001 BI 23:59:00"
	expected[-62] = "365-001 BI 23:58:58"
	expected[-3600] = "365-001 BI 23:00:00"
	return expected[i]
}

func testInput() []int64 {
	return []int64{
		-366,
		-365,
		-364,
		-1,
		0,
		1,
		364,
		365,
		366,
		729,
		730,
		731,
	}
}

func TestCalendar(t *testing.T) {
	for i := 0; i < 1730; i++ {
		d2 := NewImperial("IC", 1105*365+i, i-15)
		fmt.Println(d2)
		//fmt.Println(d2.Hours(), d2.Minutes(), d2.Seconds(), d2.Day(), d2.Year())

		i = i + 63
	}

}
