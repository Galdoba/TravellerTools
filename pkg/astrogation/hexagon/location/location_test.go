package location

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
)

type coord struct {
	x, y int
}

func input() []coord {
	cor := []coord{}
	for x := -5; x < 5; x++ {
		for y := -5; y < 5; y++ {
			cor = append(cor, coord{x, y})
		}
	}
	return cor
}

func TestOTULocation(t *testing.T) {
	for _, cor := range input() {
		hx := hexagon.New_Unsafe(hexagon.Feed_HEX, cor.x, cor.y)
		loc := New(hx, COORDINATE_STANDARD_OTU)

		fmt.Println(cor, loc)
	}
	fmt.Println("--------")
	hx := hexagon.New_Unsafe(hexagon.Feed_HEX, 50, -19)
	loc := New(hx, COORDINATE_STANDARD_OTU)
	fmt.Println(hx.AsHex(), loc)
}
