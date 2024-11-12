package repository

import (
	"20241112/model"
	"database/sql"
	"go.uber.org/zap"
)

type PaymentMethod struct {
	Db  *sql.DB
	Log *zap.Logger
}

func InitPaymentMethodRepo(db *sql.DB, log *zap.Logger) *PaymentMethod {
	return &PaymentMethod{Db: db, Log: log}
}

func (paymentMethod *PaymentMethod) All() ([]PaymentMethod, error) {
	return nil, nil
}

func (repo *PaymentMethod) Create(paymentMethod *model.PaymentMethod) error {
	query := "INSERT INTO payment_methods (name,photo,is_active) VALUES ($1, $2, $3) RETURNING id"
	if err := repo.Db.QueryRow(query, &paymentMethod.Name, &paymentMethod.Photo, &paymentMethod.IsActive).Scan(&paymentMethod.Id); err != nil {
		repo.Log.Error("Error creating payment method", zap.Error(err))
		return err
	}
	return nil
}
