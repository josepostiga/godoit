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

type TasksRepository interface {
	FindAll() ([]*Task, error)
	FindById(id int) (*Task, error)
	Create(task *Task) error
	Update(task *Task) error
	Delete(id int) error
	ToggleStatus(id int) error
}

func New() TasksRepository {
	dsl, err := url.Parse(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Could not parse DATABASE_URL: %s", err.Error())
		return nil
	}

	db, err := sql.Open("postgres", dsl.String())
	if err != nil || db.Ping() != nil {
		log.Fatalf("Could not connect to database.")
		return nil
	}

	return &dbRepository{db}
}
