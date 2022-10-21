package task

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	DefaultValue = iota
	DifficultyEasy
	DifficultyAverage
	DifficultyDifficult
	DifficultyFormidable
	DifficultyStaggering
	DifficultyHopeless
	DifficultyImposibble
	DifficultyBeyondImpossible
)

type task struct {
	descr       string
	diff        int
	timeframe   string
	assets      []taskAsset
	hasty       bool
	extraHasty  bool
	cautious    bool
	dangerous   bool
	destructive bool
	roll        []int
	resolution  string
}

type taskAsset struct {
	descr string
	value int
}

func Asset(descr string, val int) taskAsset {
	return taskAsset{descr: descr, value: val}
}

func New(descr string, diff int, assets ...taskAsset) (*task, error) {
	ts := task{}
	ts.descr = descr
	switch Difficulty(diff) {
	case "[INVALID]":
		return &ts, fmt.Errorf("difficulty %v is invalid", diff)
	}
	ts.diff = diff
	assetMap := make(map[string]int)
	for _, asst := range assets {
		assetMap[asst.descr]++
		if assetMap[asst.descr] != 1 {
			return &ts, fmt.Errorf("asset `%v` is duplicated", asst.descr)
		}
		ts.assets = append(ts.assets, asst)
	}
	ts.resolution = "Unresolved"
	return &ts, nil
}

func (ts *task) Resolve(seed ...string) {
	dp := dice.New()
	if len(seed) > 0 {
		sd := ""
		for _, s := range seed {
			sd += s
		}
		dp.SetSeed(sd)
	}
	result := 0
	for d := 0; d < ts.diff; d++ {
		ts.roll = append(ts.roll, dp.Roll("1d6").Sum())
	}
	ones := 0
	sixes := 0
	for _, r := range ts.roll {
		result += r
		switch r {
		case 1:
			ones++
		case 6:
			sixes++
		}
	}
	tn := 0
	for _, n := range ts.assets {
		tn += n.value
	}
	ts.resolution = "Failure"
	if result <= tn {
		ts.resolution = "Success"
	}
	if tn < ts.diff {
		ts.resolution = "Spectacilar Stupid"
	}
	if ones > 2 {
		ts.resolution = "Spectacilar Failure"
	}
	if sixes > 2 {
		ts.resolution = "Spectacilar Success"
	}
	if ones > 2 && sixes > 2 {
		ts.resolution = "Spectacilar Interesting"
	}
}

func Difficulty(diff int) string {
	switch diff {
	default:
		return "[INVALID]"
	case DifficultyEasy:
		return "Easy"
	case DifficultyAverage:
		return "Average"
	case DifficultyDifficult:
		return "Difficult"
	case DifficultyFormidable:
		return "Formidable"
	case DifficultyStaggering:
		return "Staggering"
	case DifficultyHopeless:
		return "Hopeless"
	case DifficultyImposibble:
		return "Imposibble"
	case DifficultyBeyondImpossible:
		return "Beyond Impossible"
	}
}

func (t *task) Phrase() string {
	return "To " + t.descr
}

func (t *task) assetSum() string {
	st := ""
	for _, asset := range t.assets {
		st += fmt.Sprintf("%v + ", asset.descr)
	}
	st = strings.TrimSuffix(st, " + ")
	st += " ("
	for _, asset := range t.assets {
		switch {
		case asset.value > -1:

			st += fmt.Sprintf("%v+", asset.value)
		default:
			st = strings.TrimSuffix(st, "+")
			st += fmt.Sprintf("%v+", asset.value)
		}
	}
	st = strings.TrimSuffix(st, "+")
	st += ")"
	return st
}

func (t *task) Statement() string {
	return fmt.Sprintf("Difficulty (%vD) <= %v", t.diff, t.assetSum())
}

func (t *task) tn() int {
	tn := 0
	for _, a := range t.assets {
		tn += a.value
	}
	return tn
}

func (t *task) Result() string {
	if t.resolution == "Unresolved" {
		return t.resolution
	}
	text := "["
	sum := 0
	for _, val := range t.roll {
		text += fmt.Sprintf("%v ", val)
		sum += val
	}
	text = strings.TrimSuffix(text, " ")
	text = text + "]" + fmt.Sprintf(" (%v) ", sum)
	switch {
	default:
		text += fmt.Sprintf("> %v\n%v", t.tn(), t.resolution)
	case sum <= t.tn():
		text += fmt.Sprintf("<= %v\n%v", t.tn(), t.resolution)
	}
	return text
}
