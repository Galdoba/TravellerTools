package sector

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/mgt2/corebook/world"
	"github.com/Galdoba/TravellerTools/mgt2/gypsy/world/starport"
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/language"
)

func TestInSector(t *testing.T) {
	sector, err := New("Test", 3, 4, 0)
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
			world.Instruction(world.KEY_NAME, name+" "),
			//Instruction(KEY_SECTOR_DENCITY, DENSITY_FORCE_PRESENT),
		)
		if hex == "0101" {
			//ishsish 0203     0203  C   X760246-3  De Lo Lt               A  G
			c.AddInstruction(world.Instruction(world.KEY_DATA, "Drinax|0101||A434507-F|Ni||G"))
		}

		w, err := c.Create()
		if w == nil {
			continue
		}
		data = append(data, w.ShortData())
		if err != nil {
			t.Errorf(err.Error())
		}
		if err := sector.AddWorld(v, w); err != nil {
			t.Errorf(err.Error())
		}

	}
	for hex, wd := range sector.byHex {
		if wd == nil {
			continue
		}
		fmt.Println(hex, wd.String())
		fmt.Println("FUEL COST", starport.FuelCost(wd.w, 1.0))
		fmt.Println("Item Price", starport.ItemProceAdjustment(wd.w))
	}
	fmt.Println("--------------")
	sector.PrintAsTable()
}
