package world

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/mgt2/corebook/sector"
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
)

func TestConstructor(t *testing.T) {
	data := [][]string{}
	for i := 100; i < 125; i++ {
		c := NewConstructor(
			Instruction(KEY_SEED, fmt.Sprintf("Test %v", i+25)),
			Instruction(KEY_NAME, ""),
			Instruction(KEY_SECTOR_DENCITY, DENSITY_FORCE_PRESENT),
		)
		w, err := c.Create()
		if w == nil {
			continue
		}
		data = append(data, w.ShortData())
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	PrintAsTable(data...)
}

func PrintAsTable(flds ...[]string) {
	lMap := make(map[int]int)
	for _, fld := range flds {
		for i, f := range fld {
			if lMap[i] < len(f) {
				lMap[i] = len(f)
			}
		}

	}
	for _, fld := range flds {
		for i, f := range fld {
			for len(f) < lMap[i] {
				f += " "
			}
			fmt.Print(f + "  ")
		}
		fmt.Print("\n")
	}
}

func TestInSector(t *testing.T) {
	sector, err := sector.New("Test", 32, 40, 0)
	if err != nil {
		t.Errorf(err.Error())
	}
	data := [][]string{}
	for _, v := range sector.GridSorted() {
		hex := hexagon.StdCoords(v)
		c := NewConstructor(
			Instruction(KEY_SEED, fmt.Sprintf("Test %v", hex)),
			Instruction(KEY_NAME, ""),
			Instruction(KEY_HEX, hex),
			//Instruction(KEY_SECTOR_DENCITY, DENSITY_FORCE_PRESENT),
		)
		w, err := c.Create()
		if w == nil {
			continue
		}
		data = append(data, w.ShortData())
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	PrintAsTable(data...)
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
