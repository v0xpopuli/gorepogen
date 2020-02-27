package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	gen "github.com/v0xpopuli/gorepogen/internal/generator"
	repo "github.com/v0xpopuli/gorepogen/internal/generator/repository"
)

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
		Flags:   g.getFlags(),
		Action:  g.generate,
	}
}

func (g generateFromEntity) generate(ctx *cli.Context) error {

	entityInfo, err := gen.NewWalker(filepath.Base(root), name).Walk(root)
	if err != nil {
		return err
	}

	repositoryFullPath, err := repo.NewGenerator(
		g.resolveOutputDir(ctx.String("output")),
		gen.NewNamesRegister(entityInfo),
	).Generate()
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

func (g generateFromEntity) resolveOutputDir(outputDir string) string {
	if outputDir == "" {
		return root
	}
	return outputDir
}
