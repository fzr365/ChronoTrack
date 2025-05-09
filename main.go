// === main.go ===
package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"chrono-track/database"
	"chrono-track/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database.ConnectDB()

	app := fiber.New()
	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
