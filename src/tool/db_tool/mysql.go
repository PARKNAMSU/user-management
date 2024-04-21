package db_tool

import (
	"errors"
	"time"

	"github.com/PARKNAMSU/user-management/configs/db_configs"
	"github.com/jmoiron/sqlx"
)

var (
	mysqlSlave  *sqlx.DB
	mysqlMaster *sqlx.DB
)

type mysqlDB struct {
	name DBName
}

func (d *mysqlDB) Connector() (*sqlx.DB, error) {
	switch d.name {
	case SlaveDB:
		{
			if mysqlSlave != nil {
				return mysqlSlave, nil
			}
			db, err := mysqlSlaveConnector()
			mysqlSlave = db
			return db, err
		}
	case MasterDB:
		{
			if mysqlMaster != nil {
				return mysqlMaster, nil
			}
			db, err := mysqlMasterConnector()
			mysqlMaster = db
			return db, err
		}
	}
	return nil, errors.New("not supported db")
}

func mysqlSlaveConnector() (*sqlx.DB, error) {
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

func mysqlMasterConnector() (*sqlx.DB, error) {
	option := db_configs.MysqlMasterOption()
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
