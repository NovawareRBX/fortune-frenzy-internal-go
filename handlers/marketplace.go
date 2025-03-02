package handlers

import (
	"ffinternal-go/handlers/marketplace"

	"github.com/gofiber/fiber/v2"
)

func GetAllItems(c *fiber.Ctx) error {
	return marketplace.GetAllItems(c)
}