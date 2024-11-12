package service

import (
	"20241112/model"
	"20241112/repository"
)

type PaymentMethodService struct {
	PaymentMethodRepo repository.PaymentMethod
}

func InitPaymentMethodService(repo repository.PaymentMethod) *PaymentMethodService {
	return &PaymentMethodService{PaymentMethodRepo: repo}
}

func (service *PaymentMethodService) Create(paymentMethod *model.PaymentMethod) error {
	if err := service.PaymentMethodRepo.Create(paymentMethod); err != nil {
		return err
	}
	return nil
}
