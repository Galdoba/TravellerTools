package world

import (
	"fmt"
	"testing"
)

func TestGenome(t *testing.T) {

	wrld, err := NewWorld(Inject(
	//KnownData(IsMainworld, FLAG_TRUE),
	//KnownData(Primary, "G2 V"),
	))
	fmt.Println("===========")
	fmt.Println(wrld, err)
	fmt.Println("===========")
	fmt.Println(wrld.profile)
	fmt.Println(wrld.UWP())

}
