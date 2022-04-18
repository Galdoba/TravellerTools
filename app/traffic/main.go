package main

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/astrogation"
	"github.com/Galdoba/TravellerTools/pkg/survey"
	"github.com/Galdoba/devtools/cli/user"
	"github.com/Galdoba/utils"
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
	fmt.Print("State your quary: ")
	searchKey, inputErr := user.InputStr()
	if inputErr != nil {
		fmt.Println(inputErr.Error())
		return
	}
	matches := []string{}
	for _, line := range utils.LinesFromTXT(dataBase) {
		if strings.Contains(strings.ToUpper(line), strings.ToUpper(searchKey)) {
			matches = append(matches, line)
		}
	}
	if len(matches) > 200 || len(matches) < 1 {
		fmt.Println(len(matches), "detected. Please make another quary.")
		return
	}
	potentialWorlds := []*survey.SecondSurveyData{}
	for _, match := range matches {
		testWorld := survey.Parse(match)
		if strings.Contains(strings.ToUpper(testWorld.String()), strings.ToUpper(searchKey)) {
			potentialWorlds = append(potentialWorlds, testWorld)
		}

	}
	names := []string{}
	for _, sWorld := range potentialWorlds {

		names = append(names, fmt.Sprintf("%v (%v)/%v %v", sWorld.MW_Name(), sWorld.MW_UWP(), sWorld.Sector(), sWorld.Hex()))
	}
	sel := 0
	if len(names) > 1 {
		sel, _ = user.ChooseOne("Select quary:", names)
	}
	sw := potentialWorlds[sel]
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
