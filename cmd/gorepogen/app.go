package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	g "github.com/v0xpopuli/gorepogen/internal/generator"
)

func createApp() *cli.App {
	return &cli.App{
		Name:      "gorepogen",
		Usage:     "tool for repositories auto generation",
		UsageText: "gorepogen [global options]",
		Version:   "1.0.0",
		Authors:   buildAuthor(),
		Flags:     buildFlags(),
		Action:    generate,
	}
}

func buildAuthor() []*cli.Author {
	return []*cli.Author{
		{
			Name:  "v0xpopuli",
			Email: "vadim.rozhkalns@gmail.com",
		},
	}
}

func buildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Aliases:     []string{"n"},
			Usage:       "Entity name",
			Destination: &name,
		},
		&cli.StringFlag{
			Name:        "root",
			Aliases:     []string{"r"},
			Usage:       "Project root",
			Destination: &root,
		},
	}
}

func generate(_ *cli.Context) error {

	if err := checkArgs(); err != nil {
		return err
	}

	entityInfo, err := g.NewWalker(filepath.Base(root), name).Walk(root)
	if err != nil {
		return err
	}

	repositoryFullPath, err := g.NewGenerator(g.NewNamesRegister(entityInfo)).Generate(root)
	if err != nil {
		return err
	}

	fmt.Println(createResultMessage(entityInfo.EntityName, repositoryFullPath))
	return nil
}

func checkArgs() error {
	if root == "" {
		root, _ = os.Getwd()
	}

	if name == "" {
		return errors.New(`Use "gorepogen -h" for help`)
	}
	return nil
}

func createResultMessage(entityName, repositoryFullPath string) string {
	return fmt.Sprintf(
		"Repository for %s generated successfully, location: %s\n",
		entityName,
		repositoryFullPath,
	)
}
