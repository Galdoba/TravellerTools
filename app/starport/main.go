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


*/

func main() {
	app := cli.NewApp()
	app.Version = "v 0.0.2"
	app.Name = "starport"
	app.Usage = "Коллекция инструментов для Рефери, чтобы узнать всякое о Звездном Порте"
	app.Commands = []cli.Command{
		//////////////////////////////////////
		//starport traffic -world [%v] -distance [3]
		{
			Name:      "traffic",
			Usage:     "Рассчитывает пассажирские, грузовые и (TODO) торговые потоки между портом и его соседями. (требует OTU базу)",
			UsageText: "Для вывода требуются флаги WorldName (имя мира по базе), максимальный радиус для расчета (дефолт 3) и RuleSet (по дефолту MGT2_Core)",
			Category:  "Информация",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "worldname",
					Usage:    "запрос, по которому надо искать Мир в Базе (ОБЯЗАТЕЛЕН)",
					Required: false,
					Value:    "",
				},
				&cli.Int64Flag{
					Name:     "reach",
					Usage:    "радиус поиска соседей (дефолтное значение 0: адаптивно - если TL12-, то 2, иначе TL-10 )",
					Required: false,
					Value:    0,
				},
				&cli.StringFlag{
					Name:     "ruleset",
					Usage:    "обеспечивает выбор набора правил",
					Required: false,
					Value:    "mgt1_mp",
				},
			},
			Action: func(c *cli.Context) error {
				return Traffic(c)
			},
		},
		{
			Name:     "info",
			Usage:    "расписывает информацию о порте",
			Category: "Информация",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "worldname",
					Usage:    "запрос, по которому надо искать Мир в Базе (ОБЯЗАТЕЛЕН)",
					Required: false,
					Value:    "",
				},
				&cli.Int64Flag{
					Name:     "reach",
					Usage:    "радиус поиска соседей (дефолтное значение 0: адаптивно - если TL12-, то 2, иначе TL-10 )",
					Required: false,
					Value:    0,
				},
			},
			Action: func(c *cli.Context) error {
				return Info(c)
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

World  : Regina (Spinward Marches 1910)
UWP    : A788899-C
TC/Rem : Ri Pa Ph An Cp (Green Zone)
---------------------------------------
Spaceport Traffic Report:
There are 20 worlds in 4 parsecs radius. [&WORLD_NAME] is not located on a Trade Route.


*/
