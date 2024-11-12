package repository

import (
	"20241112/model"
	"database/sql"
	"errors"
	"go.uber.org/zap"
)

type Order struct {
	Db  *sql.DB
	Log *zap.Logger
}

func InitOrderRepo(db *sql.DB, log *zap.Logger) *Order {
	return &Order{Db: db, Log: log}
}

func firstOrCreateCustomer(customer *model.Customer, db *sql.DB) error {
	query := "SELECT id FROM customers WHERE phone=$1"

	err := errors.New("")
	if err = db.QueryRow(query, customer.Phone).Scan(&customer.Id); err != nil {
		query = "INSERT INTO customers (name, phone) VALUES($1, $2) RETURNING id"
		err = db.QueryRow(query, customer.Name, customer.Phone).Scan(&customer.Id)
	}

	if err != nil {
		return err
	}
	return nil
}

func (repo *Order) Create(order *model.Order) error {
	customer := model.Customer{Name: order.CustomerName, Phone: order.CustomerPhone}
	if err := firstOrCreateCustomer(&customer, repo.Db); err != nil {
		return err
	}

	order.ShippingAddress = nil

	//query = "INSERT INTO orders(customer_id, total_amount, final_amount, order_date) VALUES ($1, $2, $3, $4) RETURNING id"
	//var id int
	//repo.Db.QueryRow(query, order.TotalAmount, order.FinalAmount, order.OrderDate).Scan(&id)
	return nil
}
