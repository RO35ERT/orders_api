package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ro35ert/orders_api/database"
	"github.com/ro35ert/orders_api/routes"
)


func welcome(c *fiber.Ctx) error{
	return c.SendString("Welcome buddy")
}

func setUpRoutes(app *fiber.App){
	app.Get("/api", welcome)
	app.Post("/api/users",routes.CreateUser)
}

func main(){
	database.DOnnectDb()
	app := fiber.New()

	setUpRoutes(app)
	app.Get("/",welcome)

	log.Fatal(app.Listen(":3000"))
}