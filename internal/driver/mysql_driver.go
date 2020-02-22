package driver

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type mysqlDriver struct {
	*DatabaseInfo
}

func NewMysqlDriver(info *DatabaseInfo) AbstractDriver {
	return &mysqlDriver{
		DatabaseInfo: info,
	}
}

func (d mysqlDriver) openConnection() (*gorm.DB, error) {
	return gorm.Open(d.DriverName, fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?multiStatements=true",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.DatabaseName,
	))
}

func (d mysqlDriver) FindAllTables() ([]TableInfo, error) {
	conn, err := d.openConnection()
	if err != nil {
		return nil, err
	}

	var tables []TableInfo
	err = conn.
		Table("information_schema.columns").
		Select("table_name as table_name, column_name as column_name, data_type as data_type").
		Where("table_schema = ?", d.SchemaName).
		Scan(&tables).
		Error
	if err != nil {
		return nil, err
	}

	return tables, nil
}
