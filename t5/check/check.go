package check

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	NONE = iota
	SpectacularSuccess
	SpectacularFailure
	SpectacularlyInteresting
	SpectacularlyStupid
)

const (
	Easy             = 1
	Average          = 2
	Difficult        = 3
	Formidable       = 4
	Staggering       = 5
	Hopeless         = 6
	Impossible       = 7
	BeyondImpossible = 8
)

type Asset interface {
	AssetVal(string) int
	AssetDescr() []string
}

type Resolution struct {
	RollResolved bool
	Succes       bool
	Spectaculars []int
	Comments     []string
}

type task struct {
	prequisite string
	preqNotMet bool
	diff       int
	assets     []Asset
	vocal      bool
}

type prequisite struct {
	body       string
	definition string
	val        int
	more       bool
	less       bool
}

func NewCheck(diff int, assets ...Asset) *task {
	chk := task{}
	chk.diff = diff
	chk.assets = assets
	return &chk
}

func preqIsMet(preq string, assets ...Asset) bool {
	if preq == "" {
		return true
	}
	p := prequisite{}
	splitter := " "
	switch {
	case strings.Contains(preq, "-"):
		preq = strings.TrimSuffix(preq, "-")
		p.less = true
	case strings.Contains(preq, "+"):
		preq = strings.TrimSuffix(preq, "+")
		p.more = true
	case strings.Contains(preq, "="):
		splitter = "="
	}
	pr := strings.Split(preq, splitter)
	p.body = pr[0]
	if len(pr) == 1 {
		pr = append(pr, pr[0])
	}
	p.definition = pr[1]
	p.val, _ = strconv.Atoi(p.definition)
	for _, asset := range assets {
		switch p.val {
		case 0:
			for _, def := range asset.AssetDescr() {
				if def == p.definition {
					return true
				}
			}
		default:
			if p.val == asset.AssetVal(p.body) {
				return true
			}
			if p.less {
				if p.val >= asset.AssetVal(p.body) {
					return true
				}
			}
			if p.more {
				if p.val <= asset.AssetVal(p.body) {
					return true
				}
			}
		}
	}
	return false
}

func (chk *task) WithPrequisite(preq string) *task {
	chk.prequisite = preq
	return chk
}

func (chk *task) Vocal(preq string) *task {
	chk.vocal = true
	return chk
}

func (chk *task) tn() int {
	s := 0
	for _, asset := range chk.assets {
		s += asset.AssetVal("selfval")
	}
	return s
}

func Resolve(chk *task, dice *dice.Dicepool) Resolution {
	res := Resolution{}
	if !preqIsMet(chk.prequisite, chk.assets...) {
		res.Comments = append(res.Comments, "Prequisite not met")
		return res
	}
	fmt.Println(fmt.Sprintf("%vd6", chk.diff))
	rollmap := dice.RollMap(fmt.Sprintf("%vd6", chk.diff))
	res.Comments = append(res.Comments, fmt.Sprintf("Roll%v", rollmap))
	tn := chk.tn()
	if len(rollmap) > tn {
		res.Spectaculars = append(res.Spectaculars, SpectacularlyStupid)
		res.Comments = append(res.Comments, fmt.Sprintf("Spectacularly Stupid"))
	}
	if rollmap[1] >= 3 && rollmap[6] >= 3 {
		res.Spectaculars = append(res.Spectaculars, SpectacularlyInteresting)
		res.Comments = append(res.Comments, fmt.Sprintf("Spectacularly Interesting"))
	}
	if rollmap[6] >= 3 {
		res.Spectaculars = append(res.Spectaculars, SpectacularFailure)
		res.Comments = append(res.Comments, fmt.Sprintf("Spectacular Failure"))
	}
	if rollmap[1] >= 3 {
		res.Spectaculars = append(res.Spectaculars, SpectacularSuccess)
		res.Comments = append(res.Comments, fmt.Sprintf("Spectacular Success"))
		res.Succes = true
	}
	if sum(rollmap) < tn {
		res.Succes = true
	}
	if res.Succes {
		res.Comments = append(res.Comments, fmt.Sprintf("Success"))
	}
	res.RollResolved = true
	return res
}

func sum(rollMap map[int]int) int {
	s := 0
	for k, _ := range rollMap {
		s += k
	}
	return s
}

/*
pawn.Resolve(check,dice) Resolution



*/
