package interfaces

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestHex(t *testing.T) {
	dice := dice.New()
	for i := 0; i < 50; i++ {
		h := NewHex()
		fmt.Println("created  :", h)
		h.Populate(dice)
		fmt.Println("populated:", h.Population)
		if err := h.Populate(dice); err != nil {
			t.Errorf("expected: %v", err.Error())
		}
	}

}
