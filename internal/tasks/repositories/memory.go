package repositories

import (
	"errors"
	"math/rand"
)

type memoryRepository struct{}

var tasksList []*Task

func (r *memoryRepository) findAll() ([]*Task, error) {
	return tasksList, nil
}

func (r *memoryRepository) findById(id int64) (*Task, error) {
	for _, t := range tasksList {
		if t.Id == id {
			return t, nil
		}
	}

	return nil, errors.New("Task not found")
}

func (r *memoryRepository) create(t *Task) error {
	t.Id = rand.Int63()
	tasksList = append(tasksList, t)

	return nil
}

func (r *memoryRepository) update(t *Task) error {
	for i, task := range tasksList {
		if task.Id == t.Id {
			tasksList[i] = t
			return nil
		}
	}

	return errors.New("Couldn't update Task: not found")
}

func (r *memoryRepository) delete(id int64) error {
	for i, t := range tasksList {
		if t.Id == id {
			tasksList = append(tasksList[:i], tasksList[i+1:]...)
			return nil
		}
	}
	return nil
}
