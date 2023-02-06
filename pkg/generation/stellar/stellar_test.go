package stellar

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func TestStellar(t *testing.T) {
	return
	for i := 480; i < 1000; i++ {
		dp := dice.New().SetSeed(i)
		stellar := GenerateStellar(dp)
		stars := Parse(stellar)
		fmt.Printf("stellar: '%v' [%v]\n", stellar, strings.Join(stars, "|"))
		if stellar != strings.Join(stars, " ") {
			t.Errorf("not merging := %v", fmt.Sprintf("stellar: '%v' [%v]\n", stellar, strings.Join(stars, " ")))
		}
	}

}

func TestHabitableOrbits(t *testing.T) {
	return
	for i, code := range listAllStars() {
		hz := HabitableOrbitByCode(code)
		fmt.Printf("%v	%v: %v\n", i, code, hz)
	}
}

type nsInput struct {
	typ string
	dec int
	cls string
}

func newStarInputsCORRECT() []nsInput {
	nsi := nsInput{}
	nsiArr := []nsInput{}
	for _, sType := range []string{"O", "B", "A", "F", "G", "K", "M", "L", "T", "Y", "BD"} {
		nsi.typ = sType
		for _, dc := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
			nsi.dec = dc
			for _, cl := range []string{"Ia", "Ib", "II", "III", "IV", "V", "VI", "D"} {
				nsi.cls = cl
				if sType == "A" && cl == "VI" {
					cl = "V"
				}

				nsiArr = append(nsiArr, nsi)
			}
		}
	}
	return nsiArr
}

func TestNewStar(t *testing.T) {
	return
	errorMap := make(map[string]int)
	for _, input := range newStarInputsCORRECT() {
		s, err := NewStar(input.typ, input.dec, input.cls)
		if err != nil {
			errorMap[err.Error()]++
			//t.Errorf("input %v (%v) = returned error: %v", i, input, err.Error())
		}
		if s.mass <= 0 {
			errorMap["Mass <= 0"]++
			fmt.Println(s.star)
		}
	}
	for k, v := range errorMap {
		fmt.Println(v, k)
	}
}
