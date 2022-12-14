package mainworld

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestMWGeneration(t *testing.T) {
	for i := 0; i < 30; i++ {
		dice := dice.New().SetSeed(i)
		mw := New(dice)
		fmt.Println(mw)
		fmt.Println("--------------------")
		fmt.Println("Phase 1:")
		mw.DetermineSystemData()
		fmt.Println(mw)
		fmt.Println("--------------------")
		fmt.Println("Phase 2:")
		mw.DeterminePhysicalCharacteristics()
		fmt.Println(mw, "|||", mw.ProfileUncharted())
		fmt.Println("--------------------")
		fmt.Println("Phase 3:")
		mw.DetermineSocialCharacteristics()
		fmt.Println(mw)
		fmt.Println("--------------------")
		fmt.Println("Phase 4:")
		mw.DetermineAdditionalCharacteristics()
		fmt.Println(mw)
		fmt.Println("--------------------")
		fmt.Println("  ")
		fmt.Println("  ")
	}
	fmt.Println("TODO: добавить доп коды на вывод")
}
