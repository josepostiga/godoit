package main

import (
	"github.com/go-chi/chi/v5"
	health_check "godoit/health-check"
	"godoit/tasks"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	health_check.RegisterRoutes(r)
	tasks.RegisterRoutes(r)

	log.Fatalf("Couldn't start server: %s", http.ListenAndServe(":3000", r))
}
