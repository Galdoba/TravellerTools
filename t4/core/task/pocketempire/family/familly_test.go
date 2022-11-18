package family

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/t4/core/task/pocketempire/calendar"
)

func TestFamily(t *testing.T) {
	cal := calendar.New()
	fm := New("Calarita")
	fmt.Println(fm)
	for i := 0; i < len(fm.Membrs); i++ {
		fmt.Println(fm.Membrs[i])
	}
	for len(fm.Membrs) < 10 {
		cal.Advance(calendar.NextDAY)
		fmt.Printf("%v ", cal.String())
		fm.Grow(*cal)
		fmt.Printf("%v \n", len(fm.Membrs))
	}

}
