package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ro35ert/orders_api/database"
	"github.com/ro35ert/orders_api/models"
)

type User struct{
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func CreateResponseUser(user models.User) User{
	return User{ID: user.ID,FirstName: user.FirstName,LastName: user.LastName}
}

func CreateUser(c *fiber.Ctx) error{
	var user models.User

	if err := c.BodyParser(&user); err != nil{
		return c.Status(400).JSON(err.Error())
	}
	user.CreatedAT = time.Now()
	database.Database.Db.Create(&user)
	responseUser :=CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}