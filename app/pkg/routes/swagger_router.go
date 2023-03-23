package routes

import (
	"github.com/gofiber/fiber/v2"

	swagger "github.com/gofiber/swagger"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRouter(a *fiber.App) {
	// Create routes group.
	route := a.Group("/swagger")

	route.Get("/*", swagger.HandlerDefault)
}
