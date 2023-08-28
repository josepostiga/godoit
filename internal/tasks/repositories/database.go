package repositories

import (
	"database/sql"
	"errors"
)

type dbRepository struct {
	db *sql.DB
}

func (r *dbRepository) FindAll() ([]*Task, error) {
	rows, err := r.db.Query("SELECT id, title, description FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasksList := make([]*Task, 0)
	for rows.Next() {
		t := new(Task)
		err := rows.Scan(&t.Id, &t.Title, &t.Description)
		if err != nil {
			return nil, err
		}
		tasksList = append(tasksList, t)
	}

	return tasksList, nil
}

func (r *dbRepository) FindById(id int64) (*Task, error) {
	t := new(Task)

	err := r.db.QueryRow("SELECT id, title, description FROM tasks WHERE id = $1", id).Scan(&t.Id, &t.Title, &t.Description)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *dbRepository) Create(t *Task) error {
	if t.Title == "" {
		return errors.New("Title is required")
	}

	if e := r.db.QueryRow("INSERT INTO tasks (title, description) VALUES ($1, $2) returning id", t.Title, t.Description).Scan(&t.Id); e != nil {
		return errors.New("Could not save task: " + e.Error())
	}

	return nil
}

func (r *dbRepository) Update(t *Task) error {
	if t.Title == "" {
		return errors.New("Title is required")
	}

	if _, e := r.FindById(t.Id); e != nil {
		return errors.New("Could not find task")
	}

	if row := r.db.QueryRow("UPDATE tasks set title=$1, description=$2 where id = $3", t.Title, t.Description, t.Id); row.Err() != nil {
		return errors.New("Could not save task")
	}

	return nil
}

func (r *dbRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}
