package model

import "time"

type BookReview struct {
	Id           int
	OrderId      string    `json:"orderId"`
	BookId       string    `json:"bookId"`
	CustomerName string    `json:"customerName"`
	Rating       float64   `json:"rating"`
	ReviewText   string    `json:"reviewText"`
	ReviewDate   time.Time `json:"reviewDate"`
}
