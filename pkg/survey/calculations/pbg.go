package calculations

import (
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/internal/dice"
)

func PBGvalid(pbg string, uwp string) bool {
	hex := strings.Split(pbg, "")
	if len(hex) != 3 {
		return false
	}
	if !popDigValid(hex[0], uwp) {
		return false
	}
	if !beltValid(hex[1]) {
		return false
	}
	if !ggValid(hex[2]) {
		return false
	}

	return true
}

func FixPBG(pbg, uwp, seed string) string {
	hex := strings.Split(pbg, "")
	for len(hex) < 3 {
		hex = append(hex, "")
	}
	d := dice.New().SetSeed(seed)
	if !popDigValid(hex[0], uwp) {
		hex[0] = d.Roll("1d9").SumStr()
		u := strings.Split(uwp, "")
		if u[4] == "0" {
			hex[0] = "0"
		}
	}
	if !beltValid(hex[1]) {
		b := d.Roll("1d6").Sum() - 6
		if b < 0 {
			b = 0
		}
		hex[1] = strconv.Itoa(b)
	}
	if !ggValid(hex[2]) {
		gg := d.Roll("2d6").Sum()/2 - 2
		if gg < 0 {
			gg = 0
		}
		hex[2] = strconv.Itoa(gg)
	}
	return hex[0] + hex[1] + hex[2]
}

func popDigValid(pg, uwp string) bool {
	pop := strings.Split(uwp, "")[4]
	if pg == "0" && pop != "0" {
		return true
	}
	switch pg {
	default:
		return false
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
	}
	return true
}

func beltValid(pg string) bool {
	switch pg {
	default:
		return false
	case "0", "1", "2", "3":
	}
	return true
}

func ggValid(pg string) bool {
	switch pg {
	default:
		return false
	case "0", "1", "2", "3", "4":
	}
	return true
}
