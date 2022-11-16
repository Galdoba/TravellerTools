package hexagon

import (
	"fmt"
	"math"
	"strings"
)

const (
	defaultValue = iota
	Direction_N
	Direction_NE
	Direction_SE
	Direction_S
	Direction_SW
	Direction_NW
	Feed_HEX
	Feed_CUBE
	wrongValue
)

//Hexagon - Travaller использует глобальные кординаты из Хексагонов типа "odd-q"
type Hexagon struct {
	cube cubeCoords
	hex  hexCoords
}

func Global(x, y int) Hexagon {
	return New_Unsafe(Feed_HEX, x, y)

}

func New_Unsafe(feed int, coordinates ...int) Hexagon {
	hex, _ := New(feed, coordinates...)
	return hex
}

func New(feed int, coordinates ...int) (Hexagon, error) {
	hx := Hexagon{}
	switch feed {
	default:
		return hx, fmt.Errorf("feed value unreconised")
	case Feed_HEX:
		if len(coordinates) != 2 {
			return hx, fmt.Errorf("2 coordinates expected")
		}
		hx.hex.col = coordinates[0]
		hx.hex.row = coordinates[1]
		hx.cube = hexToCube(hx.hex)
	case Feed_CUBE:
		if len(coordinates) != 3 {
			return hx, fmt.Errorf("3 coordinates expected")
		}
		hx.cube.q = coordinates[0]
		hx.cube.r = coordinates[1]
		hx.cube.s = coordinates[2]
		hx.hex = cubeToHex(hx.cube)
	}
	switch {
	case hx.cube.sum() != 0:
		return hx, fmt.Errorf("cube sum is not 0")

	}
	return hx, nil
}

type Hex interface {
	CoordX() int
	CoordY() int
}

func (hx *Hexagon) AsHex() Hex {
	return &hx.hex
}

func (c *Hexagon) CoordX() int {
	return c.hex.col
}

func (h *hexCoords) CoordX() int {
	return h.col
}

func (c *Hexagon) CoordY() int {
	return c.hex.row
}

func (h *hexCoords) CoordY() int {
	return h.row
}

func FromHex(h Hex) Hexagon {
	hex, _ := New(Feed_HEX, h.CoordX(), h.CoordY())
	return hex
}

type Cube interface {
	CoordQ() int
	CoordR() int
	CoordS() int
}

func (hx *Hexagon) AsCube() Cube {
	return &hx.cube
}

func (hx *Hexagon) CoordQ() int {
	return hx.cube.q
}

func (hx *Hexagon) CoordR() int {
	return hx.cube.r
}

func (hx *Hexagon) CoordS() int {
	return hx.cube.s
}

func FromCube(c Cube) (Hexagon, error) {
	return New(Feed_CUBE, c.CoordQ(), c.CoordR(), c.CoordS())
}

func (hx *Hexagon) String() string {
	return fmt.Sprintf("[Hex: {%v %v}, Cube{%v %v %v}]", hx.hex.col, hx.hex.row, hx.cube.q, hx.cube.r, hx.cube.s)
}

func (c *cubeCoords) sum() int {
	return c.q + c.r + c.s
}

////////////////////////////////////////////////////////////////

type cubeCoords struct {
	q int
	r int
	s int
}

func (cc *cubeCoords) CoordQ() int {
	return cc.q
}
func (cc *cubeCoords) CoordR() int {
	return cc.r
}
func (cc *cubeCoords) CoordS() int {
	return cc.s
}

func setCubeCoords(q, r, s int) cubeCoords {
	cube := cubeCoords{}
	cube.q = q //x
	cube.r = r //y
	cube.s = s //z
	return cube
}

// func evenQToCube(hex hexCoords) cubeCoords {
// 	var x = hex.col
// 	var z = hex.row - (hex.col+(hex.col&1))/2
// 	var y = -x - z
// 	return setCubeCoords(x, y, z)
// }

// func cubeToEvenq(cube cubeCoords) hexCoords {
// 	var col = cube.q
// 	var row = cube.s + (cube.q+(cube.q&1))/2
// 	return setHexCoords(col, row)
// }

type hexCoords struct {
	col int
	row int
}

func setHexCoords(c, r int) hexCoords {
	offCrds := hexCoords{}
	offCrds.col = c
	offCrds.row = r
	return offCrds
}

func cubeToHex(cube cubeCoords) hexCoords {
	col := cube.q
	row := cube.r + (cube.q-(cube.q&1))/2
	return setHexCoords(col, row)
}

func hexToCube(hex hexCoords) cubeCoords {
	q := hex.col
	r := hex.row - (hex.col-hex.col&1)/2
	return cubeCoords{q, r, -q - r}
}

//var hexDirections [][]hexCoords
var cubeDirectionVectors []cubeCoords

func init() {
	// hexDirections = [][]hexCoords{
	// 	{hexCoords{1, 0}, hexCoords{1, -1}, hexCoords{0, -1}, hexCoords{-1, -1}, hexCoords{-1, 0}, hexCoords{0, 1}},
	// 	{hexCoords{1, 1}, hexCoords{1, 0}, hexCoords{0, -1}, hexCoords{-1, 0}, hexCoords{-1, 1}, hexCoords{0, 1}},
	// }
	// hexDirections = [][]hexCoords{
	// 	{hexCoords{0, -1}, hexCoords{1, -1}, hexCoords{1, 0}, hexCoords{0, 1}, hexCoords{-1, 0}, hexCoords{-1, -1}},
	// 	{hexCoords{0, -1}, hexCoords{1, 0}, hexCoords{1, 1}, hexCoords{0, 1}, hexCoords{-1, 1}, hexCoords{-1, 0}},
	// }
	cubeDirectionVectors = []cubeCoords{
		{0, -1, 1}, // N
		{1, -1, 0}, // NE
		{1, 0, -1}, // SE
		{0, 1, -1}, // S
		{-1, 1, 0}, // SW
		{-1, 0, 1}, // NW

	}
}

func cubeDirection(direction int) cubeCoords {
	return cubeDirectionVectors[direction]
}

func cubeAdd(hex, vec cubeCoords) cubeCoords {
	return cubeCoords{hex.q + vec.q, hex.r + vec.r, hex.s + vec.s}
}

func Add(hex, vec Hexagon) Hexagon {
	newCube := cubeAdd(hex.cube, vec.cube)
	newHex, _ := New(Feed_CUBE, newCube.q, newCube.r, newCube.s)
	return newHex
}

func cubeNeighbor(cube cubeCoords, direction int) cubeCoords {
	return cubeAdd(cube, cubeDirection(direction))
}

func cubeScale(hex cubeCoords, factor int) cubeCoords {
	return cubeCoords{hex.q * factor, hex.r * factor, hex.s * factor}
}

func cubeRing(center cubeCoords, radius int) []cubeCoords {
	ring := []cubeCoords{}
	hex := cubeAdd(center, cubeScale(cubeDirection(4), radius))
	for i := 0; i < 6; i++ {
		for j := 0; j < radius; j++ {
			ring = append(ring, hex)
			hex = cubeNeighbor(hex, i)
		}
	}
	return ring
}

func cubeSpiral(center cubeCoords, radius int) []cubeCoords {
	spiral := []cubeCoords{center}
	for i := 1; i <= radius; i++ {
		spiral = append(spiral, cubeRing(center, i)...)
	}
	return spiral
}

func cubeDistance(cubeA, cubeB cubeCoords) int {
	return int((math.Abs(float64(cubeA.q-cubeB.q)) + math.Abs(float64(cubeA.r-cubeB.r)) + math.Abs(float64(cubeA.s-cubeB.s))) / 2)
}

// func hexDistance(hexA, hexB hexCoords) int {
// 	cubeA := hexToCube(hexA)
// 	cubeB := hexToCube(hexB)
// 	return cubeDistance(cubeA, cubeB)
// }

func addCube(ac, bc cubeCoords) cubeCoords {
	return setCubeCoords(ac.q+bc.q, ac.r+bc.r, ac.s+bc.s)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//////////////

func (c *Hexagon) HexValues() (int, int) {
	return c.hex.col, c.hex.row
}

func (c *Hexagon) CubeValues() (int, int, int) {
	return c.cube.q, c.cube.r, c.cube.s
}

func Match(c1, c2 Hexagon) bool {
	return MatchHex(&c1, &c2) && MatchCube(&c1, &c2)

}

func MatchHex(h1, h2 Hex) bool {
	return h1.CoordX() == h2.CoordX() && h1.CoordY() == h2.CoordY()
}

func MatchCube(c1, c2 Cube) bool {
	return c1.CoordQ() == c2.CoordQ() && c1.CoordR() == c2.CoordR() && c1.CoordS() == c2.CoordS()
}

func Distance(c1, c2 Hexagon) int {
	return cubeDistance(c1.cube, c2.cube)
}

func DistanceCube(c1, c2 Hexagon) int {
	return cubeDistance(c1.cube, c2.cube)
}

func DistanceHex(h1, h2 Hex) int {
	//hx1, _ := FromHex(h1)
	//hx2, _ := FromHex(h2)
	return cubeDistance(FromHex(h1).cube, FromHex(h2).cube)
}

func Neighbors(c Cube) ([]Hexagon, error) {
	cb := cubeCoords{c.CoordQ(), c.CoordR(), c.CoordS()}
	cbNeib := []Hexagon{}
	for dir := Direction_N; dir <= Direction_NW; dir++ {
		cc := cubeNeighbor(cb, dir)
		hx, err := FromCube(&cc)
		if err != nil {
			return nil, err
		}
		cbNeib = append(cbNeib, hx)
	}
	return cbNeib, nil
}

func Ring(cntr Cube, radius int) ([]Hexagon, error) {
	center := cubeCoords{cntr.CoordQ(), cntr.CoordR(), cntr.CoordS()}
	ringCube := cubeRing(center, radius)
	output := []Hexagon{}
	for _, cb := range ringCube {
		hx, err := New(Feed_CUBE, cb.q, cb.r, cb.s)
		if err != nil {
			return nil, err
		}
		output = append(output, hx)
	}
	return output, nil
}

func Spiral(cntr Cube, radius int) ([]Hexagon, error) {
	center := cubeCoords{cntr.CoordQ(), cntr.CoordR(), cntr.CoordS()}
	ringCube := cubeSpiral(center, radius)
	output := []Hexagon{}
	for _, cb := range ringCube {
		hx, err := New(Feed_CUBE, cb.q, cb.r, cb.s)
		if err != nil {
			return nil, err
		}
		output = append(output, hx)
	}
	return output, nil
}

func StdCoords(h Hexagon) string {
	str := []string{}
	for _, v := range []int{h.CoordX(), h.CoordY()} {
		if v < 1 {
			return ""
		}
		out := fmt.Sprintf("%v", v)
		if len(out) == 1 {
			out = "0" + out
		}
		str = append(str, out)
	}
	return strings.Join(str, "")
}
