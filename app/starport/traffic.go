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
	to := 0
	from := 0
	for _, coord := range targetWorldsCoordinates {
		targetWorld, srchErr := survey.SearchByCoordinates(coord.ValuesHEX())
		if srchErr != nil {
			return srchErr
		}
		bf, fError := traffic.BaseFreightFactor_MGT2_Core(sourceworld, targetWorld)
		if fError != nil {
			return fError
		}
		bfr, fError := traffic.BaseFreightFactor_MGT2_Core(targetWorld, sourceworld)
		if fError != nil {
			return fError
		}
		from += bf
		to += bfr
		fmt.Printf("[%v] --> [%v] = %v\n", sourceworld.MW_Name(), targetWorld.MW_Name(), bf)
		fmt.Printf("[%v] --> [%v] = %v\n", targetWorld.MW_Name(), sourceworld.MW_Name(), bfr)

	}
	fmt.Printf("TOTAL: [%v]\nArriving [%v]\nDeparting [%v]\n", sourceworld.MW_Name(), to, from)
	return fmt.Errorf("commant 'Traffic' not complete")
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
