package handlers

import (
	"ffinternal-go/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var todos []models.Todo
var nextID = 1

func GetAllTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	newTodo := new(models.Todo)
	if err := c.BodyParser(newTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	newTodo.ID = nextID
	nextID++
	todos = append(todos, *newTodo)

	return c.Status(fiber.StatusCreated).JSON(newTodo)
}

func GetTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, todo := range todos {
		if strconv.Itoa(todo.ID) == id {
			return c.JSON(todo)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Todo not found",
	})
}
