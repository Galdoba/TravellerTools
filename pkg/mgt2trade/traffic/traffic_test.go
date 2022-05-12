package traffic

import (
	"fmt"
	"testing"
)

type testInput struct {
	uwp    string
	travel string
	x      int
	y      int
	q      int
	r      int
	s      int
	rem    string
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
func (ti *testInput) CoordQ() int {
	return ti.x
}
func (ti *testInput) CoordR() int {
	return ti.y - (ti.x-ti.x%1)/2
}
func (ti *testInput) CoordS() int {
	return -ti.CoordQ() - ti.CoordR()
}
func (ti *testInput) MW_Remarks() string {
	return ti.rem
}

func input() []inputStruct {
	inp := []inputStruct{}
	inp = append(inp, inputStruct{source: testInput{"A43645A-E", "", -107, -17, 0, 0, 0, "Ni"}, dest: testInput{"B867564-6", "", -108, -17, 0, 0, 0, "Ag Ni Ga Pr O:2324"}}) //Drinax ---> Asim
	//inp = append(inp, inputStruct{source: testInput{"C543487-B", "A", 24, 21, ""}, dest: testInput{"B867564-6", "R", 22, 21, ""}})
	//inp = append(inp, inputStruct{source: testInput{"A894A96-F", "", 32, 35, ""}, dest: testInput{"B552665-B", "", 30, 32, ""}})
	//inp = append(inp, inputStruct{source: testInput{"A894A96-F", "", 32, 35, ""}, dest: testInput{"C645747-5", "A", 30, 32, ""}})
	//inp = append(inp, inputStruct{source: testInput{"C645747-5", "A", 30, 32, ""}, dest: testInput{"A894A96-F", "", 12, 35, ""}})

	return inp
}

func Test_PassengerTraffic(t *testing.T) {
	inp := input()
	for i, input := range inp {
		fmt.Println("test", i+1, input.source, input.dest)
		pf, err := BasePassengerFactor_MGT2_Core(&input.source, &input.dest)
		if err != nil {
			t.Errorf("internal error: %v", err.Error())
			continue
		}
		if pf == -1000 {
			t.Errorf("factor value was not adressed")
			continue
		}
		bf, err := BaseFreightFactor_MGT2_Core(&input.source, &input.dest)
		if err != nil {
			t.Errorf("internal error: %v", err.Error())
			continue
		}
		if bf == -1000 {
			t.Errorf("factor value was not adressed")
			continue
		}
		bfMP, err := BaseFreightFactor_MGT1_MP(&input.source, &input.dest)
		if err != nil {
			t.Errorf("internal error: %v", err.Error())
			continue
		}
		if bfMP == -1000 {
			t.Errorf("factor value was not adressed")
			continue
		}

		fmt.Println("MGT2_Core: Passenger Factor =", pf)
		fmt.Println("MGT2_Core: Freight Factor =", bf)
		fmt.Println("MGT1_MP  : Freight Factor =", bfMP)
		fmt.Println("Test PASS")
	}

}
