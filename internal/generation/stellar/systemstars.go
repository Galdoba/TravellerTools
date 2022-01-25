package stellar

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/internal/ehex"
	"github.com/Galdoba/TravellerTools/internal/helper"
	"github.com/Galdoba/TravellerTools/internal/struct/star"
	"github.com/Galdoba/TravellerTools/pkg/survey/calculations"
)

type StarNexus struct {
	StarSystems []*StarSystem
}

type StarSystem struct {
	Sun       *star.Star
	Companion *star.Star
	Body      []*planetaryBody
}

type SurveyReporter interface {
	CoordX() int
	CoordY() int
	Sector() string
	Hex() string
	MW_Name() string
	MW_UWP() string
	MW_Remarks() string
	MW_Importance() string
	MW_ImportanceInt() int
	MW_Economic() string
	MW_Cultural() string
	MW_Nobility() string
	Bases() string
	TravelZone() string
	PBG() string
	Worlds() int
	Allegiance() string
	Stellar() string
	RU() int
	SubSector() string
	SubSectorInt() int
	Quadrant() int
	BasesOld() string
	SectorAbb() string
	SubSectorName() string
	AllegianceExt() string
	NameByConvention() string
}

func NewNexus(ssd SurveyReporter) (*StarNexus, error) {
	sn := StarNexus{}
	err := fmt.Errorf("initial error was not adressed")
	name := ssd.NameByConvention()
	stellar := ssd.Stellar()
	if stellar == "" {
		stellar = calculations.GenerateNewStellar(name)
	}
	//////////////
	err = sn.placeStars(name, stellar)
	if err != nil {
		return &sn, err
	}
	for i := range sn.StarSystems {
		sn.StarSystems[i].SetOrbits()
	}
	/////////////Place MW
	sn.PlaceMainWorld(ssd)
	/////////////Place GG
	sn.PlaceGasGigants(ssd)
	/////////////Place Belts

	/////////////Place Other
	/////////////Place Satelites
	return &sn, err
}
func (sn *StarNexus) PlaceMainWorld(ssd SurveyReporter) error {
	_, _, gg, err := calculations.Decode(ssd.PBG())
	if err != nil {
		return err
	}
	habitZone := sn.StarSystems[0].Sun.HZ()
	remarks := ssd.MW_Remarks()
	//Place if NOT Satellite
	if !helper.SliceStrContains(strings.Split(remarks, " "), "Sa") {
		sn.StarSystems[0].Body[habitZone] = placeWorldTo("MainWorld", "MW", habitZone)
		return nil
	}
	//Place if IS Satellite
	switch {
	default:
		sn.StarSystems[0].Body[habitZone] = placeWorldTo("Gas Gigant w", "GG", habitZone)
	case gg < 1:
		sn.StarSystems[0].Body[habitZone] = placeWorldTo("BigWorld w", "BW", habitZone)
	}
	sn.StarSystems[0].Body[habitZone].satelites = append(sn.StarSystems[0].Body[habitZone].satelites, placeWorldTo("Mainworld", "MW_Sat", 0))
	sn.StarSystems[0].Body[habitZone].setSatteliteOrbits(-4)
	return nil
}

func (sn *StarNexus) PlaceGasGigants(ssd SurveyReporter) error {
	dp := dice.New().SetSeed(ssd.NameByConvention() + "Place GG")
	_, _, gg, err := calculations.Decode(ssd.PBG())
	if err != nil {
		return err
	}
	if gg < 1 {
		return nil
	}
	for g := 0; g < gg; g++ {

	}

	return fmt.Errorf("Not Implemented")
}

func (sn *StarNexus) hasAnyUnfilled() bool {
	for _, ss := range sn.StarSystems {
		if ss.haveUnfilledOrbits() {
			return true
		}
	}
	return false
}

func newGasGigant(seed string) string {
	dp := dice.New().SetSeed(seed)
	s := dp.Roll("2d6").DM(19).Sum()
	t := "L"
	if s < 23 {
		t = "S"
	}
	return ehex.New().Set(s).Code() + t
}

func (pb *planetaryBody) setSatteliteOrbits(zoneDM int) {
	dp := dice.New().SetSeed(pb.name + "_sat")
	concluded := false
	for !concluded {
		s := dp.Roll("1d6").DM(zoneDM).Sum()
		if s == 0 {
			pb.satelites = append(pb.satelites, placeWorldTo("Ring", "R", 0))
			continue
		}
		for i := 0; i < s; i++ {
			pb.satelites = append(pb.satelites, placeWorldTo("Sattelite", "S", 0))
		}
		concluded = true
	}
}

func (n *StarNexus) String() string {
	str := ""
	for _, st := range n.StarSystems {
		if st.Sun != nil {
			str += fmt.Sprintf("Sun: %v - %v\n", st.Sun.Name(), st.Sun.Code())
		}
		if st.Companion != nil {
			str += fmt.Sprintf("Companion: %v - %v\n", st.Companion.Name(), st.Companion.Code())
		}
		for i, b := range st.Body {
			str += fmt.Sprintf("    Body: %v - %v\n", b.Name(), b.Orbit())
			for _, s := range st.Body[i].satelites {
				str += fmt.Sprintf("        Sat: %v - %v\n", s.Name(), s.Orbit())
			}
		}
	}
	return str
}

func (stsys *StarSystem) SetOrbits() {
	maxOrb := 17
	if stsys.Sun == nil {
		return
	}
	if stsys.Sun.Orbit() > 0 {
		maxOrb = stsys.Sun.Orbit() - 3
		if maxOrb < 0 {
			maxOrb = 0
		}
	}
	for i := 0; i < maxOrb; i++ {
		stsys.Body = append(stsys.Body, EmptyOrbit(i))
	}
}

func (sn *StarNexus) placeStars(name, stellar string) error {
	sn.StarSystems = append(sn.StarSystems, &StarSystem{})
	sn.StarSystems = append(sn.StarSystems, &StarSystem{})
	sn.StarSystems = append(sn.StarSystems, &StarSystem{})
	sn.StarSystems = append(sn.StarSystems, &StarSystem{})
	starCodes, err := star.ParseStellar(stellar)
	if err != nil {
		return err
	}
	compos, err := SystemComposition(name, stellar)
	separated := separateBySystems(compos)
	codePosition := 0
	for stsys, s := range separated {
		for pos, categ := range s {
			st, _ := star.New(name+" "+greekLetter(codePosition+1), starCodes[codePosition], categ)
			st.SetOrbit()
			fmt.Println("orbit:", st.Name(), st.Orbit())
			switch pos {
			case 0:
				sn.StarSystems[stsys].Sun = st
			case 1:
				sn.StarSystems[stsys].Companion = st
			}
			codePosition++
		}
	}
	return nil
}

func (n *StarNexus) Print() {
	fmt.Println(len(n.StarSystems))
	for i := range n.StarSystems {
		if n.StarSystems[i].Sun != nil {
			n.StarSystems[i].Sun.Print()
		}
		if n.StarSystems[i].Companion != nil {
			n.StarSystems[i].Companion.Print()
		}
	}
}

func greekLetter(i int) string {
	switch i {
	case star.Category_Primary:
		return "Alpha"
	case star.Category_PrimaryCompanion:
		return "Beta"
	case star.Category_Close:
		return "Gamma"
	case star.Category_CloseCompanion:
		return "Delta"
	case star.Category_Near:
		return "Epsilon"
	case star.Category_NearCompanion:
		return "Zeta"
	case star.Category_Far:
		return "Eta"
	case star.Category_FarCompanion:
		return "Theta"
	}
	return "???"
}

type PlanetaryBody interface {
	Orbit() int //скорее всего да
	Name() string
	//Distance() float64 //скорее всего нет
}

type planetaryBody struct {
	orbit     int
	name      string
	pbType    string
	satelites []*planetaryBody
}

func (pb *planetaryBody) Orbit() int {
	return pb.orbit
}

func (pb *planetaryBody) Name() string {
	return pb.name
}

func placeWorldTo(wName, wType string, orbit int) *planetaryBody {
	return &planetaryBody{orbit, wName, wType, nil}
}

func EmptyOrbit(o int) *planetaryBody {
	return &planetaryBody{o, fmt.Sprintf("Orbit %v", o), "EMPTY", nil}
}

func separateBySystems(composition []int) [4][]int {
	sys := [4][]int{}
	for _, v := range composition {
		switch v {
		case 1, 3, 5, 7:
			sys[(v-1)/2] = []int{v}
		case 2, 4, 6, 8:
			sys[(v/2)-1] = []int{v - 1, v}
		}
	}
	return sys

}

//Distance - расчитывает расстояние тела от центра массы главной звезды в AU
func Distance(pb PlanetaryBody) (float64, error) {
	orb := pb.Orbit()
	switch {
	case orb < 0:
		return -1.0, fmt.Errorf("orbit is negaive")
	case orb > 20:
		return -1.0, fmt.Errorf("orbit is in another hex")
	default:
		return decimalOrbit(pb), nil
	}
}

func decimalOrbit(pb PlanetaryBody) float64 {
	dp := dice.New().SetSeed(pb.Name())
	fl := dp.Flux() + 5
	switch pb.Orbit() {
	case 0:
		return []float64{0.15, 0.16, 0.17, 0.18, 0.19, 0.2, 0.22, 0.24, 0.26, 0.28, 0.30}[fl]
	case 1:
		return []float64{0.30, 0.32, 0.34, 0.36, 0.38, 0.4, 0.43, 0.46, 0.49, 0.52, 0.55}[fl]
	case 2:
		return []float64{0.55, 0.58, 0.61, 0.64, 0.67, 0.7, 0.73, 0.76, 0.79, 0.82, 0.85}[fl]
	case 3:
		return []float64{0.85, 0.88, 0.91, 0.94, 0.97, 1.0, 1.06, 1.12, 1.18, 1.24, 1.30}[fl]
	case 4:
		return []float64{1.30, 1.36, 1.42, 1.48, 1.54, 1.6, 1.72, 1.84, 1.96, 2.08, 2.20}[fl]
	case 5:
		return []float64{2.20, 2.32, 2.44, 2.56, 2.68, 2.8, 3.04, 3.28, 3.52, 3.76, 4.00}[fl]
	case 6:
		return []float64{4.0, 4.2, 4.4, 4.7, 4.9, 5.2, 5.6, 6.1, 6.6, 7.1, 7.6}[fl]
	case 7:
		return []float64{7.6, 8.1, 8.5, 9.0, 9.5, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0}[fl]
	case 8:
		return []float64{15, 16, 17, 18, 19, 20, 22, 24, 26, 28, 30}[fl]
	case 9:
		return []float64{30, 32, 34, 36, 38, 40, 43, 47, 51, 54, 58}[fl]
	case 10:
		return []float64{58, 62, 65, 69, 73, 77, 84, 92, 100, 107, 115}[fl]
	case 11:
		return []float64{115, 123, 130, 138, 146, 154, 169, 184, 200, 215, 231}[fl]
	case 12:
		return []float64{231, 246, 261, 277, 292, 308, 338, 369, 400, 430, 461}[fl]
	case 13:
		return []float64{461, 492, 522, 553, 584, 615, 676, 738, 799, 861, 922}[fl]
	case 14:
		return []float64{922, 984, 1045, 1107, 1168, 1230, 1352, 1475, 1598, 1721, 1844}[fl]
	case 15:
		return []float64{1844, 1966, 2089, 2212, 2335, 2458, 2703, 2949, 3195, 3441, 3687}[fl]
	case 16:
		return []float64{3687, 3932, 4178, 4424, 4670, 4916, 5407, 5898, 6390, 6881, 7373}[fl]
	case 17:
		return []float64{7373, 7864, 8355, 8847, 9338, 9830, 10797, 11764, 12731, 13698, 14665}[fl]
	}
	return -1
}

func SystemComposition(systemName, stellarCode string) ([]int, error) {
	res := []int{}
	stars, err := star.ParseStellar(stellarCode)
	if err != nil {
		return res, err
	}
	dp := dice.New().SetSeed(systemName)
	try := 0
	for len(res) != len(stars) {
		try++
		res = []int{}
		res = append(res, star.Category_Primary)
		if dp.Flux() > 2 {
			res = append(res, star.Category_Close)
		}
		if dp.Flux() > 2 {
			res = append(res, star.Category_Near)
		}
		if dp.Flux() > 2 {
			res = append(res, star.Category_Far)
		}
		strs := res
		for _, st := range strs {
			switch st {
			case star.Category_Primary, star.Category_Close, star.Category_Near, star.Category_Far:
				if dp.Flux() > 2 {
					res = append(res, st+1)
				}
			}
		}
		//fmt.Printf("Try: %v/Res: %v (%v)\n", try, len(res), res)
	}
	//fmt.Println("tried", try, "times for", len(res), "stars")
	return res, err
}

/*
Планетарным телом может быть:

-тело
--звезда-компаньён
--Газовый Гигант
--Обычная
--Астеройдный Пояс

stellar.PlanetaryPosition(Star (Mass), Body (Distance), Date.Day())


*/

func PlanetaryPosition(mass float64, bodyDistance float64, time int64) (float64, int) {
	return 0, 0
}

func (stsys *StarSystem) haveUnfilledOrbits() bool {
	for _, orb := range stsys.Body {
		if orb.pbType == "EMPTY" {
			return true
		}
	}
	return false
}
