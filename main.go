package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	health_check "godoit/health-check"
	"godoit/tasks"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Couldn't load .env file: %s", err)
	}

	r := chi.NewRouter()
	health_check.RegisterRoutes(r)
	tasks.RegisterRoutes(r)

	log.Fatalf("Couldn't start server: %s", http.ListenAndServe(":3000", r))
}
