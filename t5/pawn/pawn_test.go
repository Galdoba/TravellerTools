package pawn

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/t5/genetics"
)

func TestPawn(t *testing.T) {
	gt := genetics.NewTemplate("SDEIES", "222222")
	chr, err := New(control_Random, gt, []string{"Ph", "Pa", "Ri"})
	fmt.Println(chr)
	fmt.Println(err)
}
