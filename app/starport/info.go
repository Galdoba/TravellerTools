package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

func Info(c *cli.Context) error {
	searchKey := c.String("worldname")

	sourceworld, err := SearchSourcePort(searchKey)
	if err != nil {
		return err
	}
	reach := 4
	fmt.Printf("Sourceworld [%v] detected...\nChecking for Trade Routs within a reach of %v parsecs...\n", sourceworld.MW_Name(), reach)
	targetWorldsCoordinates := searchNeighbours(sourceworld, reach)
	//tradePointCoordinates := append(targetWorldsCoordinates, astrogation.NewCoordinates(sourceworld.CoordX(), sourceworld.CoordY()))
	//fmt.Println(targetWorldsCoordinates)
	fmt.Printf("%v have...\n", sourceworld.MW_Name())
	/////////////////ПРОВЕРЯЕМ ТОРГОВЫЕ ПУТИ:
	for _, crds := range targetWorldsCoordinates {
		port, _ := PortByCoordinates(crds.HexValues())
		if port.TravelZone() == "R" || sourceworld.TravelZone() == "R" {
			continue
		}
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
		if cargoIN || cargoOUT {
			fmt.Println("Trade route with", port.MW_Name())
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

		//fmt.Println("IN:", cargoIN, "OUT:", cargoOUT, "---", port.MW_Name())

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
