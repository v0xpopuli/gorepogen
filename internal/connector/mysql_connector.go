package connector

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func mysqlConnector(info *DatabaseInfo) (*gorm.DB, error) {
	return gorm.Open(info.DriverName, fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?multiStatements=true",
		info.Username,
		info.Password,
		info.Host,
		info.Port,
		info.DatabaseName,
	))
}
