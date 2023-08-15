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

type repository interface {
	findAll() ([]*Task, error)
	findById(id int) (*Task, error)
	save(task *Task) error
	delete(id int) error
}

func repo() repository {
	switch os.Getenv("DATABASE_DRIVER") {
	case "memory":
		return &memoryRepository{}
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

	if repo().save(t) != nil {
		err = append(err, errors.New("Could not save task"))
	}

	if len(err) > 0 {
		return errors.Join(err...)
	}

	return nil
}

func UpdateTask(id int, t *Task) error {
	var (
		err  []error
		repo repository = repo()
	)

	if t.Title == "" {
		err = append(err, errors.New("Title is required"))
	}

	task, e := repo.findById(id)
	if e != nil {
		err = append(err, errors.New("Could not find task"))
	}

	task.Title = t.Title
	task.Description = t.Description

	if repo.save(task) != nil {
		err = append(err, errors.New("Could not save task"))
	}

	if len(err) > 0 {
		return errors.Join(err...)
	}

	return nil
}

func FindTaskById(id int) (*Task, error) {
	return repo().findById(id)
}

func DeleteTask(id int) error {
	return repo().delete(id)
}
