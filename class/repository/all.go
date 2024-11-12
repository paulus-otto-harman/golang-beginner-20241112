package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type AllRepository struct {
	BookRepo     *Book
	CustomerRepo *Customer
	OrderRepo    *Order
}

func NewAllRepository(db *sql.DB, log *zap.Logger) AllRepository {
	return AllRepository{
		BookRepo:     InitBookRepo(db, log),
		CustomerRepo: InitCustomerRepo(db, log),
		OrderRepo:    InitOrderRepo(db, log),
	}
}
