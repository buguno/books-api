package routes

import "github.com/gofiber/fiber/v2"

// NotFoundRouter func for describe 404 Error route.
func NotFoundRouter(a *fiber.App) {
	// Register new special route.
	a.Use(
		// Anonimus function.
		func(c *fiber.Ctx) error {
			// Return HTTP 404 status and JSON response.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "sorry, endpoint is not found",
			})
		},
	)
}
