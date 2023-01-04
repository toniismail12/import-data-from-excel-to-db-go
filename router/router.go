package router

import (
	"import-excel/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// app.Get("/upload", controller.Upload)
	app.Get("/", controller.Import)

}
