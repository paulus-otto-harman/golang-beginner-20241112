package handler

import (
	"20241112/model"
	"20241112/service"
	"encoding/json"
	"net/http"
)

type BookReviewHandler struct {
	BookReviewService service.BookReviewService
}

func InitBookReviewHandler(bookReviewService service.BookReviewService) BookReviewHandler {
	return BookReviewHandler{BookReviewService: bookReviewService}
}

func (handler BookReviewHandler) Create(w http.ResponseWriter, r *http.Request) {
	review := model.BookReview{}
	json.NewDecoder(r.Body).Decode(&review)

	if err := handler.BookReviewService.Create(&review); err != nil {
		json.NewEncoder(w).Encode(model.Response{
			Status:  http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Response{
		Status:  http.StatusOK,
		Message: "Order placed successfully",
		Data:    order,
	})
}
