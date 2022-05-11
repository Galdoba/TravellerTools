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
		if hxgn == nil {
			t.Errorf("input: %v\n struct not created", input)
			continue
		}
		fmt.Println(hxgn)
		if hxgn.cube.q+hxgn.cube.r+hxgn.cube.s != 0 {
			t.Errorf("input: %v cube sum error:\ncube coordinates invalid [%v %v %v] - sum must be = 0 (have %v)", input, hxgn.cube.q, hxgn.cube.r, hxgn.cube.s, hxgn.cube.q+hxgn.cube.r+hxgn.cube.s)
		}
	}
}
