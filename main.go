package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	health_check "godoit/health-check"
	"godoit/tasks"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Couldn't load .env file: %s", err)
	}

	r := chi.NewRouter()
	health_check.RegisterRoutes(r)
	tasks.RegisterRoutes(r)

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		fmt.Println("HTTP_PORT not set, using default port")
		port = "8080"
	}
	fmt.Printf("Starting server on port %s", port)

	log.Fatalf("Couldn't start server: %s", http.ListenAndServe(":"+port, r))
}
