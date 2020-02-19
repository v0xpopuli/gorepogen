package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	gen "github.com/v0xpopuli/gorepogen/internal/generator"
)

var ErrGenHelp = errors.New(`Use "gorepogen -h" for help`)

var (
	name, root string
)

type generateFromEntity struct{}

func NewGenerateFromEntity() GenerateCommand {
	return &generateFromEntity{}
}

func (g generateFromEntity) CreateCommand() *cli.Command {
	return &cli.Command{
		Name:    "gen",
		Aliases: []string{"g"},
		Usage:   "generate repository from entity name",
		Flags:   g.buildFlags(),
		Action:  g.generate,
	}
}

func (g generateFromEntity) generate(*cli.Context) error {
	if err := g.checkArgs(); err != nil {
		return err
	}

	entityInfo, err := gen.NewWalker(filepath.Base(root), name).Walk(root)
	if err != nil {
		return err
	}

	repositoryFullPath, err := gen.NewGenerator(gen.NewNamesRegister(entityInfo)).Generate(root)
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

func (generateFromEntity) buildFlags() []cli.Flag {
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

func (g generateFromEntity) checkArgs() error {
	if root == "" {
		root, _ = os.Getwd()
	}

	if name == "" {
		return ErrGenHelp
	}
	return nil
}
