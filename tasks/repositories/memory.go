package repositories

import (
	"errors"
	"math/rand"
)

type InMemoryRepository struct{}

var tasksList []*Task

func (r *InMemoryRepository) FindAll() ([]*Task, error) {
	return tasksList, nil
}

func (r *InMemoryRepository) FindById(id int) (*Task, error) {
	for _, t := range tasksList {
		if t.Id == id {
			return t, nil
		}
	}

	return nil, errors.New("Task not found")
}

func (r *InMemoryRepository) Save(t *Task) error {
	if t.Id == 0 {
		t.Id = rand.Int()
	}

	tasksList = append(tasksList, t)

	return nil
}

func (r *InMemoryRepository) Delete(id int) error {
	for i, t := range tasksList {
		if t.Id == id {
			tasksList = append(tasksList[:i], tasksList[i+1:]...)
			return nil
		}
	}
	return nil
}
