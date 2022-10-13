package systemgeneration

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/survey"
)

func printSystem(gs *GenerationState) {
	region := gs.SystemName
	data := ""
	data += fmt.Sprintf("Region: %v\n", region)
	data += fmt.Sprintf("--------------------------------------------------------------------------------\n")
	for _, star := range gs.System.Stars {
		data += fmt.Sprintf("Star %v %v\n", star.rank, star.Code())
		for _, orb := range star.orbitDistances {
			data += fmt.Sprintf(" %v   ", formatFloatOutput(orb))
			switch v := star.orbit[orb].(type) {
			case *rockyPlanet:
				data += "Planet      " + "_" + v.sizeCode + v.atmoCode + v.hydrCode + "___-_  " + v.habZone + " zone " + v.sizeType
				for _, moon := range v.moons {
					data += "\n             Moon" + "        _" + moon.sizeCode + moon.atmoCode + moon.hydrCode + "___-_  " + moon.habZone + " zone " + moon.sizeType
				}
			case *ggiant:
				data += "Gas Gigant  " + v.descr
				data += "\n             " + v.ring
				for _, moon := range v.moons {
					data += "\n             Moon" + "        _" + moon.sizeCode + moon.atmoCode + moon.hydrCode + "___-_  " + moon.habZone + " zone " + moon.sizeType
				}
			case *belt:
				data += "Belt        " + v.composition
			case *jumpZoneBorder:
				data += "Jump Border " + v.zone
			}
			data += "\n"
		}
	}
	data += fmt.Sprintf("--------------------------------------------------------------------------------")
	fmt.Println(data)
}

func formatFloatOutput(fl float64) string {
	s := fmt.Sprintf("%v", fl)
	sp := strings.Split(s, ".")
	for len(sp[0]) < 4 {
		sp[0] = " " + sp[0]
	}
	if len(sp) == 1 {
		sp = append(sp, "0")
	}

	for len(sp[1]) < 4 {
		sp[1] = sp[1] + " "
	}
	return sp[0] + "." + sp[1]
}

type StarSystemData struct {
	World survey.SecondSurveyData
}

type importData struct {
	data     string
	dataKey  string
	dataType string
}

type importer interface {
	Data() string
	DataKey() string
	DataType() string
	Read(string) []importData
}

func (gs *GenerationState) Import(data ...importer) error {
	for _, dt := range data {
		switch v := dt.(type) {
		default:
			gs.importedData = append(gs.importedData, importData{dt.Data(), dt.DataKey(), dt.DataType()})
		case *seImporter:
			gs.importedData = append(gs.importedData, importData{v.se.Stellar(), "Stellar", "STR"})
		}

	}
	return nil
}

type seImporter struct {
	se survey.SecondSurveyData
}

func InjectSecondSurveyData(se survey.SecondSurveyData) *seImporter {
	return &seImporter{se}
}

func (d seImporter) Data() string {
	return d.se.String()
}

func (d seImporter) DataType() string {
	return "SecondSurveyData"
}

func (d seImporter) DataKey() string {
	return "FULL"
}

func (d seImporter) Read(key string) []importData {
	imprt := []importData{}
	switch key {
	default:
		imprt = append(imprt, importData{dataKey: key, dataType: "UNKNOWN", data: "IMPORT ERROR"})
	case "Stellar":
		imprt = append(imprt, importData{dataKey: key, dataType: "STR", data: d.se.Stellar()})
	}
	return imprt
}

/*
Registry Name: TrojXA-202-T196C27
--------------------------------------------------------------------------------
Trojan Reach 0202
Primary: M5 V
    0.01    Jump Border      D10
    0.13    Planet           X630000-0
         14 Moon             X010000-0
         26 Moon             X010000-0
         31 Moon             X010000-0
    0.14    Planet           X010000-0
    0.15    Jump Border      D100
    0.77    Planet           X000000-0
    1.15    Planet           X110000-0
            Moon             X010000-0
            Moon             X000000-0
Secondary: M1 V (distance to Primary 12 AU)
    0.05    Jump Border      D10
    0.1     Gas Gigant       Hot Neptunian
            5 rings with width 29 km
    0.47    Jump Border      D100
--------------------------------------------------------------------------------

Region            -определяет общий тип региона (звездная система/туманность/пустота)
  Star            -звезды (если звезды нет то специальные объекты или точки привязки)
    Body          -тела вращающиеся вокруг звузд или центра региона (планеты/газовые гиганты/астеройдные пояса и скопления/блуждающие планеты)
	  Moon        -тела вращающиеся вокруг других тел (спутники)
	    Points    -навигационные точки для перелетов (точки лагранжа/гравитационные барьеры)

РЕГИОН:
Глобальные координаты
Тип

ТОЧКИ:



*/
