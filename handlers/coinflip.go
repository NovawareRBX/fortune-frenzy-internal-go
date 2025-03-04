package handlers

import (
	"ffinternal-go/handlers/coinflip"

	"github.com/gofiber/fiber/v2"
)

func CreateCoinflip(c *fiber.Ctx) error {
	return coinflip.CreateCoinflip(c)
}
