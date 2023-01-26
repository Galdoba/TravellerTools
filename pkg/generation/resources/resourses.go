package resources

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

type Resources struct {
}

func testGen() {
	size := 5
	dice := dice.New()
	//rs := 0
	rm := make(map[int]int)
	for try := 0; try < 1000; try++ {
		rf := 0
		for i := 0; i < 12; i++ {
			r := dice.Sroll("3d6-3") - size
			if r <= 0 {
				rf++
			}
		}
		rm[rf]++
	}
	for l := 0; l < 20; l++ {
		if val, ok := rm[l]; ok == true {
			fmt.Println(l, ":", val)
		}
	}
	//fmt.Println(rm)
	// Resources := (Whexes*2) + 2d6/1000 / Gravity
	/*
	   F
	   EE
	   DDD
	   СССС
	   BBBBB
	*/

}
