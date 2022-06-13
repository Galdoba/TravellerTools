package main

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/astrogation"
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/Galdoba/TravellerTools/pkg/mgt2trade/traffic/tradecodes"
	"github.com/Galdoba/TravellerTools/pkg/starport/portsec"
	"github.com/Galdoba/TravellerTools/pkg/starport/sai"
	"github.com/Galdoba/TravellerTools/pkg/starport/ssp"
	"github.com/urfave/cli"
)



func Info(c *cli.Context) error {

	searchKey := c.String("worldname")

	sourceworldMain := WorldFrom(searchKey)
	fmt.Println(sourceworldMain.String())
	reach := 4
	//////////////////////
	securityProfile, err := ssp.NewSecurityProfile(sourceworldMain)
	fmt.Println(securityProfile.Describe())
	portSecurity, err := portsec.GenerateSecurityForces(sourceworldMain)
	fmt.Println(portSecurity.String())

	//Собираем список всех портов и всех координат
	fmt.Println("Constructing J-4 map...")
	allPorts := []Port{}
	//tradeZone, _ := hexagon.Spiral(sourceWorldHex.AsCube(), 4)
	allWorldsCoordinates := append([]hexagon.Hexagon{Hexagon(sourceworldMain)}, searchNeighbours(sourceworldMain, reach)...)
	for _, hex := range allWorldsCoordinates {
		p, _ := PortByCoordinates(hex.HexValues())
		allPorts = append(allPorts, p)
	}

	fmt.Println("Evaluating Trade Routes...")
	routes := evaluateTradeRoutes(allPorts, sourceworldMain)
	shipping := []int{0, 0, 0}
	if len(routes) > 0 {
		fmt.Println("Trade Routes Detected:")
		for _, tr := range routes {

			fmt.Printf("%v: %v \n", tr.status, tr.jp.Path)
			

		}
		shipping = shippingActivityBase(sourceworldMain.MW_Name(), routes)
		fmt.Printf("Arrive/Depart/Transit cargo traffic: %v/%v/%v\n", shipping[0], shipping[1], shipping[2])

		//fmt.Printf("Average Ships in Port at any given point: %v", ((arrive+depart+transit+transit)*3)*7/2)
	} else {
		fmt.Printf("%v have no constant trade with neighbour worlds.\n", sourceworldMain.MW_Name())
	}
	fmt.Println("--------------------------------------------------------------------------------")
	sa, err := sai.NewShippingActivity(sourceworldMain, shipping)
	if err != nil {
		return err
	}
	fmt.Println(sa)

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
	jp          astrogation.Plot
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
				if !portVisited(portOfInterest, jp.Path) {
					continue
				}
				trade := tradeRoute{src, dest, "Transit", jp}
				names := strings.Split(jp.Path, " ---> ")
				for i, name := range names {
					if name == portOfInterest.MW_Name() {
						switch i {
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

func shippingActivityBase(mw_name string, routes []tradeRoute) []int {
	arrive := 0
	depart := 0
	transit := 0
	for _, tr := range routes {
		if strings.Contains(tr.jp.Path, mw_name+" --->") {
			depart++
		}
		if strings.Contains(tr.jp.Path, "---> "+mw_name) {
			arrive++
		}
		if strings.Contains(tr.jp.Path, "---> "+mw_name+" --->") {
			transit++
			arrive--
			depart--
		}
	}
	return []int{arrive, depart, transit}
}

func portVisited(p Port, path string) bool {
	for _, poi := range strings.Split(path, " ---> ") {
		if poi == p.MW_Name() {
			return true
		}
	}
	return false
}

func tradePossible(source, destination Port) bool {
	sTC, _ := tradecodes.FromUWPstr(source.MW_UWP())
	dTC, _ := tradecodes.FromUWPstr(destination.MW_UWP())
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
