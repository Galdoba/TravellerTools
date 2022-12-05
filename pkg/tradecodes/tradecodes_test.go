package tradecodes

import (
	"fmt"
	"strings"
	"testing"
)

func input() [][]tcInput {
	return [][]tcInput{
		[]tcInput{Input(KEY_uwp, "A511965-E")},
	}
}

func TestAnalize(t *testing.T) {
	for _, inp := range input() {
		//feed := Input(KEY_uwp, inp)
		tc := Analize(inp...)
		fmt.Println(inp)
		if inp[0].val == "A511965-E" {
			expect := "Ht Ic In Na"
			if expect != strings.Join(tc.confirmedByUWP, " ") {
				t.Errorf("result : %v != %v", tc.confirmedByUWP, expect)
			}
		}
	}
}
