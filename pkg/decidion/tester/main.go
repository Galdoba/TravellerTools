package main

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/decidion"
)

func main() {
	M1S := decidion.Manual_One_STR("To...", "be...", "or not to be")
	fmt.Println(M1S)

	a2, ns := decidion.Manual_One_Exclude_STR("CHOOSE", "A", "B", "C", "D")
	fmt.Println(a2)
	fmt.Println(ns)

	a3 := decidion.Manual_Few_STR(5, "label 3", "A", "B", "C", "D")
	fmt.Println(a3)

	a4, e4 := decidion.Manual_Few_Exclude_STR(3, "label 3", "A", "B", "C", "D")
	fmt.Println(a4)
	fmt.Println(e4)
}
