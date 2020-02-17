package main

import (
	"github.com/urfave/cli/v2"
	"github.com/v0xpopuli/gorepogen/internal/command"
)

func createApp() *cli.App {
	return &cli.App{
		Name:      "gorepogen",
		Usage:     "tool for repositories auto generation",
		UsageText: "gorepogen [global options]",
		Version:   "1.0.0",
		Authors:   getAuthor(),
		Commands: []*cli.Command{
			command.NewGenerateFromEntity().CreateCommand(),
			command.NewGenerateFromDatabase().CreateCommand(),
		},
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
