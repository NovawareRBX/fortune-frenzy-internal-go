package main

import (
	"ffinternal-go/middleware"
	"ffinternal-go/routes"
	"ffinternal-go/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	service.InitMariaDB()
	service.InitRedis()
	app := fiber.New()

	app.Use(middleware.RequestTimer())

	routes.SetupMarketplaceRoutes(app)
	routes.SetupCoinflipRoutes(app)

	log.Println("Server is running on port 3004")
	if err := app.Listen(":3004"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
