package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"github.com/v0xpopuli/gorepogen/internal/component"
	"github.com/v0xpopuli/gorepogen/internal/generator"
	"github.com/v0xpopuli/gorepogen/internal/param"
	"github.com/v0xpopuli/gorepogen/internal/walker"
)

var (
	name, root string

	fromEntitySuccessMessage = "Repository for %s generated successfully, location: %s\n"
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
		Flags:   g.getFlags(),
		Action:  g.generate,
	}
}

func (g generateFromEntity) generate(ctx *cli.Context) error {

	info, err := walker.New(filepath.Base(root), name).Walk(root)
	if err != nil {
		return err
	}

	components := component.New(info).GetComponents()
	params := param.NewGeneratorParams(info, ctx.String("output"))
	fullRepoPath, err := generator.New(params, components).Generate()
	if err != nil {
		return err
	}

	fmt.Printf(fromEntitySuccessMessage, info.EntityName, fullRepoPath)
	return nil
}

func (g generateFromEntity) getFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Aliases:     []string{"n"},
			Usage:       "Entity name",
			Destination: &name,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "root",
			Aliases:     []string{"r"},
			Usage:       "Project root",
			Destination: &root,
			Value:       g.getCurrentPath(),
		},
	}
}

func (g generateFromEntity) getCurrentPath() string {
	path, _ := os.Getwd()
	return path
}
