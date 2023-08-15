package tasks

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"godoit/tasks/repositories"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	tasks, _ := repositories.FindAllTasks(repositories.NewRepository())

	resp, _ := json.Marshal(struct {
		Tasks []*repositories.Task `json:"tasks" `
	}{
		Tasks: tasks,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func store(w http.ResponseWriter, r *http.Request) {
	var t *repositories.Task

	json.NewDecoder(r.Body).Decode(&t)

	err := repositories.CreateTask(t, repositories.NewRepository())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	resp, _ := json.Marshal(&t)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func update(w http.ResponseWriter, r *http.Request) {
	var t *repositories.Task
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	json.NewDecoder(r.Body).Decode(&t)

	err := repositories.UpdateTask(id, t, repositories.NewRepository())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	resp, _ := json.Marshal(&t)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func show(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	t, err := repositories.FindTaskById(id, repositories.NewRepository())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	resp, _ := json.Marshal(&t)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := repositories.DeleteTask(id, repositories.NewRepository())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
