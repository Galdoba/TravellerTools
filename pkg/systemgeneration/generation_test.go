package systemgeneration

import (
	"fmt"
	"testing"
)

func TestGeneration(t *testing.T) {
	for _, name := range []string{
		"S1",
		"S2",
		"S3",
		"S5",
		"S8o7dsf",
	} {
		gen, _ := NewGenerator(name)
		if gen.NextStep != 1 {
			t.Errorf("have nextstep=%v (expect 1)", gen.NextStep)
		}
		if gen.System == nil {
			t.Errorf("Star System not generated")
		}
		if gen.System.ObjectType != ObjectUNDEFINED && gen.System.ObjectType != ObjectNONE {
			t.Errorf("Star System object not set\n  have %v (expect %v or %v)", gen.System.ObjectType, ObjectUNDEFINED, ObjectNONE)
		}
		if err := gen.GenerateData(); err != nil {
			t.Errorf("error: %v", err.Error())
		}
		fmt.Println(gen.System)
		// gen.Step01()
		// gen.trackStatus()
		// gen.Step02()
		// gen.trackStatus()
		fmt.Println(" ")
		fmt.Println("//////////")
		fmt.Println(" ")
	}

}
