package orbitns

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

type testPair struct {
	input  float64
	output float64
}

func TestDecode(t *testing.T) {
	//4294967295
	//42 94 967295
	ui, errI := encodeUINT(4.3, 6)
	if errI != nil {
		t.Errorf("err=%v|%v", errI.Error(), ui)
	}
	prim, full, frac, err := decodeUINT(ui)
	if err != nil {
		t.Errorf("err=%v|%v:%v %v %v", err.Error(), ui, prim, full, frac)
	}

}

func TestOrbitn2AU_correct(t *testing.T) {

	data := []testPair{
		{input: 4.3, output: 1.96},
		{input: 0.09, output: 0.036},
		{input: 6.1, output: 5.68},
		{input: 12.1, output: 338.7},
	}
	for testNum, testData := range data {
		orb, _ := NewOrbit(testData.input, dice.New())
		result := OR2MKM(orb.orbN)
		if result != testData.output {
			t.Errorf("test %v:\n have =%v, expected =%v from input =%v", testNum, result, testData.output, testData.input)
		}
	}
}

func TestOrbitn2AU_incorrect(t *testing.T) {
	return
	data := []testPair{
		// {input: 4.3, output: 1.96},
		// {input: 0.09, output: 0.036},
		// {input: 6.1, output: 5.68},
		// {input: 12.2, output: 338.7},
	}
	for testNum, testData := range data {
		orb, _ := NewOrbit(testData.input, dice.New())
		result := OR2MKM(orb.orbN)
		if result == testData.output {
			t.Errorf("test %v:\n have =%v, expected other than =%v from input =%v", testNum, result, testData.output, testData.input)
		}
	}
}

func TestMap(t *testing.T) {
	dice := dice.New()
	for f := 0; f <= 200; f++ {
		fl := float64(f) / 10
		orb, err := NewOrbit(fl, dice)
		if err != nil {
			t.Errorf("fl = %v: %v", fl, err.Error())
		}
		fmt.Println(fl, orb.distanse)
	}
}
