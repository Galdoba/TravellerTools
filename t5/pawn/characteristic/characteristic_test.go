package characteristic

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestChr(t *testing.T) {
	chr := New(Dexterity, 2)
	chr.RollValue(dice.New())
	fmt.Println(chr)
	fmt.Println(chr.ValueAs(Dexterity))
	fmt.Println(chr.ValueAs(Agility))
	fmt.Println(chr.ValueAs(Grace))
}
