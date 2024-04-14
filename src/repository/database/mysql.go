package database

import (
	"errors"
	"time"

	"github.com/PARKNAMSU/user-management/configs/db_configs"
	"github.com/jmoiron/sqlx"
)

type DBName = string

var (
	SlaveDB  DBName = "SLAVE_DB"
	MasterDB DBName = "Master_DB"
)

var (
	slaveDB  *sqlx.DB
	masterDB *sqlx.DB
)

func Connector(db DBName) (*sqlx.DB, error) {
	switch db {
	case SlaveDB:
		{
			if slaveDB != nil {
				return slaveDB, nil
			}
			return slaveDBConnector()
		}
	case MasterDB:
		{
			if masterDB != nil {
				return masterDB, nil
			}
			return masterDBConnector()
		}
	}
	return nil, errors.New("not supported db")
}

func slaveDBConnector() (*sqlx.DB, error) {
	option := db_configs.MysqlSlaveOption()
	db, dbErr := sqlx.Connect(option.Engine, option.User+":"+option.Password+"@tcp("+option.Host+")/"+option.Database+"?charset=utf8mb4&parseTime=True&maxAllowedPacket=0")
	if dbErr != nil {
		return nil, dbErr
	}
	db.SetConnMaxLifetime(time.Minute)
	db.SetConnMaxIdleTime(time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	return db, nil
}

func masterDBConnector() (*sqlx.DB, error) {
	option := db_configs.MysqlSlaveOption()
	db, dbErr := sqlx.Connect(option.Engine, option.User+":"+option.Password+"@tcp("+option.Host+")/"+option.Database+"?charset=utf8mb4&parseTime=True&maxAllowedPacket=0")
	if dbErr != nil {
		return nil, dbErr
	}
	db.SetConnMaxLifetime(time.Minute)
	db.SetConnMaxIdleTime(time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	return db, nil
}
