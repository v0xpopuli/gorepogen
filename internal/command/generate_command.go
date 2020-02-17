package command

import "github.com/urfave/cli/v2"

type GenerateCommand interface {
	generate(*cli.Context) error
	buildFlags() []cli.Flag
	checkArgs() error
	CreateCommand() *cli.Command
}
