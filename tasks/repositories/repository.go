package repositories

import (
	"database/sql"
	"errors"
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
	findAll() ([]*Task, error)
	findById(id int64) (*Task, error)
	create(task *Task) error
	update(task *Task) error
	delete(id int64) error
}

func repo() repository {
	switch os.Getenv("DATABASE_DRIVER") {
	case "memory":
		return &memoryRepository{}
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

		return &dbRepository{db}
	default:
		log.Fatalf("Database driver %s not supported", os.Getenv("DATABASE_DRIVER"))
		return nil
	}
}

func FindAllTasks() ([]*Task, error) {
	return repo().findAll()
}

func CreateTask(t *Task) error {
	err := []error{}

	if t.Title == "" {
		err = append(err, errors.New("Title is required"))
	}

	if e := repo().create(t); e != nil {
		err = append(err, errors.New("Could not save task: "+e.Error()))
	}

	if len(err) > 0 {
		return errors.Join(err...)
	}

	return nil
}

func UpdateTask(t *Task) error {
	var (
		err  []error
		repo repository = repo()
	)

	if t.Title == "" {
		err = append(err, errors.New("Title is required"))
	}

	if _, e := repo.findById(t.Id); e != nil {
		err = append(err, errors.New("Could not find task"))
		return errors.Join(err...)
	}

	if e := repo.update(t); e != nil {
		err = append(err, errors.New("Could not save task: "+e.Error()))
	}

	if len(err) > 0 {
		return errors.Join(err...)
	}

	return nil
}

func FindTaskById(id int64) (*Task, error) {
	return repo().findById(id)
}

func DeleteTask(id int64) error {
	return repo().delete(id)
}
