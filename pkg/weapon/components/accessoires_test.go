package components

import (
	"fmt"
	"testing"
)

func TestAccessoires(t *testing.T) {
	a, err := NewAccessoires(ACSR_AFD_MAGAZINE_STANDARD, ACSR_SCOPE_LASER_POINTER, ACSR_OTHER_BAYONET_LUG)
	if err != nil {
		t.Errorf(err.Error())
	}
	if a.ammoFeed == 0 {
		t.Errorf("ammo feed device not asignned")
	}
	if len(a.other) == 0 {
		t.Errorf("other devices not asignned")
	}

	fmt.Println(a)
}
