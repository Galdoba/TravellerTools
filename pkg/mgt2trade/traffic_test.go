package mgt2trade

import (
	"fmt"
	"testing"
)

type testInput struct {
	uwp    string
	travel string
	x      int
	y      int
}

type inputStruct struct {
	source testInput
	dest   testInput
}

func (ti *testInput) MW_UWP() string {
	return ti.uwp
}

func (ti *testInput) TravelZone() string {
	return ti.travel
}
func (ti *testInput) CoordX() int {
	return ti.x
}
func (ti *testInput) CoordY() int {
	return ti.y
}

func input() []inputStruct {
	inp := []inputStruct{}
	inp = append(inp, inputStruct{source: testInput{"A576655-C", "", 23, 20}, dest: testInput{"B867564-6", "R", 21, 25}})
	inp = append(inp, inputStruct{source: testInput{"C543487-B", "A", 24, 21}, dest: testInput{"B867564-6", "R", 22, 21}})
	inp = append(inp, inputStruct{source: testInput{"A894A96-F", "", 32, 35}, dest: testInput{"B552665-B", "", 30, 32}})
	inp = append(inp, inputStruct{source: testInput{"A894A96-F", "", 32, 35}, dest: testInput{"C645747-5", "A", 30, 32}})
	inp = append(inp, inputStruct{source: testInput{"C645747-5", "A", 30, 32}, dest: testInput{"A894A96-F", "", 12, 35}})

	return inp
}

func Test_PassengerTraffic(t *testing.T) {
	inp := input()
	for i, input := range inp {
		fmt.Println("test", i+1, input.source, input.dest)
		pf, err := BasePassengerFactor(&input.source, &input.dest)
		if err != nil {
			t.Errorf("internal error: %v", err.Error())
			continue
		}
		if pf == -1000 {
			t.Errorf("factor value was not adressed")
			continue
		}
		bf, err := BaseFreightFactor(&input.source, &input.dest)
		if err != nil {
			t.Errorf("internal error: %v", err.Error())
			continue
		}
		if bf == -1000 {
			t.Errorf("factor value was not adressed")
			continue
		}
		fmt.Println("Passenger Factor =", pf)
		fmt.Println("Freight Factor =", bf)
		fmt.Println("Test PASS")
	}

}
