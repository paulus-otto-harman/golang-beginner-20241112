package model

import "time"

type Order struct {
	Id              string           `json:"orderId,omitempty"`
	CustomerId      int              `json:"customerId,omitempty"`
	CustomerName    string           `json:"customerName"`
	CustomerPhone   string           `json:"customerPhone,omitempty"`
	ShippingAddress *ShippingAddress `json:"shippingAddress,omitempty"`
	OrderItems      []OrderItem      `json:"orderItems"`
	PaymentMethod   string           `json:"paymentMethod,omitempty"`
	TotalAmount     int              `json:"totalAmount"`
	FinalAmount     int              `json:"finalAmount,omitempty"`
	OrderDate       time.Time        `json:"orderDate"`
	Status          string           `json:"status"`
}
