package routes

import (
	"ffinternal-go/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupMarketplaceRoutes(app *fiber.App) {
	app.Get("/marketplace/items", handlers.GetAllItems)
}