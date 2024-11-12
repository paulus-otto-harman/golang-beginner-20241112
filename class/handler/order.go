package handler

import (
	"20241112/model"
	"20241112/service"
	"encoding/json"
	"net/http"
)

type OrderHandler struct {
	OrderService service.OrderService
}

func InitOrderHandler(orderService service.OrderService) OrderHandler {
	return OrderHandler{OrderService: orderService}
}

func (handler OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	order := model.Order{}
	json.NewDecoder(r.Body).Decode(&order)

	if err := handler.OrderService.Create(&order); err != nil {
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
