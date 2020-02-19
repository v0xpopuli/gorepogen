package driver

type postgresDriver struct {
	*DatabaseInfo
}

func NewPostgresDriver(info *DatabaseInfo) AbstractDriver {
	return &postgresDriver{
		DatabaseInfo: info,
	}
}
