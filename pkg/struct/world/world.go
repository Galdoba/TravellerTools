package world

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/generation/star"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/devtools/errmaker"
)

/*
World. A planet or satellite.
-----------------------------
Profile Keys:
Starport [ABCDEFGHXY]
Size [0-K]
Atmo [0-F]
Hydr [0-A]
HZvar [0-F] h
SateliteCode [0-2] h
SateliteOrb [0-F] h
Climate [0-6] h
Pops [0-F]
Govr [0-F]
Laws [0-J]
Tech [0-N] s
PopDigit [0-9] h
OtherBelts [0-3] h
GasG [0-4] h
LifeFactor [0-A] h
LifeCompatability [0-A] h (количество совпадений в геноме рандомного софонта с человеком минус 2)

*/

type World struct {
	Alias    string //Самоназвание
	Catalog  string //номенклатурное название (121-311 Stargos A III 2d)
	profile  profile.Profile
	HomeStar star.StarBody
	Flag     map[string]bool
}

type knownData struct {
	key  string
	val  string
	used bool
}

const (
	FLAG_TRUE  = "True"
	FLAG_FALSE = "False"
	Alias      = "Alias"
	Catalog    = "Catalog Name"
	Primary    = "Primary Star"
	Companion  = "Companion Star"
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
	switch kd.key {
	default:
		w.profile.Inject(kd.key, kd.val)
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
		if w.profile.Data(key) != nil {
			injected++
			continue
		}
		switch key {
		case profile.KEY_SIZE:

		}
	}
	if injected == checkList {
		return ErrFullyGenerated
	}
	return nil
}
