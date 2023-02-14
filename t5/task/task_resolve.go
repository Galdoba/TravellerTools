package task

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/dice/probability"
	"github.com/Galdoba/TravellerTools/t5/pawn/characteristic"
)

type resolution struct {
	dice     *dice.Dicepool
	resolved bool
	rolled   []int
	tn       int
	prob     float64
	text     string
}

func CheckCharacteristicAverage(chr *characteristic.Frame) *resolution {
	res := resolution{}
	diceNum := chr.Genes()
	res.tn = chr.Actual()
	tsk, err := New(
		Difficulty(diceNum),
		Char(chr.Name()),
		Phrase("check "+chr.Name()),
	)
	if err != nil {
		panic(err.Error())
	}
	res.text = tsk.toString()
	code := fmt.Sprintf("%vd6", diceNum)
	res.prob, err = probability.Calculate(probability.RESULT_IS_SAME_OR_LESS, code, res.tn)
	return &res
}
