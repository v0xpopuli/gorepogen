package connector

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func postgresConnector(info *DatabaseInfo) (*gorm.DB, error) {
	return gorm.Open(info.DriverName, fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		info.Username,
		info.Password,
		info.Host,
		info.Port,
		info.DatabaseName,
	))
}
