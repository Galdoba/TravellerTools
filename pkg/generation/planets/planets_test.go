package planets

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/generation/stellar"
)

func TestPlanet(t *testing.T) {
	try := 1
	nill := 0
	resMap := make(map[int]int)
	for {
		try++
		if try >= 1000 {
			break
		}
		dice2 := dice.New().SetSeed(try)
		star := stellar.GenerateStellarOneStar(dice2)
		hz := dice2.Flux()
		p := PhysicalData_T5(dice2, hz, star)
		//fmt.Println("try", try, ":", p.size.Code(), p.atmo.Code(), p.hydr.Code(), p.hz, p.comment)
		if p.BaseResources() > -1 && p.life.Value() == 10 {
			fmt.Printf("%v %v              \n", try, p.String())
			resMap[p.BaseResources()]++
		}

	}
	fmt.Println("")
	fmt.Println(nill)
	//TODO: Исправить генерацию звезд *RGG и подобных в stellar.GenerateOneStar()
	for i := 0; i < 99; i++ {
		if val, ok := resMap[i]; ok == true {
			fmt.Println(i, val)
		}
	}
}
