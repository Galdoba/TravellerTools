package tradecodes

import (
	"fmt"
	"testing"
)

func input() []string {
	return []string{
		"X000000-0",
	}
}

func TestAnalize(t *testing.T) {
	for _, inp := range input() {
		feed := Input(KEY_uwp, inp)
		tc := Analize(feed)
		fmt.Println(tc)
	}
}
