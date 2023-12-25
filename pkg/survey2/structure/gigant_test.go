package structure

import (
	"fmt"
	"testing"
)

func TestGG(t *testing.T) {
	return
	for i := 0; i < 10000; i++ {
		fmt.Printf("test GasGigant %v/10000\r", i+1)
		gg := NewGasGigant(fmt.Sprintf("%v", i))
		switch gg.Size {
		case "GS":
			if gg.Diameter < 2 {
				t.Errorf("diameter '%v' (expect 2-6 for size GS)", gg.Diameter)
			}
			if gg.Diameter > 6 {
				t.Errorf("diameter '%v' (expect 2-6 for size GS)", gg.Diameter)
			}
			if gg.Mass < 10 {
				t.Errorf("mass '%v' (expect 10-35 for size GS)", gg.Diameter)
			}
			if gg.Mass > 35 {
				t.Errorf("mass '%v' (expect 10-35 for size GS)", gg.Diameter)
			}
		case "GM":
			if gg.Diameter < 6 {
				t.Errorf("diameter '%v' (expect 6-12 for size GM)", gg.Diameter)
			}
			if gg.Diameter > 12 {
				t.Errorf("diameter '%v' (expect 6-12 for size GM)", gg.Diameter)
			}
			if gg.Mass < 40 {
				t.Errorf("mass '%v' (expect 40-340 for size GM)", gg.Diameter)
			}
			if gg.Mass > 340 {
				t.Errorf("mass '%v' (expect 40-340 for size GM)", gg.Diameter)
			}
		case "GL":
			if gg.Diameter < 8 {
				t.Errorf("diameter '%v' (expect 8-18 for size GL)", gg.Diameter)
			}
			if gg.Diameter > 18 {
				t.Errorf("diameter '%v' (expect 8-18 for size GL)", gg.Diameter)
			}
			if gg.Mass < 350 {
				t.Errorf("mass '%v' (expect 350-4000 for size GL)", gg.Diameter)
			}
			if gg.Mass > 4000 {
				t.Errorf("mass '%v' (expect 350-4000 for size GL)", gg.Diameter)
			}
		default:
			t.Errorf("unknown size '%v'", gg.Size)
		}
	}
}
