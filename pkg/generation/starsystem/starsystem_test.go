package starsystem

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestSSD(t *testing.T) {
	for i := 0; i < 200; i++ {
		dice := dice.New().SetSeed(i)
		sts := New(dice)
		fmt.Println("======")
		fmt.Println(sts)
		fmt.Println("======")
		for i := -10; i < 95000; i++ {
			if bod, ok := sts.reservedFor[i]; ok {
				if bod == nil {
					continue
				}
				fmt.Println(i, bod)
				//fmt.Println("", sts.starPlanetSys.ByCode(i).AU(), "au")
			}
		}
	}
}
