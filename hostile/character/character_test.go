package character

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestGenerate(t *testing.T) {
	gen := NewGenerator()
	gen.dice = dice.New()
	chr, err := gen.Generate()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Println(chr)
}
