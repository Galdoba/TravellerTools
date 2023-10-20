package rolltable

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

type Table struct {
	Name     string
	DiceCode string
	DMs      []string
	Line     []Outcome
}

type Outcome struct {
	Value         string
	ResultS       string
	NextTableName string
}

func defineMethod(s string) int {
	fmt.Println("DEFINE", s)
	if strings.HasSuffix(s, "+") {
		_, err := strconv.Atoi(strings.TrimSuffix(s, "+"))
		if err == nil {
			return 1
		}
	}
	if strings.HasSuffix(s, "-") {
		_, err := strconv.Atoi(strings.TrimSuffix(s, "-"))
		if err == nil {
			return 2
		}
	}
	_, err := strconv.Atoi(s)
	if err == nil {
		return 3
	}
	data := strings.Split(s, "|")
	for _, d := range data {
		_, err := strconv.Atoi(d)
		if err != nil {
			return -1
		}
	}
	if len(data) == 2 {
		return 4
	}
	return -1
}

func resultMatch(r int, check string) int {
	switch defineMethod(check) {
	default:
		return -1
	case 1:
		val, _ := strconv.Atoi(check)
		if r >= val {
			fmt.Println("result match", check, r, val)
			return 1
		}
	case 2:
		val, _ := strconv.Atoi(check)
		if r <= val {
			fmt.Println("result match", check, r, val)
			return 1
		}
	case 3:
		data := strings.Split(check, "|")
		d1, _ := strconv.Atoi(data[0])
		d2, _ := strconv.Atoi(data[1])
		if r >= d1 && r <= d2 {
			fmt.Println("result match", check, r)
			return 1
		}
	case 4:
		val, _ := strconv.Atoi(check)
		if r == val {
			fmt.Println("result match", check, r)
			return 1
		}
	}
	fmt.Println("result NOT match", check, r)
	return 0
}

type TableSetup struct {
	StartName string
	Tbl       map[string]Table
}

func SetupNew(start string, tables ...Table) TableSetup {
	ts := TableSetup{}
	ts.StartName = start
	ts.Tbl = make(map[string]Table)
	for _, table := range tables {
		ts.Tbl[table.Name] = table
	}
	return ts
}

func NewTable(name string, dicecode string, lines ...Outcome) Table {
	t := Table{}
	t.Name = name
	t.DiceCode = dicecode
	t.Line = append(t.Line, lines...)
	return t
}

func NewOutcome(check string, resS string, nextTable string) Outcome {
	return Outcome{check, resS, nextTable}
}

func getOutcome(r int, tbl Table) (Outcome, error) {
	for _, out := range tbl.Line {
		switch resultMatch(r, out.Value) {
		case -1:
			return Outcome{}, fmt.Errorf("bad table value: %v", out.Value)
		case 0:
			fmt.Println(out, "rejected")
			continue
		case 1:
			return out, nil
		}
	}
	return Outcome{}, fmt.Errorf("no result match")
}

type roller struct {
	dice     *dice.Dicepool
	DMs      map[string]int
	outcomes []Outcome
}

func NewRoller(dice *dice.Dicepool) *roller {
	rl := roller{}
	rl.dice = dice
	rl.DMs = make(map[string]int)
	return &rl
}

func (rl *roller) AddDM(key string, val int) {
	rl.DMs[key] = val
}

func (rl *roller) countDMs(dms ...string) int {
	dm := 0
	for _, d := range dms {
		dm += rl.DMs[d]
	}
	return dm
}

func (rl *roller) Roll(ts TableSetup) error {
	nextTable := ts.StartName
	out := []Outcome{}
	r := -1
	for nextTable != "" {
		fmt.Println("ROLL", nextTable)
		r = rl.dice.Sroll(ts.Tbl[nextTable].DiceCode)
		fmt.Println("TRY", r)
		o, err := getOutcome(r, ts.Tbl[nextTable])
		if err != nil {
			return err
		}
		out = append(out, o)
		nextTable = o.NextTableName
		fmt.Println(o, r)
	}
	rl.outcomes = out
	return nil
}

func (rl *roller) Outcome() []Outcome {
	return rl.outcomes
}

type Roller interface {
	AddDM(string, int)
	Roll(TableSetup) error
	Outcome() []Outcome
}
