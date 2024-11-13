package service

import (
	"20241112/model"
	"20241112/repository"
)

type BookReviewService struct {
	BookReviewRepo repository.BookReview
}

func InitBookReviewService(repo repository.BookReview) *BookReviewService {
	return &BookReviewService{BookReviewRepo: repo}
}

func (bookReviewService BookReviewService) Create(review *model.BookReview) error {
	return bookReviewService.BookReviewRepo.Create(review)
}
