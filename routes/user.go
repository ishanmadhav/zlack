package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ishanmadhav/zlack/controllers"
)

func UserRoutes(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:id", controllers.GetUserById)
	app.Get("/users", controllers.GetAllUsers)
	app.Delete("/user/:id", controllers.DeleteUserById)

}
