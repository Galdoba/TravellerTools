package main

import (
	"github.com/Galdoba/TravellerTools/pkg/decidion"
)

func main() {
	options := []string{"A", "B", "C", "D", "E", "F"}
	label := "label"
	decidion.Manual_Few_Exclude(3, label, options...)
}
