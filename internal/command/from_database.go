package command

import (
	"github.com/urfave/cli/v2"
	"github.com/v0xpopuli/gorepogen/internal/connector"
	"github.com/v0xpopuli/gorepogen/internal/connector/mapper"
)

var (
	drvName, username, password, host, port, dbName, schema string
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
		Flags:   g.getFlags(),
		Action:  g.generate,
	}
}

func (g generateFromDatabase) generate(ctx *cli.Context) error {

	conn, err := connector.NewConnector(g.getDbInfo())
	if err != nil {
		return err
	}

	tables, err := conn.FindAllTables()
	if err != nil {
		return err
	}

	_, err = mapper.MapTablesToEntityDefinition(tables)
	if err != nil {
		return err
	}

	return nil
}

func (generateFromDatabase) getFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "drvname",
			Aliases:     []string{"r"},
			Usage:       "Driver name (mysql or postgres)",
			Destination: &drvName,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "username",
			Aliases:     []string{"u"},
			Usage:       "Database username",
			Destination: &username,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "password",
			Aliases:     []string{"w"},
			Usage:       "Database password",
			Destination: &password,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "host",
			Aliases:     []string{"h"},
			Usage:       "Database host",
			Destination: &host,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "port",
			Aliases:     []string{"p"},
			Usage:       "Database port",
			Destination: &port,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "dbname",
			Aliases:     []string{"d"},
			Usage:       "Database name",
			Destination: &dbName,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "schema",
			Aliases:     []string{"s"},
			Usage:       "Schema name",
			Destination: &schema,
			Required:    true,
		},
	}
}

func (g generateFromDatabase) getDbInfo() *connector.DatabaseInfo {
	return &connector.DatabaseInfo{
		SchemaName:   schema,
		DriverName:   drvName,
		Username:     username,
		Password:     password,
		Host:         host,
		Port:         port,
		DatabaseName: dbName,
	}
}
