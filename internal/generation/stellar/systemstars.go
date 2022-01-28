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
	ssd         SurveyReporter
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
	GenerationSeed() string
	NameByConvention() string
}

func NewNexus(ssd SurveyReporter) (*StarNexus, error) {
	sn := StarNexus{}
	sn.ssd = ssd
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
	sn.PlaceMainWorld()
	/////////////Place GG
	sn.PlaceGasGigants(ssd)
	/////////////Place Belts
	sn.PlaceBelts(ssd)
	/////////////Place Other
	sn.PlaceOther(ssd)
	/////////////Place Satelites
	sn.PlaceSatellites(ssd)
	cleanSN := Clean(sn)
	return cleanSN, err
}

func Clean(sn StarNexus) *StarNexus {
	snClean := StarNexus{}
	for s, syst := range sn.StarSystems {
		if syst.Sun == nil {
			continue
		}
		cleanSystem := &StarSystem{}
		cleanSystem.Sun = sn.StarSystems[s].Sun
		cleanSystem.Companion = sn.StarSystems[s].Companion
		for p, body := range sn.StarSystems[s].Body {
			if body.pbType != "EMPTY" {
				cleanSystem.Body = append(cleanSystem.Body, sn.StarSystems[s].Body[p])
			}
		}
		snClean.StarSystems = append(snClean.StarSystems, cleanSystem)
	}
	return &snClean
}

func (sn *StarNexus) PlaceSatellites(ssd SurveyReporter) {
	dp := dice.New().SetSeed(ssd.GenerationSeed() + "Satelites")
	for s, sts := range sn.StarSystems {

		sun := sn.StarSystems[s].Sun
		if sun == nil {
			continue
		}
		hz := sun.HZ()
		for orb, pb := range sts.Body {
			if pb.pbType == "EMPTY" {
				continue
			}
			hzPlace := "inner"
			dm := -5
			switch {
			case pb.pbType == "Hospitable":
				dm = -4
			case pb.pbType == "GG":
				dm = -1
			case orb-hz > 1:
				dm = -3
				hzPlace = "outer"
			}
			sats := rollSats(dp, dm)
			for snum, sat := range sats {
				switch sat {
				case "S":
					sType := satType(hzPlace, dp.Roll("1d6").Sum())
					sn.StarSystems[s].Body[orb].satelites = append(sn.StarSystems[s].Body[orb].satelites, placeWorldTo(fmt.Sprintf("Sat %v (%v)", snum, sType), sType, snum))
				case "R":
					sn.StarSystems[s].Body[orb].satelites = append(sn.StarSystems[s].Body[orb].satelites, placeWorldTo(fmt.Sprintf("Ring %v", snum), "Ring", snum))
				}
			}
		}
	}
}

func satType(place string, index int) string {
	switch place {
	case "inner":
		switch index {
		case 1:
			return "Inferno"
		case 2:
			return "InnerWorld"
		case 3:
			return "BigWorld"
		case 4:
			return "StormWorld"
		case 5:
			return "RadWorld"
		case 6:
			return "Hospitable"
		}
	case "outer":
		switch index {
		case 1:
			return "Worldlet"
		case 2:
			return "IceWorld"
		case 3:
			return "BigWorld"
		case 4:
			return "StormWorld"
		case 5:
			return "RadWorld"
		case 6:
			return "IceWorld"
		}
	}
	return "satellite type error"
}

func rollSats(dp *dice.Dicepool, dm int) []string {
	sats := []string{}
	for dp.Roll("1d6").DM(dm).Sum() == 0 {
		sats = append(sats, "R")
	}
	for i := 0; i < dp.Sum(); i++ {
		sats = append(sats, "S")
	}
	return sats
}

func (sn *StarNexus) PlaceMainWorld() error {
	_, _, gg, err := calculations.Decode(sn.ssd.PBG())
	if err != nil {
		return err
	}
	habitZone := sn.StarSystems[0].Sun.HZ()
	remarks := sn.ssd.MW_Remarks()
	//Place if NOT Satellite
	if !helper.SliceStrContains(strings.Split(remarks, " "), "Sa") {
		sn.StarSystems[0].Body[habitZone] = placeWorldTo("MainWorld", "MW", habitZone)
		return nil
	}
	//Place if IS Satellite
	switch {
	default:
		ggSize, ggType := newGasGigantData(sn.ssd.GenerationSeed() + "MW_GG")
		sn.StarSystems[0].Body[habitZone] = placeWorldTo(fmt.Sprintf("Gas Gigant 0 (%v-%v)", ggSize, ggType), "GG", habitZone)
	case gg < 1:
		sn.StarSystems[0].Body[habitZone] = placeWorldTo("BigWorld w", "BW", habitZone)
	}
	sn.StarSystems[0].Body[habitZone].satelites = append(sn.StarSystems[0].Body[habitZone].satelites, placeWorldTo("Mainworld", "MW_Sat", 0))
	//fmt.Println(sn.StarSystems[0].Body[habitZone].satelites[0].Name(), 0, habitZone, 0)
	return nil
}

func (sn *StarNexus) PlaceGasGigants(ssd SurveyReporter) error {
	dp := dice.New().SetSeed(ssd.GenerationSeed() + "Place GG")
	_, _, gg, err := calculations.Decode(ssd.PBG())
	gg = gg - sn.ggPlaced()
	if err != nil {
		return err
	}
	if gg < 1 {
		return nil
	}
	for g := 0; g < gg; g++ {
		ggSize, ggType := newGasGigantData(ssd.GenerationSeed() + fmt.Sprintf("_gas gigant %v", g))
		placed := false
		for !placed {
			i := dp.Roll("1d" + fmt.Sprintf("%v", len(sn.StarSystems))).DM(-1).Sum()
			if !sn.StarSystems[i].haveUnfilledOrbits() {
				continue
			}
			tryOrbit := -1
			switch ggType {
			case "LGG":
				tryOrbit = rollLGGplacement(dp) + sn.StarSystems[i].Sun.HZ()
			case "SGG":
				r := dp.Roll("1d2").Sum()
				switch r {
				case 1:
					tryOrbit = rollSGGplacement(dp) + sn.StarSystems[i].Sun.HZ()
				case 2:
					ggType = "IG"
					tryOrbit = rollIGGplacement(dp) + sn.StarSystems[i].Sun.HZ()
				}

			}
			if tryOrbit < 0 {
				continue
			}
			if len(sn.StarSystems[i].Body)-1 < tryOrbit {
				continue
			}
			if sn.StarSystems[i].Body[tryOrbit].pbType != "EMPTY" {
				continue
			}
			sn.StarSystems[i].Body[tryOrbit] = placeWorldTo(fmt.Sprintf("Gas Gigant %v (%v-%v)", g, ggSize, ggType), "GG", tryOrbit)
			placed = true
		}
	}
	return fmt.Errorf("Not Implemented")
}

func (sn *StarNexus) PlaceBelts(ssd SurveyReporter) error {
	dp := dice.New().SetSeed(ssd.GenerationSeed() + "Place GG")
	_, belts, _, err := calculations.Decode(ssd.PBG())
	if err != nil {
		return err
	}
	if belts < 1 {
		return nil
	}
	for belt := 0; belt < belts; belt++ {
		//ggSize, ggType := newGasGigantData(ssd.GenerationSeed() + fmt.Sprintf("_Belt %v", belt))
		placed := false
		for !placed {
			i := dp.Roll("1d" + fmt.Sprintf("%v", len(sn.StarSystems))).DM(-1).Sum()
			if !sn.StarSystems[i].haveUnfilledOrbits() {
				continue
			}
			tryOrbit := -1
			tryOrbit = rollBeltsPlacement(dp) + sn.StarSystems[i].Sun.HZ()
			if tryOrbit < 0 {
				continue
			}
			if len(sn.StarSystems[i].Body)-1 < tryOrbit {
				continue
			}
			if sn.StarSystems[i].Body[tryOrbit].pbType != "EMPTY" {
				continue
			}
			sn.StarSystems[i].Body[tryOrbit] = placeWorldTo(fmt.Sprintf("Belt %v", belt), "Belt", tryOrbit)
			placed = true
		}
	}

	return fmt.Errorf("Not Implemented")
}

func (sn *StarNexus) PlaceOther(ssd SurveyReporter) error {
	dp := dice.New().SetSeed(ssd.GenerationSeed() + "Other")
	_, b, g, _ := calculations.Decode(ssd.PBG())
	worlds := ssd.Worlds() - b - g - 1
	if worlds < 1 {
		return nil
	}
	for w := 0; w < worlds; w++ {
		//ggSize, ggType := newGasGigantData(ssd.GenerationSeed() + fmt.Sprintf("_Belt %v", belt))
		placed := false
		tr := 0
		for !placed {
			tr++
			i := dp.Roll("1d" + fmt.Sprintf("%v", len(sn.StarSystems))).DM(-1).Sum()
			if !sn.StarSystems[i].haveUnfilledOrbits() {
				continue
			}
			tryOrbit := -1
			switch {
			default:
				tryOrbit = rollOtherPlacement(dp) + tr/100
			case w+1 == worlds:
				tryOrbit = rollOther2Placement(dp)
			}
			if tryOrbit < 0 {
				//fmt.Println("err low orbit", tryOrbit, w, worlds)
				continue
			}
			if len(sn.StarSystems[i].Body)-1 < tryOrbit {
				//fmt.Println("err High orbit", tryOrbit, w, worlds)
				continue
			}
			if sn.StarSystems[i].Body[tryOrbit].pbType != "EMPTY" {
				//fmt.Println("err filled orbit", tryOrbit, w, worlds)
				continue
			}
			worldType := ""
			switch {
			case tryOrbit-sn.StarSystems[i].Sun.HZ() > 1:
				worldType = innerWorldType(dp.Roll("1d6").Sum())
			default:
				worldType = outerWorldType(dp.Roll("1d6").Sum())
			}

			sn.StarSystems[i].Body[tryOrbit] = placeWorldTo(fmt.Sprintf("Other %v", w), worldType, tryOrbit)
			placed = true
		}
	}

	return fmt.Errorf("Not Implemented")
}

func innerWorldType(i int) string {
	switch i {
	case 1:
		return "Inferno"
	case 2:
		return "InnerWorld"
	case 3:
		return "BigWorld"
	case 4:
		return "StormWorld"
	case 5:
		return "RadWorld"
	case 6:
		return "Hospitable"
	}
	return "world type error"
}

func outerWorldType(i int) string {
	switch i {
	case 1:
		return "Worldlet"
	case 2:
		return "IceWorld"
	case 3:
		return "BigWorld"
	case 4:
		return "IceWorld"
	case 5:
		return "RadWorld"
	case 6:
		return "Iceworld"
	}
	return "world type error"
}

func (sn *StarNexus) ggPlaced() int {
	ggp := 0
	for _, stm := range sn.StarSystems {
		for _, b := range stm.Body {
			if b.pbType == "GG" {
				ggp++
			}
		}
	}
	return ggp
}

func rollLGGplacement(dp *dice.Dicepool) int {
	r := dp.Roll("2d6").Sum()
	return r - 5
}

func rollSGGplacement(dp *dice.Dicepool) int {
	r := dp.Roll("2d6").Sum()
	return r - 4
}

func rollIGGplacement(dp *dice.Dicepool) int {
	r := dp.Roll("2d6").Sum()
	return r - 11
}

func rollBeltsPlacement(dp *dice.Dicepool) int {
	r := dp.Roll("2d6").Sum()
	return r - 4
}

func rollOtherPlacement(dp *dice.Dicepool) int {
	r := dp.Roll("2d6").Sum()
	orb := []int{10, 8, 6, 4, 2, 0, 1, 3, 5, 7, 9}
	return orb[r-2]
}

func rollOther2Placement(dp *dice.Dicepool) int {
	r := dp.Roll("2d6").Sum()
	orb := []int{17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7}
	return orb[r-2]
}

func (sn *StarNexus) hasAnyUnfilled() bool {
	for _, ss := range sn.StarSystems {
		if ss.haveUnfilledOrbits() {
			return true
		}
	}
	return false
}

func newGasGigantData(seed string) (string, string) {
	dp := dice.New().SetSeed(seed)
	s := dp.Roll("2d6").DM(19).Sum()
	t := "LGG"
	if s < 23 {
		t = "SGG"
	}

	return ehex.New().Set(s).Code(), t
}

func (n *StarNexus) String() string {
	str := ""
	for i, st := range n.StarSystems {
		if st.Sun != nil {
			str += fmt.Sprintf("Sun: %v - %v\n", st.Sun.Name(), st.Sun.Code())
		}
		if st.Companion != nil {
			str += fmt.Sprintf("Companion: %v - %v\n", st.Companion.Name(), st.Companion.Code())
		}
		for k, b := range n.StarSystems[i].Body {
			if b.pbType == "EMPTY" {
				continue
			}
			str += fmt.Sprintf("    Body: %v - %v (%v)\n", b.Name(), b.Orbit(), n.StarSystems[i].Body[k].pbType)
			for _, s := range n.StarSystems[i].Body[k].satelites {
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
			//fmt.Println("orbit:", st.Name(), st.Orbit())
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
	//fmt.Println(len(n.StarSystems))
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
