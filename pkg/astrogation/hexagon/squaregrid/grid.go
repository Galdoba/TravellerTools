package squaregrid

import (
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
)

type SGrid struct {
	hex  map[coords]hexagon.Hexagon
	maxX int
	maxY int
}
type coords struct {
	x int
	y int
}

func Coords(x, y int) coords {
	return coords{x, y}
}

func New(xMax, yMax int) SGrid {
	sg := SGrid{}
	for x := 0; x <= xMax; x++ {
		for y := 0; y <= yMax; y++ {
			sg.hex[Coords(x, y)] = hexagon.Global(x, y)
		}
	}
	return sg
}

func GridContains(grid SGrid, coord coords) bool {
	if _, ok := grid.hex[coord]; ok {
		return ok
	}
	return false
}

func Neibours(sg SGrid, coord coords) []hexagon.Hexagon {
	approvedNeib := []hexagon.Hexagon{}
	hex := sg.hex[coord]
	neib, _ := hexagon.Neighbors(hex.AsCube())
	for _, h := range neib {
		if GridContains(sg, Coords(h.CoordX(), h.CoordY())) {
			approvedNeib = append(approvedNeib, h)
		}
	}
	return approvedNeib
}
