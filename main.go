package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ro35ert/orders_api/database"
)


func welcome(c *fiber.Ctx) error{
	return c.SendString("Welcome buddy")
}

func main(){
	database.DOnnectDb()
	app := fiber.New()

	app.Get("/",welcome)

	log.Fatal(app.Listen(":3000"))
}