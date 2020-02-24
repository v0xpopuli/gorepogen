package dialect

import (
	"fmt"
	"strings"

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

func (d mysqlDriver) FindAllTables() (map[string][]Field, error) {
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
		Select("table_name as table_name, column_name as column_name, data_type as data_type").
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

var (
	boolTypes    = []string{"tinyint"}
	decimalTypes = []string{"decimal", "float", "double"}
	integerTypes = []string{"smallint", "mediumint", "int", "bigint", "bit"}
	timeTypes    = []string{"date", "time", "datetime", "timestamp", "year"}
	stringTypes  = []string{"char", "varchar", "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob", "tinytext", "text", "mediumtext", "longtext"}
)

func (d mysqlDriver) mapDBTypeToVarType(dType string) string {
	t := strings.ToLower(dType)
	switch {
	case containsType(boolTypes, t):
		return "bool"
	case containsType(decimalTypes, t):
		return "float32"
	case containsType(integerTypes, t):
		return "int32"
	case containsType(stringTypes, t):
		return "string"
	case containsType(timeTypes, t):
		return "time.Time"
	default:
		return "interface{}"
	}
}
