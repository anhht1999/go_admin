package main

import (
	"github.com/anhht1999/go_admin/database"
	"github.com/anhht1999/go_admin/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
  //database
  database.Connect()

  app := fiber.New()

  app.Use(cors.New(cors.Config{
      AllowCredentials: true,
  }))

  routers.Setup(app)

  app.Listen(":3000")
}