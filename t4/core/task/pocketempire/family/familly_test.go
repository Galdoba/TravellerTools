package family

import (
	"fmt"
	"testing"
	"time"

	"github.com/Galdoba/TravellerTools/pkg/calendar"
)

func TestFamily(t *testing.T) {
	date := calendar.SetDate(24, 58)
	fm := Create("Calarita", date)
	fmt.Println(fm)
	// fm := New("Calarita")
	// fmt.Println(fm)
	for i := 0; i < len(fm.Membrs); i++ {
		fmt.Println(fm.Membrs[i].String())
	}
	fm.AddMember(fm.Membrs[0], date, spouse)
	for len(fm.Membrs) < 5 {
		date.Advance(calendar.NEXT_WEEK)
		fmt.Printf("Date: %v \r", date)
		time.Sleep(time.Millisecond * 250)
		fm.AddMember(fm.Membrs[0], date, child)
		//fm.Grow(*cal)
		//break
		//fmt.Printf("%v \n", len(fm.Membrs))
	}
	for i := 0; i < len(fm.Membrs); i++ {
		fmt.Println(fm.Membrs[i].GeneticsCard())
	}

}
