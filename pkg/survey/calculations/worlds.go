package calculations

import (
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

func WorldsValid(worlds int, pbg string) bool {

	hex := strings.Split(pbg, "")
	b := ehex.New().Set(hex[1]).Value()
	g := ehex.New().Set(hex[2]).Value()
	if worlds < 1+b+g+2 {
		return false
	}
	if worlds > 1+b+g+12 {
		return false
	}
	return true
}

func FixWorlds(pbg, seed string) int {
	hex := strings.Split(pbg, "")
	b := ehex.New().Set(hex[1]).Value()
	g := ehex.New().Set(hex[2]).Value()
	d := dice.New().SetSeed(seed)
	return 1 + b + g + d.Roll("2d6").Sum()
}
