package weapon

import (
	"fmt"
	"testing"
)

func TestAccessoires(t *testing.T) {

	a, err := newAccessoires(accessoire_SUPPRESSOR_ABSENT, accessoire_AFD_MAGAZINE_STANDARD, accessoire_SCOPE_ABSENT)
	if err != nil {
		t.Errorf(err.Error())
	}
	if a.suppressor == 0 {
		t.Errorf("supperssor not asignned")
	}
	if a.ammoFeed == 0 {
		t.Errorf("ammo feed device not asignned")
	}
	if a.sighting == 0 {
		t.Errorf("sighting device not asignned")
	}
	if len(a.other) == 0 {
		t.Errorf("other devices not asignned")
	}
	fmt.Println("")
}
