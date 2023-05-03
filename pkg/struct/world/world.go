package world

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TravellerTools/pkg/classifications"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/generation/planets"
	"github.com/Galdoba/TravellerTools/pkg/generation/star"
	"github.com/Galdoba/TravellerTools/pkg/generation/stellar"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/t5/genetics"
	"github.com/Galdoba/devtools/errmaker"
)

/*
World. A planet or satellite.
-----------------------------
Profile Keys:
OrbitPair [A-D] - не имеет значения в генерации
OrbitNum  [0-X]
OrbitSat  [0-Y]
HZvar [0-F] h
//SateliteCode [0-2] h } уходят и переносятся в блок Орбиты
//SateliteOrb [0-F] h  }

PlanetType [A-J]
Size [0-K]
Atmo [0-F]
Hydr [0-A]
//Climate [0-6] h - не нужно. заменено HZvar
Life [0-A] h
Life_Compat [0-A] h

Pops [0-F]
PopDigit [0-9] h
Govr [0-F]
Laws [0-J]
Tech [0-N] s


Starport [ABCDEFGHXY]

OtherBelts [0-3] h - нужно только в рамках системы или для MW
GasG [0-4] h  - нужно только в рамках системы или для MW

*/

type World struct {
	Alias           string //Самоназвание
	Catalog         string //номенклатурное название (121-311 Stargos A III 2d)
	profile         profile.Profile
	HomeStar        star.StarBody
	prim            string
	comp            string
	nativeGenome    genetics.Genome
	Flag            map[string]bool
	classifications []int
}

func (w *World) String() string {
	orbS := " "
	satS := " "
	satHex := w.Data(profile.KEY_SATELITE_ORBIT)
	if satHex == "*" {
		orbS = w.Data(profile.KEY_PLANETARY_ORBIT)
	} else {
		satS = satHex
	}

	uwp := ""
	uwp += w.Data(profile.KEY_PORT)
	uwp += w.Data(profile.KEY_SIZE)
	uwp += w.Data(profile.KEY_ATMO)
	uwp += w.Data(profile.KEY_HYDR)
	uwp += w.Data(profile.KEY_POPS)
	uwp += w.Data(profile.KEY_GOVR)
	uwp += w.Data(profile.KEY_LAWS)
	uwp += "-"
	uwp += w.Data(profile.KEY_TL)
	tc := ""
	for _, code := range w.classifications {
		tc += classifications.Call(code).String() + " "
	}
	str := fmt.Sprintf("%v	%v	%v	%v", orbS, satS, uwp, tc)
	return str
}

type knownData struct {
	key  string
	val  string
	used bool
}

const (
	FLAG_TRUE          = "True"
	FLAG_FALSE         = "False"
	Alias              = "Alias"
	Catalog            = "Catalog Name"
	Primary            = "Primary Star"
	Companion          = "Companion Star"
	IsMainworld        = "MW"
	IsNotMainworld     = "Not-MW"
	IsPlanet           = "Planet"
	IsCloseSat         = "Close Satelite"
	IsFarSat           = "Far Satelite"
	IsNotColonized     = "Not Colonized"
	IsAmberZone        = "Amber Zone"
	IsRedZone          = "Red Zone"
	IsMilitaryRule     = "Military Rule"
	IsResearchLab      = "Research Lab"
	IsSubsectorCapital = "Subsector Capital"
	IsSectorCapital    = "Sector Capital"
	IsCapital          = "Capital"
	IsColony           = "Colony"
	IsDataRepository   = "Data Repository"
	IsAntientSite      = "Antient Site"
)

/*
New(Inject, Flags... bool) - Создаёт объект
	Inject([]KnownData{key string, data string}) - обязательно! Добавляет известные заранее данные
		KnownData(key string, val string)
		Flags - Добавляет заранее известные флаги для планеты (например MW). используем Инекцию со значением FLAG_TRUE/FLAG_FALSE
*/

func NewWorld(kndt []*knownData, flags ...bool) (*World, error) {
	w := World{}
	//w.Catalog = cat
	w.profile = profile.New()
	w.Flag = make(map[string]bool)
	for _, kd := range kndt {
		if err := w.injectData(kd); err != nil {
			return &w, errmaker.ErrorFrom(err, kd)
		}
	}
	dice := dice.New().SetSeed(w.Catalog)
	if w.prim == "" {
		w.prim = stellar.GenerateStellarOneStar(dice)
	}
	pair, err := star.NewPair(w.prim, w.comp)
	if err != nil {
		return &w, errmaker.ErrorFrom(err, w.prim, w.comp)
	}
	w.HomeStar = pair
	w.Generate(dice)
	w.classifications = classifications.Evaluate(&w)
	return &w, nil
}

func (w *World) injectData(kd *knownData) error {
	if kd.used {
		return errmaker.ErrorFrom(fmt.Errorf("data was already used"), kd.key, kd.val)
	}
	kd.used = true
	if kd.val == FLAG_TRUE || kd.val == FLAG_FALSE {
		flag, err := strconv.ParseBool(kd.val)
		if err != nil {
			return errmaker.ErrorFrom(err, kd)
		}
		w.Flag[kd.key] = flag
		return nil
	}
	if err := w.defineMW(); err != nil {
		return errmaker.ErrorFrom(err)
	}
	switch kd.key {
	default:
		w.profile.Inject(kd.key, kd.val)
	case Primary:
		w.prim = kd.val
	case Companion:
		w.comp = kd.val
	case Catalog:
		w.SetCatalogueName(kd.val)
	}
	return nil
}

func Inject(kndt ...*knownData) []*knownData {
	kd := []*knownData{}
	kd = append(kd, kndt...)
	if len(kd) == 0 {
		return nil
	}
	return kd
}

func KnownData(key, val string) *knownData {
	return &knownData{key, val, false}
}

func (w *World) SetCatalogueName(cat string) {
	w.Catalog = cat
}

var ErrFullyGenerated = fmt.Errorf("World generation complete")

func (w *World) Generate(dice *dice.Dicepool) error {
	checkList := len(profile.UWPkeys())
	injected := 0
	for _, key := range profile.UWPkeys() {
		var err error
		if w.profile.Data(key) != nil {
			fmt.Println("key", key, "was injected or generated")
			injected++
			continue
		}
		fmt.Println("generating key:", key)
		switch key {
		default:
			fmt.Println(key, "generation is NOT IMPLEMENTED")
			continue
		case profile.KEY_PLANETARY_ORBIT:
			err = w.generateOrbitAndHZvar(dice)
		case profile.KEY_SATELITE_ORBIT:
			err = w.generateSateliteOrbit(dice)
		case profile.KEY_WORLDTYPE:
			err = w.generateWorldType(dice)
		case profile.KEY_LIMIT_size:
			err = w.generateSizeLimit(dice)
		case profile.KEY_SIZE:
			err = w.generateSize(dice)
		case profile.KEY_ATMO:
			err = w.generateAtmo(dice)
		case profile.KEY_HYDR:
			err = w.generateHydr(dice)
		case profile.KEY_LIFE_FACTOR:
			err = w.generateDominantLife(dice)
		case profile.KEY_LIFE_COMPATABILITY:
			err = w.determineLifeCompatability(dice)
		case profile.KEY_LIMIT_pops:
			w.SetDefaultPopulationLimit()
		case profile.KEY_POPS:
			err = w.generatePops(dice)
		case profile.KEY_POP_DIGIT:
			err = w.generatePopDigit(dice)
		case profile.KEY_GOVR:
			err = w.generateGovr(dice)
		case profile.KEY_LAWS:
			err = w.generateLaws(dice)
		case profile.KEY_LIMIT_tl:
			err = w.setupMinimumTL()
		case profile.KEY_PORT:
			err = w.generatePort(dice)
		case profile.KEY_TL:
			err = w.generateTL(dice)
		case profile.KEY_BASES:
			err = w.generateBases(dice)
		}
		if err != nil {
			return errmaker.ErrorFrom(err, key)
		}
		fmt.Println("conclude generating", key, w.profile.Data(key).Code())
	}
	if injected == checkList {
		return ErrFullyGenerated
	}
	return nil
}

func (w *World) generateOrbitAndHZvar(dice *dice.Dicepool) error {
	starHZ := w.HomeStar.HabitableZone()
	switch w.Flag[IsMainworld] {
	case true:
		baseOrb := starHZ
		hzVar := w.profile.Data(profile.KEY_HABITABLE_ZONE_VAR)
		if hzVar == nil {
			hzVar = planets.RollMainworldHabitableZone(dice)
		}
		w.profile.Inject(profile.KEY_HABITABLE_ZONE_VAR, hzVar)
		realOrbit := baseOrb + w.profile.Data(profile.KEY_HABITABLE_ZONE_VAR).Value() - 10
		w.profile.Inject(profile.KEY_PLANETARY_ORBIT, realOrbit)
	case false:
		baseOrb := planets.GeneratePlanetOrbit(dice)
		hzVar := planets.HabitableZoneVar(starHZ, baseOrb.Value())
		w.profile.Inject(profile.KEY_HABITABLE_ZONE_VAR, hzVar)
		w.profile.Inject(profile.KEY_PLANETARY_ORBIT, baseOrb)
	}
	if w.profile.Data(profile.KEY_HABITABLE_ZONE_VAR) == nil {
		return fmt.Errorf("profile.KEY_HABITABLE_ZONE_VAR was not generated")
	}
	if w.profile.Data(profile.KEY_PLANETARY_ORBIT) == nil {
		return fmt.Errorf("profile.KEY_PLANETARY_ORBIT was not generated")
	}
	return nil
}

func (w *World) generateSateliteOrbit(dice *dice.Dicepool) error {
	flagsConfirmed := 0
	flagVal := -1
	for i, b := range []bool{w.Flag[IsPlanet], w.Flag[IsCloseSat], w.Flag[IsFarSat]} {
		if b {
			flagsConfirmed++
			flagVal = i
			if flagsConfirmed > 1 {
				err := fmt.Errorf("flag input controversy")
				return errmaker.ErrorFrom(err, IsPlanet, w.Flag[IsPlanet], IsCloseSat, w.Flag[IsCloseSat], IsFarSat, w.Flag[IsFarSat])
			}
		}
	}
	w.profile.Inject(profile.KEY_SATELITE_ORBIT, planets.GenerateSateliteOrbit(dice, flagVal))
	return nil
}

func (w *World) generateWorldType(dice *dice.Dicepool) error {
	satOrb := w.profile.Data(profile.KEY_SATELITE_ORBIT)
	hzVar := w.profile.Data(profile.KEY_HABITABLE_ZONE_VAR)
	if satOrb == nil {
		return fmt.Errorf("profile.KEY_SATELITE_ORBIT undefined")
	}
	if hzVar == nil {
		return fmt.Errorf("profile.KEY_HABITABLE_ZONE_VAR undefined")
	}
	w.profile.Inject(profile.KEY_WORLDTYPE, planets.GeneratePlanetType(dice, satOrb, hzVar))
	return nil
}

func (w *World) generateSizeLimit(dice *dice.Dicepool) error {
	if w.profile.Data(profile.KEY_SATELITE_ORBIT).Code() == "*" {
		w.profile.Inject(profile.KEY_LIMIT_size, ehex.New().Set(19))
		return nil
	}
	lim := dice.Sroll("2d6+6")
	w.profile.Inject(profile.KEY_LIMIT_size, ehex.New().Set(lim))
	return nil
}

func (w *World) generateSize(dice *dice.Dicepool) error {
	worldtype := w.profile.Data(profile.KEY_WORLDTYPE)
	if worldtype == nil {
		return fmt.Errorf("profile.KEY_WORLDTYPE undefined")
	}
	sizeLim := w.profile.Data(profile.KEY_LIMIT_size)
	if sizeLim == nil {
		return fmt.Errorf("profile.KEY_LIMIT_size undefined")
	}
	size := planets.GenerateSize(dice, worldtype)
	if size == nil {
		return errmaker.ErrorFrom(fmt.Errorf("can't generate size"), worldtype)
	}
	if size.Value() > sizeLim.Value() {
		size = sizeLim
	}
	w.profile.Inject(profile.KEY_SIZE, size)
	return nil
}

func (w *World) generateAtmo(dice *dice.Dicepool) error {
	worldtype := w.profile.Data(profile.KEY_WORLDTYPE)
	if worldtype == nil {
		return fmt.Errorf("profile.KEY_WORLDTYPE undefined")
	}
	size := w.profile.Data(profile.KEY_SIZE)
	if size == nil {
		return fmt.Errorf("profile.KEY_SIZE undefined")
	}
	atmo := planets.GenerateAtmo(dice, size, worldtype)
	w.profile.Inject(profile.KEY_ATMO, atmo)
	return nil
}

func (w *World) generateHydr(dice *dice.Dicepool) error {
	worldtype := w.profile.Data(profile.KEY_WORLDTYPE)
	if worldtype == nil {
		return fmt.Errorf("profile.KEY_WORLDTYPE undefined")
	}
	habzone := w.profile.Data(profile.KEY_HABITABLE_ZONE_VAR)
	if habzone == nil {
		return fmt.Errorf("profile.KEY_HABITABLE_ZONE_VAR undefined")
	}
	size := w.profile.Data(profile.KEY_SIZE)
	if size == nil {
		return fmt.Errorf("profile.KEY_SIZE undefined")
	}
	atmo := w.profile.Data(profile.KEY_ATMO)
	if atmo == nil {
		return fmt.Errorf("profile.KEY_ATMO undefined")
	}
	hydr := planets.GenerateHydr(dice, size, atmo, habzone, worldtype)
	w.profile.Inject(profile.KEY_HYDR, hydr)
	return nil
}

func (w *World) generateDominantLife(dice *dice.Dicepool) error {
	atmo := w.profile.Data(profile.KEY_ATMO)
	if atmo == nil {
		return fmt.Errorf("profile.KEY_ATMO undefined")
	}
	hydr := w.profile.Data(profile.KEY_HYDR)
	if hydr == nil {
		return fmt.Errorf("profile.KEY_HYDR undefined")
	}
	habzone := w.profile.Data(profile.KEY_HABITABLE_ZONE_VAR)
	if habzone == nil {
		return fmt.Errorf("profile.KEY_HABITABLE_ZONE_VAR undefined")
	}
	life := planets.GenerateDominantLife(dice, atmo, hydr, habzone, w.prim)
	w.profile.Inject(profile.KEY_LIFE_FACTOR, life)

	return nil
}

func (w *World) determineLifeCompatability(dice *dice.Dicepool) error {
	life := w.profile.Data(profile.KEY_LIFE_FACTOR)
	if life == nil {
		return fmt.Errorf("profile.KEY_LIFE_FACTOR undefined")
	}
	comp := 0
	switch life.Value() {
	default:
		w.nativeGenome = genetics.RollGenome(dice)
		comp = genetics.GenomeCompatability(w.nativeGenome, genetics.GeneTemplateHuman())
	case 0:
	}
	compat := ehex.New().Set(comp)
	w.profile.Inject(profile.KEY_LIFE_COMPATABILITY, compat)
	return nil
}

func (w *World) SetDefaultPopulationLimit() {
	defaultLim := ehex.New().Set(15)
	w.profile.Inject(profile.KEY_LIMIT_pops, defaultLim)
}

func (w *World) generatePops(dice *dice.Dicepool) error {
	if w.Flag[IsNotColonized] {
		w.profile.Inject(profile.KEY_POPS, 0)
		return nil
	}
	limit := w.profile.Data(profile.KEY_LIMIT_pops)
	if limit == nil {
		return fmt.Errorf("profile.KEY_LIMIT_pops undefined")
	}
	worldtype := w.profile.Data(profile.KEY_WORLDTYPE)
	if worldtype == nil {
		return fmt.Errorf("profile.KEY_WORLDTYPE undefined")
	}
	pops := planets.GeneratePops(dice, worldtype, limit)
	w.profile.Inject(profile.KEY_POPS, pops)
	return nil
}

func (w *World) generatePopDigit(dice *dice.Dicepool) error {
	pops := w.profile.Data(profile.KEY_POPS)
	if pops == nil {
		return fmt.Errorf("profile.KEY_POPS undefined")
	}
	if pops.Value() == 0 {
		w.profile.Inject(profile.KEY_POP_DIGIT, 0)
		return nil
	}
	w.profile.Inject(profile.KEY_POP_DIGIT, planets.PopulationDigit(dice))
	return nil
}

func (w *World) generateGovr(dice *dice.Dicepool) error {
	if w.Flag[IsNotColonized] {
		w.profile.Inject(profile.KEY_GOVR, 0)
		return nil
	}
	pops := w.profile.Data(profile.KEY_POPS)
	if pops == nil {
		return fmt.Errorf("profile.KEY_POPS undefined")
	}
	worldtype := w.profile.Data(profile.KEY_WORLDTYPE)
	if worldtype == nil {
		return fmt.Errorf("profile.KEY_WORLDTYPE undefined")
	}
	govr := planets.GenerateGoverment(dice, pops, worldtype)
	w.profile.Inject(profile.KEY_GOVR, govr)
	return nil
}

func (w *World) generateLaws(dice *dice.Dicepool) error {
	if w.Flag[IsNotColonized] {
		w.profile.Inject(profile.KEY_LAWS, 0)
		return nil
	}
	govr := w.profile.Data(profile.KEY_GOVR)
	if govr == nil {
		return fmt.Errorf("profile.KEY_GOVR undefined")
	}
	worldtype := w.profile.Data(profile.KEY_WORLDTYPE)
	if worldtype == nil {
		return fmt.Errorf("profile.KEY_WORLDTYPE undefined")
	}
	laws := planets.GenerateLaws(dice, govr, worldtype)
	w.profile.Inject(profile.KEY_LAWS, laws)
	return nil
}

func (w *World) setupMinimumTL() error {
	atmo := w.profile.Data(profile.KEY_ATMO)
	if atmo == nil {
		return fmt.Errorf("profile.KEY_ATMO undefined")
	}
	tlLimit := 0
	switch atmo.Value() {
	default:
		tlLimit = 8
	case 5, 6, 8:
		tlLimit = 24
	case 4, 7, 9:
		tlLimit = 3
	case 10:
		tlLimit = 8
	case 11:
		tlLimit = 9
	case 12:
		tlLimit = 10
	case 13, 14:
		tlLimit = 5
	case 15:
		tlLimit = 8
	}
	w.profile.Inject(profile.KEY_LIMIT_tl, ehex.New().Set(tlLimit))
	return nil
}

func (w *World) generatePort(dice *dice.Dicepool) error {
	worldtype := w.profile.Data(profile.KEY_WORLDTYPE)
	if worldtype == nil {
		return fmt.Errorf("profile.KEY_WORLDTYPE undefined")
	}
	pops := w.profile.Data(profile.KEY_POPS)
	if pops == nil {
		return fmt.Errorf("profile.KEY_POPS undefined")
	}
	port := planets.GeneratePort(dice, pops, worldtype)
	switch port.Value() {
	default:
		return fmt.Errorf("pseudoPortValue incorect")
	case 1:
		port = ehex.New().Set("F")
		if w.Flag[IsMainworld] {
			port = ehex.New().Set("A")
		}
	case 2:
		port = ehex.New().Set("F")
		if w.Flag[IsMainworld] {
			port = ehex.New().Set("B")
		}
	case 3:
		port = ehex.New().Set("F")
		if w.Flag[IsMainworld] {
			port = ehex.New().Set("C")
		}
	case 4:
		port = ehex.New().Set("G")
		if w.Flag[IsMainworld] {
			port = ehex.New().Set("D")
		}
	case 5:
		port = ehex.New().Set("H")
		if w.Flag[IsMainworld] {
			port = ehex.New().Set("E")
		}
	case 6:
		port = ehex.New().Set("Y")
		if w.Flag[IsMainworld] {
			port = ehex.New().Set("X")
		}
	}
	w.profile.Inject(profile.KEY_PORT, port)
	return nil
}

func (w *World) generateTL(dice *dice.Dicepool) error {
	worldtype := w.profile.Data(profile.KEY_WORLDTYPE)
	if worldtype == nil {
		return fmt.Errorf("profile.KEY_WORLDTYPE undefined")
	}
	port := w.profile.Data(profile.KEY_PORT)
	if port == nil {
		return fmt.Errorf("profile.KEY_PORT undefined")
	}
	size := w.profile.Data(profile.KEY_SIZE)
	if size == nil {
		return fmt.Errorf("profile.KEY_SIZE undefined")
	}
	atmo := w.profile.Data(profile.KEY_ATMO)
	if atmo == nil {
		return fmt.Errorf("profile.KEY_ATMO undefined")
	}
	hydr := w.profile.Data(profile.KEY_HYDR)
	if hydr == nil {
		return fmt.Errorf("profile.KEY_HYDR undefined")
	}
	pops := w.profile.Data(profile.KEY_POPS)
	if pops == nil {
		return fmt.Errorf("profile.KEY_POPS undefined")
	}
	govr := w.profile.Data(profile.KEY_GOVR)
	if govr == nil {
		return fmt.Errorf("profile.KEY_GOVR undefined")
	}

	tl := planets.GenerateTL(dice, port, size, atmo, hydr, pops, govr)
	w.profile.Inject(profile.KEY_TL, tl)
	return nil
}

func (w *World) UWP() string {
	return profile.UWP(w.profile)
}

func (w *World) generateBases(dice *dice.Dicepool) error {
	port := w.profile.Data(profile.KEY_PORT)
	if port == nil {
		return fmt.Errorf("profile.KEY_PORT undefined")
	}

	bases := planets.GenerateBases(dice, port)
	w.profile.Inject(profile.KEY_BASES, bases)
	return nil
}

func (w *World) defineMW() error {
	w.profile.Inject(profile.KEY_MAINWORLD, "*")
	if w.Flag[IsMainworld] && w.Flag[IsNotMainworld] {
		return errmaker.ErrorFrom(fmt.Errorf("flags '%v' and '%v' active at same time", IsMainworld, IsNotMainworld))
	}
	if w.Flag[IsMainworld] {
		w.profile.Inject(profile.KEY_MAINWORLD, "1")
	}
	if w.Flag[IsNotMainworld] {
		w.profile.Inject(profile.KEY_MAINWORLD, "0")
	}
	return nil
}

func (w *World) Data(key string) string {
	data := w.profile.Data(key)
	if data != nil {
		return data.Code()
	}
	if val, ok := w.Flag[key]; ok {
		if val {
			fmt.Println("Has Flag", val, key)
			return key
		}
	}
	return ""
}
