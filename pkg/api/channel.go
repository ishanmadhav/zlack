package api

import "github.com/gofiber/fiber/v2"

func SetupChannelRoutes(app *fiber.App) {

	//Get Channel
	app.Get("/channels/:id", GetChannelByID)
	app.Get("/channels", GetAllChannels)
	app.Get("/channels/:id/messages", GetAllMessagesByChannelID)

	//Post Channel
	app.Post("/channel", CreateChannel)

	//Put Channel
	app.Put("/channel/:id", UpdateChannel)

	//Delete Channel
	app.Delete("/channel/:id", DeleteChannelByID)

}

//Function will make an gRPC call to the Channel service to get a channel by ID
func GetChannelByID(c *fiber.Ctx) error {
	return c.SendString("Get Channel By ID")
}

//Function will make an gRPC call to the Channel service to get all channels
func GetAllChannels(c *fiber.Ctx) error {
	return c.SendString("Get All Channels")
}

//Function will make an gRPC call to the Channel service to create a channel
func CreateChannel(c *fiber.Ctx) error {
	return c.SendString("Create Channel")
}

//Function will make an gRPC call to the Channel service to update a channel
func UpdateChannel(c *fiber.Ctx) error {
	return c.SendString("Update Channel")
}

//Function will make an gRPC call to the Channel service to delete a channel
func DeleteChannelByID(c *fiber.Ctx) error {
	return c.SendString("Delete Channel")
}

//Function will make an gRPC call to the Channel service to get all messages by channel ID
func GetAllMessagesByChannelID(c *fiber.Ctx) error {
	return c.SendString("Get All Messages By Channel ID")
}
