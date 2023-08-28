package repositories

import (
	"errors"
	"math/rand"
)

type memoryRepository struct{}

var tasksList []*Task

func (r memoryRepository) FindAll() ([]*Task, error) {
	return tasksList, nil
}

func (r memoryRepository) FindById(id int64) (*Task, error) {
	for _, t := range tasksList {
		if t.Id == id {
			return t, nil
		}
	}

	return nil, errors.New("Task not found")
}

func (r memoryRepository) Create(t *Task) error {
	err := []error{}

	if t.Title == "" {
		err = append(err, errors.New("Title is required"))
	}

	t.Id = rand.Int63()
	tasksList = append(tasksList, t)

	if len(err) > 0 {
		return errors.Join(err...)
	}

	return nil
}

func (r memoryRepository) Update(t *Task) error {
	var err []error

	if t.Title == "" {
		err = append(err, errors.New("Title is required"))
	}

	for i, task := range tasksList {
		if task.Id == t.Id {
			tasksList[i] = t
			return nil
		}

		err = append(err, errors.New("Could not find task"))
	}

	if len(err) > 0 {
		return errors.Join(err...)
	}

	return nil
}

func (r memoryRepository) Delete(id int64) error {
	var err []error

	for i, t := range tasksList {
		if t.Id == id {
			tasksList = append(tasksList[:i], tasksList[i+1:]...)
			return nil
		}
		err = append(err, errors.New("Could not find task"))
	}

	if len(err) > 0 {
		return errors.Join(err...)
	}

	return nil
}
