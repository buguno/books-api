package main

import (
	"books-api/app/pkg/configs"
	"books-api/app/pkg/middleware"
	"books-api/app/pkg/routes"
	"books-api/app/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRouter(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)   // Register a public routes for app.
	routes.PrivateRoutes(app)  // Register a private routes for app.
	routes.NotFoundRouter(app) // Register route for 404 Error.

	// Start server (with graceful shutdown).
	utils.StartServerWithGracefulShutdown(app)
}
