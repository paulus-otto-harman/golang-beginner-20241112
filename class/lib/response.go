package lib

import (
	"20241112/model"
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter) Response {
	return Response{W: w}
}

type Response struct {
	W http.ResponseWriter
}

func (response Response) Fail(status int, message string) {
	r := model.Response{
		Status:  status,
		Message: message,
	}

	json.NewEncoder(response.W).Encode(r)
}

func (response Response) Success(status int, message string, data interface{}) {
	r := model.Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	json.NewEncoder(response.W).Encode(r)
}
