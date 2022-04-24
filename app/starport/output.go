package main

import (
	"fmt"
)

const (
	DEFAULT_VALUE = iota
	TRAFFIC_PASSENGERS_DEPART
	TRAFFIC_PASSENGERS_ARRIVE
	TRAFFIC_FREIGHT_DEPART
	TRAFFIC_FREIGHT_ARRIVE
	WRONG_INSTRUCTION
)

type TrafficData struct {
	port        Port
	freightD    map[Port]freightInfo
	passengersD map[Port]passengerInfo
	freightA    map[Port]freightInfo
	passengersA map[Port]passengerInfo
	status      string
	neibhours   int
	reach       int
}

func NewTrafficData(source Port, reach int) *TrafficData {
	td := TrafficData{}
	td.port = source
	td.freightD = make(map[Port]freightInfo)
	td.passengersD = make(map[Port]passengerInfo)
	td.freightA = make(map[Port]freightInfo)
	td.passengersA = make(map[Port]passengerInfo)
	td.reach = reach
	return &td
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
	rep += fmt.Sprintf("There are %v worlds in %v parsecs radius. ", len(td.freightD), td.reach)
	//rep += fmt.Sprintf("%v serves average %v passengers and %v tons of freight cargo per week.\n", td.port.MW_Name(), td.sumOf(TRAFFIC_PASSENGERS_ARRIVE)+td.sumOf(TRAFFIC_PASSENGERS_DEPART),
	rep += fmt.Sprintf("%v serves average %v passengers per week. %v of them arriving and %v are departing from the world. \n", td.port.MW_Name(), td.sumOf(TRAFFIC_PASSENGERS_ARRIVE)+td.sumOf(TRAFFIC_PASSENGERS_DEPART),
		td.sumOf(TRAFFIC_PASSENGERS_ARRIVE), td.sumOf(TRAFFIC_PASSENGERS_DEPART))
	rep += fmt.Sprintf("Freight traffic is about %v dTons per week. %v of them comming from the outer space and %v is leaving the Port. \n", td.sumOf(TRAFFIC_FREIGHT_ARRIVE)+td.sumOf(TRAFFIC_FREIGHT_DEPART),
		td.sumOf(TRAFFIC_FREIGHT_ARRIVE), td.sumOf(TRAFFIC_FREIGHT_DEPART))
	return rep
}

func (td *TrafficData) sumOf(instr int) int {
	sum := 0
	switch instr {
	case TRAFFIC_FREIGHT_ARRIVE:
		for _, v := range td.freightA {
			sum += v.total
		}
	case TRAFFIC_FREIGHT_DEPART:
		for _, v := range td.freightD {
			sum += v.total
		}
	case TRAFFIC_PASSENGERS_ARRIVE:
		for _, v := range td.passengersA {
			sum += v.total
		}
	case TRAFFIC_PASSENGERS_DEPART:
		for _, v := range td.passengersD {
			sum += v.total
		}
	}
	return sum
}

/*RESULT DATA

World  : Regina (Spinward Marches 1910)
UWP    : A788899-C
TC/Rem : Ri Pa Ph An Cp (Green Zone)
---------------------------------------
Spaceport Traffic Report:
There are 20 worlds in 4 parsecs radius. [&WORLD_NAME] is not located on a Trade Route.


*/
