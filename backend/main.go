package main

import (
	"log"

	"posttest/backend/config"
	"posttest/backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	routes.Register(app)

	port := config.GetEnv("PORT", "3000")
	log.Printf("server running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
