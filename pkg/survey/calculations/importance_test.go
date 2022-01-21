package calculations

import (
	"fmt"
	"testing"
)

type dataSet struct {
	data []string
}

func extrapolateInput() []dataSet {
	var allData []dataSet
	var data []string
	for _, stprt := range []string{"A", "B", "C", "D", "E", "X", "Y", "F", "G", "H"} {
		for _, tl := range []string{"7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H"} {
			for _, pops := range []string{"4", "5", "6", "7", "8"} {
				for _, bs := range []string{"KV", "KW", "W", "NS", "R"} {
					for _, rem := range []string{"Ba Lo Ni", "Ag Ni", "Ag Ni Ri", "Hi In Na Po"} {
						data = []string{}
						uwp := stprt + "SAH" + pops + "GL-" + tl
						data = append(data, uwp)
						data = append(data, bs)
						data = append(data, rem)
						allData = append(allData, dataSet{data})
					}
				}
			}
		}
	}
	return allData
}

func TestImportance(t *testing.T) {
	impMap := make(map[int]int)
	for _, test := range extrapolateInput() {
		im := Importance(test.data[0], test.data[1], test.data[2])
		fmt.Println(test)
		impMap[im]++
	}
	for k, v := range impMap {
		fmt.Println("imp", k, v)
	}

}
