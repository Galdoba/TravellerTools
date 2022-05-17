package trvdb

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/survey"
	"github.com/Galdoba/devtools/cli/user"
	"github.com/Galdoba/utils"
)

const (
	dataBase = "c:\\Users\\Public\\TrvData\\cleanedData.txt"
)

func WorldByName(quarry ...string) (*survey.SecondSurveyData, error) {

	searchKey := "" //
	switch len(quarry) {
	default:
		searchKey = quarry[0]
		if searchKey == "" {
			fmt.Print("State your quary: ")
			searchKey, _ = user.InputStr()

		}
	case 0:
		fmt.Print("State your quary: ")
		searchKey, _ = user.InputStr()
	}
	// if len(searchKey) < 3 {
	// 	return nil, fmt.Errorf("quarry must me at least 3 characters [%v]", searchKey)
	// }
	matches := []string{}
	db := utils.LinesFromTXT(dataBase)
	for _, line := range db {
		if strings.Contains(strings.ToUpper(line), "|"+strings.ToUpper(searchKey)+"|") {
			matches = append(matches, line)
		}
	}
	if len(matches) > 1300 {
		return nil, fmt.Errorf("matches limit exided (%v)", len(matches))
	}
	if len(matches) < 1 {
		return nil, fmt.Errorf("world '%v' not found", searchKey)
	}
	if len(matches) < 1 {
		return nil, fmt.Errorf("no matches on '%v' in database", searchKey)
	}
	potentialWorlds := []*survey.SecondSurveyData{}
	for _, match := range matches {
		testWorld := survey.Parse(match)
		if strings.ToUpper(testWorld.MW_Name()) == strings.ToUpper(searchKey) {
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
	return sw, nil
}
