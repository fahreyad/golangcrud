package sqlite

import (
	"database/sql"

	"github.com/fahreyad/golangcrud/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cnf *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cnf.StoragePath)

	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
	 	name VARCHAR(100),
		email VArCHAR(100), 
		age INTEGER
	 )`)
	if err != nil {
		return nil, err
	}
	return &Sqlite{Db: db}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	result, err := s.Db.Prepare("INSERT INTO students(name, email, age) VALUES(?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer result.Close()
	res, err := result.Exec(name, email, age)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
