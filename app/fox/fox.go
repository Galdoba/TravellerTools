package main

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	Часы     = "Часы"
	Шарф     = "Шарф"
	Трость   = "Трость"
	Портфель = "Портфель"
	Монокль  = "Монокль"
	Пенсне   = "Пенсне"
	Цилиндр  = "Цилиндр"
	Бусы     = "Бусы"
	Плащ     = "Плащ"
	Цветок   = "Цветок"
	Зонт     = "Зонт"
	Перчатки = "Перчатки"
)

type Gear struct {
	Have   []bool
	posses map[string]bool
}

type Коварный_Лис struct {
	Имя  string
	Gear Gear
}

func (кл *Коварный_Лис) String() string {
	return fmt.Sprintf("%v:	%v	%v	%v	%v	%v	%v	%v	%v	%v	%v	%v	%v", longname(кл.Имя),
		кл.Gear.posses[Часы],
		кл.Gear.posses[Шарф],
		кл.Gear.posses[Трость],
		кл.Gear.posses[Портфель],
		кл.Gear.posses[Монокль],
		кл.Gear.posses[Пенсне],
		кл.Gear.posses[Цилиндр],
		кл.Gear.posses[Бусы],
		кл.Gear.posses[Плащ],
		кл.Gear.posses[Цветок],
		кл.Gear.posses[Зонт],
		кл.Gear.posses[Перчатки],
	)
}

func longname(s string) string {
	for len(s) < 7 {
		s = s + " "
	}
	return s
}

func Все_Имена() []string {
	return []string{
		"Тэд",
		"Оливер",
		"Анна",
		"Ева",
		"Мэри",
		"Дэйзи",
		"Клер",
		"Патрик",
		"Элис",
		"Вера",
		"Ральф",
		"Нил",
		"Люси",
		"Лили",
		"Кевин",
	}
}

func НовыйЛис(Имя string) *Коварный_Лис {
	fox := Коварный_Лис{}
	fox.Имя = Имя
	fox.Gear.posses = make(map[string]bool)
	switch Имя {
	case "Джулия":
		fox.Gear.posses["Трость"] = true
		fox.Gear.posses["Монокль"] = true
		fox.Gear.posses["Цветок"] = true
	case "Тэд":
		fox.Gear.posses["Цилиндр"] = true
		fox.Gear.posses["Шарф"] = true
		fox.Gear.posses["Часы"] = true
	case "Оливер":
		fox.Gear.posses["Цилиндр"] = true
		fox.Gear.posses["Часы"] = true
		fox.Gear.posses["Монокль"] = true
	case "Анна":
		fox.Gear.posses["Зонт"] = true
		fox.Gear.posses["Перчатки"] = true
		fox.Gear.posses["Шарф"] = true
	case "Ева":
		fox.Gear.posses["Бусы"] = true
		fox.Gear.posses["Часы"] = true
		fox.Gear.posses["Плащ"] = true
	case "Мэри":
		fox.Gear.posses["Пенсне"] = true
		fox.Gear.posses["Перчатки"] = true
		fox.Gear.posses["Зонт"] = true
	case "Дэйзи":
		fox.Gear.posses["Бусы"] = true
		fox.Gear.posses["Плащ"] = true
		fox.Gear.posses["Трость"] = true
	case "Клер":
		fox.Gear.posses["Зонт"] = true
		fox.Gear.posses["Перчатки"] = true
		fox.Gear.posses["Бусы"] = true
	case "Патрик":
		fox.Gear.posses["Цилиндр"] = true
		fox.Gear.posses["Трость"] = true
		fox.Gear.posses["Пенсне"] = true
	case "Элис":
		fox.Gear.posses["Монокль"] = true
		fox.Gear.posses["Плащ"] = true
		fox.Gear.posses["Портфель"] = true
	case "Вера":
		fox.Gear.posses["Цветок"] = true
		fox.Gear.posses["Шарф"] = true
		fox.Gear.posses["Портфель"] = true
	case "Ральф":
		fox.Gear.posses["Цилиндр"] = true
		fox.Gear.posses["Шарф"] = true
		fox.Gear.posses["Портфель"] = true
	case "Нил":
		fox.Gear.posses["Пенсне"] = true
		fox.Gear.posses["Плащ"] = true
		fox.Gear.posses["Портфель"] = true
	case "Люси":
		fox.Gear.posses["Монокль"] = true
		fox.Gear.posses["Перчатки"] = true
		fox.Gear.posses["Зонт"] = true
	case "Лили":
		fox.Gear.posses["Бусы"] = true
		fox.Gear.posses["Трость"] = true
		fox.Gear.posses["Цветок"] = true
	case "Кевин":
		fox.Gear.posses["Пенсне"] = true
		fox.Gear.posses["Часы"] = true
		fox.Gear.posses["Цветок"] = true
	default:
		fox.Имя = "Некто"
	}
	return &fox
}

func Взять_Предмет() int {
	return dice.New().Roll("1d12").DM(-1).Sum()
}

func main() {
	for _, name := range Все_Имена() {
		fox := НовыйЛис(name)
		fmt.Println(fox)
	}

	// Suspect := &Коварный_Лис{}
	// for i := 0; i < 12; i++ {
	// 	Suspect.Взять_Предмет()
	// 	for _, name := range Все_Имена() {
	// 		Busted := НовыйЛис(name)
	// 		if Check(Busted, Suspect) {
	// 			fmt.Println(Busted)
	// 		}
	// 	}
	// }
}

func (кл *Коварный_Лис) Взять_Предмет() {
	dice := dice.New()
	Haveет := dice.Roll("1d12").DM(-1).Sum()
	for кл.Gear.Have[Haveет] == true {
		Haveет = dice.Roll("1d12").DM(-1).Sum()
	}
	кл.Gear.Have[Haveет] = true
}

func Check(fox1, fox2 *Коварный_Лис) bool {
	match := 0
	for Have := range fox1.Gear.Have {
		if fox1.Gear.Have[Have] == true && fox2.Gear.Have[Have] == true {
			match++
		}
	}
	if match >= 3 {
		return true
	}
	return false
}
