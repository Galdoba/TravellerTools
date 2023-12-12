package star2

import (
	"fmt"
	"testing"
)

func TestStar(t *testing.T) {
	st, err := FromProfile("Ab-0.09-0.11-G8 V-0.907-0.957-0.681")
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	fmt.Println(st)
}
