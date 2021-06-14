package routers

import (
	"github.com/anhht1999/go_admin/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App)  {
	
	app.Get("/", controller.Hello)
}