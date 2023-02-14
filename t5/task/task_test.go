package task

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/calendar"
	"github.com/Galdoba/TravellerTools/t5/pawn/characteristic"
)

func instructionSets() map[int][]instruction {
	insMap := make(map[int][]instruction)
	insMap[1] = []instruction{
		Phrase("convince a buyer that goods are acceptable"),
		Duration(Ignored, calendar.Day),
		Difficulty(3),
		Char(characteristic.SocialStanding),
		Skill("Broker", USAGE_SkillOnly),
		Modifier("Quality - 5", 1, true),
	}

	return insMap
}

func TestTask(t *testing.T) {
	instructionSet := instructionSets()

	for i := 1; i <= len(instructionSet); i++ {
		tsk, err := New(instructionSet[i]...)
		t.Logf("Task %v", i)
		if tsk == nil {
			t.Errorf("func returned no object")
		} else {

			t.Logf("%v", tsk)
			t.Logf("\n%v", tsk.toString())
			fmt.Println(tsk.char)
		}
		if err != nil {
			t.Errorf("func returned error: %v", err.Error())
		}
		if strings.Contains(tsk.toString(), "error") {
			t.Errorf("object has error")
		}

		t.Logf("==================")
	}

}

func TestCharacteristic(t *testing.T) {
	chr := characteristic.New(characteristic.Agility, 3)
	chr.SetValue(12)
	res := CheckCharacteristicAverage(chr)
	fmt.Println(res)
}
