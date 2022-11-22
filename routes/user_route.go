package routes

import (
	"rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)

	app.Get("/user/:userId", controllers.GetSingleUser)

	app.Get("/users", controllers.GetAllUsers)

	app.Put("/user/:userId", controllers.EditSingleUser)

	app.Get("/userTime/:resourceId", controllers.TimeEntryData)
}
