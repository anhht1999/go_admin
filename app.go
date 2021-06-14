package main

import (
  "github.com/anhht1999/go_admin/database"
  "github.com/anhht1999/go_admin/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
  //database
  database.Connect()

  app := fiber.New()

  routers.Setup(app)

  app.Listen(":3000")
}