package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ishanmadhav/zlack/database"
	"github.com/ishanmadhav/zlack/models"
	"github.com/ishanmadhav/zlack/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=mysecretpassword dbname=zlack port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	database.DBConn.AutoMigrate(&models.User{})
	database.DBConn.AutoMigrate(&models.Message{})
	database.DBConn.AutoMigrate(&models.Channel{})
	if err != nil {
		panic("failed to connect database")
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Up and running")
	})

	routes.UserRoutes(app)
	routes.ChannelRoutes(app)

	app.Listen(":5000")

}
