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

func TestTime(t *testing.T) {
	for i := -5000000; i < 50000000; i = i + 59 {
		code := int64(i)
		t1, err := decodeTime(code)
		if err != nil {
			t.Errorf("internal Error: %v | %v - %v", err.Error(), i, t1.String())
			break
		}
		fmt.Print(t1.String()+"|", code, "\n")
		if testVals(code) != "" && testVals(code) != t1.String() {
			t.Errorf("Halt: expect '%v', but have '%v' - %v (%v)", testVals(code), t1.String(), t1, i)
			break
		}
		if t1.Second > 59 {
			t.Errorf("seconds can't be more than 59 | %v - %v", i, t1.String())
			break
		}
		if t1.Minute > 59 {
			t.Errorf("minutes can't be more than 59 | %v - %v", i, t1.String())
			break
		}
		if t1.Hour > 23 {
			t.Errorf("hours can't be more than 23 | %v - %v", i, t1.String())
			break
		}
		if t1.Day > 365 {
			t.Errorf("Days can't be more than 365 | %v - %v", i, t1.String())
			break
		}
		if t1.Second < 0 {
			t.Errorf("seconds can't be less than 0 | %v - %v", i, t1.String())
			break
		}
		if t1.Minute < 0 {
			t.Errorf("minutes can't be less than 0 | %v - %v", i, t1.String())
			break
		}
		if t1.Hour < 0 {
			t.Errorf("hours can't be less than 0 | %v - %v", i, t1.String())
			break
		}
		if t1.Day < 1 {
			t.Errorf("Days can't be less than 1 | %v - %v", i, t1.String())
			break
		}

	}

}