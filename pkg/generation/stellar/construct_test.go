package stellar

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/utils"
)

func TestStarConstruction(t *testing.T) {
	return
	smap, list := mapStars()
	l := len(smap)
	for i := 116; i < 125; i++ {
		stel, err := ConstructNew(CONSTRUCTOR_PARADIGM_T5, dice.New())
		if stel == nil {
			t.Errorf("func returned no object")
		}
		if err != nil {
			t.Errorf("func returned error: %v", err.Error())
		}
		done := 0
		for _, s := range stel.systemstars {
			if s.mass > 0 {
				smap[s.star]++
			}
		}

		//fmt.Println(list)
		for _, ls := range list {
			if smap[ls] == 0 {
				done++
			}
		}
		//fmt.Print("___")
		//time.Sleep(time.Second)

		fmt.Printf("%v/%v (%v)\r", done, l, i)
		//utils.ClearScreen()
		if done >= l {
			fmt.Println("TEST PASS")
			break
		}

		for pos := -1; pos < 11; pos++ {
			if str, ok := stel.systemstars[pos]; ok == true {
				fmt.Println(str)
			}
		}
		fmt.Println("=========")
		fmt.Println(stel)
		fmt.Println("---------")
		//t.Logf("==============")
		//t.Logf("%v", stel)
	}

}

func mapStars() (map[string]int, []string) {
	smap := make(map[string]int)
	list := []string{}
	for _, s := range newStarInputsCORRECT() {
		code := fixedStarCode(s.typ, s.dec, s.cls)
		smap[code] = 0
		list = utils.AppendUniqueStr(list, code)
	}
	return smap, list
}
