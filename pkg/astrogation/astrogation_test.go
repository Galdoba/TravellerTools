package astrogation

import (
	"fmt"
	"testing"
)

func inputCube() [][]int {
	return [][]int{
		{0, 0},
		{0, 5},
		{0, -1},
		{1, 0},
		{1, -2},
		{1, 2},
		{-1, 0},
		{-1, -2},
		{-1, 2},
	}
}

func Test_Coordinates(t *testing.T) {
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
		for _, val := range cubeDirectionVectors {
			fmt.Println(cubeNeighbor(cc, val))
		}
		fmt.Println("=================================")

	}

}

func cubeSum(c cubeCoords) int {
	return c.q + c.r + c.s
}
