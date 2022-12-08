package location

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/astrogation/charts/otu/sectors"
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
)

const (
	COORDINATE_STANDARD_OTU = "OTU"
)

type location struct {
	standard      string
	sectorX       int
	sectorY       int
	sectrorName   string
	sectorAbb     string
	sectorHex     string
	subsectorCode string
	presentation  string
}

func New(hx hexagon.Hexagon, st string) location {
	loc := location{}
	loc.standard = st
	switch loc.standard {
	default:
		loc.presentation = "Unknown"
	case COORDINATE_STANDARD_OTU:
		lcX, lcY, sX, sY := recalculateToUTUlocal(hx.CoordX(), hx.CoordY())
		loc.sectorHex = sectorHex(lcX, lcY)
		loc.sectorX = sX
		loc.sectorY = sY
		loc.sectrorName, loc.sectorAbb = sectors.NameAbb(loc.sectorX, loc.sectorY)
		switch loc.sectorAbb {
		case "Unch":
			loc.presentation = loc.sectrorName + "-" + loc.sectorHex
		default:
			loc.presentation = loc.sectorAbb + " " + loc.sectorHex
		}
	}
	return loc
}

func (loc *location) String() string {
	return loc.presentation
}

func recalculateToUTUlocal(glX, glY int) (lcX, lcY, sX, sY int) {
	lcX = glX + 1
	lcY = glY + 40
	for lcX > 32 {
		lcX -= 32
		sX++
	}
	for lcX < 1 {
		lcX += 32
		sX--
	}
	for lcY > 40 {
		lcY -= 40
		sY++
	}
	for lcY < 1 {
		lcY += 40
		sY--
	}
	return lcX, lcY, sX, sY
}

func sectorHex(x, y int) string {
	s := ""
	if x < 10 {
		s += "0"
	}
	s += fmt.Sprintf("%v", x)
	if y < 10 {
		s += "0"
	}
	s += fmt.Sprintf("%v", y)
	return s
}

/*
0 0 => 1 40
3 0 => 4 40
3 -1=> 4 39

lcX = GlX + 1
lcY = GlY + 40
+----------+
|s p tw    |
|          |
|          |
|   xxxx   |
+----------+

*/
