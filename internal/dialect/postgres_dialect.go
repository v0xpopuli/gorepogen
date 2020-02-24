package dialect

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type postgresDriver struct {
	*DatabaseInfo
}

func NewPostgresDriver(info *DatabaseInfo) AbstractDriver {
	return &postgresDriver{
		DatabaseInfo: info,
	}
}

func (d postgresDriver) openConnection() (*gorm.DB, error) {
	return gorm.Open(d.DriverName, fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		d.Host,
		d.Port,
		d.Username,
		d.DatabaseName,
		d.Password,
	))
}

func (d postgresDriver) FindAllTables() (map[string][]Field, error) {
	conn, err := d.openConnection()
	if err != nil {
		return nil, err
	}

	var (
		tName, cName, dType string
		tables              = make(map[string][]Field, 0)
	)
	rows, err := conn.
		Table("information_schema.columns").
		Select("table_name, column_name, data_type").
		Where("table_schema = ?", d.SchemaName).
		Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&tName, &cName, &dType)
		tables[tName] = append(tables[tName], Field{
			name:  cName,
			vType: d.mapDBTypeToVarType(dType),
		})
		if err != nil {
			return nil, err
		}
	}

	return tables, nil
}

func (d postgresDriver) mapDBTypeToVarType(string) string {
	panic("implement me")
}
