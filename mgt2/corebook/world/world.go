package world

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/language"
	"github.com/Galdoba/TravellerTools/pkg/mgt2trade/traffic/tradecodes"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

const (
	KEY_SEED      = "Seed"
	KEY_NAME      = "Name"
	DEFAULT_VALUE = iota
	TEMP_FROZEN
	TEMP_COLD
	TEMP_TEMPERATE
	TEMP_HOT
	TEMP_BOILING
)

type constructor struct {
	dice    *dice.Dicepool
	options map[string]string
}

type instruction struct {
	key   string
	value string
}

func Instruction(k, v string) instruction {
	return instruction{k, v}
}

func NewConstructor(options ...instruction) *constructor {
	c := constructor{}
	c.options = make(map[string]string)
	for _, inst := range options {
		c.addInstruction(inst)
	}

	seed := c.options[KEY_SEED]
	c.dice = dice.New().SetSeed(seed)
	return &c
}

func (c *constructor) addInstruction(i instruction) {
	c.options[i.key] = i.value
}

func (c *constructor) callInstruction(key string) string {
	return c.options[key]
}

type world struct {
	name        string
	location    string      //позже интерфейс "Координаты"
	statistics  []ehex.Ehex //позже интерфейс "Универсальный Планетарный Профайл"
	uwp         uwp.UWP
	bases       string
	tradeCodes  []string
	travelCode  string
	gasGigants  string
	temperature int
	nlife       bool
	pbg         string
}

func (c *constructor) Create() (*world, error) {
	w := &world{}
	w.name = c.callInstruction(KEY_NAME)
	if w.name == "" {
		lang, _ := language.New("VILANI")
		word := language.NewWord(c.dice, lang, 0)
		w.name = word
	}
	for _, err := range []error{
		w.rollSize(c),
		w.rollAtmo(c),
		w.rollTemp(c),
		w.rollHydr(c),
		w.rollPops(c),
		w.rollGovr(c),
		w.rollLaws(c),
		w.rollStprt(c),
		w.rollTL(c),
		w.enviromentalLimits(c),
		w.getTradeCodes(c),
		w.getTravelCode(c),
		w.setPBG(c),
	} {
		if err != nil {
			fmt.Printf("failed\n   %v", err.Error())
			return w, err
		}
	}
	return w, nil
}

func (w *world) rollSize(c *constructor) error {
	if len(w.statistics) > 0 {
		return fmt.Errorf("Size already rolled")
	}
	s := c.dice.Sroll("2d6-2")
	for {
		if s < 10 { //если размер 0-9 оставляем как есть
			break
		}
		if c.dice.Sroll("1d6") < 4 { //иначе увеличиваем размер на 1 с 50% вероятностью
			break
		}
		s++
	}
	sz := ehex.New().Set(s)
	w.statistics = append(w.statistics, sz)
	return nil
}

func (w *world) rollAtmo(c *constructor) error {
	if len(w.statistics) > 1 {
		return fmt.Errorf("Atmo already rolled")
	}
	a := c.dice.Sroll("2d6-7") + w.statistics[0].Value()
	switch w.statistics[0].Value() {
	case 0, 1:
		a = c.dice.Sroll("2d6-12") + w.statistics[0].Value()
	case 3, 4:
		switch a {
		case 4:
			a = 2
		case 5:
			a = 3
		case 6:
			a = 5
		case 7:
			a = 4
		case 8:
			a = 6
		}
	}
	if a < 0 {
		a = 0
	}
	if a > 15 {
		a = 15
	}
	w.statistics = append(w.statistics, ehex.New().Set(a))
	return nil
}

func (w *world) rollTemp(c *constructor) error {
	dm := 0
	extreme := false
	if len(w.statistics) < 2 {
		return fmt.Errorf("Can not roll Temperature as there no Atmo Stat")
	}
	switch w.statistics[1].Code() {
	case "0", "1":
		extreme = true
	case "2", "3":
		dm = -2
	case "4", "5", "E":
		dm = -1
	case "6", "7":
		dm = 0
	case "8", "9":
		dm = 1
	case "A", "D", "F":
		dm = 2
	case "B", "C":
		dm = 6
	}
	r := c.dice.Sroll("2d6") + dm
	switch r {
	case 3, 4:
		w.temperature = TEMP_COLD
		if extreme {
			w.temperature = TEMP_FROZEN
		}
	case 5, 6, 7, 8, 9:
		w.temperature = TEMP_TEMPERATE
		if extreme {
			switch c.dice.Sroll("1d2") {
			case 1:
				w.temperature = TEMP_FROZEN
			case 2:
				w.temperature = TEMP_BOILING
			}
		}
	case 10, 11:
		w.temperature = TEMP_HOT
		if extreme {
			w.temperature = TEMP_BOILING
		}

	}
	if r <= 2 {
		w.temperature = TEMP_FROZEN
	}
	if r >= 12 {
		w.temperature = TEMP_BOILING
	}
	if w.temperature == 0 {
		return fmt.Errorf("temperature was not asigned")
	}
	return nil
}

func (w *world) rollHydr(c *constructor) error {
	dm := 0
	if len(w.statistics) < 2 {
		return fmt.Errorf("Can not roll Hydro as there no Atmo Stat")
	}
	switch w.statistics[1].Value() {
	case 0, 1, 10, 11, 12, 13, 14, 15:
		dm = -4
	}
	switch w.statistics[1].Value() {
	case 0, 1, 10, 11, 12, 14:
		dm = -4
	case 13, 15:
		dm = -4
		if w.temperature == TEMP_HOT {
			dm += -2
		}
		if w.temperature == TEMP_COLD {
			dm += -6
		}
	}
	h := c.dice.Sroll("2d6-7") + dm
	switch w.statistics[0].Value() {
	case 0, 1:
		h = 0
	}
	if h < 0 {
		h = 0
	}
	if h > 10 {
		h = 10
	}
	w.statistics = append(w.statistics, ehex.New().Set(h))
	return nil
}

func (w *world) rollPops(c *constructor) error {
	if len(w.statistics) > 3 {
		return fmt.Errorf("Pops already rolled")
	}
	p := c.dice.Sroll("2d6-2")
	for {
		if p < 10 { //если размер 0-9 оставляем как есть
			break
		}
		if c.dice.Sroll("1d6") < 4 { //иначе увеличиваем размер на 1 с 50% вероятностью
			break
		}
		p++
	}
	pop := ehex.New().Set(p)
	w.statistics = append(w.statistics, pop)
	return nil
}

func (w *world) rollGovr(c *constructor) error {
	if len(w.statistics) < 4 {
		return fmt.Errorf("No data on Pops")
	}
	g := c.dice.Sroll("2d6-7") + w.statistics[3].Value()
	if g < 0 {
		g = 0
	}
	if g > 15 {
		g = 15
	}
	w.statistics = append(w.statistics, ehex.New().Set(g))
	return nil
}

func (w *world) rollLaws(c *constructor) error {
	if len(w.statistics) < 5 {
		return fmt.Errorf("No data on Govr")
	}
	l := c.dice.Sroll("2d6-7") + w.statistics[4].Value()
	if l < 0 {
		l = 0
	}
	w.statistics = append(w.statistics, ehex.New().Set(l))
	return nil
}

func (w *world) rollStprt(c *constructor) error {
	if len(w.statistics) < 6 {
		return fmt.Errorf("No data on Laws")
	}
	dm := 0
	switch w.statistics[3].Value() {
	case 0, 1, 2:
		dm = -2
	case 3, 4:
		dm = -1
	case 8, 9:
		dm = 1
	}
	if w.statistics[3].Value() > 9 {
		dm = 1 + ((w.statistics[3].Value() - 10) / 3)
	}
	code := "X"
	r := c.dice.Sroll("2d6") + dm
	switch r {
	case 3, 4:
		code = "E"
	case 5, 6:
		code = "D"
	case 7, 8:
		code = "C"
	case 9, 10:
		code = "B"
	}
	if r >= 11 {
		code = "A"
	}
	w.statistics = append(w.statistics, ehex.New().Set(code))
	return nil
}

func (w *world) rollTL(c *constructor) error {
	if len(w.statistics) < 7 {
		return fmt.Errorf("No data on Starport")
	}
	tl := c.dice.Sroll("1d6")
	switch w.statistics[0].Value() {
	case 0, 1:
		tl = tl + 2
	case 2, 3, 4:
		tl++
	}
	switch w.statistics[1].Value() {
	case 0, 1, 2, 3, 10, 11, 12, 13, 14, 15:
		tl++
	}
	switch w.statistics[2].Value() {
	case 0, 9:
		tl++
	case 10:
		tl = tl + 2
	}
	switch w.statistics[3].Value() {
	case 1, 2, 3, 4, 5, 8:
		tl++
	case 9:
		tl = tl + 2
	case 10:
		tl = tl + 4
	case 11:
		tl = tl + 3
	case 12:
		tl = tl + 2
	case 13:
		tl = tl + 1
	}
	switch w.statistics[4].Value() {
	case 0, 5:
		tl++
	case 7:
		tl = tl + 2
	case 13, 14:
		tl = tl - 2
	}
	switch w.statistics[6].Code() {
	case "X":
		tl = tl - 4
	case "C":
		tl = tl + 2
	case "B":
		tl = tl + 4
	case "A":
		tl = tl + 6
	}
	if tl < 0 {
		tl = 0
	}
	w.statistics = append(w.statistics, ehex.New().Set(tl))
	return nil
}

func (w *world) enviromentalLimits(c *constructor) error {
	if len(w.statistics) < 8 {
		return fmt.Errorf("No data on TL")
	}
	if c.dice.Sroll("2d6") == 12 {
		switch w.statistics[1].Value() {
		case 4, 5, 6, 7, 8, 9, 13, 14, 10, 11, 12, 15:
			w.nlife = true
		}
	}
	tlLimit := 0
	switch w.statistics[1].Value() {
	case 5, 6, 8:
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
	if w.nlife {
		tlLimit = 0
	}
	if w.statistics[7].Value() < tlLimit {
		w.statistics[3] = ehex.New().Set(0)
		w.statistics[4] = ehex.New().Set(0)
		w.statistics[5] = ehex.New().Set(0)
		w.statistics[6] = ehex.New().Set("X")
		w.statistics[7] = ehex.New().Set(0)
	}
	switch {
	case w.statistics[7].Value() < 5 || w.statistics[3].Value() < 3:
		switch w.statistics[6].Code() {
		case "A", "B", "C", "D":
			w.statistics[6] = ehex.New().Set("E")
		}
	case w.statistics[7].Value() < 7 || w.statistics[3].Value() < 5:
		switch w.statistics[6].Code() {
		case "A", "B", "C":
			w.statistics[6] = ehex.New().Set("D")
		}
	case w.statistics[7].Value() < 8 || w.statistics[3].Value() < 6:
		switch w.statistics[6].Code() {
		case "A", "B":
			w.statistics[6] = ehex.New().Set("C")
		}
	case w.statistics[7].Value() < 9 || w.statistics[3].Value() < 7:
		switch w.statistics[6].Code() {
		case "A":
			w.statistics[6] = ehex.New().Set("B")
		}
	}
	uwpSTR := w.statistics[6].Code() + w.statistics[0].Code() + w.statistics[1].Code() + w.statistics[2].Code() +
		w.statistics[3].Code() + w.statistics[4].Code() + w.statistics[5].Code() + "-" + w.statistics[7].Code()
	w.uwp = uwp.Inject(uwpSTR)
	return nil
}

func (w *world) getTradeCodes(c *constructor) error {
	tc, err := tradecodes.FromUWP(w.uwp)
	if err != nil {
		return err
	}
	switch w.temperature {
	case TEMP_FROZEN:
		tc = append(tc, "Fr")
	case TEMP_COLD:
		tc = append(tc, "Co")
	case TEMP_HOT:
		tc = append(tc, "Ho")
	case TEMP_BOILING:
		tc = append(tc, "Bo")
	}
	w.tradeCodes = tc
	return nil
}

func (w *world) getTravelCode(c *constructor) error {
	warn := 0
	if w.uwp.Atmo() >= 10 {
		warn++
	}
	switch w.uwp.Govr() {
	case 0, 7, 10:
		if w.uwp.Laws() == 0 || w.uwp.Laws() >= 9 {
			warn++
		}
	}
	if w.uwp.Govr()+w.uwp.Laws() >= 20 {
		warn++
	}
	if w.uwp.Govr()+w.uwp.Laws() >= 22 {
		warn++
	}
	if w.uwp.Starport() == "X" {
		warn++
	}
	switch warn {
	default:
	case 1, 2:
		w.travelCode = "A"
	case 3, 4, 5:
		w.travelCode = "R"
	}
	return nil
}

func (w *world) setPBG(c *constructor) error {
	w.pbg = "G"
	if c.dice.Sroll("2d6") > 9 {
		w.pbg = ""
	}
	return nil
}

func listTC(sl []string) string {
	s := ""
	for _, tc := range sl {
		s += tc + " "
	}
	s = strings.TrimSuffix(s, " ")
	return s
}

func (w *world) ShortData() []string {
	fields := []string{}
	fields = append(fields, w.name)
	fields = append(fields, w.location)
	fields = append(fields, w.bases)
	fields = append(fields, fmt.Sprintf("%v", w.uwp))
	fields = append(fields, listTC(w.tradeCodes))
	fields = append(fields, w.travelCode)
	fields = append(fields, w.pbg)
	return fields
}
