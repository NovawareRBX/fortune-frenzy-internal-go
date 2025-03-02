package handlers

import (
	marketplace "ffinternal-go/handlers/todo"

	"github.com/gofiber/fiber/v2"
)

func GetAllItems(c *fiber.Ctx) error {
	return marketplace.GetAllItems(c)
}