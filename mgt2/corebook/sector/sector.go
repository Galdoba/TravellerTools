package sector

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
)

//SectorMap - Коллекция Хексов вызявается по кубическим координатам
type SectorMap struct {
	sectorName string
	w          int
	h          int
	r          int
	values     map[int]int
	hex        map[globalCoords]hexagon.Hexagon
	world      map[globalCoords]World
}

type World interface {
	Name() string
	Location() string
	Bases() string
	Statistics() string
	TradeCodes() string
	TravelCode() string
	Allegiance() string
	GG() string
}

func (sm *SectorMap) Name() string {
	return sm.sectorName
}

const (
	lowX  = 1
	lowY  = 2
	highX = 3
	highY = 4
)

type globalCoords struct {
	x, y int
}

func New(name string, w, h, r int) (*SectorMap, error) {
	sm := SectorMap{}
	sm.hex = make(map[globalCoords]hexagon.Hexagon)
	sm.sectorName = name
	switch {
	default:
		return &sm, fmt.Errorf("cannot create sector: bad input: w=%v h=%v r=%v", w, h, r)
	case w == 0 && h == 0 && r > 0:
		sm.r = r
		center := hexagon.New_Unsafe(hexagon.Feed_CUBE, 0, 0, 0)
		spiral, err := hexagon.Spiral(center.AsCube(), sm.r)
		if err != nil {
			return &sm, err
		}
		for _, hx := range spiral {
			x, y := hx.HexValues()
			sm.hex[Coords(x, y)] = HexagonOnGrid(x, y)
		}
		//create by spiral
	case w > 0 && h > 0 && r == 0:
		sm.w = w
		sm.h = h
		for y := 1; y <= h; y++ {
			for x := 1; x <= w; x++ {
				sm.hex[Coords(x, y)] = HexagonOnGrid(x, y)
			}
		}
	}
	sm.values = make(map[int]int)
	lX, lY, hX, hY := 1000000, 1000000, -1000000, -1000000
	for _, v := range sm.hex {
		x, y := v.HexValues()
		if x > hX {
			hX = x
		}
		if y > hY {
			hY = y
		}
		if x < lX {
			lX = x
		}
		if y < lY {
			lY = y
		}
	}
	sm.values[lowX] = lX
	sm.values[lowY] = lY
	sm.values[highX] = hX
	sm.values[highY] = hY
	return &sm, nil
}

func Coords(x, y int) globalCoords {
	return globalCoords{x, y}
}

func HexagonOnGrid(x, y int) hexagon.Hexagon {
	h, _ := hexagon.New(hexagon.Feed_HEX, x, y)
	return h
}

func HexagonOnCube(q, r, s int) hexagon.Hexagon {
	h, _ := hexagon.New(hexagon.Feed_CUBE, q, r, s)
	return h
}

func (sm *SectorMap) Grid() []hexagon.Hexagon {
	grid := []hexagon.Hexagon{}
	for _, v := range sm.hex {
		grid = append(grid, v)
	}
	return grid
}

func (sm *SectorMap) GridSorted() []hexagon.Hexagon {
	grid := []hexagon.Hexagon{}
	for y := sm.values[lowY]; y <= sm.values[highY]; y++ {
		for x := sm.values[lowX]; x <= sm.values[highX]; x++ {
			grid = append(grid, HexagonOnGrid(x, y))
		}
	}
	return grid
}

func (sm *SectorMap) AddWorld(w World) {
	glCo := Coords(1, 1)

}
