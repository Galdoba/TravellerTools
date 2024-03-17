package characteristic

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestCharacteristic(t *testing.T) {
	dice := dice.New()
	chrSet, err := NewCharSet(Human()...)
	if err != nil {
		t.Errorf(err.Error())

	}
	chrSet.Roll(dice)
	fmt.Println(chrSet.String())
}
