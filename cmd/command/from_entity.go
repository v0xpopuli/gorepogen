package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	g "github.com/v0xpopuli/gorepogen/internal/generator"
)

var name string
var root string

func GenerateRepositoryFromEntity() *cli.Command {
	return &cli.Command{
		Name:   "gen",
		Usage:  "generate repository from entity name",
		Flags:  buildFlags(),
		Action: generateRepositoryFromEntity,
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

func checkArgs() error {
	if root == "" {
		root, _ = os.Getwd()
	}

	if name == "" {
		return errors.New(`Use "gorepogen -h" for help`)
	}
	return nil
}

func generateRepositoryFromEntity(_ *cli.Context) error {

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

	fmt.Printf(
		"Repository for %s generated successfully, location: %s\n",
		entityInfo.EntityName,
		repositoryFullPath,
	)
	return nil
}
