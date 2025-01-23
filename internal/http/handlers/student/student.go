package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/fahreyad/golangcrud/internal/types"
	"github.com/fahreyad/golangcrud/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("new student creating")
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.ResponseError(err))
			return
		}

		if err != nil {
			response.WriteJSON(w, http.StatusInternalServerError, response.ResponseError(err))
			return
		}

		//request validation
		err1 := validator.New().Struct(student)
		if err1 != nil {
			validationErr := err1.(validator.ValidationErrors)
			response.WriteJSON(w, http.StatusBadRequest, response.ValidationErrors(validationErr))
			return
		}

		response.WriteJSON(w, http.StatusCreated, map[string]string{"success": "ok"})
	}
}
