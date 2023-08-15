package repositories

import (
	"errors"
	"log"
	"os"
)

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Repository interface {
	FindAll() ([]*Task, error)
	FindById(id int) (*Task, error)
	Save(task *Task) error
	Delete(id int) error
}

func NewRepository() Repository {
	switch os.Getenv("DATABASE_DRIVER") {
	case "memory":
		return &InMemoryRepository{}
	default:
		log.Fatalf("Database driver %s not supported", os.Getenv("DATABASE_DRIVER"))
		return nil
	}
}

func FindAllTasks(repo Repository) ([]*Task, error) {
	return repo.FindAll()
}

func CreateTask(t *Task, repo Repository) error {
	err := []error{}

	if t.Title == "" {
		err = append(err, errors.New("Title is required"))
	}

	if repo.Save(t) != nil {
		err = append(err, errors.New("Could not save task"))
	}

	if len(err) > 0 {
		return errors.Join(err...)
	}

	return nil
}

func UpdateTask(id int, t *Task, repo Repository) error {
	var err []error

	if t.Title == "" {
		err = append(err, errors.New("Title is required"))
	}

	task, e := repo.FindById(id)
	if e != nil {
		err = append(err, errors.New("Could not find task"))
	}

	task.Title = t.Title
	task.Description = t.Description

	if repo.Save(task) != nil {
		err = append(err, errors.New("Could not save task"))
	}

	if len(err) > 0 {
		return errors.Join(err...)
	}

	return nil
}

func FindTaskById(id int, repo Repository) (*Task, error) {
	return repo.FindById(id)
}

func DeleteTask(id int, repo Repository) error {
	return repo.Delete(id)
}
