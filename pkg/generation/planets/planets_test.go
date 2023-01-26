package planets

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestPlanet(t *testing.T) {
	p := New(dice.New())

	if err := p.GenerateBasic(); err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(p)
	//TODO: Исправить генерацию звезд *RGG и подобных в stellar.GenerateOneStar()

}
