package survey

import (
	"fmt"
	"os"
	"testing"

	"github.com/Galdoba/utils"
)

func TestParseClean(t *testing.T) {
	return
	lines := utils.LinesFromTXT("c:\\Users\\Public\\TrvData\\cleanedData.txt")
	//tab := []*SecondSurveyData{}
	uwpMap := make(map[string]int)
	for i, line := range lines {
		ssd := Parse(line)
		uwpMap[ssd.MW_UWP()]++
		fmt.Print(len(uwpMap), i, "\r")
	}

	un := 0
	for _, v := range uwpMap {
		if v == 1 {
			un++
		}
	}
	fmt.Println("\nDone", un)
}

func TestParcing(t *testing.T) {
	return
	f, err := os.Create("c:\\Users\\Public\\TrvData\\cleanedData.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	wwritenn := 0
	lines := utils.LinesFromTXT("c:\\Users\\Public\\TrvData\\formattedData.txt")
	lenLines := len(lines) - 2
	errFound := 0
	errMap := make(map[string]int)
	dataMap := make(map[string]int)
	toTable := []*SecondSurveyData{}
	for i, input := range lines {

		fmt.Printf("checking world data: %v/%v (errors found: %v) - worlds Written: %v\r", i-1, lenLines, errFound, wwritenn)

		if i < 2 {
			continue
		}

		ssd := Parse(input)
		dataMap[input]++
		if dataMap[input] > 1 {
			continue
		}
		block := true
		if !ssd.containsErrors() && block {
			wwritenn++
			//cleaned := strings.ReplaceAll(ssd.String(), "   ", "|")
			cleaned := ssd.Compress()
			f.WriteString(cleaned + "\n")

			if ssd.Allegiance() != "XXXX" {
				toTable = append(toTable, ssd)
			}
		}
		//errFound++
		//fmt.Println(ssd)
		//dataMap[ssd.allegiance]++
		for _, err := range ssd.errors {
			if err != nil {
				//fmt.Println(err.Error())
				errFound++
				errMap[err.Error()]++

			}
		}

		// if i > 29480 {
		// 	return
		// }
	}
	fmt.Println("\n----------------------------------------")

	// for k, v := range dataMap {
	// 	//if v > 0 {
	// 	fmt.Println(k, ":", v)
	// 	//}
	// }

	fmt.Println("\n----------------------------------------")
	for _, lns := range ListOf(toTable) {
		fmt.Println(lns)
	}

}

func TestSearch(t *testing.T) {
	return
	found, err := Search("Earth")
	for _, v := range found {
		fmt.Println(v)
		fmt.Println(v.errors)
	}

	fmt.Println(err)
}
