package repository

import (
	"database/sql"
	"go.uber.org/zap"
)

type Book struct {
	Db  *sql.DB
	Log *zap.Logger
}

func InitBookRepo(db *sql.DB, log *zap.Logger) *Book {
	return &Book{Db: db, Log: log}
}
