package decidion

import (
	"fmt"
	"testing"
)

func TestDesidio(t *testing.T) {
	options := []string{"one", "two", "three", "four", "five"}
	fmt.Println("options:", options)
	// dice := dice.New()
	// answers:=[]string{}
	// answer := Random_One(dice, options...)
	// fmt.Println("Random_One:", answer)
	// answers := Random_Few(9, dice, options...)
	// fmt.Println("Random_Few:", answers)
	// answer, options = Random_One_Exclude(dice, options...)
	// fmt.Println("options:", options)

	// fmt.Println("Random_One_exclude:", answer)
	// answers, options = Random_Few_Exclude(7, dice, options...)

	// fmt.Println("Random_Few_Exclude:", answers)
	// fmt.Println("options:", options)
	answer := Manual_One("label", options...)
	fmt.Println("-----")
	fmt.Println(answer)
}
