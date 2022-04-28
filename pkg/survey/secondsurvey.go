package survey

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/survey/calculations"

	"github.com/Galdoba/utils"
)

const (
	reserved = iota
	MW_Name
	Hex
	MW_UWP
	PBG
	TravelZone
	Bases
	Allegiance
	Stellar
	SubSector
	MW_Importance
	MW_ImportanceInt
	MW_Economic
	MW_Cultural
	MW_Nobility
	Worlds
	RU
	SubSectorInt
	Quadrant
	CoordX
	CoordY
	MW_Remarks
	BasesOld
	Sector
	SubSectorName
	SectorAbb
	AllegianceExt
	Seed
	END_OF_SURVEY_DATA
	cleanedDataPath = "c:\\Users\\Public\\TrvData\\cleanedData.txt"
)

type SecondSurveyData struct {
	coordX           int
	coordY           int
	sector           string
	hex              string
	mw_Name          string
	mw_UWP           string
	mw_Remarks       string
	mw_Importance    string
	mw_ImportanceInt int
	mw_Economic      string
	mw_Cultural      string
	mw_Nobility      string
	bases            string
	travelZone       string
	pbg              string
	worlds           int
	allegiance       string
	stellar          string
	ru               int
	input            string //temp
	subSector        string
	subSectorInt     int
	quadrant         int
	basesOld         string
	sectorAbb        string
	subSectorName    string
	allegianceExt    string
	seed             string
	errors           []error
}

func Parse(input string) *SecondSurveyData {
	ssd := SecondSurveyData{}
	//ssd.input = input
	data := strings.Split(input, "|")
	for i := range data {
		data[i] = strings.TrimSpace(data[i])
	}
	for len(data) < END_OF_SURVEY_DATA+1 {
		data = append(data, "")
	}
	ssd.mw_Name = data[1]
	ssd.hex = data[2]
	ssd.mw_UWP = data[3]
	ssd.pbg = data[4]
	ssd.travelZone = data[5]
	ssd.bases = data[6]
	ssd.allegiance = data[7]
	ssd.stellar = data[8]
	ssd.subSector = data[9]
	ssd.mw_Importance = data[10]
	impInt, errImp := strconv.Atoi(data[11])
	if errImp != nil {
		ssd.errors = append(ssd.errors, errImp)
	}
	ssd.mw_ImportanceInt = impInt
	ssd.mw_Economic = data[12]
	ssd.mw_Cultural = data[13]
	ssd.mw_Nobility = data[14]
	worlds, errWorlds := strconv.Atoi(data[15])
	if errWorlds != nil {
		ssd.errors = append(ssd.errors, errWorlds)
	}
	ssd.worlds = worlds
	ru, errRu := strconv.Atoi(data[16])
	if errRu != nil {
		ssd.errors = append(ssd.errors, errRu)
	}
	ssd.ru = ru
	ssInt, errssInt := strconv.Atoi(data[17])
	if errssInt != nil {
		ssd.errors = append(ssd.errors, errssInt)
	}
	ssd.subSectorInt = ssInt
	ssQuad, errQuad := strconv.Atoi(data[18])
	if errQuad != nil {
		ssd.errors = append(ssd.errors, errQuad)
	}
	ssd.quadrant = ssQuad
	xCoord, errXcoord := strconv.Atoi(data[19])
	if errXcoord != nil {
		ssd.errors = append(ssd.errors, errXcoord)
	}
	ssd.coordX = xCoord
	yCoord, errYcoord := strconv.Atoi(data[20])
	if errYcoord != nil {
		ssd.errors = append(ssd.errors, errYcoord)
	}
	ssd.coordY = yCoord
	ssd.mw_Remarks = data[21]
	ssd.basesOld = data[22]
	ssd.sector = data[23]
	ssd.subSectorName = data[24]
	ssd.sectorAbb = data[25]
	ssd.allegianceExt = data[26]
	ssd.seed = data[27]
	ssd.verify()
	return &ssd
}

func (ssd *SecondSurveyData) Compress() string {
	compressed := "|"
	compressed += fmt.Sprintf("%v|", ssd.mw_Name)
	compressed += fmt.Sprintf("%v|", ssd.hex)
	compressed += fmt.Sprintf("%v|", ssd.mw_UWP)
	compressed += fmt.Sprintf("%v|", ssd.pbg)
	compressed += fmt.Sprintf("%v|", ssd.travelZone)
	compressed += fmt.Sprintf("%v|", ssd.bases)
	compressed += fmt.Sprintf("%v|", ssd.allegiance)
	compressed += fmt.Sprintf("%v|", ssd.stellar)
	compressed += fmt.Sprintf("%v|", ssd.subSector)
	compressed += fmt.Sprintf("%v|", ssd.mw_Importance) //10
	compressed += fmt.Sprintf("%v|", ssd.mw_ImportanceInt)
	compressed += fmt.Sprintf("%v|", ssd.mw_Economic)
	compressed += fmt.Sprintf("%v|", ssd.mw_Cultural)
	compressed += fmt.Sprintf("%v|", ssd.mw_Nobility)
	compressed += fmt.Sprintf("%v|", ssd.worlds)
	compressed += fmt.Sprintf("%v|", ssd.ru)
	compressed += fmt.Sprintf("%v|", ssd.subSectorInt) //17
	compressed += fmt.Sprintf("%v|", ssd.quadrant)
	compressed += fmt.Sprintf("%v|", ssd.coordX)
	compressed += fmt.Sprintf("%v|", ssd.coordY)
	compressed += fmt.Sprintf("%v|", ssd.mw_Remarks) //21
	compressed += fmt.Sprintf("%v|", ssd.basesOld)   //22
	compressed += fmt.Sprintf("%v|", ssd.sector)     //23
	compressed += fmt.Sprintf("%v|", ssd.subSectorName)
	compressed += fmt.Sprintf("%v|", ssd.sectorAbb)
	compressed += fmt.Sprintf("%v", ssd.allegianceExt)

	return compressed
}

func (ssd *SecondSurveyData) containsErrors() bool {
	for _, val := range ssd.errors {
		if val != nil {
			return true
		}
	}
	return false
}

func (ssd *SecondSurveyData) verify() {
	if ssd.mw_Name == "" {
		ssd.mw_Name = ssd.NameByConvention()
	}
	if ssd.stellar == "" {
		ssd.stellar = calculations.GenerateNewStellar(ssd.GenerationSeed())
	}
	if !calculations.UWPvalid(ssd.mw_UWP) {
		ssd.mw_UWP = calculations.FixUWP(ssd.mw_UWP, ssd.GenerationSeed())
	}
	if !calculations.PBGvalid(ssd.pbg, ssd.mw_UWP) {
		ssd.pbg = calculations.FixPBG(ssd.pbg, ssd.mw_UWP, ssd.GenerationSeed())
	}
	if ssd.mw_Importance == "{+?}" {
		ssd.mw_Importance = importanceToString(ssd.mw_ImportanceInt)
		calc := calculations.Importance(ssd.mw_UWP, ssd.bases, ssd.mw_Remarks)
		if calc != ssd.mw_ImportanceInt && ssd.mw_ImportanceInt == 0 {
			ssd.mw_Importance = importanceToString(calc)
			ssd.mw_ImportanceInt = calc
		}
	}
	if importanceToInt(ssd.mw_Importance) != ssd.mw_ImportanceInt {
		ssd.mw_Importance = importanceToString(ssd.mw_ImportanceInt)
	}
	if !calculations.ExValid(ssd.mw_Economic) {
		ssd.mw_Economic = calculations.FixEconomicExtention(ssd.mw_Economic, ssd.mw_UWP, ssd.pbg, ssd.GenerationSeed(), ssd.mw_ImportanceInt)
	}
	if calculations.RU(ssd.mw_Economic) != ssd.ru {
		ssd.ru = calculations.RU(ssd.mw_Economic)
	}
	if !calculations.CxValid(ssd.mw_Cultural, ssd.mw_UWP) {
		ssd.mw_Cultural = calculations.Cultural(ssd.mw_UWP, ssd.GenerationSeed(), ssd.mw_ImportanceInt)
	}
	if !calculations.WorldsValid(ssd.worlds, ssd.pbg) {
		ssd.worlds = calculations.FixWorlds(ssd.pbg, ssd.GenerationSeed())
	}
	if len(calculations.NobilityErrors(ssd.mw_Nobility, strings.Fields(ssd.mw_Remarks), ssd.mw_ImportanceInt)) != 0 {
		ssd.mw_Nobility = calculations.FixNobility(strings.Fields(ssd.mw_Remarks), ssd.mw_ImportanceInt)
	}
	if !calculations.CxValid(ssd.mw_Cultural, ssd.mw_UWP) {
		fmt.Println("invalid culture data:", ssd.mw_Cultural)
	}
	if calculations.AllegianceFull(ssd.allegiance) == "UNKNOWN SHORTFORM" {
		ssd.allegiance = "XXXX"
		ssd.allegianceExt = calculations.AllegianceFull(ssd.allegiance)
	}
	switch {
	default:
		return
	case ssd.mw_Name == "":
		ssd.errors = append(ssd.errors, fmt.Errorf("mainworld name missing (fixed)"))
	case ssd.stellar == "":
		ssd.errors = append(ssd.errors, fmt.Errorf("stellar data missing (f)"))
	case ssd.hex == "":
		ssd.errors = append(ssd.errors, fmt.Errorf("hex data missing"))
	case !calculations.PBGvalid(ssd.pbg, ssd.mw_UWP):
		ssd.errors = append(ssd.errors, fmt.Errorf(" pbg data not valid"))
	case ssd.mw_Importance == "{+?}":
		ssd.errors = append(ssd.errors, fmt.Errorf("importance data does not present correctly (fixable)"))
	case ssd.mw_Economic == "(???+?)":
		ssd.errors = append(ssd.errors, fmt.Errorf("economic Not calculated"))
	case !calculations.ExValid(ssd.mw_Economic):
		ssd.errors = append(ssd.errors, fmt.Errorf("economic Not Valid"))
	case ssd.mw_Economic == "":
		ssd.errors = append(ssd.errors, fmt.Errorf("economic data missing"))
	case ssd.mw_Cultural == "[????]":
		ssd.errors = append(ssd.errors, fmt.Errorf("cultural data Not calculated"))
	case importanceToInt(ssd.mw_Importance) != ssd.mw_ImportanceInt:
		ssd.errors = append(ssd.errors, fmt.Errorf("importance data does not match"))
	case calculations.RU(ssd.mw_Economic) != ssd.ru:
		ssd.errors = append(ssd.errors, fmt.Errorf("projected Ru does not match actual"))
	case !calculations.CxValid(ssd.mw_Cultural, ssd.mw_UWP):
		ssd.errors = append(ssd.errors, fmt.Errorf("culture data invalid"))
	case !calculations.WorldsValid(ssd.worlds, ssd.pbg):
		ssd.errors = append(ssd.errors, fmt.Errorf("world number incorrect (have %v)", ssd.worlds))
	case len(calculations.NobilityErrors(ssd.mw_Nobility, strings.Fields(ssd.mw_Remarks), ssd.mw_ImportanceInt)) != 0:
		ssd.errors = append(ssd.errors, calculations.NobilityErrors(ssd.mw_Nobility, strings.Fields(ssd.mw_Remarks), ssd.mw_ImportanceInt)...)
	case calculations.AllegianceFull(ssd.allegiance) == "UNKNOWN SHORTFORM":
		ssd.errors = append(ssd.errors, fmt.Errorf("allegiance unknown"))
	}
}

func (ssd *SecondSurveyData) NameByConvention() string {
	x := ssd.coordX
	pX := "S"
	if x < 0 {
		x = x * -1
		pX = "T"
	}
	y := ssd.coordY
	pY := "R"
	if y < 0 {
		y = y * -1
		pY = "C"
	}
	return fmt.Sprintf("%v %v/%v%v-%v%v", ssd.sector, ssd.hex, pX, x, pY, y)
}

func (ssd *SecondSurveyData) GenerationSeed() string {
	x := ssd.coordX
	pX := "S"
	if x < 0 {
		x = x * -1
		pX = "T"
	}
	y := ssd.coordY
	pY := "R"
	if y < 0 {
		y = y * -1
		pY = "C"
	}
	return fmt.Sprintf("%v%v%v%v%v%v%v%v", ssd.sector, ssd.hex, pX, x, pY, y, ssd.mw_Name, ssd.seed)
}

func importanceToInt(str string) int {
	switch str {
	default:
		return -999
	case "{ -5 }":
		return -5
	case "{ -4 }":
		return -4
	case "{ -3 }":
		return -3
	case "{ -2 }":
		return -2
	case "{ -1 }":
		return -1
	case "{ +0 }":
		return 0
	case "{ +1 }":
		return 1
	case "{ +2 }":
		return 2
	case "{ +3 }":
		return 3
	case "{ +4 }":
		return 4
	case "{ +5 }":
		return 5
	case "{ +6 }":
		return 6
	}
}

func importanceToString(i int) string {
	switch i {
	default:
		return "{+?}"
	case -5:
		return "{ -5 }"
	case -4:
		return "{ -4 }"
	case -3:
		return "{ -3 }"
	case -2:
		return "{ -2 }"
	case -1:
		return "{ -1 }"
	case 0:
		return "{ +0 }"
	case 1:
		return "{ +1 }"
	case 2:
		return "{ +2 }"
	case 3:
		return "{ +3 }"
	case 4:
		return "{ +4 }"
	case 5:
		return "{ +5 }"
	case 6:
		return "{ +6 }"
	}
}

func (ssd *SecondSurveyData) String() string {
	rep := ssd.sectorAbb + "   "
	rep += ssd.hex + "   "
	rep += ssd.mw_Name + "   "
	rep += ssd.mw_UWP + "   "
	rep += ssd.mw_Remarks + "   "
	rep += ssd.mw_Importance + "   "
	rep += ssd.mw_Economic + "   "
	rep += ssd.mw_Cultural + "   "
	rep += ssd.mw_Nobility + "   "
	rep += ssd.bases + "   "
	rep += ssd.travelZone + "   "
	rep += ssd.pbg + "   "
	rep += strconv.Itoa(ssd.worlds) + "   "
	rep += ssd.allegiance + "   "
	rep += ssd.stellar
	return rep
}

func ListOf(ssds []*SecondSurveyData) []string {
	if len(ssds) < 1 {
		return nil
	}
	sample := ssds[0].String()
	fields := strings.Split(sample, "   ")
	colMap := make(map[int]int)
	for f := range fields {
		for _, ssd := range ssds {
			testFields := strings.Split(ssd.String(), "   ")
			if colMap[f] < len(testFields[f]) {
				colMap[f] = len(testFields[f])
			}
		}
	}
	table := []string{}
	for _, ssd := range ssds {
		newFields := strings.Split(ssd.String(), "   ")
		line := "|"
		for n, fld := range newFields {
			for len(fld) < colMap[n] {
				fld += " "
			}
			line += fld + "|"
		}
		table = append(table, line)
	}
	return table
}

func Search(key string) ([]*SecondSurveyData, error) {
	var ssdArr []*SecondSurveyData
	lines := utils.LinesFromTXT(cleanedDataPath)
	for _, val := range lines {
		if strings.Contains(val, key) {
			ssdArr = append(ssdArr, Parse(val))
		}
	}
	if len(ssdArr) < 1 {
		return nil, fmt.Errorf("nothing was found")
	}
	return ssdArr, nil
}

func SearchByCoordinates(x, y int) (*SecondSurveyData, error) {
	lines := utils.LinesFromTXT(cleanedDataPath)
	for _, val := range lines {
		if !strings.Contains(val, fmt.Sprintf("|%v|%v|", x, y)) {
			continue
		}
		ssd := Parse(val)
		if ssd.CoordX() == x && ssd.CoordY() == y {
			return ssd, nil
		}

	}
	return nil, fmt.Errorf("no entry on set coordinates")
}

func (ssd *SecondSurveyData) CoordX() int {
	return ssd.coordX
}
func (ssd *SecondSurveyData) CoordY() int {
	return ssd.coordY
}
func (ssd *SecondSurveyData) Sector() string {
	return ssd.sector
}
func (ssd *SecondSurveyData) Hex() string {
	return ssd.hex
}
func (ssd *SecondSurveyData) MW_Name() string {
	return ssd.mw_Name
}
func (ssd *SecondSurveyData) MW_UWP() string {
	return ssd.mw_UWP
}
func (ssd *SecondSurveyData) MW_Remarks() string {
	return ssd.mw_Remarks
}
func (ssd *SecondSurveyData) MW_Importance() string {
	return ssd.mw_Importance
}
func (ssd *SecondSurveyData) MW_ImportanceInt() int {
	return ssd.mw_ImportanceInt
}
func (ssd *SecondSurveyData) MW_Economic() string {
	return ssd.mw_Economic
}
func (ssd *SecondSurveyData) MW_Cultural() string {
	return ssd.mw_Cultural
}
func (ssd *SecondSurveyData) MW_Nobility() string {
	return ssd.mw_Nobility
}
func (ssd *SecondSurveyData) Bases() string {
	return ssd.bases
}
func (ssd *SecondSurveyData) TravelZone() string {
	return ssd.travelZone
}
func (ssd *SecondSurveyData) PBG() string {
	return ssd.pbg
}
func (ssd *SecondSurveyData) Worlds() int {
	return ssd.worlds
}
func (ssd *SecondSurveyData) Allegiance() string {
	return ssd.allegiance
}
func (ssd *SecondSurveyData) Stellar() string {
	return ssd.stellar
}
func (ssd *SecondSurveyData) RU() int {
	return ssd.ru
}
func (ssd *SecondSurveyData) SubSector() string {
	return ssd.subSector
}
func (ssd *SecondSurveyData) SubSectorInt() int {
	return ssd.subSectorInt
}
func (ssd *SecondSurveyData) Quadrant() int {
	return ssd.quadrant
}
func (ssd *SecondSurveyData) BasesOld() string {
	return ssd.basesOld
}
func (ssd *SecondSurveyData) SectorAbb() string {
	return ssd.sectorAbb
}
func (ssd *SecondSurveyData) SubSectorName() string {
	return ssd.subSectorName
}
func (ssd *SecondSurveyData) AllegianceExt() string {
	return ssd.allegianceExt
}
