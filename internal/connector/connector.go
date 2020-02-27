package connector

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var ErrUnsupportableDriver = errors.New("unsupportable driver provided")

const (
	mysqlDriverName    = "mysql"
	postgresDriverName = "postgres"
)

type Connector struct {
	connection *gorm.DB
	schemaName string
}

type DatabaseInfo struct {
	DriverName   string
	Username     string
	Password     string
	Host         string
	Port         string
	DatabaseName string
	SchemaName   string
}

type Table struct {
	TableName  string
	ColumnName string
	ColumnType string
}

func NewConnector(info *DatabaseInfo) (*Connector, error) {
	switch info.DriverName {
	case postgresDriverName:
		conn, err := postgresConnector(info)
		return &Connector{connection: conn, schemaName: info.SchemaName}, err
	case mysqlDriverName:
		conn, err := mysqlConnector(info)
		return &Connector{connection: conn, schemaName: info.SchemaName}, err
	default:
		return nil, ErrUnsupportableDriver
	}
}

func (c Connector) FindAllTables() ([]Table, error) {
	var tables []Table

	err := c.connection.
		Table("information_schema.columns").
		Select("table_name as table_name, column_name as column_name, data_type as column_type").
		Where("table_schema = ?", c.schemaName).
		Find(&tables).
		Error

	return tables, err
}
