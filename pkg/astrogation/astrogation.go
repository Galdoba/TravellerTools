package astrogation

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/Galdoba/utils"
)

const (
	directionSE = iota
	directionNE
	directionN
	directionNW
	directionSW
	directionS
)

type stellarHex struct {
	hexStr string
	hex    hexCoords
	cube   cubeCoords
}

func Hex(hexStr string) stellarHex {
	h := stellarHex{}
	h.hexStr = hexStr
	bts := []byte(hexStr)
	col, err := strconv.Atoi(string(bts[0]) + string(bts[1]))
	if err != nil {
		fmt.Println(err)
	}
	row, err := strconv.Atoi(string(bts[2]) + string(bts[3]))
	if err != nil {
		fmt.Println(err)
	}
	h.hex = setHexCoords(col, row)
	h.cube = hexToCube(h.hex)
	return h
}

type cubeCoords struct {
	q int
	r int
	s int
}

// func cubeCoordsStr(cube cubeCoords) string {
// 	fmt.Println(cube)
// 	xStr := coordNumToStr("X", cube.q)
// 	yStr := coordNumToStr("Y", cube.r)
// 	zStr := coordNumToStr("Z", cube.s)
// 	output := xStr + " " + yStr + " " + zStr
// 	return output
// }

// func coordNumToStr(coordName string, x int) string {
// 	xStr := coordName
// 	if x < 0 {
// 		xStr += "-"
// 		x = x * -1
// 	} else {
// 		xStr += " "
// 	}
// 	fmt.Println("1:", xStr)
// 	if x < 10 && x > -10 {
// 		xStr += "0"
// 		xStr += strconv.Itoa(x)
// 	} else {
// 		xStr += strconv.Itoa(x)
// 	}
// 	return xStr
// }

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

// func oddQToCube(hex hexCoords) cubeCoords {
// 	x := hex.col
// 	z := hex.row - (hex.col-(hex.col&1))/2
// 	y := -x - z
// 	return setCubeCoords(x, y, z)
// }

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
		{1, 0, -1}, // SE
		{1, -1, 0}, // NE
		{0, -1, 1}, // N
		{-1, 0, 1}, // NW
		{-1, 1, 0}, // SW
		{0, 1, -1}, // S
	}
}

func cubeDirection(direction int) cubeCoords {
	return cubeDirectionVectors[direction]
}

func cubeAdd(hex, vec cubeCoords) cubeCoords {
	return cubeCoords{hex.q + vec.q, hex.r + vec.r, hex.s + vec.s}
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

// func hexNeighbor(hex hexCoords, direction int) hexCoords {
// 	parity := hex.col & 1
// 	dir := hexDirections[parity][direction]
// 	return setHexCoords(hex.col+dir.col, hex.row+dir.row)
// }

func cubeDistance(cubeA, cubeB cubeCoords) int {
	return int((math.Abs(float64(cubeA.q-cubeB.q)) + math.Abs(float64(cubeA.r-cubeB.r)) + math.Abs(float64(cubeA.s-cubeB.s))) / 2)
}

func hexDistance(hexA, hexB hexCoords) int {
	cubeA := hexToCube(hexA)
	cubeB := hexToCube(hexB)
	return cubeDistance(cubeA, cubeB)
}

func JumpDistance(h1, h2 string) int {
	sh1 := Hex(h1)
	sh2 := Hex(h2)
	return cubeDistance(sh1.cube, sh2.cube)
}

func evenQToStr(hx hexCoords) string {
	hexStr := ""
	if hx.col < 10 {
		hexStr += "0"
	}
	hexStr += strconv.Itoa(hx.col)
	if hx.row < 10 {
		hexStr += "0"
	}
	hexStr += strconv.Itoa(hx.row)
	return hexStr
}

func addCube(ac, bc cubeCoords) cubeCoords {
	return setCubeCoords(ac.q+bc.q, ac.r+bc.r, ac.s+bc.s)
}

//JumpCoordinatesFrom - дает перечень всех хексов в радиусе j
//требует координатов секторной карты в формате "XXYY"
func JumpCoordinatesFrom(initHex string, j int) []string {
	var coords []string
	start := Hex(initHex)
	for x := -j; x <= j; x++ {
		for y := utils.Max(-j, -x-j); y <= utils.Min(j, -x+j); y++ {
			z := -x - y
			cb := addCube(setCubeCoords(x, y, z), start.cube)
			hx := cubeToHex(cb)
			coords = append(coords, evenQToStr(hx))
		}
	}
	//fmt.Println(coords)
	sort.Strings(coords)
	return coords

}

func JumpCoordinatesAll() []string {
	var jc []string
	for x := 1; x <= 32; x++ {
		for y := 1; y <= 40; y++ {
			hexCoord := setHexCoords(x, y)
			hex := evenQToStr(hexCoord)
			jc = append(jc, hex)
		}
	}
	return jc
}

//////////////
type Coordinator interface {
	CoordX() int
	CoordY() int
}

type Coordinates struct {
	hex  hexCoords
	cube cubeCoords
}

func (c *Coordinates) CoordX() int {
	return c.hex.col
}

func (c *Coordinates) CoordY() int {
	return c.hex.row
}

func CoordinatesOf(coord Coordinator) Coordinates {
	coords := Coordinates{}
	coords.hex = setHexCoords(coord.CoordX(), coord.CoordY())
	coords.cube = hexToCube(coords.hex)
	return coords
}

func (c *Coordinates) HexValues() (int, int) {
	return c.hex.col, c.hex.row
}

func (c *Coordinates) CubeValues() (int, int, int) {
	return c.cube.q, c.cube.r, c.cube.s
}

func NewCoordinates(x, y int) Coordinates {
	coords := Coordinates{}
	coords.hex = setHexCoords(x, y)
	coords.cube = hexToCube(coords.hex)
	return coords
}

func isSame(c1, c2 Coordinates) bool {
	if c1.cube.q != c2.cube.q {
		return false
	}
	if c1.cube.r != c2.cube.r {
		return false
	}
	if c1.cube.s != c2.cube.s {
		return false
	}
	return true
}

func coordsFromCube(cc cubeCoords) Coordinates {
	coord := Coordinates{}
	coord.cube = cc
	coord.hex = cubeToHex(cc)
	return coord
}

func Distance(c1, c2 Coordinates) int {
	return cubeDistance(c1.cube, c2.cube)
}

func DistanceRaw(x1, y1, x2, y2 int) int {
	c1 := NewCoordinates(x1, y1)
	c2 := NewCoordinates(x2, y2)
	return cubeDistance(c1.cube, c2.cube)
}

//JumpCoordinatesFrom - дает перечень всех хексов в радиусе j
//требует координатов секторной карты в формате "XXYY"
func JumpMap(center Coordinator, radius int) []Coordinates {
	var coords []Coordinates
	startHex := hexCoords{center.CoordX(), center.CoordY()}
	start_cube := hexToCube(startHex)
	spiral_cube := cubeSpiral(start_cube, radius)
	for _, cc := range spiral_cube {
		coords = append(coords, coordsFromCube(cc))
	}
	return coords
}
