package handler

import (
	"20241112/lib"
	"20241112/model"
	"20241112/service"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module/web"
	"go.uber.org/zap"
	"log"
	"net/http"
	"reflect"
)

type PaymentMethodHandler struct {
	PaymentMethodService service.PaymentMethodService
	Log                  *zap.Logger
}

func InitPaymentMethodHandler(paymentMethodService service.PaymentMethodService, log *zap.Logger) PaymentMethodHandler {
	return PaymentMethodHandler{PaymentMethodService: paymentMethodService, Log: log}
}

func (handler PaymentMethodHandler) Create(w http.ResponseWriter, r *http.Request) {
	file := gola.StoreUploadedFile("photo", true, r)
	if file.Error != nil {
		handler.Log.Error(fmt.Sprintf("%s_%s", reflect.TypeOf(handler).Name(), "Create"), zap.Error(file.Error.Default))
		lib.JsonResponse(w).Fail(http.StatusUnprocessableEntity, "Photo processing failed")
		return
	}

	name := r.FormValue("name")
	//if err := validate(name).Required().Min(5).Max(15); err != nil {
	//	handler.Log.Error(fmt.Sprintf("%s_%s", reflect.TypeOf(handler).Name(), "Create"), zap.Error(err))
	//}
	isActive := r.FormValue("is_active") == "1"
	//validate(isActive).Required()

	paymentMethod := &model.PaymentMethod{
		Name:     name,
		IsActive: isActive,
		Photo:    file.Uploaded.FullUrl,
	}

	if err := handler.PaymentMethodService.Create(paymentMethod); err != nil {
		log.Println(err)
		lib.JsonResponse(w).Fail(http.StatusInternalServerError, "Unable to create item")
		return
	}
	lib.JsonResponse(w).Success(0, "Metode Pembayaran berhasil ditambahkan", paymentMethod)
}

func (handler PaymentMethodHandler) All(w http.ResponseWriter, r *http.Request) {

}

func (handler PaymentMethodHandler) Get(w http.ResponseWriter, r *http.Request) {

}

func (handler PaymentMethodHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (handler PaymentMethodHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
