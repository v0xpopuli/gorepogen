package main

import (
	"fmt"
	g "gorepogen/internal/generator"
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

func main() {

	app := &cli.App{
		Name:      "gorepogen",
		Usage:     "tool for repositories auto generation",
		UsageText: "gorepogen [entity name]",
		Version:   "1.0.0",
		Author:    "v0xpopuli",
		Action:    generate,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func generate(c *cli.Context) error {

	cd, _ := os.Getwd()
	args := c.Args()
	if len(args) == 0 {
		return errors.New("provide entity name please")
	}

	entityInfo, err := g.Search(cd, args.Get(0))
	if err != nil {
		return err
	}

	namesRegistry := g.CreateNamesRegistry(entityInfo)
	components := g.AssignNamesToComponents(namesRegistry)

	repositoryFullPath, err := g.Generate(components, namesRegistry, cd)
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
