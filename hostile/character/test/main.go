package main

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/hostile/character"
)

func main() {
	gen := character.NewGenerator(character.Option(character.KeyManual, "YES"))
	// gen.dice = dice.New()
	chr, err := gen.Generate_By_Events()
	if err != nil {
		fmt.Errorf("%v", err)
	}
	fmt.Println("=========")
	chr.FlushScreen()
}
