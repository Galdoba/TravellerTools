package world

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
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
	bases       string
	tradeCodes  []string
	travelCode  string
	gasGigants  string
	temperature int
}

func (c *constructor) Create() (*world, error) {
	w := &world{}
	w.name = c.callInstruction(KEY_NAME)
	if w.name == "" {
		w.name = "[No Name]"
	}
	for i, err := range []error{
		w.rollSize(c),
		w.rollAtmo(c),
		w.rollTemp(c),
		w.rollHydr(c),
		w.rollPops(c),
	} {
		fmt.Printf("DEBUG: Step %v: ", i+1)
		if err != nil {
			fmt.Printf("failed\n   %v", err.Error())
			return w, err
		}
		fmt.Printf("ok\n")
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
