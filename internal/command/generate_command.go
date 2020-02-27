package command

import "github.com/urfave/cli/v2"

type GenerateCommand interface {
	generate(*cli.Context) error
	getFlags() []cli.Flag
	CreateCommand() *cli.Command
}
