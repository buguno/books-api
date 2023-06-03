package main

import (
	"books-api/pkg/configs"
	"books-api/pkg/middleware"
	"books-api/pkg/routes"
	"books-api/pkg/utils"

	"github.com/gofiber/fiber/v2"

	_ "books-api/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title Books API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name MIT License
// @license.url https://mit-license.org/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
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
