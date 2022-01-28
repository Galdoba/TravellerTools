package uwp

import (
	"fmt"
	"testing"
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
			err := uwpS.Set(asp, val)
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
	uwpS.Set(Port, 10)
	uwpS.Set(Size, 2)
	uwpS.Set(Atmo, 2)
	uwpS.Set(Hydr, 0)
	uwpS.Set(Pops, 6)
	uwpS.Set(Govr, 5)
	uwpS.Set(Laws, 16)
	uwpS.Set(TL, 8)
	fmt.Println(uwpS)
	for asp := range uwpS.aspect.Data {
		fmt.Println(uwpS.Describe(asp))
	}
	fmt.Println(uwpS.Describe("All"))
	fmt.Println(uwpS.Describe("error test"))
	fmt.Println(uwpS.Starport())
	fmt.Println(uwpS.Size())
	fmt.Println(uwpS.Atmo())
	fmt.Println(uwpS.Hydr())
	fmt.Println(uwpS.Pops())
	fmt.Println(uwpS.Govr())
	fmt.Println(uwpS.Laws())
	fmt.Println(uwpS.TL())
}

func TestUWPinput(t *testing.T) {
	uwpS := New()
	err := uwpS.SetString("X620000-0")
	fmt.Println(err)
	fmt.Println(uwpS.Describe("All"))
}
