package stellar

import (
	"fmt"
	"testing"

	"github.com/Galdoba/utils"
)

func TestStarConstruction(t *testing.T) {
	smap, list := mapStars()
	l := len(smap)
	for i := 20; i < 46; i++ {
		df1 := Instruction("Hex", fmt.Sprintf("%v", i))
		df2 := Instruction(KEY_MW, "Reginaasd")
		stel, err := ConstructNew(CONSTRUCTOR_PARADIGM_T5, df1, df2)
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
		fmt.Println(stel)
		for pos := -1; pos < 11; pos++ {
			if str, ok := stel.systemstars[pos]; ok == true {
				fmt.Println(str)
			}
		}
		fmt.Println("=========")
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
