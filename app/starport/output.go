package main

import (
	"fmt"
)

type TrafficData struct {
	port      Port
	status    string
	neibhours int
	reach     int
}

func (td *TrafficData) String() string {
	lenght := 0
	l1 := fmt.Sprintf("World  : %v (%v %v)", td.port.MW_Name(), td.port.Sector(), td.port.Hex())
	if lenght < len(l1) {
		lenght = len(l1)
	}
	l2 := fmt.Sprintf("UWP    : %v", td.port.MW_UWP())
	if lenght < len(l2) {
		lenght = len(l2)
	}
	zone := ""
	switch td.port.TravelZone() {
	default:
		zone = "Green Zone"
	case "A":
		zone = "Amber Zone"
	case "R":
		zone = "Red Zone"
	}
	l3 := fmt.Sprintf("TC/Rem : %v (%v)", TradeCodes(td.port), zone)
	if lenght < len(l3) {
		lenght = len(l3)
	}
	sep := ""
	for len(sep) < lenght {
		sep += "-"
	}
	rep := fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n", sep, l1, l2, l3, sep)
	rep += "Spaceport Traffic Report:\n"
	rep += fmt.Sprintf("There are %v worlds in %v parsecs radius. ", td.neibhours, td.reach)

	return rep
}

/*RESULT DATA

World  : Regina (Spinward Marches 1910)
UWP    : A788899-C
TC/Rem : Ri Pa Ph An Cp (Green Zone)
---------------------------------------
Spaceport Traffic Report:
There are 20 worlds in 4 parsecs radius. [&WORLD_NAME] is not located on a Trade Route.


*/
