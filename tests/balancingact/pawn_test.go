package balancingact

import (
	"fmt"
	"testing"
	"time"
)

type input struct {
	name string
	role int
}

func testInput() []input {
	return []input{
		{"Leader 1", Leader},
		{"Leader 2", Leader},
		{"Agent 3", Agent},
		{"Agent 4", Agent},
		//{"Unknown", 0}, //expected to fail
	}
}

func TestPawn(t *testing.T) {

	for i, inp := range testInput() {
		fmt.Print("Creating pawn ", i)
		p, err := createPawn(inp.name, inp.role)
		if err != nil {
			t.Errorf("internal error: %v", err.Error())
		}
		if p == nil {
			continue
		}
		if p.position != Leader && p.position != Agent {
			t.Errorf("unexpected value pawn.position = %v (expect %v or %v)", p.position, Leader, Agent)
		}
		if len(p.chars) != 4 {
			t.Errorf("pawn.chars have %v entries (expected 4)", len(p.chars))
		}
		for _, v := range []int{WIL, INT, EDU, CHR} {
			if _, ok := p.chars[v]; !ok {
				t.Errorf("pawn.chars[%v] expected to present)", v)
			}
		}
		if len(p.skills) != 5 {
			t.Errorf("pawn.skills have %v entries (expected 5)", len(p.skills))
		}
		for _, v := range []int{Administration, CovertOps, Economics, Politics, Military} {
			if _, ok := p.skills[v]; !ok {
				t.Errorf("pawn.skills[%v] expected to present)", v)
			}
		}
		//fmt.Println("pawn", i, "created:")
		//fmt.Println(p)
		time.Sleep(time.Millisecond * 15)
		//fmt.Println("  ")
		fmt.Println(" SUCCESS!")
	}

}
