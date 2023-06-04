package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ishanmadhav/zlack/database"
	"github.com/ishanmadhav/zlack/models"
)

type UserBody struct {
	Username string `json:"username"`
	Channel  uint   `json:"channelId"`
}

type responseMessage struct {
	message string
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DBConn
	// Parse the request body into a User struct
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	fmt.Println(user)
	// Insert the user into the database
	err := db.Create(&user).Error
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func GetUserById(c *fiber.Ctx) error {
	fmt.Println("Get User by id")
	db := database.DBConn
	var user models.User
	result := db.First(&user, c.Params("id"))
	if result.Error != nil {
		return c.JSON(result.Error)
	}
	return c.JSON(user)
}

func GetAllUsers(c *fiber.Ctx) error {
	fmt.Println("Get All Users")
	db := database.DBConn
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		return c.JSON(result.Error)
	}
	return c.JSON(users)
}

func DeleteUserById(c *fiber.Ctx) error {
	fmt.Println("Delete User by Id")
	db := database.DBConn
	result := db.Delete(&models.User{}, c.Params("id"))
	if result.Error != nil {
		return c.JSON(result.Error)
	}
	response := responseMessage{message: "User successfully deleted"}
	return c.JSON(response)
}

func EditUserById(c *fiber.Ctx) error {
	fmt.Println("Edit User by ID")
	db := database.DBConn
	var user models.User
	db.First(&user, c.Params("id"))
	tmpUser := new(models.User)
	err := c.BodyParser(&tmpUser)
	if err != nil {
		fmt.Println(err)
		return c.JSON(err)
	}
	user.ChannelID = tmpUser.ChannelID
	user.Username = tmpUser.Username
	result := db.Save(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return c.JSON(result.Error)
	}
	return c.JSON(user)
}
