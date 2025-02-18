package svc

import (
	"coupon/conf"
	"database/sql"
)

type ServiceContext struct {
	Config conf.Config
	DB     *sql.DB
}

func NewServiceContext(c conf.Config) *ServiceContext {

	db, err := sql.Open("mysql", c.DSN)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
