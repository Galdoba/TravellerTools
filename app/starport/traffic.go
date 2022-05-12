package main

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/app/modules/trvdb"
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/Galdoba/TravellerTools/pkg/mgt2trade/traffic"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"github.com/Galdoba/TravellerTools/pkg/survey"
	"github.com/urfave/cli"
)

func Traffic(c *cli.Context) error {
	searchKey := c.String("worldname")

	sourceworld, err := SearchSourcePort(searchKey)
	if err != nil {
		return err
	}

	reach := processReach(c.Int("reach"), sourceworld.MW_UWP())
	fmt.Printf("Sourceworld [%v] detected...\nChecking for neighbours within a reach of %v parsecs...\n", sourceworld.MW_Name(), reach)
	targetWorldsCoordinates := searchNeighbours(sourceworld, reach)

	fmt.Println("Gathering traffic data:")
	tradeData := NewTrafficData(sourceworld, reach)
	for _, coord := range targetWorldsCoordinates {
		freightFS := freightInfo{}
		freightFT := freightInfo{}
		passengersFS := passengerInfo{}
		passengersFT := passengerInfo{}
		targetWorld, srchErr := PortByCoordinates(coord.HexValues())
		if srchErr != nil {
			return srchErr
		}

		if targetWorld.CoordX() == sourceworld.CoordX() && targetWorld.CoordY() == sourceworld.CoordY() {
			continue
		}
		switch c.String("ruleset") {
		default:
			return fmt.Errorf("ruleset is not defined")
		case "mgt2_core":
			pf, pError := traffic.BasePassengerFactor_MGT2_Core(sourceworld, targetWorld)
			if pError != nil {
				return nil
			}
			passengersFS.addAveragePassengers_MGT2_Core(pf)
			passengersFT.addAveragePassengers_MGT2_Core(pf)
			ff, fError := traffic.BaseFreightFactor_MGT2_Core(sourceworld, targetWorld)
			if fError != nil {
				return fError
			}
			freightFS.addAverageFreight_MGT2_Core(ff)
			freightFT.addAverageFreight_MGT2_Core(ff)
			tradeData.freightD[targetWorld] = freightFS
			tradeData.freightA[targetWorld] = freightFT
			tradeData.passengersD[targetWorld] = passengersFS
			tradeData.passengersA[targetWorld] = passengersFT
		case "mgt1_mp":
			pfd, pError := traffic.BasePassengerFactor_MGT1_MP(sourceworld, targetWorld)
			if pError != nil {
				return nil
			}
			pfa, pError := traffic.BasePassengerFactor_MGT1_MP(targetWorld, sourceworld)
			if pError != nil {
				return nil
			}
			passengersFS.addAveragePassengers_MGT1_MP(pfd)
			passengersFT.addAveragePassengers_MGT1_MP(pfa)
			ffd, fError := traffic.BaseFreightFactor_MGT1_MP(sourceworld, targetWorld)
			if fError != nil {
				return fError
			}
			ffa, fError := traffic.BaseFreightFactor_MGT1_MP(targetWorld, sourceworld)
			if fError != nil {
				return fError
			}
			freightFS.addAverageFreight_MGT1_MP(ffd)
			freightFT.addAverageFreight_MGT1_MP(ffa)
			tradeData.freightD[targetWorld] = freightFS
			tradeData.freightA[targetWorld] = freightFT
			tradeData.passengersD[targetWorld] = passengersFS
			tradeData.passengersA[targetWorld] = passengersFT
		}
		fmt.Printf("[%v] <--> [%v]   Passengers (D/A): %v/%v   Freight (D/A): %v/%v\n", sourceworld.MW_Name(), targetWorld.MW_Name(), tradeData.passengersD[targetWorld].total, tradeData.passengersA[targetWorld].total, tradeData.freightD[targetWorld].total, tradeData.freightA[targetWorld].total)

	}
	//tradeData := TrafficData{sourceworld, "Soureworld", len(targetWorldsCoordinates), 4}

	fmt.Print(tradeData.String())
	return nil
}

func processReach(flagValue int, uwpS string) int {
	if flagValue > 0 {
		return flagValue
	}
	u, _ := uwp.FromString(uwpS)
	if u.TL() < 11 {
		return 2
	}
	return u.TL() - 9
}

func (fi *freightInfo) addAverageFreight_MGT2_Core(bfv int) {
	mjLotsDice, _ := traffic.FreightTrafficValues_MGT2_Core(bfv + 7 - 4)
	mnLotsDice, _ := traffic.FreightTrafficValues_MGT2_Core(bfv + 7)
	inLotsDice, _ := traffic.FreightTrafficValues_MGT2_Core(bfv + 7 + 2)
	fi.mjLots = mjLotsDice * 4
	fi.mnLots = mnLotsDice * 4
	fi.inLots = inLotsDice * 4
	for i := 0; i < fi.mjLots; i++ {
		fi.total += 40
	}
	for i := 0; i < fi.mnLots; i++ {
		fi.total += 20
	}
	for i := 0; i < fi.inLots; i++ {
		fi.total += 4
	}
}

func (fi *freightInfo) addAverageFreight_MGT1_MP(bfv int) {
	mjLotsDice, aMj := traffic.FreightTrafficValues_MGT1_MP(bfv, traffic.Lot_Major)
	mnLotsDice, aMn := traffic.FreightTrafficValues_MGT1_MP(bfv, traffic.Lot_Minor)
	inLotsDice, aIn := traffic.FreightTrafficValues_MGT1_MP(bfv, traffic.Lot_Incidental)
	fi.mjLots = mjLotsDice*4 + aMj
	fi.mnLots = mnLotsDice*4 + aMn
	fi.inLots = inLotsDice*4 + aIn
	for i := 0; i < fi.mjLots; i++ {
		fi.total += 40
	}
	for i := 0; i < fi.mnLots; i++ {
		fi.total += 20
	}
	for i := 0; i < fi.inLots; i++ {
		fi.total += 4
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

func (pi *passengerInfo) addAveragePassengers_MGT1_MP(bpv int) {
	lpas, _ := traffic.PassengerTrafficValues_MGT1_MP(bpv, traffic.Passage_Low)
	bpas, _ := traffic.PassengerTrafficValues_MGT1_MP(bpv, traffic.Passage_Basic)
	mpas, _ := traffic.PassengerTrafficValues_MGT1_MP(bpv, traffic.Passage_Middle)
	hpas, _ := traffic.PassengerTrafficValues_MGT1_MP(bpv, traffic.Passage_High)
	pi.low += lpas * 4
	pi.basic += bpas * 4
	pi.middle += mpas * 4
	pi.high += hpas * 4
	pi.total = pi.low + pi.basic + pi.middle + pi.high
}

func searchNeighbours(sourceworld Port, distance int) []hexagon.Hexagon {
	//jcoord := astrogation.JumpMap(sourceworld, distance)
	jcoord, _ := hexagon.Spiral(sourceworld, distance)
	coords := []hexagon.Hexagon{}
	for i, v := range jcoord {
		fmt.Printf("Search %v/%v    \r", i, len(jcoord))
		neighbour, err := PortByCoordinates(v.HexValues())
		if err != nil {
			continue
		}
		if hexagon.MatchHex(sourceworld, neighbour) {
			continue
		}
		coords = append(coords, hexagon.FromHex(neighbour))
	}
	//	utils.ClearScreen()
	return coords
}

type freightInfo struct {
	mjLots int
	mnLots int
	inLots int
	total  int
}

type passengerInfo struct {
	low    int
	basic  int
	middle int
	high   int
	total  int
}

type Port interface {
	MW_Name() string
	MW_UWP() string
	MW_Remarks() string
	TravelZone() string
	Hex() string
	Sector() string
	CoordX() int
	CoordY() int
	CoordQ() int
	CoordR() int
	CoordS() int
	hexagon.Cube
	hexagon.Hex
}

func TradeCodes(p Port) string {
	uwpS := p.MW_UWP()
	u, _ := uwp.FromString(uwpS)
	res := ""
	if u.TL() >= 12 {
		res = "Ht "
	}
	if u.TL() <= 5 && u.Pops() >= 1 {
		res = "Lt "
	}
	return res + p.MW_Remarks()
}

func PortByCoordinates(x, y int) (Port, error) {
	return survey.SearchByCoordinates(x, y)

}

func Hexagon(p Port) hexagon.Hexagon {
	hex, _ := hexagon.New(hexagon.Feed_HEX, p.CoordX(), p.CoordY())
	return hex
}

func SearchSourcePort(searchKey string) (Port, error) {
	return trvdb.WorldByName(searchKey)
}
