package balancingact

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/survey"
	"github.com/Galdoba/utils"
)

type plInput struct {
	name string
	uwp  string
	eX   string
	cX   string
	x    int
	y    int
}

func testPlanets() []plInput {

	return []plInput{
		{"Capital", "A586A98-F", "(H9G+4)", "[AE5F]", 20, -22},
	}
}

func parsedInput() []SurveyData {
	lines := utils.LinesFromTXT(`C:\Users\Public\TrvData\testSector.txt`)
	sd := []SurveyData{}
	for _, ln := range lines {
		sd = append(sd, survey.Parse(ln))
	}
	return sd
}

func TestPlanet(t *testing.T) {
	planets := []*planet{}
	for v, inp := range parsedInput() {
		if v > 5 {
			//break
		}
		fmt.Printf("Test planet %v  \r", v)
		pl, err := ImportPlanet(inp)
		planets = append(planets, pl)
		if err != nil {
			t.Errorf("internal error: %v", err.Error())
		}
		if pl == nil {
			t.Errorf("planet was not created")
		}
		for k, v := range pl.atrib {
			if v == -1 {
				t.Errorf("value %v is untouched", k)
			}
		}
		for key, val := range pl.atrib {
			continue
			if val > 10 {
				switch key {
				case Solidarity:
					t.Errorf("Solidarity exided 10: %v", val)
				case Wealth:
					t.Errorf("Wealth exided 10: %v", val)
				case Expansion:
					t.Errorf("Expansion exided 10: %v", val)
				case Might:
					t.Errorf("Might exided 10: %v", val)
				case Development:
					t.Errorf("Development exided 10: %v", val)
				}
			}
		}
		fmt.Println(pl)
	}
	//fmt.Printf("Trade: %v <--> %v is %v MCr per Year", planets[0].name, planets[1].name, TradeRouteValue(planets[0], planets[1]))
}
