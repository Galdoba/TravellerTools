package task

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/calendar"
	"github.com/Galdoba/TravellerTools/pkg/dice"
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
			fmt.Println(tsk.charCode)
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
	tsk, err := New(
		Phrase("test check"),
		Char(chr.Name()),
		Difficulty(chr.Genes()),
	)
	fmt.Println(err)
	tsk.Resolve(dice.New(), TargetNumber(chr))
	fmt.Println(tsk.toString())
	fmt.Println(tsk.rr.String(), tsk.rr)
}

func TestResults(t *testing.T) {
	res := []int{}
	suc := 0
	fail := 0
	interest := 0
	total := 0
	for _, d1 := range []int{1, 2, 3, 4, 5, 6} {
		for _, d2 := range []int{1, 2, 3, 4, 5, 6} {
			res = append(res, d1, d2)
			res = fixResult(res)
			s := false
			f := false
			if hasSSuc(res) {
				suc++
				s = true
			}
			if hasSFail(res) {
				fail++
				f = true
			}
			if s && f {
				interest++
				fmt.Printf("Rolls: %v | Result: %v | suc/fail/interest = %v/%v/%v         \r", total, res, suc, fail, interest)
			}
			total++

			res = nil
		}
	}
	fmt.Printf("Rolls: %v | Result: %v | suc/fail/interest = %v/%v/%v         \r", total, res, suc, fail, interest)
	fmt.Println("")
	//Rolls:      216 | Result: [] | suc/fail/interest =        1       /1      /0
	//Rolls:     1296 | Result: [] | suc/fail/interest =       21      /21      /0
	//Rolls:     7776 | Result: [] | suc/fail/interest =      276     /276      /0
	//Rolls:    46656 | Result: [] | suc/fail/interest =     2906    /2906     /20
	//Rolls:   279936 | Result: [] | suc/fail/interest =    26811   /26811    /630
	//Rolls:  1679616 | Result: [] | suc/fail/interest =   226491  /226491  /11382
	//Rolls: 10077696 | Result: [] | suc/fail/interest =  1796446 /1796446 /154812
	//Rolls: 60466176 | Result: [] | suc/fail/interest = 13591176/13591176/1761552

}

func fixResult(sl []int) []int {
	rf := []int{}
	for _, r := range sl {
		if r != 0 {
			rf = append(rf, r)
		}
	}
	return rf
}

func hasSSuc(sl []int) bool {
	if len(sl) < 3 {
		return false
	}
	ones := 0
	for _, d := range sl {
		if d == 1 {
			ones++
		}
		if ones > 2 {
			return true
		}
	}
	return false
}

func hasSFail(sl []int) bool {
	if len(sl) < 3 {
		return false
	}
	sixes := 0
	for _, d := range sl {
		if d == 6 {
			sixes++
		}
		if sixes > 2 {
			return true
		}
	}
	return false
}
