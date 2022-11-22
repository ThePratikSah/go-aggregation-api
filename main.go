package main

import (
	"rest/configs"
	"rest/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// connect to database
	configs.ConnectDB()

	// routes middleware
	routes.UserRouter(app)
	routes.CalculateRoute(app)

	app.Listen(":80")
}
