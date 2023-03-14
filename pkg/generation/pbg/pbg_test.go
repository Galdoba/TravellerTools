package pbg

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func Test_PBG(t *testing.T) {
	dice := dice.New().SetSeed(1)
	p := New(dice)
	fmt.Println(p.String())
}
