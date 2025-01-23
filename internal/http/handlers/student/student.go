package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/fahreyad/golangcrud/internal/storage"
	"github.com/fahreyad/golangcrud/internal/types"
	"github.com/fahreyad/golangcrud/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
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
		err = validator.New().Struct(student)
		if err != nil {
			validationErr := err.(validator.ValidationErrors)
			response.WriteJSON(w, http.StatusBadRequest, response.ValidationErrors(validationErr))
			return
		}
		//db operation
		id, err := storage.CreateStudent(student.Name, student.Email, student.Age)
		if err != nil {
			response.WriteJSON(w, http.StatusInternalServerError, response.ResponseError(err))
			return
		}
		slog.Info("new student created", slog.Int64("id", id))

		response.WriteJSON(w, http.StatusCreated, map[string]int64{"id": id})
	}
}

func List(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("student list")
		id := r.PathValue("id")
		fmt.Println(id)
		if id == "" {
			response.WriteJSON(w, http.StatusBadRequest, response.ResponseError(errors.New("id is required")))
			return
		}
		intID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.ResponseError(err))
			return
		}
		studentID, err := storage.GetStudents(intID)
		if err != nil {
			response.WriteJSON(w, http.StatusInternalServerError, response.ResponseError(err))
			return
		}
		response.WriteJSON(w, http.StatusOK, studentID)
	}
}
