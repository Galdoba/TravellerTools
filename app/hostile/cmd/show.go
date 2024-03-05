package cmd

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/hostile/world"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/urfave/cli/v2"
)

func SurveySystem() *cli.Command {
	// cfg := &config.Config{}
	return &cli.Command{
		Name:  "survey_system",
		Usage: "Creates layout for Random System using Solo rules p.134-138",
		Flags: []cli.Flag{
			//-short
			&cli.BoolFlag{
				Name:        "nomenclature",
				Usage:       "astronomical name of the system (used as seed for reocurance)",
				Aliases:     []string{"n"},
				DefaultText: "random if left blank",
			},
		},
		Before: func(c *cli.Context) error {
			// cfg, _ = config.Load(c.App.Name)
			return nil
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Start survey_system")
			fmt.Println("End   survey_system")

			return nil
		},
	}
}

type StarSystemData struct {
	SystemName   string           `json:"System"`
	SurveyPoints int              `json:"Survey Points,omitempty"`
	Orbit        []PlanetaryOrbit `json:"Orbits,omitempty"`
	stars        int
}

type PlanetaryOrbit struct {
	Orbit       int
	Body        string
	Zone        string
	Notes       string
	SAH         string
	Temperature string
}

func NewSystem(name string) (*StarSystemData, error) {
	ssd := StarSystemData{}
	ssd.SystemName = name
	dice := dice.New().SetSeed(ssd.SystemName)
	// dice.Vocal()
	addStars := 0
	switch dice.Sroll("2d6") {
	default:
		ssd.stars = 1
	case 9, 10:
		ssd.stars = 2
	case 11, 12:
		ssd.stars = 3
	}
	fillOrbits := dice.Sroll("2d6+2")
	for i := 0; i < fillOrbits; i++ {
		ssd.Orbit = append(ssd.Orbit, PlanetaryOrbit{})
	}
	//Add Stars
	for i := 0; i < ssd.stars; i++ {
		switch i {
		case 0:
			ssd.Orbit[0].Body = "Star"
		default:
			switch dice.Sroll("1d2") {
			case 1:
				for o, orb := range ssd.Orbit {
					if orb.Body == "Star" {
						continue
					}
					ssd.Orbit[o].Body = "Star"
					break
				}
			case 2:
				addStars++
			}
		}
	}
	//Place Mainworld
	mwo := dice.Sroll("1d3+2")
	set := false
	for !set {
		for i, bod := range ssd.Orbit {
			if bod.Body != "" {
				continue
			}
			mwo--
			if mwo == 0 {
				ssd.Orbit[i].Body = "MW"
				set = true
				mwo = i
				break
			}
		}
	}
	//Set GG
	hgg := 0
	gg := (dice.Sroll("2d6") / 2) - 2
	if gg > 0 && dice.Sroll("1d6") == 6 {
		gg--
		hgg++
	}
	if hgg > 0 {
		for o, _ := range ssd.Orbit {
			if ssd.Orbit[o].Body != "" {
				continue
			}
			ssd.Orbit[o].Body = "GG"
			ssd.Orbit[o].Notes = "Hot Jupiter"
			gg--
			break
		}
	}
	for i := 0; i < gg; i++ {
		r := dice.Sroll("1d6") + mwo - 1
		ssd.placeGG(r, dice)
	}
	//SET ASTEROIDS
	for i := range ssd.Orbit {
		if ssd.Orbit[i].Body != "" {
			continue
		}
		if dice.Sroll("1d6") == 6 {
			ssd.Orbit[i].Body = "BELT"
		}
	}
	//REST OF THE PLANETS
	inner := true

	for i := range ssd.Orbit {

		switch ssd.Orbit[i].Body {
		default:
			if strings.Contains(ssd.Orbit[i].Body, "MW") {
				inner = false
				ssd.Orbit[i].Zone = "habitable"
			}
			continue
		case "":
			switch inner {
			case true:
				ssd.Orbit[i].Body = innerPlanetType(dice) //inner
				ssd.Orbit[i].Zone = "inner"
			case false:
				ssd.Orbit[i].Body = outerPlanetType(dice) //inner
				ssd.Orbit[i].Zone = "outer"
			}
		}
		ssd.Orbit[i].Orbit = i
	}
	// Outer Stars
	for i := 0; i < addStars; i++ {
		ssd.Orbit = append(ssd.Orbit, PlanetaryOrbit{Body: "Star"})
	}
	for i := range ssd.Orbit {
		ssd.Orbit[i].detailPlanet(dice)
	}
	return &ssd, nil
}

func innerPlanetType(dice *dice.Dicepool) string {
	switch dice.Sroll("2d6") {
	case 2, 3, 4, 5, 6, 7, 8:
		return "Rock"
	case 9, 10:
		return "Hellhole"
	case 11, 12:
		return "Desert"
	}
	return "Unknown"
}

func outerPlanetType(dice *dice.Dicepool) string {
	switch dice.Sroll("2d6") {
	case 2, 3, 4, 5, 6, 7, 8, 9:
		return "Iceball"
	case 10:
		return "Desert"
	case 11, 12:
		return "Hellhole"
	}
	return "Unknown"
}
func (ssd *StarSystemData) placeGG(o int, dice *dice.Dicepool) error {
	if o > len(ssd.Orbit)-1 {
		ssd.Orbit = append(ssd.Orbit, PlanetaryOrbit{})
		return ssd.placeGG(o, dice)
	}
	switch ssd.Orbit[o].Body {
	case "MW":
		switch dice.Sroll("1d2") {
		case 1:
			ssd.Orbit[o].Body = "MW"
			ssd.Orbit[o].Notes = "orbits GG"
		case 2:
			return ssd.placeGG(o+1, dice)
		}
	case "GG":
		return ssd.placeGG(o+1, dice)
	case "":
		ssd.Orbit[o].Body = "GG"
	default:
		return fmt.Errorf("unknown err placeGG(%v, dice)", o)
	}
	return nil
}

func (po *PlanetaryOrbit) detailPlanet(dice *dice.Dicepool) {
	body := po.Body
	if strings.Contains(body, "MW") {
		body = mwType(dice.Sroll("2d6"))
	}
	switch body {
	case "Rock":
		po.detailRock(dice)
	}

}

func (po *PlanetaryOrbit) detailRock(dice *dice.Dicepool) {
	size := world.NewSize(dice)
	atmoDM := 0
	switch dice.Sroll("1d3") {
	case 1, 2:
		for size.Value() != 1 && size.Value() != 2 {
			size = world.NewSize(dice)
		}
		atmoDM = -1
	case 3:
		for size.Value() < 3 || size.Value() > 6 {
			size = world.NewSize(dice)
		}
	}
	atmo := world.NewAtmosphere(dice, size.Value())
	hydrDM := 0
	switch dice.Sroll("1d6") + atmoDM {
	case 0, 1, 2, 3, 4:
		for atmo.Value() > 2 {
			atmo = world.NewAtmosphere(dice, size.Value())
		}
	case 5, 6:
		for atmo.Value() > 5 {
			atmo = world.NewAtmosphere(dice, size.Value())
		}
		hydrDM = 1
	}
	hydr := world.NewHydrographics(dice, size.Value(), atmo.Value())
	switch dice.Sroll("1d6") + hydrDM {
	case 1, 2, 3, 4, 5, 6:
		hydr = ehex.New().Set(0)
	case 7:
		for hydr.Value() > 2 {
			hydr = world.NewHydrographics(dice, size.Value(), atmo.Value())
		}
	}
	if po.Zone == "inner" {
		po.Temperature = "Hot"
		if po.Orbit == 1 || po.Orbit == 2 {
			po.Temperature = "Inferno"
		}
	}
	if po.Zone == "habitable" {
		po.Temperature = "Temperate"
	}
	po.SAH = size.Code() + atmo.Code() + hydr.Code()
}

func (po *PlanetaryOrbit) detailHellhole(dice *dice.Dicepool) {
	size := world.NewSize(dice)
	atmo := world.NewAtmosphere(dice, size.Value())
	for atmo.Value() < 11 {
		size = world.NewSize(dice)
		atmo = world.NewAtmosphere(dice, size.Value())
	}
	hydr := world.NewHydrographics(dice, size.Value(), atmo.Value())
	switch dice.Sroll("1d6") {
	case 1, 2, 3:
		hydr = ehex.New().Set(0)
	case 7:
		for hydr.Value() > 2 {
			hydr = world.NewHydrographics(dice, size.Value(), atmo.Value())
		}
	}
	if po.Zone == "inner" {
		po.Temperature = "Hot"
		if po.Orbit == 1 || po.Orbit == 2 {
			po.Temperature = "Inferno"
		}
	}
	if po.Zone == "habitable" {
		po.Temperature = "Temperate"
	}
	po.SAH = size.Code() + atmo.Code() + hydr.Code()
}
func mwType(i int) string {
	switch i {
	case 2, 3, 4, 5:
		return "Rock"
	case 6:
		return "Hellhole"
	case 7:
		return "Desert"
	case 8, 9, 10, 11:
		return "Garden"
	case 12:
		return "Waterworld"
	}
	return "error"
}
