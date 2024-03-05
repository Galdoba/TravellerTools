package main

import (
	"fmt"
	"os"

	"github.com/Galdoba/TravellerTools/hostile/struct/world"
	"github.com/Galdoba/TravellerTools/hostile/trade/tradecodes"
	"github.com/Galdoba/TravellerTools/hostile/uwp"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/urfave/cli/v2"
)

/*
run


*/

var configPath string

const (
	programName = "hostile"
)

func main() {

	app := cli.NewApp()
	app.Version = "v 0.0.1"
	app.Name = programName
	app.Usage = "набор генераторов для настолки Hostile"
	app.Flags = []cli.Flag{}

	//ДО НАЧАЛА ДЕЙСТВИЯ
	app.Before = func(c *cli.Context) error {

		return nil
	}
	app.Commands = []*cli.Command{

		{
			Name:      "world_data",
			Usage:     "генерит случайный мир",
			ArgsUsage: "--  --",
			Flags:     []cli.Flag{},
			Action: func(c *cli.Context) error {
				dice := dice.New()
				wrld := world.NewMainWorld()
				// fmt.Println(wrld)
				ruleset := make(map[string]string)
				data := uwp.GenerateMainWorldUWP(dice, ruleset)
				wrld.Feed(data)
				wrld.UWP = uwp.HostileUWP(data)
				wrld.Trade_Codes = tradecodes.ParseTradeCodes(data)
				wrld.GenerateDetails(dice)
				// fmt.Println(wrld)
				wrld.MarshalJson(`test.json`)

				return nil
			},
		},
		cmd.NewCharacter(),
	}

	//ПО ОКОНЧАНИЮ ДЕЙСТВИЯ
	app.After = func(c *cli.Context) error {
		return nil
	}
	args := os.Args
	if err := app.Run(args); err != nil {
		errOut := fmt.Sprintf("%v error: %v", programName, err.Error())
		println(errOut)
		os.Exit(1)
	}

}

/*
tgnotyfier send -t "--------------------" -m "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum." -ps "PS: Владыка, услышь меня!"
*/

func assertNoError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

////////////////////////////
