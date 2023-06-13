package worldmap

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/struct/world"
)

func TestTri(t *testing.T) {
	return
	for i := 5; i < 6; i++ {
		grid := newGrid(i)
		fmt.Println("grid", i)
		dataMap := make(map[string]int)
		coordMap := make(map[coordinates]int)
		reverseCoordMap := make(map[int]int)
		for k, v := range grid {
			coordMap[k]++
			//	fmt.Println(k)
			//	fmt.Println(v.neiboirs)
			dataMap[fmt.Sprintf("have %v nodes", len(v.neiboirs))]++
			dataMap[fmt.Sprintf("total hexes")]++

		}
		for _, v := range coordMap {
			reverseCoordMap[v]++
		}
		fmt.Println("------------")
		for k, v := range dataMap {
			fmt.Println(k, v)
		}
		fmt.Println("------------")
		for k, v := range coordMap {

			fmt.Println(k, v)

		}

	}

}

func TestMap(t *testing.T) {
	wrld, _ := world.NewWorld()
	wrld.GenerateBasic(dice.New())
	wm := New(wrld)
	fmt.Println(wm)
}
