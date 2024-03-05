package world

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestNew(t *testing.T) {
	dice := dice.New()
	size := NewSize(dice)
	atmo := NewAtmosphere(dice, size.Value())
	hydr := NewHydrographics(dice, size.Value(), atmo.Value())
	fmt.Println(size, atmo, hydr)
}
