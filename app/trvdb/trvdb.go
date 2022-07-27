package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

const (
	dataBase = "c:\\Users\\Public\\TrvData\\cleanedData.txt"
)

/*
1 найти мир
	запрос:
		0 найдено
			END
		n найдено
			выбрать мир
			GO TO 2
		50+ найдено
			END
2 найти соседей мира
3 расчитать трафик
curl -v --proxy "http://proxy.local:3128" https://docs.google.com/spreadsheets/d/1Waa58usrgEal2Da6tyayaowiWujpm0rzd06P5ASYlsg/gviz/tq?tqx=out:csv -k --output content.csv
curl -v "traveller.com/data"

*/

func main() {
	app := cli.NewApp()
	app.Version = "v 0.0.1"
	app.Name = "trvdb"
	app.Usage = "Менеджер базы данных вселенной Traveller"
	app.Commands = []cli.Command{
		//////////////////////////////////////
		//starport traffic -world [%v] -distance [3]
		{
			Name:      "update",
			Usage:     "Обновляет базу данных из интернета",
			UsageText: "Скачивает данные по секторам с сайта http://travellermap.com используя кучу curl запросов, после чего форматирует полученную информацию в отдельный файл и форматирует в формат `T5 Second Survey`",
			Category:  "Информация",
			Flags:     []cli.Flag{},
			Action: func(c *cli.Context) error {
				//основное тело
				return nil
			},
		},

		//////////////////////////////////////
		{
			Name:        "TEMPLATE",
			ShortName:   "",
			Aliases:     []string{},
			Usage:       "Create today's work directories and daily files",
			UsageText:   "",
			Description: "",
			ArgsUsage:   "",
			Category:    "",
			BashComplete: func(*cli.Context) {
			},
			Before: func(*cli.Context) error {
				return nil
			},
			After: func(*cli.Context) error {
				return nil
			},
			Action: func(c *cli.Context) error {
				return nil
			},
			OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
				return nil
			},
			Subcommands:            []cli.Command{},
			Flags:                  []cli.Flag{},
			SkipFlagParsing:        false,
			SkipArgReorder:         false,
			HideHelp:               false,
			Hidden:                 false,
			UseShortOptionHandling: false,
			HelpName:               "",
			CustomHelpTemplate:     "",
		},
		//////////////////////////////////////
	}
	args := os.Args
	if len(args) < 2 {
		args = append(args, "help") //Принудительно зовем помощь если нет других аргументов
	}
	if err := app.Run(args); err != nil {
		fmt.Println(err.Error())
	}
}

/*RESULT DATA

World  : Regina / Spinward Marches 1910
UWP    : A788899-C
TC/Rem : Ri Pa Ph An Cp (Green Zone)
---------------------------------------
Spaceport Traffic Report:
There are 20 worlds in 4 parsecs radius. [&WORLD_NAME] is not located on a Trade Route.


World  : Mirak / Ea / Reaver's Deep 1127
UWP    : C766763-8
TC/Rem : Ag Ga Ri Pz Mr (Amber Zone)
---------------------------------------
Spaceport Traffic Report:
There are 17 worlds in 4 parsecs radius. [&WORLD_NAME] is not located on a Trade Route.


*/
