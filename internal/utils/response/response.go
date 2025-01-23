package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

const (
	StatusSuccess = "success"
	StatusError   = "error"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func ResponseError(err error) Response {
	return Response{
		Status:  StatusError,
		Message: err.Error(),
	}
}

func ValidationErrors(err validator.ValidationErrors) Response {
	var errMsg []string

	for _, e := range err {

		switch e.ActualTag() {
		case "required":
			errMsg = append(errMsg, fmt.Sprintf("%s is required", e.Field()))

		case "email":
			errMsg = append(errMsg, fmt.Sprintf("%s email is not valid", e.Field()))

		default:
			errMsg = append(errMsg, fmt.Sprintf("%s email is not valid", e.Field()))
		}
	}
	msg := strings.Join(errMsg, ", ")
	return Response{
		Status:  StatusError,
		Message: msg,
	}
}
