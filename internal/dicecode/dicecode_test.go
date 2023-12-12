package dicecode

import (
	"fmt"
	"testing"
)

func TestDetectValue(t *testing.T) {
	codes := []string{
		"2-",
		"3",
		"4",
		"5...8",
		"9 10 12",
		"11",
		"13+",
	}
	for i := -5; i < 15; i++ {
		code, err := DetectMatch(i, codes)
		fmt.Println("roll", i, "==>", code)
		if err != nil {
			t.Errorf("r:%v err:%v", i, err.Error())
		}
	}
}
