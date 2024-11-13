package repository

import (
	"20241112/model"
	"database/sql"
)

type BookReview struct {
	Db *sql.DB
}

func InitBookReviewRepo(db *sql.DB) *BookReview {
	return &BookReview{Db: db}
}

func (repo *BookReview) Create(review *model.BookReview) error {
	query := "INSERT INTO reviews (order_id, book_id, customer_name,rating,review_text) VALUES($1, $2, $3, $4, $5) RETURNING id"
	if err := repo.Db.QueryRow(query, &review.OrderId, &review.BookId, &review.CustomerName, &review.Rating, &review.ReviewText).Scan(&review.Id); err != nil {
		return err
	}
	return nil
}
