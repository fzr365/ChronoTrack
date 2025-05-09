package routes

import (
	"ChronoTrack/controllers"
	"ChronoTrack/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	task := api.Group("/tasks", middleware.JWTProtected())
	task.Post("/", controllers.CreateTask)
	task.Get("/", controllers.GetTasks)
	task.Put("/:id", controllers.UpdateTask)
	task.Delete("/:id", controllers.DeleteTask)
}
