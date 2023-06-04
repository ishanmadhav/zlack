package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ishanmadhav/zlack/database"
	"github.com/ishanmadhav/zlack/models"
)

func CreateChannel(c *fiber.Ctx) error {
	fmt.Println("Create channel route was hit")
	db := database.DBConn
	channel := new(models.Channel)
	err := c.BodyParser(&channel)
	if err != nil {
		fmt.Println("Error in body parsing")
		fmt.Print(err)
		return c.JSON(err)
	}
	result := db.Create(&channel)
	if result.Error != nil {
		fmt.Println("Error in creation")
		fmt.Println(result.Error)
		return c.JSON(result.Error)
	}
	return c.JSON(channel)
}

func GetUsersInChannel(c *fiber.Ctx) error {
	fmt.Println("Get all Users in a channel")
	db := database.DBConn
	var channel models.Channel
	err := db.Model(&models.Channel{}).Preload("Users").First(&channel, c.Params("id")).Error
	if err != nil {
		fmt.Print(err)
		return c.JSON(err)
	}
	c.Status(fiber.StatusFound)
	return c.JSON(channel)
}

func GetAllChannels(c *fiber.Ctx) error {
	fmt.Println("Get All Channels")
	db := database.DBConn
	var channels []models.Channel
	err := db.Model(&models.Channel{}).Preload("Users").Find(&channels).Error
	if err != nil {
		fmt.Print(err)
		return c.JSON(err)
	}
	return c.JSON(channels)
}
