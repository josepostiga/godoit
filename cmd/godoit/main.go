package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	health_check "github.com/josepostiga/godoit/internal/health-check"
	"github.com/josepostiga/godoit/internal/tasks"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Couldn't load .env file: %s", err)
	}

	app := fiber.New()

	health_check.RegisterRoutes(app)
	tasks.RegisterRoutes(app)

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		fmt.Println("HTTP_PORT not set, using default port")
		port = "8080"
	}
	fmt.Printf("Started server on port %s", port)

	log.Fatalf("Couldn't start server: %s", app.Listen(":"+port))
}
