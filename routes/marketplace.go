package routes

import (
	"ffinternal-go/handlers"
	// "ffinternal-go/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupMarketplaceRoutes(app *fiber.App) {
	marketplace := app.Group("/marketplace")

	marketplace.Get("/items",
		// middleware.Authorization(middleware.AuthTypeServerKey),
		handlers.GetAllItems,
	)

	marketplace.Get("/items/:id",
		// middleware.Authorization(middleware.AuthTypeServerKey),
		handlers.GetItemByID,
	)
}
