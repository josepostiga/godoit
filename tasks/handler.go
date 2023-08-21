package tasks

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"godoit/tasks/repositories"
	"net/http"
	"strconv"
)

func respond(w http.ResponseWriter, resp []byte, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(resp)
}

func index(w http.ResponseWriter, r *http.Request) {
	tasks, _ := repositories.FindAllTasks()

	resp, _ := json.Marshal(struct {
		Tasks []*repositories.Task `json:"tasks" `
	}{
		Tasks: tasks,
	})

	respond(w, resp, http.StatusOK)
}

func store(w http.ResponseWriter, r *http.Request) {
	var t *repositories.Task

	json.NewDecoder(r.Body).Decode(&t)

	err := repositories.CreateTask(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	resp, _ := json.Marshal(&t)

	respond(w, resp, http.StatusCreated)
}

func update(w http.ResponseWriter, r *http.Request) {
	var t *repositories.Task
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	json.NewDecoder(r.Body).Decode(&t)
	t.Id = id

	err := repositories.UpdateTask(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	resp, _ := json.Marshal(&t)

	respond(w, resp, http.StatusOK)
}

func show(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	t, err := repositories.FindTaskById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	resp, _ := json.Marshal(&t)

	respond(w, resp, http.StatusOK)
}

func delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	err := repositories.DeleteTask(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	respond(w, nil, http.StatusNoContent)
}
