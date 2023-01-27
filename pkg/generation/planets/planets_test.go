package planets

import (
	"fmt"
	"testing"
	"time"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestPlanet(t *testing.T) {
	try := 0
	lifemap := make(map[string]int)
	for {

		try++
		if try >= 10000 {
			break
		}
		p := New(dice.New())

		if err := p.GenerateBasic(); err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println("try", try, ":", p)
		lifemap[p.dominantLife.Code()]++
		time.Sleep(time.Microsecond * 18)

	}
	fmt.Println("")
	//TODO: Исправить генерацию звезд *RGG и подобных в stellar.GenerateOneStar()

}
