package charset

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/t5/genetics"
	"github.com/Galdoba/TravellerTools/t5/pawn/characteristic"
)

func TestUPP(t *testing.T) {
	pass := 0
	fail := 0
	total := 0
	dice := dice.New()
	total++
	t.Logf("test %d\n", total)
	gd, err := genetics.GeneTemplateManual("SDEIES", "222222", "")
	if err != nil {
		t.Errorf("GeneDataManual error: %v", err.Error())
		fail++
	}
	up := NewCharSet(dice, gd.Profile(), gd.Variations())
	t.Logf("up: %v\n", up.String())
	pass++
	t.Logf("Result: %v/%v/%v (PASS/FAIL/TOTAL)\n", pass, fail, total)

	human := NewCharSet(dice, genetics.GeneTemplateHuman().Profile(), genetics.GeneTemplateHuman().Variations())
	fmt.Println(human)
}

func TestHumanSet(t *testing.T) {
	dice := dice.New()
	for i := 0; i < 20; i++ {
		human := NewCharSet(dice, genetics.GeneTemplateHuman().Profile(), genetics.GeneTemplateHuman().Variations())
		fmt.Println(human)
		fmt.Println(human.ValueOf(characteristic.Agility))
	}
}
