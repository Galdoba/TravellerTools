package weapon

import (
	"fmt"
	"testing"

	. "github.com/Galdoba/TravellerTools/pkg/weapon/components"
)

func TestWeapon(t *testing.T) {
	input := []int{RCVR_TYPE_LIGHT_SUPPORT_WEAPON, CALLIBRE_LONGARM_Rifle_AntiMaterialHeavy, MECHANISM_SEMI_AUTOMATIC, BRL_len_HANDGUN, FRNTR_STOCKLESS, AMMUNITION_CAPACITY_50PCT_LESS}
	expected := WpnSheet{
		penetration:    -1,
		tl:             6,
		effectiveRange: 10,
		damageDice:     3,
		damageMod:      -3,
		weight:         0.75,
		cost:           121,
		magazine:       6,
		magazineCost:   5,
		quickdraw:      4,
		traits:         []string{},
	}
	w, err := New(input...)
	fmt.Println(w)
	if w == nil {
		t.Errorf("Weapon was not created!")

	}
	if err != nil {
		t.Errorf("designer error: %v", err.Error())
	}
	if w == nil {
		t.FailNow()
	}
	if w.penetration != expected.penetration {
		t.Errorf(" have penetration = %v, but expect %v", w.penetration, expected.penetration)
	}
	if w.tl != expected.tl {
		t.Errorf(" have tl = %v, but expect %v", w.tl, expected.tl)
	}
	if w.effectiveRange != expected.effectiveRange {
		t.Errorf(" have effectiveRange = %v, but expect %v", w.effectiveRange, expected.effectiveRange)
	}
	if w.damageDice != expected.damageDice {
		t.Errorf(" have damageDice = %v, but expect %v", w.damageDice, expected.damageDice)
	}
	if w.damageMod != expected.damageMod {
		t.Errorf(" have damageMod = %v, but expect %v", w.damageMod, expected.damageMod)
	}
	if w.weight != expected.weight {
		t.Errorf(" have weight = %v, but expect %v", w.weight, expected.weight)
	}
	if w.cost != expected.cost {
		t.Errorf(" have cost = %v, but expect %v", w.cost, expected.cost)
	}
	if w.magazine != expected.magazine {
		t.Errorf(" have magazine = %v, but expect %v", w.magazine, expected.magazine)
	}
	if w.magazineCost != expected.magazineCost {
		t.Errorf(" have magazineCost = %v, but expect %v", w.magazineCost, expected.magazineCost)
	}
	if w.quickdraw != expected.quickdraw {
		t.Errorf(" have quickdraw = %v, but expect %v", w.quickdraw, expected.quickdraw)
	}

	if len(w.traits) != len(expected.traits) {
		t.Errorf(" have traits = %v, but expect %v", w.traits, expected.traits)
	}

	fmt.Println("------------------")
	fmt.Println(w.Summary())
	fmt.Println("------------------")
}
