package weapon

import (
	"fmt"
	"testing"
)

func TestBarrel(t *testing.T) {
	brl, err := newBarrel(brl_len_ASSAULT, brl_weight_HEAVY)
	if err != nil {
		t.Errorf("error: %v", err.Error())
	}
	fmt.Println("barrel struct:", brl)
	if brl.lenght == _UNDEFINED_ {
		t.Errorf("barrel lentgh is undefined")
	}
	if !isBarrelLentgh(brl.lenght) {
		t.Errorf("barrel lentgh value incorect (%v)", brl.lenght)
	}
}
