package main

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/app/modules/trvdb"
	"github.com/Galdoba/TravellerTools/pkg/astrogation"
	"github.com/Galdoba/TravellerTools/pkg/mgt2trade/traffic"
	"github.com/Galdoba/TravellerTools/pkg/survey"
	"github.com/urfave/cli"
)

func Traffic(c *cli.Context) error {
	searchKey := c.String("worldname")
	sourceworld, err := trvdb.WorldByName(searchKey)
	if err != nil {
		return err
	}
	targetWorldsCoordinates := searchNeighbours(sourceworld, 3)
	fmt.Println("base freight factor:")
	freightFS := freightInfo{}
	freightFT := freightInfo{}
	passengersFS := passengerInfo{}
	passengersFT := passengerInfo{}
	for _, coord := range targetWorldsCoordinates {
		targetWorld, srchErr := survey.SearchByCoordinates(coord.ValuesHEX())
		if srchErr != nil {
			return srchErr
		}

		switch c.String("ruleset") {
		default:
			ff, fError := traffic.BaseFreightFactor_MGT2_Core(sourceworld, targetWorld)
			if fError != nil {
				return fError
			}
			freightFS.addAverageFreight_MGT2_Core(ff)
			freightFT.addAverageFreight_MGT2_Core(ff)
			pf, pError := traffic.BasePassengerFactor_MGT2_Core(sourceworld, targetWorld)
			if pError != nil {
				return nil
			}
			passengersFS.addAveragePassengers_MGT2_Core(pf)
			passengersFT.addAveragePassengers_MGT2_Core(pf)
		}

	}
	fmt.Printf("TOTAL FREIGHT: [%v]\nArriving [%v] tons of cargo\nDeparting [%v] tons of cargo\n", sourceworld.MW_Name(), freightFT.totalTons, freightFS.totalTons)
	fmt.Printf("TOTAL PASSENGERS: [%v]\nArriving passengers[%v]\nDeparting passengers[%v]\n", sourceworld.MW_Name(), passengersFT.total, passengersFS.total)
	return nil
}

func (fi *freightInfo) addAverageFreight_MGT2_Core(bfv int) {
	mjLotsDice, _ := traffic.FreightTrafficValues_MGT2_Core(bfv + 7 - 4)
	mnLotsDice, _ := traffic.FreightTrafficValues_MGT2_Core(bfv + 7)
	inLotsDice, _ := traffic.FreightTrafficValues_MGT2_Core(bfv + 7 + 2)
	fi.mjLots = mjLotsDice * 4
	fi.mnLots = mnLotsDice * 4
	fi.inLots = inLotsDice * 4
	for i := 0; i < fi.mjLots; i++ {
		fi.totalTons += 40
	}
	for i := 0; i < fi.mnLots; i++ {
		fi.totalTons += 20
	}
	for i := 0; i < fi.inLots; i++ {
		fi.totalTons += 4
	}
}

func (pi *passengerInfo) addAveragePassengers_MGT2_Core(bpv int) {
	lpas, _ := traffic.PassengerTrafficValues_MGT2_Core(bpv + 7 + 1)
	bpas, _ := traffic.PassengerTrafficValues_MGT2_Core(bpv + 7)
	mpas, _ := traffic.PassengerTrafficValues_MGT2_Core(bpv + 7)
	hpas, _ := traffic.PassengerTrafficValues_MGT2_Core(bpv + 7 - 4)
	pi.low += lpas * 4
	pi.basic += bpas * 4
	pi.middle += mpas * 4
	pi.high += hpas * 4
	pi.total = pi.low + pi.basic + pi.middle + pi.high
}

func searchNeighbours(sw *survey.SecondSurveyData, distance int) []astrogation.Coordinates {
	jcoord := astrogation.JumpFromCoordinates(astrogation.NewCoordinates(sw.CoordX(), sw.CoordY()), distance)
	coords := []astrogation.Coordinates{}
	for i, v := range jcoord {
		fmt.Printf("Search %v/%v\r", i, len(jcoord))
		nWorld, err := survey.SearchByCoordinates(v.ValuesHEX())

		if err != nil {
			//x, y := v.ValuesHEX()
			//fmt.Println(x, y, err.Error())♦
			continue
		}
		if nWorld.CoordX() == sw.CoordX() && nWorld.CoordY() == sw.CoordY() {
			continue
		}
		fmt.Println(fmt.Sprintf("%v (%v)/%v %v", nWorld.MW_Name(), nWorld.MW_UWP(), nWorld.Sector(), nWorld.Hex()))
		coords = append(coords, astrogation.NewCoordinates(nWorld.CoordX(), nWorld.CoordY()))
	}
	fmt.Println("                               \r")
	return coords
}

type freightInfo struct {
	mjLots    int
	mnLots    int
	inLots    int
	totalTons int
}

type passengerInfo struct {
	low    int
	basic  int
	middle int
	high   int
	total  int
}
