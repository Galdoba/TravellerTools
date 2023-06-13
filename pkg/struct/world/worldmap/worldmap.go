package worldmap

import (
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/pkg/struct/world"
)

type worldmap struct {
	sizeCode     int
	hydroPer     int
	seismoStress int
	worldX       int
	worldY       int
	WorldHex     map[coordinates]*WorldHex
}

type coordinates struct {
	x int
	y int
}

func newCoords(x, y int) coordinates {
	return coordinates{x, y}
}

func (c *coordinates) CoordX() int {
	return c.x
}
func (c *coordinates) CoordY() int {
	return c.y
}
func (c *coordinates) Coords() (int, int) {
	return c.x, c.y
}

type coords interface {
	CoordX() int
	CoordY() int
	Coords() (int, int)
}

type WorldHex struct {
	wx             int
	wy             int
	overallTerrain []int
	neiboirs       []coordinates
}

func presentOnMap(wm *worldmap, hex hexagon.Hexagon) bool {
	if hex.CoordX() < 0 || hex.CoordX() >= wm.worldX {
		return false
	}
	if hex.CoordY() < 0 || hex.CoordY() >= wm.worldY {
		return false
	}
	return true
}

func newWorldHex(x, y int, neib []coordinates) *WorldHex {
	wh := WorldHex{}
	wh.wx = x
	wh.wy = y
	wh.neiboirs = neib
	return &wh
}

type WorldCoords struct {
	x int
	y int
}

func New(world *world.World) *worldmap {
	wm := worldmap{}
	size := world.Profile().Data(profile.KEY_SIZE).Value()

	wm.WorldHex = make(map[coordinates]*WorldHex)
	wm.WorldHex = newGrid(size)
	return &wm
}

func newGrid(size int) map[coordinates]*WorldHex {
	if size == 0 {
		return nil
	}
	grid := make(map[coordinates]*WorldHex)
	topRows := []int{}
	middleRows := []int{}
	bottomRows := []int{}
	row := 0
	for row <= size {
		topRows = append(topRows, row)
		row++
	}
	maxWidth := rowLen(topRows[len(topRows)-1])
	for row > size && row < size*2 {
		middleRows = append(middleRows, row)
		row++
	}
	for row >= size*2 && row < size*3+1 {
		bottomRows = append(bottomRows, row)
		row++
	}
	//fmt.Println(topRows, middleRows, bottomRows)
	switch size {
	case 1:
		grid[newCoords(0, 0)] = newWorldHex(0, 0, []coordinates{
			newCoords(0, 1), newCoords(1, 1), newCoords(2, 1), newCoords(3, 1), newCoords(4, 1)})
		for i := 0; i < 5; i++ {
			r := i + 1
			l := i - 1
			if r == 5 {
				r = 0
			}
			if l == -1 {
				l = 4
			}
			grid[newCoords(i, 1)] = newWorldHex(i, 1, []coordinates{
				newCoords(0, 0), newCoords(r, 1), newCoords(l, 1), newCoords(i, 2), newCoords(r, 2)})
		}
		for i := 0; i < 5; i++ {
			r := i + 1
			l := i - 1
			if r == 5 {
				r = 0
			}
			if l == -1 {
				l = 4
			}
			grid[newCoords(i, 2)] = newWorldHex(i, 2, []coordinates{
				newCoords(0, 3), newCoords(l, 1), newCoords(i, 1), newCoords(r, 2), newCoords(l, 2)})
		}
		grid[newCoords(0, 3)] = newWorldHex(0, 3, []coordinates{
			newCoords(0, 3), newCoords(1, 3), newCoords(2, 3), newCoords(3, 3), newCoords(4, 3)})

	// 	grid[newCoords(0, 0)] = newWorldHex(0, 0, []coordinates{
	// 		newCoords(0, 1), newCoords(1, 1), newCoords(2, 1), newCoords(3, 1), newCoords(4, 1)})
	// 	grid[newCoords(0, 1)] = newWorldHex(0, 1, []coordinates{
	// 		newCoords(0, 0), newCoords(1, 1), newCoords(0, 2), newCoords(1, 2), newCoords(4, 1)})
	// 	grid[newCoords(1, 1)] = newWorldHex(1, 1, []coordinates{
	// 		newCoords(0, 0), newCoords(0, 1), newCoords(2, 1), newCoords(0, 2), newCoords(2, 2)})
	// 	grid[newCoords(2, 1)] = newWorldHex(2, 1, []coordinates{
	// 		newCoords(0, 0), newCoords(1, 1), newCoords(3, 1), newCoords(1, 2), newCoords(2, 2)})
	// 	grid[newCoords(3, 1)] = newWorldHex(3, 1, []coordinates{
	// 		newCoords(0, 0), newCoords(2, 1), newCoords(4, 1), newCoords(2, 2), newCoords(3, 2)})
	default:
		for i := 0; i <= 3*size; i++ {
			width := rowLen(i)
			if maxWidth < width {
				width = maxWidth

			}
			if i > 2*size-1 {
				width = rowLen((3 * size) - i)
			}
			for x := 0; x < width; x++ {
				coord := newCoords(x, i)
				nb := defineNeibours(coord, topRows, middleRows, bottomRows)
				//fmt.Println("+", nb, coord, "------", len(nb))
				grid[coord] = newWorldHex(coord.x, coord.y, nb)
			}

		}

	}
	return grid
}

func addNeibbour(neib []coordinates, coords coordinates, newNeib coordinates) []coordinates {

	neib = append(neib, newNeib)
	return neib
}

func defineNeibours(coord coordinates, top, mid, bot []int) []coordinates {
	upperRow := defineRow(coord.y-1, top, mid, bot)
	currentRow := defineRow(coord.y, top, mid, bot)
	lowerRow := defineRow(coord.y+1, top, mid, bot)
	neib := []coordinates{}
	reverseTop := []int{}
	for _, v := range top {
		reverseTop = append(reverseTop, v)
	}
	allWid := append([]int{}, top...)
	for i := range mid {
		k := i + 2
		k++
		allWid = append(allWid, mid[0])
	}

	for i, j := 0, len(reverseTop)-1; i < j; i, j = i+1, j-1 {
		reverseTop[i], reverseTop[j] = reverseTop[j], reverseTop[i]
	}
	for i := range reverseTop {
		reverseTop[i] = reverseTop[i] * -1
	}
	allWid = append(allWid, reverseTop...)
	maxWidth := top[len(top)-1] * 5
	// track := false
	// if coord.x == 2 && coord.y == 5 {
	// 	fmt.Println(allWid, upperRow, currentRow, lowerRow, coord)
	// 	track = true
	// }
	// if coord.x == 3 && coord.y == 5 {
	// 	fmt.Println(allWid, upperRow, currentRow, lowerRow, coord)
	// 	track = true
	// }
	switch currentRow {
	case "NP":
		for i := 0; i < 5; i++ {
			neib = append(neib, newCoords(i, coord.y+1))
		}
	case "top":
		isNode := false
		for _, nodeVal := range rowNodes(coord.y) {
			if nodeVal == coord.x {
				isNode = true
			}
		}
		switch upperRow {
		case "NP":
			neib = append(neib, newCoords(0, 0))
		case "top":
			switch isNode {
			case true:
				thisNode := -1
				for i, n := range rowNodes(coord.y) {
					if n == coord.x {
						thisNode = i
						break
					}
				}
				want := rowNodes(coord.y - 1)[thisNode]
				neib = append(neib, newCoords(want, coord.y-1))
			case false:
				nodes := rowNodes(coord.y)
				maxOffset := nodes[1] - nodes[0]
				thisNode := -1
				thisOffset := -1
				for n, node := range nodes {
					for of := 0; of < maxOffset; of++ {
						if node+of == coord.x {
							thisNode = n
							thisOffset = of
						}
					}
				}
				if thisNode == -1 {
					thisNode = 4
					thisOffset = maxOffset/2 + coord.x
				}
				want := rowNodes(coord.y - 1)[thisNode] + thisOffset

				want2 := -999
				switch {

				case coord.y%2 == 1:
					// if track {
					// 	fmt.Println("ROW odd", coord.y)
					// }
					want--
					want2 = want + 1
				case coord.y%2 == 0:
					want2 = want - 1
					// if track {
					// 	fmt.Println("ROW even", coord.y)
					// }
				}
				if want > rowLen(coord.y-1)-1 {
					want -= rowLen(coord.y - 1)
				}
				if want < 0 {
					want += rowLen(coord.y - 1)
				}
				if want2 > rowLen(coord.y-1)-1 {
					want2 -= rowLen(coord.y - 1)
				}
				if want2 < 0 {
					want2 += rowLen(coord.y - 1)
				}

				neib = append(neib, newCoords(want, coord.y-1))
				neib = append(neib, newCoords(want2, coord.y-1))
			}
		}
		///////////////
		rX := coord.x + 1
		if rX > rowLen(coord.y)-1 {
			rX -= rowLen(coord.y)
		}
		lX := coord.x - 1
		if lX < 0 {
			lX += rowLen(coord.y)
		}
		neib = append(neib, newCoords(rX, coord.y))
		neib = append(neib, newCoords(lX, coord.y))

		///////////////////
		switch lowerRow {
		case "top":
			switch isNode {
			case true:
				thisNode := -1
				for i, n := range rowNodes(coord.y) {
					if n == coord.x {
						thisNode = i
						break
					}
				}
				want0 := rowNodes(coord.y + 1)[thisNode]
				for _, off := range []int{-1, 0, 1} {
					want := want0 + off
					if want < 0 {
						want += rowLen(coord.y + 1)
					}
					if want > rowLen(coord.y+1)-1 {
						want -= rowLen(coord.y + 1)
					}
					neib = append(neib, newCoords(want, coord.y+1))

					//neib = append(neib, newCoords(want+1, coord.y+1))
				}

			case false:
				nodes := rowNodes(allWid[coord.y])
				maxOffset := nodes[1] - nodes[0]
				thisNode := -1
				thisOffset := -1
				for n, node := range nodes {
					for of := 0; of < maxOffset; of++ {
						if node+of == coord.x {
							thisNode = n
							thisOffset = of
						}
					}
				}
				if thisNode == -1 {
					thisNode = 4
					thisOffset = maxOffset/2 + coord.x
				}
				lowerNode := rowNodes(allWid[coord.y+1])
				want := lowerNode[thisNode] + thisOffset
				if want < 0 {
					want += rowLen(allWid[coord.y+1])
				}
				if want > rowLen(allWid[coord.y+1])-1 {
					want -= rowLen(allWid[coord.y+1])
				}
				want2 := want + 1
				if want2 > rowLen(allWid[coord.y+1])-1 {
					want2 -= rowLen(allWid[coord.y+1])
				}

				neib = append(neib, newCoords(want, coord.y+1))
				neib = append(neib, newCoords(want2, coord.y+1))
				// for _, add := range []int{0, 1} {
				// 	want := coord.x + offset + add
				// 	if want < 0 {
				// 		want += rowLen(coord.y + 1)
				// 	}
				// 	if want > rowLen(coord.y+1)-1 {
				// 		want -= rowLen(coord.y + 1)
				// 	}
				// 	if track {
				// 		fmt.Println("ADD BOT", newCoords(want, coord.y+1))
				// 	}
				// 	neib = append(neib, newCoords(want, coord.y+1))
				// }
			}
		case "mid":
			for _, offset := range []int{0, -1} {
				want := coord.x + offset
				if want < 0 {
					want += rowLen(coord.y)
				}
				neib = append(neib, newCoords(want, coord.y+1))
			}
		}
	case "mid":
		for _, n := range []int{-1, 0, 1} {
			rX := coord.x + 1

			if rX > (maxWidth)-1 { //    rowLen(top[len(top)])-1 {
				rX -= maxWidth // rowLen(top[len(top)])
			}
			lX := coord.x
			if n == 0 {
				lX--
			}

			if lX < 0 {
				lX += maxWidth //rowLen(top[len(top)])
			}

			neib = append(neib, newCoords(rX, coord.y+n))
			neib = append(neib, newCoords(lX, coord.y+n))
		}
	case "bot":
		isNode := false
		for _, nodeVal := range rowNodes(allWid[coord.y]) {
			if nodeVal == coord.x {
				isNode = true
			}
		}
		switch upperRow {
		case "mid":
			urX := coord.x
			ulX := coord.x - 1
			if ulX < 0 {
				ulX += maxWidth
			}
			neib = append(neib, newCoords(urX, coord.y-1))
			neib = append(neib, newCoords(ulX, coord.y-1))

			///////////////

		case "bot":
			nodes := rowNodes(allWid[coord.y])
			maxOffset := nodes[1] - nodes[0]
			thisNode := -1
			thisOffset := -1
			for n, node := range nodes {
				for of := 0; of < maxOffset; of++ {
					if node+of == coord.x {
						thisNode = n
						thisOffset = of
					}
				}
			}
			upperNodes := rowNodes(allWid[coord.y-1])
			switch isNode {
			case true:
				want2 := upperNodes[thisNode]
				want3 := want2 + 1
				want := want2 - 1
				if want < 0 {
					want += rowLen(allWid[coord.y-1])
				}
				if want3 > rowLen(allWid[coord.y-1])-1 {
					want3 -= rowLen(allWid[coord.y-1])
				}
				neib = append(neib, newCoords(want, coord.y-1))
				neib = append(neib, newCoords(want2, coord.y-1))
				neib = append(neib, newCoords(want3, coord.y-1))
			case false:
				//fmt.Print("top = 2")
				want := upperNodes[thisNode] + thisOffset
				want2 := want + 1
				if want2 > rowLen(allWid[coord.y-1])-1 {
					want2 -= rowLen(allWid[coord.y-1])
				}
				neib = append(neib, newCoords(want, coord.y-1))
				neib = append(neib, newCoords(want2, coord.y-1))
			}
		}

		switch lowerRow {
		case "SP":
			//fmt.Print(" bot = 1")
			neib = append(neib, newCoords(0, coord.y+1))

		case "bot":
			switch isNode {
			case true:
				thisNode := -1
				for i, n := range rowNodes(allWid[coord.y]) {
					if n == coord.x {
						thisNode = i
						break
					}
				}
				want := rowNodes(allWid[coord.y+1])[thisNode]
				neib = append(neib, newCoords(want, coord.y+1))
			case false:
				nodes := rowNodes(allWid[coord.y])
				maxOffset := nodes[1] - nodes[0]
				thisNode := -1
				thisOffset := -1
				for n, node := range nodes {
					for of := 0; of < maxOffset; of++ {
						if node+of == coord.x {
							thisNode = n
							thisOffset = of
						}
					}
				}
				lowerNode := rowNodes(allWid[coord.y+1])

				want := lowerNode[thisNode] + thisOffset
				if want > rowLen(allWid[coord.y+1])-1 {
					want -= rowLen(allWid[coord.y+1])
				}
				if want < 0 {
					want += rowLen(allWid[coord.y+1])
				}
				if thisNode == -1 {
					thisNode = 4
					thisOffset = maxOffset/2 + coord.x
				}
				if want > rowLen(allWid[coord.y+1])-1 {
					want -= rowLen(allWid[coord.y+1])
				}
				if want < 0 {
					want += rowLen(allWid[coord.y+1])
				}
				want2 := want - 1
				if want2 < 0 {
					want2 += rowLen(allWid[coord.y+1])
				}

				neib = append(neib, newCoords(want, coord.y+1))
				neib = append(neib, newCoords(want2, coord.y+1))
			}
		}
		rX := coord.x + 1
		if rX > rowLen(allWid[coord.y])-1 {
			rX -= rowLen(allWid[coord.y])
		}
		lX := coord.x - 1
		if lX < 0 {
			lX += rowLen(allWid[coord.y])
		}
		neib = append(neib, newCoords(rX, coord.y))
		neib = append(neib, newCoords(lX, coord.y))

	case "SP":
		for i := 0; i < 5; i++ {
			neib = append(neib, newCoords(i, coord.y-1))
		}
	}
	// if track {
	// 	fmt.Println(neib)
	// 	fmt.Println("------")
	// }

	return neib
}

// func mapTopNeibB(currentWidth int) map[int]int {
// 	width := currentWidth * 5
// 	nMap := make(map[int]int)
// 	nodes := []int{}
// 	for i := 0; i < 5; i++ {
// 		nodes = append(nodes, currentWidth*i)
// 	}
// 	off := -1

// 	for i := 0; i < width; i++ {
// 		node := false
// 		if i/currentWidth != off || i == 0 {
// 			off = i / currentWidth
// 			node = true
// 		}
// 		nMap[i] = i + off
// 		if node {
// 			nMap[i] = nMap[i] * -1
// 		}
// 	}
// 	return nMap
// }

// func mapTopNeibC(currentWidth int) map[int]int {
// 	width := currentWidth * 5
// 	nMap := make(map[int]int)
// 	nodes := []int{}
// 	for i := 0; i < 5; i++ {
// 		nodes = append(nodes, currentWidth*i)
// 	}
// 	off := -1

// 	for i := 0; i < width; i++ {
// 		node := false
// 		if i/currentWidth != off || i == 0 {
// 			off = i / currentWidth
// 			node = true
// 		}
// 		nMap[i] = i - off
// 		if node {
// 			nMap[i] = nMap[i] * -1
// 		}
// 	}
// 	return nMap
// }

func reverseMap(originalMap map[int]int) map[int]int {
	newmap := make(map[int]int)
	for k, v := range originalMap {
		newmap[v] = k
	}
	return newmap
}

func defineRow(row int, top, mid, bot []int) string {
	if row == top[0] {
		return "NP"
	}
	if row < 0 {
		return "nNP"
	}
	if row == bot[len(bot)-1] {
		return "SP"
	}
	if row > bot[len(bot)-1] {
		return "nSP"
	}
	for _, v := range top {
		if v == row {
			return "top"
		}
	}
	for _, v := range mid {
		if v == row {
			return "mid"
		}
	}
	for _, v := range bot {
		if v == row {
			return "bot"
		}
	}
	return "err"
}

func rowNodes(r int) []int {
	nodes := []int{}
	switch r <= 0 {
	case false:
		nodes = append([]int{}, r/2)
		nodes = append(nodes, nodes[0]+r, nodes[0]+r+r, nodes[0]+r+r+r, nodes[0]+r+r+r+r)
	case true:
		r = r * -1
		nodes = append(nodes, 0)
		nodes = append(nodes, nodes[0]+r, nodes[0]+r+r, nodes[0]+r+r+r, nodes[0]+r+r+r+r)
	}

	return nodes
}

func rowLen(r int) int {
	switch r {
	case 0:
		return 1
	default:
		if r < 0 {
			r = r * -1
		}
		return r * 5
	}
}

// func mapTopNeib(current int) map[int]int {
// 	nMap := make(map[int]int)
// 	switch current {
// 	case 1:
// 		for i := 0; i < rowLen(current); i++ {
// 			nMap[i] = 0
// 		}
// 		return nMap
// 	case 2:
// 		for i := 0; i < rowLen(current); i++ {
// 			n := i / 2
// 			for n >= 5 {
// 				n -= 5
// 			}
// 			nMap[i] = n
// 		}
// 	default:
// 		topRow := []int{}
// 		for i := 0; i < rowLen(current-1); i++ {
// 			topRow = append(topRow, -1)
// 		}
// 		topNodes := rowNodes(current - 1)
// 		//topRow[len(topRow)-1] = 0
// 		thisNodes := rowNodes(current)
// 		for i := 0; i < rowLen(current); i++ {
// 			nodeNum, off := minNode(thisNodes, i)
// 			n := -999
// 			if nodeNum < 0 {
// 				nodeNum = 4
// 				off = (current/2 - i - 1) * -1
// 				switch current % 2 {
// 				case 0:
// 					off--
// 				case 1:
// 				}
// 				n = thisNodes[0] + off
// 			} else {
// 				n = topNodes[nodeNum] + off
// 			}
// 			if n >= rowLen(current-1) {
// 				n -= rowLen(current - 1)
// 			}
// 			nMap[i] = n
// 		}
// 	}
// 	return nMap
// }

// func minNode(nodes []int, x int) (nodeNum int, offset int) {
// 	nodeNum = -1
// 	for i, n := range nodes {
// 		if x < n {
// 			continue
// 		}
// 		nodeNum = i
// 		offset = x - n
// 	}
// 	if nodeNum == -1 {
// 		nm := nodes[0] - (nodes[1] - nodes[0])
// 		offset = x + nm
// 	}
// 	return nodeNum, offset
// }

/*
GENERATING THE WORLD MAP
As instructed, mark the specific terrain type in directed World Hexes on the World Map. Within the limits of the instructions, terrain may be placed in any available World Hex. If the specific world is too small for the terrain called for, restrict the number placed to what the specific hex will contain.
1. Select a blank World Map based on World Size.
	//2. Resources. Determine Resources from the Economic Extension. Subtract system GG and Belts: place the resulting number of Resource Hexes one per Triangle.
3. Mountains. Place 1D Mountains in each Triangle.
4. Chasms. Place World Size x Chasms Sets (1D per Triangle).
5. Precipices. Place World Size x Precipices one per Triangle.
If Di
6. Die-Back. Place 1D Ruins in each Triangle.
If Va
7. Vacuum Plain. Place Craters (1D per Triangle).
If De
8. Desert. Mark all unmarked hexes Desert.
9. Oceans.
Randomly select Hyd x 2 Triangles as Oceans. Consolidate Ocean Triangles that share sides. Enclose Oceans with Shore lines (which may run through any type terrain). Non-Ocean Triangles are Continents (they are not consolidated; treat each Triangle as a separate Continent).
10. Seas. Randomly select Hyd Continents and place a one-hex Ocean (Sea) in each.
Surround each with Shore in all adjacent hexes.
11. Islands. Convert each Mountain Hex in Ocean to Islands.
12. Ice-Caps. If HZ or greater, mark the top and bottom Hyd/2 rows as Ice Cap (if Hyd less than 2, no Ice Caps).
If Ic
13. More Ice Cap. Add 1D rows to each Ice Cap.
If Fr
14. Frozen. Mark Ocean as Ice Field and Land as Frozen Lands (except under Ice Cap).
If Tu
15. Tundra. Mark a line 1D hexes from each Pole. Between each line and its Pole, mark  Ocean as Ice Field and Land as Frozen lands (except under Ice Cap).
If Ag
16. Agricultural. Place 2D Cropland in each Continent.
If Fa
17. Farming. Place 1D Cropland in each Continent.
If Lo
18. Low Population. Place one Town. Skip to 22.
If Ni
19. Non-Industrial. Place one Town. Skip to 22.
20. Cities. Place Cities equal to Pop, one per Continent.
If Atm=0-1, A-C, or E+ = Domed if not NIL.
If Hi
21. High Population. Place total Pop/2 Arcologies.
22. Rural. Mark clear hexes within Pop hexes of City as Rural.
23. Starport. Place the World Starport (or Spaceport).
If Tz
24. Create A Twilight Zone. Select one Pole Triangle and draw a vertical line directly down. Shift 2.5 times World
Size hexes to one side and draw a parallel line: this is the one-World-Hex-wide Twilight Zone.
If Tz
25. Create Two Hemispheres For A Twilight World. Mark one side of the Twilight Zone as Baked Lands and the
other side as Frozen Lands (overlaying existing terrain). Terrain in the Twilight Zone remains as previously
created. Convert Ocean in Baked Lands to Desert. Convert Ocean in Frozen Lands to Ice Field.
If Pe
26. Penal Colony. Mark Pop x Penal (one per Triangle).
27. Wasteland. If TL>5, mark 1D adjacent hexes in one Triangle Wasteland.
28. Exotic. Place one Exotic hex in one Triangle.
29. Noble Lands. Place one Noble Lands estate.
30. All other terrain remains Clear.
*/
