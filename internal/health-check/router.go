package health_check

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	r.Route("/health-check", func(r chi.Router) {
		r.Get("/", show)
	})
}
