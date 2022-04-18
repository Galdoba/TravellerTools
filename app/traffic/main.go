package main

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/app/modules/trvdb"
	"github.com/Galdoba/TravellerTools/pkg/astrogation"
	"github.com/Galdoba/TravellerTools/pkg/survey"
)

const (
	dataBase = "c:\\Users\\Public\\TrvData\\cleanedData.txt"
)

/*
1 найти мир
	запрос:
		0 найдено
			END
		n найдено
			выбрать мир
			GO TO 2
		50+ найдено
			END
2 найти соседей мира
3 расчитать трафик


*/

func main() {

	wrlds, err := trvdb.WorldByName()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Found", wrlds)
	sw := wrlds

	fmt.Println("Source World Chosen:", fmt.Sprintf("%v (%v)/%v %v\n", sw.MW_Name(), sw.MW_UWP(), sw.Sector(), sw.Hex()))
	jcoord := astrogation.JumpFromCoordinates(astrogation.NewCoordinates(sw.CoordX(), sw.CoordY()), 3)
	fmt.Println("Trade capable worlds:")
	for i, v := range jcoord {
		fmt.Printf("Search %v/%v\r", i, len(jcoord))
		nWorld, err := survey.SearchByCoordinates(v.ValuesHEX())

		if err != nil {
			//x, y := v.ValuesHEX()
			//fmt.Println(x, y, err.Error())
			continue
		}
		if nWorld.CoordX() == sw.CoordX() && nWorld.CoordY() == sw.CoordY() {
			continue
		}
		fmt.Println(fmt.Sprintf("%v (%v)/%v %v", nWorld.MW_Name(), nWorld.MW_UWP(), nWorld.Sector(), nWorld.Hex()))

	}
	fmt.Println("                               \r")
}
