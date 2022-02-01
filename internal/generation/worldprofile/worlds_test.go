package worldprofile

import (
	"fmt"
	"testing"
)

func TestNums(t *testing.T) {
	drArr := []int{1, 2, 3, 4, 5, 6}
	try1 := 0
	port1 := ""
	dm := 0
	portMap := make(map[string]int)
	for _, d1 := range drArr {
		for _, d2 := range drArr {
			for _, d3 := range drArr {
				for _, d4 := range drArr {
					try1++
					pp := d1 + d2 - 2
					switch pp {
					case 8, 9:
						dm = 1
					case 3, 4:
						dm = -1
					case 10:
						dm = 2
					case 0, 1, 2:
						dm = -2
					}

					switch d3 + d4 + dm {
					case 3, 4:
						port1 = "E"
					case 5, 6:
						port1 = "D"
					case 7, 8:
						port1 = "C"
					case 9, 10:
						port1 = "B"
					}
					if d3+d4+dm > 10 {
						port1 = "A"
					}
					if d3+d4+dm < 3 {
						port1 = "X"
					}
					portMap[port1]++
				}
			}
		}
	}
	fmt.Println(portMap)
}
