package check

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/t5/pawn"
	"github.com/Galdoba/TravellerTools/t5/pawn/characteristic"
	"github.com/Galdoba/TravellerTools/t5/pawn/skill"
)

func TestCheck(t *testing.T) {
	dice := dice.New()
	chr := characteristic.NewChar(characteristic.CHAR_VIGOR)
	chr.SetValue(7)
	skl, _ := skill.New(skill.ID_Admin)
	skl.Learn()
	skl.Learn()
	skl.Learn()
	pawn, _ := pawn.New(dice, 0, []int{})
	fmt.Println(pawn.String())
	chk := NewCheck(Average, chr, skl).WithPrequisite(skill.NameByID(skill.ID_Admin) + " 2+")
	//chk := NewCheck(Average, chr, skl, check.Mod("Edu 10+", 2))
	res := Resolve(chk, dice)
	fmt.Println(res)
}
