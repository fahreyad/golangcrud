package storage

import "github.com/fahreyad/golangcrud/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudents(id int64) (types.Student, error)
}
