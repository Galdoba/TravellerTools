package pawn

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/t5/genetics"
)

func TestPawn(t *testing.T) {
	gt := genetics.NewTemplate("SDEIES", "832222")
	chr, err := New(gt)
	fmt.Println(chr)
	fmt.Println(err)
	fmt.Println(chr.chrSet.String())
}
