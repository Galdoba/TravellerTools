package sector

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
)

//SectorMap - Коллекция Хексов вызывается по кубическим координатам
type SectorMap struct {
	sectorName string
	w          int
	h          int
	r          int
	values     map[int]int
	byHex      map[hexagon.Hexagon]*WorldData
}

func (sm *SectorMap) gridType() int {
	if sm.w == 0 && sm.h == 0 && sm.r > 0 {
		return Location_Type_Spiral
	}
	if sm.w > 0 && sm.h > 0 && sm.r == 0 {
		return Location_Type_Square
	}
	return defaultValue
}

type World interface {
	Name() string
	//Location() string
	Bases() string
	Statistics() string
	TradeCodes() string
	TravelCode() string
	//Allegiance() string
	PBG() string
}

type WorldData struct {
	w   World
	loc string
}

func NewWorldData(hex hexagon.Hexagon, w World, sector_grid_type int) WorldData {
	wd := WorldData{}
	wd.w = w
	wd.loc = Location(hex, sector_grid_type)
	return wd
}

func (wd *WorldData) String() string {
	switch {
	case wd.w == nil:
		return ""
	}
	return fmt.Sprintf("%v  %v  %v  %v  %v  %v  %v",
		wd.w.Name(),
		wd.loc,
		wd.w.Bases(),
		wd.w.Statistics(),
		wd.w.TradeCodes(),
		wd.w.TravelCode(),
		wd.w.PBG(),
	)
}

func Location(hex hexagon.Hexagon, lType int) string {
	loc := ""
	switch lType {
	case Location_Type_Spiral:
		loc = fmt.Sprintf("[%v %v %v]", hex.CoordQ(), hex.CoordR(), hex.CoordS())
	case Location_Type_Square:
		x, y := fmt.Sprintf("%v", hex.CoordX()), fmt.Sprintf("%v", hex.CoordY())
		if len(x) == 1 {
			x = "0" + x
		}
		if len(y) == 1 {
			y = "0" + y
		}
		loc = x + y
	}

	return loc
}

func (sm *SectorMap) Name() string {
	return sm.sectorName
}

const (
	defaultValue = iota
	Location_Type_Spiral
	Location_Type_Square
	lowX
	lowY
	highX
	highY
)

func New(name string, w, h, r int) (*SectorMap, error) {
	sm := SectorMap{}
	sm.byHex = make(map[hexagon.Hexagon]*WorldData)
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
			sm.byHex[HexagonOnGrid(x, y)] = nil
		}
		//create by spiral
	case w > 0 && h > 0 && r == 0:
		sm.w = w
		sm.h = h
		for y := 1; y <= h; y++ {
			for x := 1; x <= w; x++ {
				sm.byHex[HexagonOnGrid(x, y)] = nil
			}
		}
	}
	sm.values = make(map[int]int)
	lX, lY, hX, hY := 1000000, 1000000, -1000000, -1000000
	for v, _ := range sm.byHex {
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

func (sm *SectorMap) AddWorld(hex hexagon.Hexagon, w World) error {
	if sm.gridType() == defaultValue {
		return fmt.Errorf("sector grid type undefined")
	}
	wd := NewWorldData(hex, w, sm.gridType())
	sm.byHex[hex] = &wd
	return nil
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
	for v, _ := range sm.byHex {
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

func (sm *SectorMap) PrintAsTable() {
	data := [][]string{}
	for _, h := range sm.GridSorted() {
		wd := sm.byHex[h]
		if wd == nil {
			continue
		}
		dataline := strings.Split(wd.String(), "  ")
		data = append(data, dataline)
	}
	printAsTable(data...)
}

func printAsTable(flds ...[]string) {
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
