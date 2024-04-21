package db_tool

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type DBEngine = string
type DBName = string

var (
	MYSQL    DBEngine = "mysql"
	POSTGRES DBEngine = "postgres"
)

var (
	SlaveDB  DBName = "SLAVE_DB"
	MasterDB DBName = "Master_DB"
)

type ConnectOption struct {
	Engine   DBEngine
	Database DBName
}

// 특정 DB에 종속되지 않게 connector interface로 구현
type connectorInterface interface {
	Connector() (*sqlx.DB, error)
}

func DBConnect(option ConnectOption) *sqlx.DB {
	var conn connectorInterface
	// engine 에 따라 처리
	switch option.Engine {
	case MYSQL:
		conn = &mysqlDB{
			name: option.Database,
		}
	case POSTGRES:
		conn = &postgresDB{
			name: option.Database,
		}
	default:
		log.Panicln("not support db engine")
	}
	db, err := conn.Connector()
	if err != nil {
		log.Panicln(err)
	}
	return db
}
