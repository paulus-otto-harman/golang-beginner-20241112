package model

type PaymentMethod struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Photo    string `json:"photo"`
	IsActive bool   `json:"is_active"`
}
