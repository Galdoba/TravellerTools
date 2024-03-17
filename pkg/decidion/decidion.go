package decidion

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/manifoldco/promptui"
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

func Manual_One(label string, silent bool, options ...string) string {
	prompt := promptui.Select{
		Label:        label,
		Items:        append([]string{}, options...),
		Size:         20,
		HideSelected: true,
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("prompt failed: %v\n", err)
		return ""
	}
	if !silent {
		fmt.Printf("%v: %v\n", label, result)
	}
	return result
}

func Manual_One_Exclude(label string, silent bool, options ...string) (string, []string) {
	prompt := promptui.Select{
		Label:        label,
		Items:        append([]string{}, options...),
		Size:         20,
		HideSelected: true,
	}
	i, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("prompt failed: %v\n", err)
		return "", options
	}
	options = exclude(options, i)
	if !silent {
		fmt.Printf("%v: %v\n", label, result)
	}
	return result, options
}
func Manual_Few(n int, label string, options ...string) []string {
	answers := []string{}
	fmt.Printf("%v\r", label)
	for i := 0; i < n; i++ {
		answer := Manual_One(label, true, options...)
		answers = append(answers, answer)
		fmt.Printf("%v %v of %v: %v\n", label, i+1, n, answer)
	}

	return answers
}

func Manual_Few_Exclude(n int, label string, options ...string) ([]string, []string) {
	answers := []string{}
	answer := ""
	fmt.Printf("%v\r", label)
	for i := 0; i < n; i++ {
		answer, options = Manual_One_Exclude(label, true, options...)
		answers = append(answers, answer)
		fmt.Printf("%v %v of %v: %v\n", label, i+1, n, answer)
	}

	return answers, options
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
