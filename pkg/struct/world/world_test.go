package world

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/struct/world/details/sizerelated"
)

func TestGenome(t *testing.T) {

	wrld, err := NewWorld(Inject(
		KnownData(IsMainworld, FLAG_TRUE),
		KnownData(Primary, "G2 V"),
	))
	fmt.Println("===========")
	fmt.Println(wrld)
	fmt.Println(err)
	fmt.Println("===========")
	fmt.Println(wrld.profile)
	fmt.Println(wrld.UWP())
	sDetails := sizerelated.New()
	sDetails.GenerateDetails(dice.New(), wrld.profile)
	fmt.Println(sDetails)
}
