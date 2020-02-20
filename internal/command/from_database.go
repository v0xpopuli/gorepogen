package command

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"github.com/v0xpopuli/gorepogen/internal/driver"
)

var (
	drvName, username, password, host, port, dbName, schema string

	ErrGenDBHelp = errors.New(`Use "gorepogen gendb -h" for help`)
)

type generateFromDatabase struct{}

func NewGenerateFromDatabase() GenerateCommand {
	return &generateFromDatabase{}
}

func (g generateFromDatabase) CreateCommand() *cli.Command {
	return &cli.Command{
		Name:    "gendb",
		Aliases: []string{"gd"},
		Usage:   "generate repositories from database",
		Flags:   g.buildFlags(),
		Action:  g.generate,
	}
}

func (g generateFromDatabase) generate(*cli.Context) error {

	if err := g.checkArgs(); err != nil {
		return err
	}

	// TODO: reminder - pass driver to generator
	drv, err := driver.Get(g.getDbInfo())
	if err != nil {
		return err
	}

	tables, err := drv.FindAllTables()
	fmt.Println(tables, err)

	return nil
}

func (generateFromDatabase) buildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "drvName",
			Aliases:     []string{"dr"},
			Usage:       "Driver name (mysql or postgres)",
			Destination: &drvName,
		},
		&cli.StringFlag{
			Name:        "username",
			Aliases:     []string{"us"},
			Usage:       "Database username",
			Destination: &username,
		},
		&cli.StringFlag{
			Name:        "password",
			Aliases:     []string{"pw"},
			Usage:       "Database password",
			Destination: &password,
		},
		&cli.StringFlag{
			Name:        "host",
			Aliases:     []string{"hs"},
			Usage:       "Database host",
			Destination: &host,
		},
		&cli.StringFlag{
			Name:        "port",
			Aliases:     []string{"pr"},
			Usage:       "Database port",
			Destination: &port,
		},
		&cli.StringFlag{
			Name:        "dbName",
			Aliases:     []string{"db"},
			Usage:       "Database name",
			Destination: &dbName,
		},
		&cli.StringFlag{
			Name:        "schema",
			Aliases:     []string{"sc"},
			Usage:       "Schema name",
			Destination: &schema,
		},
	}
}

func (g generateFromDatabase) checkArgs() error {
	if drvName == "" || username == "" || password == "" || host == "" || port == "" || dbName == "" || schema == "" {
		return ErrGenDBHelp
	}
	return nil
}

func (g generateFromDatabase) getDbInfo() *driver.DatabaseInfo {
	return &driver.DatabaseInfo{
		SchemaName:   schema,
		DriverName:   drvName,
		Username:     username,
		Password:     password,
		Host:         host,
		Port:         port,
		DatabaseName: dbName,
	}
}
