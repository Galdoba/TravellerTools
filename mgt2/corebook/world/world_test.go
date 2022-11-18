package world

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	data := [][]string{}
	for i := 0; i < 20; i++ {
		c := NewConstructor(
			Instruction(KEY_SEED, fmt.Sprintf("Test %v", i+25)),
			//Instruction(KEY_NAME, ""),
			Instruction(KEY_SECTOR_DENCITY, DENSITY_FORCE_PRESENT),
		)
		w, err := c.Create()
		if w == nil {
			continue
		}
		fmt.Println(i, w.ShortData())
		data = append(data, w.ShortData())
		if err != nil {
			t.Errorf(err.Error())
		}
	}
}

/*
UWP: BA7068A-B
     ||||||| +- Tech Level : 11
     ||||||+--- Law Level  : 10
     |||||+---- Goverment  : 8 (Civil Service Burocracy)
     ||||+----- Population : 6 (700,000)
     |||+------ Hydrosphere: 0 (2%)
     ||+------- Atmosphere : 7 (Standard, Tainted)
     |+-------- Size       : A (16,492 km)
     +--------- Starport   : B (Good, 1000 Cr, Highport)

*/
