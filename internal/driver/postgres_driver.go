package driver

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type postgresDriver struct {
	*DatabaseInfo
}

func NewPostgresDriver(info *DatabaseInfo) AbstractDriver {
	return &postgresDriver{
		DatabaseInfo: info,
	}
}

func (d postgresDriver) openConnection() *gorm.DB {
	return nil
}

func (d postgresDriver) SelectAllTableNames() (tableNames []string, err error) {
	return tableNames, d.openConnection().
		Table("information_schema.tables ").
		Select("table_name").
		Where("table_schema = ?", "public").
		Scan(&tableNames).
		Error
}

func (d postgresDriver) SelectColumnsByTableName(tableName string) {

}

func (d postgresDriver) selectColumnsByTableName(tableName string) string {
	return fmt.Sprintf(
		`select table_name, column_name, data_type 
                  from information_schema.columns 
				  where table_schema = 'public' and table_name = '%s'`, tableName,
	)
}
