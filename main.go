package main

import (
	"rest/configs"
	"rest/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.CalculateRoute(app)

	app.Listen(":80")
}
