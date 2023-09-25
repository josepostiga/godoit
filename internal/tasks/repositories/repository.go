package repositories

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
	"net/url"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Task struct {
	Id           int         `json:"id"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	Created_at   time.Time   `json:"created_at"`
	Updated_at   time.Time   `json:"updated_at"`
	Completed_at pq.NullTime `json:"completed_at"`
}

type repository interface {
	FindAll() ([]*Task, error)
	FindById(id int) (*Task, error)
	Create(task *Task) error
	Update(task *Task) error
	Delete(id int) error
	ToggleStatus(id int) error
}

var repo repository

func NewRepository(driver string) repository {
	if repo != nil {
		return repo
	}

	switch driver {
	case "memory":
		repo = memoryRepository{}
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

		repo = dbRepository{db}
	default:
		log.Fatalf("Database driver %s not supported", os.Getenv("DATABASE_DRIVER"))
		return nil
	}

	return repo
}
