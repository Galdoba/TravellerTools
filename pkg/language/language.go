package language

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	InitialConsonant = iota
	Vowel
	FinalConsonant
)

type Sound struct {
	sType          int
	sound          string
	freq           int
	pronuansiation string
}

type Language struct {
	Name                       string
	ConsonantsInitial          []Sound
	Vowels                     []Sound
	ConsonantsFinal            []Sound
	BasicGenerationTable       map[string]string
	AlternativeGenerationTable map[string]string
	InitialConsonantTable      map[string]string
	VowelTable                 map[string]string
	FinalConsonantTable        map[string]string
}

func New(name string) (*Language, error) {
	l := Language{}
	l.Name = name
	l.ConsonantsInitial, l.Vowels, l.ConsonantsFinal,
		l.BasicGenerationTable, l.AlternativeGenerationTable,
		l.InitialConsonantTable, l.VowelTable, l.FinalConsonantTable = callTables(name)
	return &l, nil
}

type word struct {
	struc []string
	spell []string
	text  string
}

func NewWord(dice *dice.Dicepool, lang *Language, lenght int) string {
	if lenght < 1 {
		lenght = dice.Sroll("1d6")
	}
	word := word{}
	for i := 0; i < lenght; i++ {
		switch {
		case i == 0:
			word.struc = append(word.struc, lang.BasicGenerationTable[fmt.Sprintf("%v%v", dice.Sroll("1d6"), dice.Sroll("1d6"))])
		default:
			switch word.struc[i-1] {
			case "VC", "CVC":
				word.struc = append(word.struc, lang.AlternativeGenerationTable[fmt.Sprintf("%v%v", dice.Sroll("1d6"), dice.Sroll("1d6"))])
			case "V", "CV":
				word.struc = append(word.struc, lang.BasicGenerationTable[fmt.Sprintf("%v%v", dice.Sroll("1d6"), dice.Sroll("1d6"))])
			}
		}
	}
	for _, syl := range word.struc {
		sp := strings.Split(syl, "")
		for j, s := range sp {
			code := fmt.Sprintf("%v%v%v", dice.Sroll("1d6"), dice.Sroll("1d6"), dice.Sroll("1d6"))
			switch {
			case j == 0 && s == "C":
				word.spell = append(word.spell, lang.InitialConsonantTable[code])
			case s == "V":
				word.spell = append(word.spell, lang.VowelTable[code])
			case j == 1 && s == "C":
				word.spell = append(word.spell, lang.FinalConsonantTable[code])
			case j == 2 && s == "C":
				word.spell = append(word.spell, lang.FinalConsonantTable[code])
			}
		}
	}

	word.text = strings.Join(word.spell, "")
	word.text = strings.ToLower(word.text)
	return word.text
}
