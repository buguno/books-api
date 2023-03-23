package routes

import (
	"books-api/app/controllers"
	"books-api/app/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook)
	route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook)
	route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook)
}
