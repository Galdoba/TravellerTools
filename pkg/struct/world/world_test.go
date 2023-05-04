package world

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/classifications"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/struct/world/details/sizerelated"
)

func TestGenome(t *testing.T) {
	for i := 0; i < 200; i++ {

		wrld, err := NewWorld(
		//KnownData(IsMainworld, FLAG_TRUE),
		//KnownData(Primary, "G2 V"),
		)
		dice := dice.New().SetSeed(fmt.Sprintf("%v", i))
		err = wrld.GenerateBasic(dice)
		if err != nil {
			t.Errorf(err.Error())
		}
		// fmt.Println("===========")
		// fmt.Println(wrld)
		// fmt.Println(err)
		// fmt.Println("===========")
		// fmt.Println(wrld.profile)
		// fmt.Println(wrld.UWP())
		wrld.classifications = classifications.Evaluate(wrld)
		sDetails := sizerelated.New()
		err = sDetails.GenerateDetails(dice, wrld.Profile(), wrld.HomeStar)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Printf("testing world %v: %v	%v           \n", i, wrld, sDetails)
	}
	fmt.Println("")
}
