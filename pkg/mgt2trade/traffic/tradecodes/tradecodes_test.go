package tradecodes

import (
	"fmt"
	"testing"
)

func input() []string {
	uwps := []string{}
	i := 0
	for _, st := range []string{"A"} {
		for _, s := range []string{"0", "1", "6", "7", "8", "9"} {
			for _, a := range []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"} {
				for _, h := range []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A"} {
					for _, p := range []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B"} {
						for _, g := range []string{"0", "1", "4", "5", "6", "7", "8", "9", "A"} {
							for _, l := range []string{"0", "1", "2"} {
								for _, tl := range []string{"0", "1", "2", "3", "4", "5", "6", "7", "B", "C"} {
									i++
									uwp := fmt.Sprintf("%v%v%v%v%v%v%v-%v", st, s, a, h, p, g, l, tl)
									tc, err := FromUWPstr(uwp)
									fmt.Printf("test %v: input %v have trade codes %v error: %v\n", i, uwp, tc, err)
								}
							}
						}
					}
				}
			}
		}

	}
	return uwps
}

func Test_TC(t *testing.T) {
	input()
	uwp := "A01733E-4"
	tc, err := FromUWPstr(uwp)
	fmt.Printf("input %v have trade codes %v error: %v\n", uwp, tc, err)

}
