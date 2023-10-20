package main

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/internal/rolltable"
	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func main() {
	rlr := rolltable.NewRoller(dice.New())
	tableSet := rolltable.SetupNew("1",
		rolltable.NewTable("1", "1d6",
			rolltable.NewOutcome("4-", "1 1", "2"),
			rolltable.NewOutcome("5+", "1 2", "3"),
		),
		rolltable.NewTable("2", "2d6",
			rolltable.NewOutcome("3-", "first 2", ""),
			rolltable.NewOutcome("4|10", "second 2", "1"),
			rolltable.NewOutcome("11+", "third 2", ""),
		),
		rolltable.NewTable("3", "2d6",
			rolltable.NewOutcome("2+", "333 2", ""),
		),
	)
	err := rlr.Roll(tableSet)
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	for i, o := range rlr.Outcome() {
		fmt.Println(i, o)
	}
}
