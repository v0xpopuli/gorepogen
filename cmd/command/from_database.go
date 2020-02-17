package command

import "github.com/urfave/cli/v2"

func GenerateRepositoriesFromDatabase() *cli.Command {
	return &cli.Command{
		Name:   "gendb",
		Usage:  "generate repositories from database",
		Action: generateRepositoriesFromDatabase,
	}
}

func generateRepositoriesFromDatabase(_ *cli.Context) error {
	return nil
}
