package routes

import (
	"rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func CalculateRoute(app *fiber.App) {
	app.Get("/calculate", controllers.Calculate)

	app.Get("/calculate-without-agg", controllers.CalculateWithoutAgg)
}
