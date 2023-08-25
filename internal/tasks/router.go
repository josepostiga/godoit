package tasks

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", index)
		r.Post("/", store)
		r.Patch("/{id}", update)
		r.Get("/{id}", show)
		r.Delete("/{id}", delete)
	})
}
