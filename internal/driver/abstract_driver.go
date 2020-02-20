package driver

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var ErrUnsupportableDriver = errors.New("unsupportable driver provided")

const (
	mysqlDriverName    = "mysql"
	postgresDriverName = "postgres"
)

type AbstractDriver interface {
	openConnection() (*gorm.DB, error)
	FindAllTables() ([]TableInfo, error)
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

type TableInfo struct {
	TableName  string
	ColumnName string
	DataType   string
}

func Get(info *DatabaseInfo) (AbstractDriver, error) {
	switch info.DriverName {
	case postgresDriverName:
		return NewPostgresDriver(info), nil
	case mysqlDriverName:
		return NewMysqlDriver(info), nil
	default:
		return nil, ErrUnsupportableDriver
	}
}
