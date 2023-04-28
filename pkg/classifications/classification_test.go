package classifications

import (
	"fmt"
	"strings"
	"testing"
)

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
// )

// func TestSimple(t *testing.T) {
// 	st := []string{"A", "X"}                                                             //, "C", "D", "E", "B"}
// 	sz := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B"}           //, "C", "D", "E", "F"}
// 	at := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D"} //, "E", "F"}
// 	hd := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A"}
// 	pp := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B"}           //, "C", "D", "E", "F"}
// 	gv := []string{"0", "1", "2", "3", "4", "5", "6", "7"}                               //, "8", "9", "A", "B", "C", "D", "E", "F"}
// 	lw := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A"}                //, "B", "C", "D", "E", "F"}
// 	tl := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D"} //, "E", "F"}
// 	res := ""
// 	i := 1
// 	maxTest := len(st) * len(sz) * len(at) * len(hd) * len(pp) * len(gv) * len(lw) * len(tl)
// 	for _, val1 := range st {
// 		for _, val2 := range sz {
// 			for _, val3 := range at {
// 				for _, val4 := range hd {
// 					for _, val5 := range pp {
// 						for _, val6 := range gv {
// 							for _, val7 := range lw {
// 								for _, val8 := range tl {
// 									res = fmt.Sprintf("Test %v/%v:", i, maxTest)
// 									uwps := val1 + val2 + val3 + val4 + val5 + val6 + val7 + "-" + val8
// 									u, err := uwp.FromString(uwps)
// 									if err != nil {
// 										t.Errorf("bad input '%v'", uwps)
// 									}
// 									tc, err := FromUWP(u)
// 									if err != nil {
// 										t.Errorf("bad output '%v'", uwps)
// 									}
// 									fmt.Print(uwps, " :")
// 									for _, t := range tc {
// 										res = res + t.code + " "
// 									}
// 									fmt.Print(res + "                       \r")
// 									i++
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }

func TestRequirements(t *testing.T) {
	fmt.Println(strings.Contains("", "A"))
	fmt.Println(strings.Contains("A", "A"))
	fmt.Println(strings.Contains("AB", "A"))
	fmt.Println(strings.Contains("A", "BA"))
}
