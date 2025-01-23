package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/fahreyad/golangcrud/internal/types"
	"github.com/fahreyad/golangcrud/internal/utils/response"
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

		response.WriteJSON(w, http.StatusCreated, map[string]string{"success": "ok"})
	}
}
