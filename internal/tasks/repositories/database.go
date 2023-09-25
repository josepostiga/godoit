package repositories

import (
	"database/sql"
	"errors"
	"time"
)

type dbRepository struct {
	db *sql.DB
}

func (r dbRepository) FindAll() ([]*Task, error) {
	rows, err := r.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasksList := make([]*Task, 0)
	for rows.Next() {
		t := new(Task)
		err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Created_at, &t.Updated_at, &t.Completed_at)
		if err != nil {
			return nil, err
		}
		tasksList = append(tasksList, t)
	}

	return tasksList, nil
}

func (r dbRepository) FindById(id int) (*Task, error) {
	t := new(Task)

	err := r.db.QueryRow("SELECT * FROM tasks WHERE id = $1", id).
		Scan(&t.Id, &t.Title, &t.Description, &t.Created_at, &t.Updated_at, &t.Completed_at)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r dbRepository) Create(t *Task) error {
	if t.Title == "" {
		return errors.New("Title is required")
	}

	if e := r.db.QueryRow("INSERT INTO tasks (title, description) VALUES ($1, $2) returning id", t.Title, t.Description).Scan(&t.Id); e != nil {
		return errors.New("Could not save task: " + e.Error())
	}

	return nil
}

func (r dbRepository) Update(t *Task) error {
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

func (r dbRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}

func (r dbRepository) ToggleStatus(id int) error {
	t, err := r.FindById(id)
	if err != nil {
		return err
	}

	if t.Completed_at.Valid {
		t.Completed_at.Scan(nil)
	} else {
		t.Completed_at.Scan(time.Now())
	}

	if _, e := r.db.Exec("UPDATE tasks set completed_at=$1 where id = $2", t.Completed_at, t.Id); e != nil {
		return errors.New("Could not save task")
	}

	return nil
}
