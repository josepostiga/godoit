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

func (r *memoryRepository) findById(id int) (*Task, error) {
	for _, t := range tasksList {
		if t.Id == id {
			return t, nil
		}
	}

	return nil, errors.New("Task not found")
}

func (r *memoryRepository) save(t *Task) error {
	if t.Id == 0 {
		t.Id = rand.Int()
	}

	tasksList = append(tasksList, t)

	return nil
}

func (r *memoryRepository) delete(id int) error {
	for i, t := range tasksList {
		if t.Id == id {
			tasksList = append(tasksList[:i], tasksList[i+1:]...)
			return nil
		}
	}
	return nil
}
