package sector

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/mgt2/corebook/world"
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/language"
)

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
	sector, err := New("Test", 8, 10, 0)
	fmt.Println(sector.Name())
	if err != nil {
		t.Errorf(err.Error())
	}
	data := [][]string{}
	for _, v := range sector.GridSorted() {
		hex := hexagon.StdCoords(v)
		lang, _ := language.New("VILANI")
		name := language.NewWord(dice.New(), lang, 0)
		c := world.NewConstructor(
			world.Instruction(world.KEY_HEX, hex),
			world.Instruction(world.KEY_NAME, name+" "+hex),
			//Instruction(KEY_SECTOR_DENCITY, DENSITY_FORCE_PRESENT),
		)
		if hex == "0505" {
			//ishsish 0203     0203  C   X760246-3  De Lo Lt               A  G
			c.AddInstruction(world.Instruction(world.KEY_DATA, "Earth|0505|MNSR|A876A79-B|Ga Hi||G"))
			fmt.Println("-----")
		}
		if hex == "0506" {
			//ishsish 0203     0203  C   X760246-3  De Lo Lt               A  G
			c.AddInstruction(world.Instruction(world.KEY_DATA, "Lun Escarpa|0506|N|B400310-C|Ht Lo Va||G"))
			fmt.Println("+++++")
		}
		w, err := c.Create()
		if w == nil {
			continue
		}
		data = append(data, w.ShortData())
		if err != nil {
			t.Errorf(err.Error())
		}
		sector.AddWorld(w)
	}
	PrintAsTable(data...)
}
