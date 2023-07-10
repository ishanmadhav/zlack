package api

import "github.com/gofiber/fiber/v2"

func SetupWorkspaceRoutes(app *fiber.App) {

	//Get Routes
	app.Get("/workspaces", GetAllWorkspaces)
	app.Get("/workspaces/:id", GetWorkspaceByID)

	//Post Routes
	app.Post("/workspace", CreateWorkspace)

	//Put Routes
	app.Put("/workspace/:id", UpdateWorkspace)

	//Delete Routes
	app.Delete("/workspace/:id", DeleteWorkspaceByID)
	//Utility Route, should be delete later
	app.Delete("/workspaces", DeleteAllWorkspaces)
}

//Function will make an gRPC call to the Workspace service to get all workspaces
func GetAllWorkspaces(c *fiber.Ctx) error {
	return c.SendString("Get All Workspaces")
}

//Function will make an gRPC call to the Workspace service to get a workspace by ID
func GetWorkspaceByID(c *fiber.Ctx) error {
	return c.SendString("Get Workspace By ID")
}

//Function will make an gRPC call to the Workspace service to create a workspace
func CreateWorkspace(c *fiber.Ctx) error {
	return c.SendString("Create Workspace")
}

//Function will make an gRPC call to the Workspace service to update a workspace
func UpdateWorkspace(c *fiber.Ctx) error {
	return c.SendString("Update Workspace")
}

//Function will make an gRPC call to the Workspace service to delete a workspace
func DeleteWorkspaceByID(c *fiber.Ctx) error {
	return c.SendString("Delete Workspace")
}

//Function will make an gRPC call to the Workspace service to delete all workspaces
func DeleteAllWorkspaces(c *fiber.Ctx) error {
	return c.SendString("Delete All Workspaces")
}
