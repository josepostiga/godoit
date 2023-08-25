package repositories

import (
	"database/sql"
)

type dbRepository struct {
	db *sql.DB
}

func (r *dbRepository) findAll() ([]*Task, error) {
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

func (r *dbRepository) findById(id int64) (*Task, error) {
	t := new(Task)

	err := r.db.QueryRow("SELECT id, title, description FROM tasks WHERE id = $1", id).Scan(&t.Id, &t.Title, &t.Description)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *dbRepository) create(t *Task) error {
	return r.db.QueryRow("INSERT INTO tasks (title, description) VALUES ($1, $2) returning id", t.Title, t.Description).Scan(&t.Id)
}

func (r *dbRepository) update(t *Task) error {
	row := r.db.QueryRow("UPDATE tasks set title=$1, description=$2 where id = $3", t.Title, t.Description, t.Id)
	return row.Err()
}

func (r *dbRepository) delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}
