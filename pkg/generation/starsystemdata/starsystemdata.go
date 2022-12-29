package starsystemdata

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/generation/belts"
	"github.com/Galdoba/TravellerTools/pkg/generation/gasgigants"
	"github.com/Galdoba/TravellerTools/pkg/generation/stellar"
)

// type orbiter struct {
// }

// func (o *orbiter) SystemPosition() (int, int, int) {
// 	return 0, 0, 0
// }

/*
__ __  -------*-  -- -- -- -- -- --

*/

type Orbiter interface {
	SystemPosition() (int, int, int)
}

func PositionCode(orb Orbiter) string {
	_, pla, sat := orb.SystemPosition()

	plaS := fmt.Sprintf("%v", pla)
	if pla < 0 {
		plaS = ""
	}
	satS := fmt.Sprintf("%v", sat)
	if sat < 0 {
		satS = ""
	}
	code := plaS
	for len(code) < 3 {
		code += " "
	}
	code += satS
	for len(code) < 5 {
		code += " "
	}
	return code
}

type form7 struct {
	dateOfPreparation string
	systemName        string
	hexLocation       string
	sector            string
	subsector         string
	data              []string
}

type systemLocation struct {
	parentStar int
	pOrbit     int
	sOrbit     int
}

type data struct {
	loc     systemLocation
	upp     []ehex.Ehex
	remarks []string
}

func NewStarSystemData(sector string, hex string) *form7 {
	dice := dice.New().SetSeed(sector + hex)
	stellarData := stellar.GenerateStellar(dice)
	ggData := gasgigants.Generate(dice)
	bltData := belts.Generate(dice)
	stars := stellar.Parse(stellarData)

	//totalWorlds := 1 + dice.Sroll("2d6")
	//totalBodies := totalWorlds + len(ggData) + len(bltData)
	strNum := len(stars)
	diceCode := "1d" + fmt.Sprintf("%v", strNum)

	//TODO:сгенерировать шаблон планет

	// for _, star := range stars {

	// }
	return nil
}

/*

 */
