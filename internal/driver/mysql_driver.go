package driver

import (
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

func (d mysqlDriver) openConnection() *gorm.DB {
	panic("implement me")
}

func (d mysqlDriver) SelectAllTableNames() ([]string, error) {
	panic("implement me")
}

func (d mysqlDriver) SelectColumnsByTableName(string) {
	panic("implement me")
}
