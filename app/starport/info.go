package main

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/astrogation"
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/urfave/cli"
)

func Info(c *cli.Context) error {
	searchKey := c.String("worldname")

	sourceworldMain, err := SearchSourcePort(searchKey)
	if err != nil {
		return err
	}
	reach := 4
	fmt.Printf("Sourceworld [%v] detected...\nChecking for Trade Routs within a reach of %v parsecs...\n", sourceworldMain.MW_Name(), reach)
	transitWorldsCoordinates := searchNeighbours(sourceworldMain, 2)
	transitWorldsCoordinates = append(transitWorldsCoordinates, hexagon.New_Unsafe(hexagon.Feed_HEX, sourceworldMain.CoordX(), sourceworldMain.CoordY()))
	targetWorldsCoordinates := searchNeighbours(sourceworldMain, 4)

	//tradePointCoordinates := append(targetWorldsCoordinates, astrogation.NewCoordinates(sourceworld.CoordX(), sourceworld.CoordY()))
	//fmt.Println(targetWorldsCoordinates)
	//fmt.Printf("%v have is a part of tra...\n", sourceworldMain.MW_Name())
	/////////////////ПРОВЕРЯЕМ ТОРГОВЫЕ ПУТИ:
	trafficMap := make(map[string]int)
	for _, sourceworldHex := range transitWorldsCoordinates {
		sourceworld, _ := PortByCoordinates(sourceworldHex.HexValues())
		for _, crds := range targetWorldsCoordinates {
			port, _ := PortByCoordinates(crds.HexValues())
			fmt.Println("Evaluating Trade Route from", sourceworld.MW_Name(), "to", port.MW_Name()) //, port.MW_Name())
			if hexagon.DistanceHex(port, sourceworld) == 0 {
				fmt.Println("REJECTED: self-trade imposibble")
				continue
			}
			if hexagon.DistanceHex(port, sourceworldMain) > 4 {
				fmt.Println("REJECTED: not in a Trade Zone of", sourceworldMain.MW_Name())
			}
			if sourceworld.MW_Name() != sourceworldMain.MW_Name() && hexagon.DistanceHex(port, sourceworld) > 2 {
				fmt.Println("REJECTED:", sourceworldMain.MW_Name(), "is not a transit World")
				continue
			}
			if hexagon.DistanceHex(sourceworldMain, sourceworld) > 2 {
				fmt.Println("REJECTED:", sourceworld.MW_Name(), "is not a transit World")
				continue
			}
			if port.TravelZone() == "R" {
				fmt.Println("REJECTED:", port.MW_Name(), "is a RED ZONE")
				continue
			}
			if sourceworld.TravelZone() == "R" {
				fmt.Println("REJECTED:", sourceworld.MW_Name(), "is a RED ZONE")
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
			if (cargoIN || cargoOUT) && astrogation.TradeRouteExist(Hexagon(sourceworld), Hexagon(port), targetWorldsCoordinates) {

				src := Hexagon(sourceworld)
				dest := Hexagon(port)
				path := astrogation.PlotCource(src.AsHex(), dest.AsHex(), 2, 1)
				fmt.Print(path)
				if !strings.Contains(path, sourceworldMain.MW_Name()) {
					fmt.Println(" REJECTED:", sourceworldMain.MW_Name(), "recieves no traffic")
					continue
				}
				portNames := strings.Split(path, " ---> ")
				for _, pn := range portNames {
					trafficMap[pn]++
				}
				fmt.Println(" ACCEPTED!")
				for i := 62; i > -1; i-- {
					for k, v := range trafficMap {
						if i == v {
							fmt.Println(k, v)
						}
					}
				}
			} else {
				fmt.Println("REJECTED: no trade")
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
	}
	fmt.Println("Gathering traffic data:")
	for i := 62; i > -1; i-- {
		for k, v := range trafficMap {
			if i == v {
				fmt.Println(k, v)
			}
		}
	}

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
