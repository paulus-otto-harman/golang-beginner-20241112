package repository

import (
	"database/sql"
	"go.uber.org/zap"
)

type Customer struct {
	Db  *sql.DB
	Log *zap.Logger
}

func InitCustomerRepo(db *sql.DB, log *zap.Logger) *Customer {
	return &Customer{Db: db, Log: log}
}
