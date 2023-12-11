package tradecodes

import (
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

func ParseTradeCodes(data map[string]ehex.Ehex) string {
	tc := ""
	size := data["size"].Value()
	atmo := data["atmo"].Value()
	hydr := data["hydr"].Value()
	pops := data["pops"].Value()
	if inBounds(atmo, 4, 9) && inBounds(hydr, 4, 8) && inBounds(pops, 5, 7) {
		tc += " Ag"
	}
	if inBounds(size, 0, 0) && inBounds(atmo, 0, 0) && inBounds(hydr, 0, 0) {
		tc += " As"
	}
	if inBounds(atmo, 2, 16) && inBounds(hydr, 0, 0) {
		tc += " De"
	}
	if inBounds(atmo, 10, 16) && inBounds(hydr, 1, 99) {
		tc += " Fl"
	}
	if inBounds(hydr, 4, 9) && inBounds(pops, 4, 8) {
		switch atmo {
		case 5, 6, 8:
			tc += " Ga"
		}
	}
	if inBounds(atmo, 0, 1) && inBounds(hydr, 1, 99) {
		tc += " Ic"
	}
	if inBounds(pops, 9, 99) {
		switch atmo {
		case 0, 1, 2, 4, 7, 9:
			tc += " In"
		}

	}
	if inBounds(atmo, 0, 3) && inBounds(hydr, 0, 3) && inBounds(pops, 6, 99) {
		tc += " Na"
	}
	if inBounds(pops, 1, 6) {
		tc += " Ni"
	}
	if inBounds(atmo, 2, 5) && inBounds(hydr, 0, 3) {
		tc += " Po"
	}
	if inBounds(pops, 6, 8) {
		switch atmo {
		case 6, 8:
			tc += " Ri"
		}

	}
	if inBounds(hydr, 10, 10) {
		tc += " Wa"
	}
	if inBounds(atmo, 0, 0) && !strings.Contains(tc, " As") {
		tc += " Va"
	}
	return tc
}

func inBounds(i, min, max int) bool {
	if i >= min && i <= max {
		return true
	}
	return false
}
