package main

import (
	"fmt"
	"os"

	g "github.com/v0xpopuli/gorepogen/internal/generator"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var name string
var root string

func main() {

	app := &cli.App{
		Name:      "gorepogen",
		Usage:     "tool for repositories auto generation",
		UsageText: "gorepogen [global options]",
		Version:   "1.0.0",
		Authors: []*cli.Author{
			{
				Name:  "v0xpopuli",
				Email: "vadim.rozhkalns@gmail.com",
			},
		},
		Flags: []cli.Flag{
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
		},
		Action: generate,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func generate(_ *cli.Context) error {

	if root == "" {
		root, _ = os.Getwd()
	}

	if name == "" {
		return errors.New(`use "gorepogen -h" for help`)
	}

	entityInfo, err := g.Search(root, name)
	if err != nil {
		return err
	}

	namesRegistry := g.CreateNamesRegistry(entityInfo)
	components := g.AssignNamesToComponents(namesRegistry)

	repositoryFullPath, err := g.Generate(components, namesRegistry, root)
	if err != nil {
		return err
	}

	fmt.Printf(
		"repository for %s generated successfully, location: %s\n",
		namesRegistry.EntityName,
		repositoryFullPath,
	)

	return nil
}
