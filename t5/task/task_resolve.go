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
	rr       rollResult
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

type rollResult struct {
	rollState    []int
	success      bool
	sSuccess     bool
	sFailure     bool
	sInteresting bool
	sStupid      bool
}

func newResult(rollState []int, tn int) rollResult {
	rr := rollResult{}
	rr.rollState = rollState
	sum := 0
	ones := 0
	sixes := 0
	for i, d := range rollState {
		if i+1 > tn {
			rr.sStupid = true
		}
		if d == 1 {
			ones++
		}
		if d == 6 {
			sixes++
		}
		sum += d
	}
	if ones >= 3 {
		rr.sSuccess = true
	}
	if sixes >= 3 {
		rr.sFailure = true
	}
	rr.sInteresting = rr.sSuccess && rr.sFailure
	if rr.sInteresting {
		rr.sSuccess = false
		rr.sFailure = false
	}
	if sum <= tn {
		rr.success = true
	}
	if rr.sSuccess {
		rr.success = true
	}
	if rr.sFailure {
		rr.success = false
	}
	return rr
}

func (rr *rollResult) String() string {
	s := ""
	switch {
	case rr.sInteresting == true:
		s = "Spectaculary Interesting"
	case rr.sSuccess == true:
		s = "Spectacular Success"
	case rr.sFailure == true:
		s = "Spectacular Failure"
	case rr.success == false:
		s = "Failure"
	case rr.success == true:
		s = "Success"
	}
	if rr.sStupid {
		s += " and it was Spectaculary Stupid"
	}
	return s
}
