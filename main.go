package main

import (
	"books-api/pkg/configs"
	"books-api/pkg/middleware"
	"books-api/pkg/routes"
	"books-api/pkg/utils"

	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
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
