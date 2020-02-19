package driver

type AbstractDriver interface {
}

type DatabaseInfo struct {
	username string
	password string
	host     string
	port     string
	dbName   string
}
