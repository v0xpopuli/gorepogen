package main

import (
	"os"

	"github.com/urfave/cli/v2"
	"github.com/v0xpopuli/gorepogen/internal/command"
)

var output string

func createApp() *cli.App {
	return &cli.App{
		Name:     "gorepogen",
		Usage:    "tool for repositories auto generation",
		Version:  "1.0.0",
		Authors:  getAuthor(),
		Commands: getCommands(),
		Flags:    getGlobalFlags(),
		Before:   globalFlagsValidation,
	}
}

func getAuthor() []*cli.Author {
	return []*cli.Author{
		{
			Name:  "v0xpopuli",
			Email: "vadim.rozhkalns@gmail.com",
		},
	}
}

func getCommands() []*cli.Command {
	return []*cli.Command{
		command.NewGenerateFromEntity().CreateCommand(),
		command.NewGenerateFromDatabase().CreateCommand(),
	}
}

func getGlobalFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "Output directory",
			Destination: &output,
			Required:    false,
		},
	}
}

func globalFlagsValidation(_ *cli.Context) error {
	if output != "" {
		if _, err := os.Stat(output); os.IsNotExist(err) {
			return cli.Exit(err, 1)
		}
	}
	return nil
}
