package cmd

import "github.com/urfave/cli/v2"

func NewCharacter() *cli.Command {
	cmnd := &cli.Command{
		Name:        "new_character",
		Aliases:     []string{},
		Usage:       "generate new HOSTILE charater with Random Creation",
		UsageText:   "",
		Description: "",
		ArgsUsage:   "",
		Category:    "",
		BashComplete: func(*cli.Context) {
		},
		// Before: func(*cli.Context) error {
		// },
		// After: func(*cli.Context) error {
		// },
		Action: func(*cli.Context) error {
			return nil
		},
		// OnUsageError: func(cCtx *cli.Context, err error, isSubcommand bool) error {
		// },
		Subcommands:            []*cli.Command{},
		Flags:                  []cli.Flag{},
		SkipFlagParsing:        false,
		HideHelp:               false,
		HideHelpCommand:        false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
	return cmnd
}
