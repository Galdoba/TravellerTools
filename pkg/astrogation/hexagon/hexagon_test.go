package hexagon

import (
	"fmt"
	"testing"
)

type hexagonInput struct {
	feed   int
	values []int
}

func goHexagonInput() []hexagonInput {
	return []hexagonInput{
		{},
		//{defaultValue, []int{0}},
		{defaultValue, []int{0, 1}},
		{defaultValue, []int{0, 1, 0}},
		//{defaultValue, []int{0, 1, 0, 1}},
		//{wrongValue, []int{0}},
		{wrongValue, []int{0, 1}},
		{wrongValue, []int{0, 1, 0}},
		//{wrongValue, []int{0, 1, 0, 1}},
		//{Feed_HEX, []int{0}},
		{Feed_HEX, []int{0, 1}},
		{Feed_HEX, []int{0, 1, 0}},
		//{Feed_HEX, []int{0, 1, 0, 1}},
		//{Feed_CUBE, []int{0}},
		{Feed_CUBE, []int{0, 1}},
		{Feed_CUBE, []int{0, 1, 0}},
		//{Feed_CUBE, []int{0, 1, 0, 1}},
		{Feed_CUBE, []int{1, 1, 1}},
		{Feed_CUBE, []int{1, 0, -1}},
		{Feed_HEX, []int{1, 0}},
		{Feed_HEX, []int{23, -33}},
	}

}

func TestHexagon(t *testing.T) {
	for _, input := range goHexagonInput() {

		hxgn, err := New(input.feed, input.values...)

		if err != nil {
			if err.Error() == "feed value unreconised" {
				continue
			}
			t.Errorf("creation error: %v", err.Error())
		}

		fmt.Println(hxgn)
		if hxgn.cube.q+hxgn.cube.r+hxgn.cube.s != 0 {
			t.Errorf("input: %v cube sum error:\ncube coordinates invalid [%v %v %v] - sum must be = 0 (have %v)", input, hxgn.cube.q, hxgn.cube.r, hxgn.cube.s, hxgn.cube.q+hxgn.cube.r+hxgn.cube.s)
		}
	}
}

func inputCube() [][]int {
	return [][]int{
		{0, 0},
	}
}

func Test_Coordinates(t *testing.T) {
	t.Skip()
	for _, set := range inputCube() {
		if len(set) != 2 {
			t.Errorf("expect input [%v] to have 2 integers, have %v", set, len(set))
			continue
		}
		hc := hexCoords{set[0], set[1]}
		cube := hexToCube(hc)
		if cubeSum(cube) != 0 {
			t.Errorf("coords %v wrong: expect sum = 0, but have %v\ninput [%v]", cube, cubeSum(cube), set)
			continue
		}
		oddq := cubeToHex(cube)
		if oddq.row != hc.row {
			t.Errorf("conversion wrong row: %v != %v", hc, oddq)
		}
		if oddq.col != hc.col {
			t.Errorf("conversion wrong col: %v != %v", hc, oddq)
		}
		cc := hexToCube(oddq)
		if cubeSum(cube) != 0 {
			t.Errorf("back conversion %v wrong: expect sum = 0, but have %v\ninput [%v]", cc, cubeSum(cube), set)
		}
		if hc.col == 0 && hc.row == 0 {
			//continue
		}
		fmt.Println(cc, oddq)
		fmt.Println("---------------------------------")
		for _, dir := range []int{Direction_N, Direction_NE, Direction_SE, Direction_S, Direction_SW, Direction_NW} {
			neib := cubeNeighbor(cc, dir)
			fmt.Println(neib, cubeToHex(neib))
		}
		fmt.Println("RING:")
		for _, cb := range cubeRing(cc, 2) {
			fmt.Println(cb, cubeToHex(cb))
		}
		fmt.Println("Spiral:")
		for _, cb := range cubeSpiral(cc, 2) {
			fmt.Println(cb, cubeToHex(cb))
		}
		fmt.Println("=================================")

	}

}

func cubeSum(c cubeCoords) int {
	return c.q + c.r + c.s
}
