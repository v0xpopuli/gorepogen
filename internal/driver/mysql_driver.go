package driver

type mysqlDriver struct {
	*DatabaseInfo
}

func NewMysqlDriver(info *DatabaseInfo) AbstractDriver {
	return &mysqlDriver{
		DatabaseInfo: info,
	}
}
