package api

import "github.com/gofiber/fiber/v2"

//api endpoint for everything from user side except the messaging and realtime notifications
//which will be handled by websocket

type Routes struct {
}

type API struct {
	App *fiber.App
}

func NewAPI() *API {
	return &API{}
}

func (api *API) Start() {
	app := fiber.New()
	api.App = app
	api.Routes()
	api.App.Listen(":3000")
}

func (api *API) Routes() {
	SetupUserRoutes(api.App)
	SetupWorkspaceRoutes(api.App)
	SetupChannelRoutes(api.App)
	SetupWorkspaceRoutes(api.App)
	SetupMessagingRoutes(api.App)
}
