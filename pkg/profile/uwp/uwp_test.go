package uwp

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/internal/ehex"
)

func TestSetUWP(t *testing.T) {
	return
	testAspects := []string{Port, Size, Atmo, Hydr, Pops, "wrongAspect", Govr, Laws, TL}
	testVals := []int{}
	for i := -1; i < 35; i++ {
		testVals = append(testVals, i)
	}
	uwpS := New()
	testsCompleted := 0
	testsSuccsseded := 0
	errorsDetected := 0
	for _, asp := range testAspects {
		for _, val := range testVals {
			err := uwpS.Encode(asp, ehex.New().Set(val))
			switch {
			default:
			case err != nil:
				errorsDetected++
				fmt.Printf("Tests %v: input (%v - %v) Error: '%v'\n", testsCompleted+1, asp, val, err.Error())
			case err == nil:
				testsSuccsseded++
				fmt.Printf("Tests %v: input (%v - %v) SUCCESS!\n", testsCompleted+1, asp, val)
			}
			testsCompleted++

		}
	}
	if testsCompleted != testsSuccsseded+errorsDetected {
		t.Errorf("tests not validated")
	}
}

func TestUWPcall(t *testing.T) {
	return
	uwpS := New()
	uwpS.Encode(Port, ehex.New().Set(10))
	uwpS.Encode(Size, ehex.New().Set(2))
	uwpS.Encode(Atmo, ehex.New().Set(2))
	uwpS.Encode(Hydr, ehex.New().Set(0))
	uwpS.Encode(Pops, ehex.New().Set(6))
	uwpS.Encode(Govr, ehex.New().Set(5))
	uwpS.Encode(Laws, ehex.New().Set(16))
	uwpS.Encode(TL, ehex.New().Set(8))
	fmt.Println(uwpS)

}

func TestUWPinput(t *testing.T) {
	return
	u, err := FromString("X620000-0")
	fmt.Println(err)
	fmt.Println(u.Describe("All"))
	fmt.Println(u.Describe(Port))
}

func TestUWPfromString(t *testing.T) {
	u, err := FromString("AAAAAAA-A")
	fmt.Println(u)
	fmt.Println(err)
}
