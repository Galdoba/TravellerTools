package main

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/astrogation"
	"github.com/urfave/cli"
)

func Info(c *cli.Context) error {
	searchKey := c.String("worldname")
	reach := c.Int("reach")
	sourceworld, err := SearchSourcePort(searchKey)
	if err != nil {
		return err
	}
	fmt.Printf("Sourceworld [%v] detected...\nChecking for neighbours within a reach of %v parsecs...\n", sourceworld.MW_Name(), reach)
	targetWorldsCoordinates := searchNeighbours(sourceworld, 2)
	tradePointCoordinates := append(targetWorldsCoordinates, astrogation.NewCoordinates(sourceworld.CoordX(), sourceworld.CoordY()))
	fmt.Println(targetWorldsCoordinates)
	/////////////////ПРОВЕРЯЕМ ТОРГОВЫЕ ПУТИ:
	for _, crds := range tradePointCoordinates {
		port, _ := PortByCoordinates(crds.ValuesHEX())
		tcodes := strings.Fields(TradeCodes(port))
		cargoIN := false
		cargoOUT := false
		for _, tc := range tcodes {
			switch tc {
			case "In", "Ht", "Hi", "Ri":
				cargoIN = true
			case "As", "De", "Ic", "Ni", "Ag", "Ga", "Wa":
				cargoOUT = true
			}
		}
		// if cargoIN {
		// 	fmt.Println(port.MW_Name(), "receiving Cargo")
		// 	continue
		// }

		// if cargoOUT && cargoIN {
		// 	fmt.Println(port.MW_Name(), "is NEXUS?")
		// 	continue
		// }

		// if cargoOUT {
		// 	fmt.Println(port.MW_Name(), "sending Cargo")
		// }

		fmt.Println("IN:", cargoIN, "OUT:", cargoOUT, "---", port.MW_Name())

		// if !cargoOUT && !cargoIN {
		// 	fmt.Println(port.MW_Name(), "is BACKWATER?")
		// }
	}

	fmt.Println("Gathering traffic data:")
	return nil
}

/*
+Camoran
+Khusai
+Kteiroa
+Asim
+Marduk
+Iroioah
+Tyokh
+Torpol
+The World
+Pourne
+Clarke
+Sink
+Paal
+Hilfer
+Blue
+Ergo
+Tech-World
+Drinax

*/
