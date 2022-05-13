package main

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/astrogation"
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/urfave/cli"
)

func Info0(c *cli.Context) error {
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
	trafficArriving := make(map[string]int)
	trafficTransit := make(map[string]int)
	trafficDepart := make(map[string]int)
	fmt.Println("Evaluating Trade Routes: ")
	for _, sourceworldHex := range transitWorldsCoordinates {
		sourceworld, _ := PortByCoordinates(sourceworldHex.HexValues())
		for _, crds := range targetWorldsCoordinates {
			fmt.Print("\r                                                                                                    \r")
			port, _ := PortByCoordinates(crds.HexValues())
			fmt.Print(sourceworld.MW_Name(), " ------>> ", port.MW_Name(), " ") //, port.MW_Name())
			if hexagon.DistanceHex(port, sourceworld) == 0 {
				fmt.Print("REJECTED: self-trade imposibble")
				continue
			}
			if hexagon.DistanceHex(port, sourceworldMain) > 4 {
				fmt.Print("REJECTED: not in a Trade Zone of", sourceworldMain.MW_Name())
				continue
			}
			if sourceworld.MW_Name() != sourceworldMain.MW_Name() && hexagon.DistanceHex(port, sourceworld) > 2 {
				fmt.Print("REJECTED:", sourceworldMain.MW_Name(), "is not a transit World")
				continue
			}
			// if hexagon.DistanceHex(sourceworldMain, sourceworld) > 2 {
			// 	fmt.Print("REJECTED:", sourceworld.MW_Name(), "is not a transit World")
			// 	continue
			// }
			if port.TravelZone() == "R" {
				fmt.Print("REJECTED:", port.MW_Name(), "is a RED ZONE")
				continue
			}
			if sourceworld.TravelZone() == "R" {
				fmt.Print("REJECTED:", sourceworld.MW_Name(), "is a RED ZONE")
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
			if cargoIN || cargoOUT { // && astrogation.TradeRouteExist(Hexagon(sourceworld), Hexagon(port), targetWorldsCoordinates) {

				src := Hexagon(sourceworld)
				dest := Hexagon(port)
				path, err := astrogation.PlotCource(src.AsHex(), dest.AsHex(), 2, 1)
				if path.Cost > 1000000 {
					fmt.Println(" REJECTED: no jumppoints connection")
					continue
				}
				if err != nil {
					return err
				}
				fmt.Print("[", path.Path, "]")
				if !strings.Contains(path.Path, sourceworldMain.MW_Name()) {
					fmt.Print(" REJECTED:", sourceworldMain.MW_Name(), "recieves no traffic")
					continue
				}
				portNames := strings.Split(path.Path, " ---> ")
				for pos, pn := range portNames {
					switch pos {
					default:
						trafficTransit[pn]++
					case 0:
						trafficDepart[pn]++
					case len(portNames) - 1:
						trafficArriving[pn]++
					}
				}
				fmt.Println(" ACCEPTED!")

			} else {
				fmt.Print("REJECTED: no trade")
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
	fmt.Println("Departing Ports:")
	for i := 62; i > -1; i-- {
		for k, v := range trafficDepart {
			if i == v {
				fmt.Println(k, v)
			}
		}
	}
	fmt.Println("Transit Ports:")
	for i := 62; i > -1; i-- {
		for k, v := range trafficTransit {
			if i == v {
				fmt.Println(k, v)
			}
		}
	}
	fmt.Println("Receiver Ports:")
	for i := 62; i > -1; i-- {
		for k, v := range trafficArriving {
			if i == v {
				fmt.Println(k, v)
			}
		}
	}

	return nil
}

func Info(c *cli.Context) error {
	searchKey := c.String("worldname")

	sourceworldMain, err := SearchSourcePort(searchKey)
	if err != nil {
		return err
	}
	reach := 4
	//////////////////////
	//Собираем список всех портов и всех координат
	allPorts := []Port{}
	//tradeZone, _ := hexagon.Spiral(sourceWorldHex.AsCube(), 4)
	allWorldsCoordinates := append([]hexagon.Hexagon{Hexagon(sourceworldMain)}, searchNeighbours(sourceworldMain, reach)...)
	for _, hex := range allWorldsCoordinates {
		p, _ := PortByCoordinates(hex.HexValues())
		allPorts = append(allPorts, p)
	}
	routes := evaluateTradeRoutes(allPorts, sourceworldMain)
	fmt.Println("Trade possible:")
	for _, tr := range routes {
		fmt.Printf("%v --> %v (%v)\n", tr.source.MW_Name(), tr.destination.MW_Name(), tr.status)
	}
	/////////////////ПРОВЕРЯЕМ ТОРГОВЫЕ ПУТИ:
	//состовляем все пары портов и проверяем их на возможность торговли по торговым кодам

	//tradePointCoordinates := append(targetWorldsCoordinates, astrogation.NewCoordinates(sourceworld.CoordX(), sourceworld.CoordY()))
	//fmt.Println(targetWorldsCoordinates)
	//fmt.Printf("%v have is a part of tra...\n", sourceworldMain.MW_Name())

	return nil
}

type tradeRoute struct {
	source      Port
	destination Port
	status      string
	//tradePossible bool
}

func evaluateTradeRoutes(ports []Port, portOfInterest Port) []tradeRoute {
	routes := []tradeRoute{}
	allChecks := len(ports) * len(ports)
	checksCompleted := 0
	trFound := 0
	for _, src := range ports {
		for _, dest := range ports {
			checksCompleted++
			fmt.Printf("\rChecks completed %v/%v (found: %v)", checksCompleted, allChecks, trFound)
			switch {
			case hexagon.Distance(Hexagon(src), Hexagon(dest)) > 4:
				continue
			case hexagon.Distance(Hexagon(src), Hexagon(dest)) == 0:
				continue
			// case src.TravelZone() == "R":
			// 	continue
			// case dest.TravelZone() == "R":
			// 	continue
			case tradePossible(src, dest):
				jp, _ := astrogation.PlotCource(src, dest, 2, 1)
				if !strings.Contains(jp.Path, portOfInterest.MW_Name()) {
					continue
				}
				trade := tradeRoute{src, dest, ""}
				names := strings.Split(jp.Path, " ---> ")
				for i, name := range names {
					if name == portOfInterest.MW_Name() {
						switch i {
						default:
							trade.status = "Transit"
						case 0:
							trade.status = "Export"
						case len(names) - 1:
							trade.status = "Import"
						}
					}

				}
				trFound++
				routes = append(routes, trade)
			}

		}
	}
	fmt.Println("")
	return routes
}

func tradePossible(source, destination Port) bool {
	sTC := strings.Fields(TradeCodes(source))
	dTC := strings.Fields(TradeCodes(destination))
	for _, tcS := range sTC {
		switch tcS {
		case "As", "De", "Ic", "Ni":
			for _, tcD := range dTC {
				switch tcD {
				case "In", "Ht":
					return true
				}
			}
		case "Ag", "Ga", "Wa":
			for _, tcD := range dTC {
				switch tcD {
				case "Hi", "Ri":
					return true
				}
			}
		}
	}
	return false
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
