package routes

import (
	"github.com/gofiber/fiber/v2"

	"posttest/backend/handlers"
)

// Register attaches API routes to the Fiber app.
func Register(app *fiber.App) {
	api := app.Group("/api")

	polylines := api.Group("/polylines")
	polylines.Post("/", handlers.CreatePolyline)
	polylines.Get("/", handlers.GetPolylines)
	polylines.Get("/:id", handlers.GetPolyline)
	polylines.Put("/:id", handlers.UpdatePolyline)
	polylines.Delete("/:id", handlers.DeletePolyline)
}
