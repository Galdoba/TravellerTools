package weapon

import (
	"fmt"
	"testing"

	. "github.com/Galdoba/TravellerTools/pkg/weapon/components"
)

func TestWeapon(t *testing.T) {
	input := []int{RCVR_TYPE_HANDGUN, CALLIBRE_HANDGUN_Medium, MECHANISM_REPEATER, AMMUNITION_CAPACITY_20_MORE, BRL_len_HANDGUN, FRNTR_STOCKLESS, TECH_CONVENTIONAL}
	w, err := New(input...)
	fmt.Println(w)
	if err != nil {
		t.Errorf("designer error: %v", err.Error())
	}
}
