package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ishanmadhav/zlack/controllers"
)

func ChannelRoutes(app *fiber.App) {
	app.Post("/channel", controllers.CreateChannel)
	app.Get("/channel/:id", controllers.GetUsersInChannel)
	app.Get("/channels", controllers.GetAllChannels)
}
