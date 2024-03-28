package decidion

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/charmbracelet/huh"
)

/*
	decidion.Random(dice, options...) string
	decidion.Random_Exclude(dice, options) (string, []string)
	decidion.Manual(options...) string
	decidion.Manual_Exclude(options) (string, []string)
*/

func Random_One(dice *dice.Dicepool, options ...string) string {
	picked := autoPick(options, dice)
	if picked == -1 {
		return ""
	}
	return options[picked]
}

func Random_Few(n int, dice *dice.Dicepool, options ...string) []string {
	answers := []string{}
	if n < 0 {
		return answers
	}
	for i := 0; i < n; i++ {
		answers = append(answers, Random_One(dice, options...))
	}
	return answers
}

func Random_One_Exclude(dice *dice.Dicepool, options ...string) (string, []string) {
	picked := autoPick(options, dice)
	if picked == -1 {
		return "", exclude(options, picked)
	}
	return options[picked], exclude(options, picked)
}

func Random_Few_Exclude(n int, dice *dice.Dicepool, options ...string) ([]string, []string) {
	pick := ""
	answers := []string{}
	if n > len(options) {
		return options, answers
	}
	for i := 0; i < n; i++ {
		pick, options = Random_One_Exclude(dice, options...)
		answers = append(answers, pick)
		if len(options) == 0 {
			break
		}
	}
	return answers, options
}

func Manual_One(label string, options ...string) string {
	if len(options) == 1 {
		return options[0]
	}
	answer := -1
	selectComponent := huh.NewSelect[int]()
	selectComponent = selectComponent.Title(label)
	opts := []huh.Option[int]{}
	for i, opt := range options {
		opts = append(opts, huh.NewOption(opt, i))
	}
	selectComponent = selectComponent.Options(opts...)
	selectComponent = selectComponent.Value(&answer)
	form := huh.NewForm(huh.NewGroup(selectComponent))
	err := form.Run()
	if err != nil {
		panic(err.Error())
	}
	return options[answer]
}

func Manual_Few(n int, label string, options ...string) []string {
	answers := []string{}
	currentLabel := label
	for i := 0; i < n; i++ {
		currentAnswer := Manual_One(currentLabel, options...)
		answers = append(answers, currentAnswer)
		currentLabel = ""
		if len(answers) > 0 {
			for _, a := range answers {
				currentLabel += fmt.Sprintf("%v\n", a)
			}
		}
		currentLabel = fmt.Sprintf("%v%v", currentLabel, label)
	}
	return answers
}

func Manual_Few_Exclude(n int, label string, options ...string) ([]string, []string) {
	answers := []string{}
	validOptions := options
	currentLabel := label
	currentAnswer := ""
	for i := 0; i < n; i++ {
		currentAnswer, validOptions = Manual_One_Exclude(currentLabel, validOptions...)
		answers = append(answers, currentAnswer)
		currentLabel = ""
		if len(answers) > 0 {
			for _, a := range answers {
				currentLabel += fmt.Sprintf("%v\n", a)
			}
		}
		currentLabel = fmt.Sprintf("%v%v", currentLabel, label)
	}
	return answers, validOptions
}

func Manual_One_Exclude(label string, options ...string) (string, []string) {
	answer := -1
	if len(options) == 1 {
		return options[0], nil
	}
	notSelected := []string{}
	selectComponent := huh.NewSelect[int]()
	selectComponent = selectComponent.Title(label)
	opts := []huh.Option[int]{}
	for i, opt := range options {
		opts = append(opts, huh.NewOption(opt, i))
	}
	selectComponent = selectComponent.Options(opts...)
	selectComponent = selectComponent.Value(&answer)
	form := huh.NewForm(huh.NewGroup(selectComponent))
	err := form.Run()
	if err != nil {
		panic(err.Error())
	}
	for i := range options {
		if i == answer {
			continue
		}
		notSelected = append(notSelected, options[i])
	}
	return options[answer], notSelected
}

////////////////////////////////////

func autoPick(sl []string, dice *dice.Dicepool) int {
	l := len(sl)
	if l == 0 {
		return -1
	}
	return dice.Sroll(fmt.Sprintf("1d%v", l)) - 1
}

func exclude(sl []string, n int) []string {
	if n < 0 {
		return sl
	}
	if n >= len(sl) {
		return sl
	}
	leftover := []string{}
	for i := range sl {
		if i == n {
			continue
		}
		leftover = append(leftover, sl[i])
	}
	return leftover
}
