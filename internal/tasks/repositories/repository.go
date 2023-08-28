package repositories

import (
	"database/sql"
	"log"
	"net/url"
	"os"

	_ "github.com/lib/pq"
)

type Task struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type repository interface {
	FindAll() ([]*Task, error)
	FindById(id int64) (*Task, error)
	Create(task *Task) error
	Update(task *Task) error
	Delete(id int64) error
}

var initiatedRepo repository

func NewRepository(driver string) repository {
	if initiatedRepo != nil {
		return initiatedRepo
	}

	switch driver {
	case "memory":
		initiatedRepo = memoryRepository{}
	case "database":
		dsl, err := url.Parse(os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatalf("Could not parse DATABASE_URL: %v", err)
			return nil
		}

		db, err := sql.Open("postgres", dsl.String())
		if err != nil {
			log.Fatalf("Could not open database: %v", err)
			return nil
		}

		initiatedRepo = dbRepository{db}
	default:
		log.Fatalf("Database driver %s not supported", os.Getenv("DATABASE_DRIVER"))
		return nil
	}

	return initiatedRepo
}
