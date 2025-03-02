package routes

import (
	"ffinternal-go/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupTodoRoutes(app *fiber.App) {
	app.Get("/todos", handlers.GetAllTodos)
	app.Post("/todos", handlers.CreateTodo)
	app.Get("/todos/:id", handlers.GetTodoByID)
}