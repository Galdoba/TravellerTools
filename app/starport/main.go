package main

import (
	"fmt"
	"os"

	"github.com/Galdoba/TravellerTools/app/modules/trvdb"
	"github.com/Galdoba/TravellerTools/pkg/astrogation"
	"github.com/Galdoba/TravellerTools/pkg/survey"
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
					Name:        "worldname",
					Usage:       "запрос, по которому надо искать Мир в Базе (ОБЯЗАТЕЛЕН)",
					Required:    false,
					Value:       "",
					Destination: new(string),
				},
				&cli.StringFlag{
					Name:        "ruleset",
					Usage:       "обеспечивает выбор набора правил",
					Required:    false,
					Value:       "mgt2_core",
					Destination: new(string),
				},
			},
			Action: func(c *cli.Context) error {

				return Traffic(c)
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

func main0() {

	wrlds, err := trvdb.WorldByName()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Found", wrlds)
	sw := wrlds

	fmt.Println("Source World Chosen:", fmt.Sprintf("%v (%v)/%v %v\n", sw.MW_Name(), sw.MW_UWP(), sw.Sector(), sw.Hex()))
	jcoord := astrogation.JumpFromCoordinates(astrogation.NewCoordinates(sw.CoordX(), sw.CoordY()), 3)
	fmt.Println("Trade capable worlds:")
	for i, v := range jcoord {
		fmt.Printf("Search %v/%v\r", i, len(jcoord))
		nWorld, err := survey.SearchByCoordinates(v.ValuesHEX())

		if err != nil {
			//x, y := v.ValuesHEX()
			//fmt.Println(x, y, err.Error())♦
			continue
		}
		if nWorld.CoordX() == sw.CoordX() && nWorld.CoordY() == sw.CoordY() {
			continue
		}
		fmt.Println(fmt.Sprintf("%v (%v)/%v %v", nWorld.MW_Name(), nWorld.MW_UWP(), nWorld.Sector(), nWorld.Hex()))

	}
	fmt.Println("                               \r")
}
