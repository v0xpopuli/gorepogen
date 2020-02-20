package driver

import (
	"fmt"

	"github.com/jinzhu/gorm"
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
	// TODO: implement and test MySQL part
	panic("implement me")
}
