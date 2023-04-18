package systemgeneration

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"github.com/Galdoba/TravellerTools/pkg/survey"
)

func printSystem(gs *GenerationState) {
	region := gs.SystemName
	data := ""
	data += fmt.Sprintf("Region: %v\n", region)
	data += fmt.Sprintf("--------------------------------------------------------------------------------\n")
	for s, star := range gs.System.Stars {
		data += fmt.Sprintf("Star %v %v", star.rank, star.Code())
		if s != 0 {
			data += fmt.Sprintf(" (distance to Primary: %v)", star.distanceFromPrimaryAU)
		}
		data += "\n"
		for _, orb := range star.orbitDistances {
			data += fmt.Sprintf(" %v   ", formatFloatOutput(orb))
			switch v := star.orbit[orb].(type) {
			case *rockyPlanet:
				data += "Planet      " + v.uwpStr + "  " + v.comment
				for _, moon := range v.moons {
					data += "\n             Moon" + "        " + moon.uwpStr + "  " + moon.comment
				}
			case *ggiant:
				data += "Gas Gigant  " + v.descr
				data += "\n             " + v.ring
				for _, moon := range v.moons {
					data += "\n             Moon" + "        " + moon.uwpStr + "  " + moon.comment
				}
			case *belt:
				data += "Belt        " + v.uwpStr + "  " + v.comment
				data += "\n             " + v.composition
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
			gs.importedData = append(gs.importedData, importData{v.se.PBG(), "PBG", "STR"})
			gs.importedData = append(gs.importedData, importData{v.se.MW_UWP(), "MW_UWP", "STR"})
			gs.importedData = append(gs.importedData, importData{v.se.MW_Name(), "MW_NAME", "STR"})
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
	case "PBG":
		imprt = append(imprt, importData{dataKey: key, dataType: "STR", data: d.se.PBG()})
	case "MW_UWP":
		imprt = append(imprt, importData{dataKey: key, dataType: "STR", data: d.se.MW_UWP()})
	case "MW_NAME":
		imprt = append(imprt, importData{dataKey: key, dataType: "STR", data: d.se.MW_Name()})

	}
	return imprt
}

func (gs *GenerationState) callImport(key string) error {
	for _, imported := range gs.importedData {
		if imported.dataKey != key {

			continue
		}
		switch key {
		case "Stellar":
			if err := gs.injectStellar(imported.data); err != nil {
				return err
			}
		case "PBG":
			if err := gs.injectPBG(imported.data); err != nil {
				return err
			}
			// case "MW_UWP":
			// 	if err := gs.injectMW_UWP(imported.data); err != nil {
			// 		return err
			// 	}
			// case "MW_NAME":
			// 	if err := gs.injectMW_Seed(imported.data); err != nil {
			// 		return err
			// 	}
		}
	}
	return nil
}

type worldPosition struct {
	star  int
	orbit float64
}

func (gs *GenerationState) SuggestWorldPosition() worldPosition {
	wp := worldPosition{-1, -1.0}
	world, err := uwp.FromString0(gs.System.MW_UWP)
	if err != nil {
		return wp
	}
	switch world.Data(profile.KEY_SIZE).Value() {
	case 0:
		for s, star := range gs.System.Stars {
			orbits := star.orbitDistances
			l := len(orbits)
			r := gs.Dice.Roll(fmt.Sprintf("1d%v", l)).DM(-1).Sum()

			return worldPosition{s, orbits[r]}
		}
	default:
		for s, star := range gs.System.Stars {
			orbits := star.orbitDistances
			sugest := 0.0
			n := 0
			for o, orb := range orbits {
				sugest = orb
				n = o
				if orb > star.habitableHigh {
					break
				}
			}
			if n > 0 {
				return worldPosition{s, sugest}
			}
		}
	}
	return wp
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


Факторы биологии:
Атмосфера			UWP
	Кислород		UWP
Вода			UWP
Температура
Расстояние от звезды орбита
Цвет главной звезды Объект

средняя температура поверхнеости = температура тела * (Альбедо) * Парниковый эффект
Температура тела = Яркость * Расстояние
Парниковый Эффект = Давление на поверхности * Гравитация на поверхности
Альбедо = UWP
Давление на поверхности = UWP
Гравитация на поверхности = Масса * UWP
масса = UWP / плотность
плотность = Орбита / звезда / UWP




*/
