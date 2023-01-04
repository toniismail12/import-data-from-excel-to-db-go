package main

import (
	
	"import-excel/config"
	"import-excel/database"
	"import-excel/router"

	"github.com/gofiber/fiber/v2"

)


func main() {
	
	// connection database
	database.Connect()
	
	app := fiber.New()
	
	// router
	router.Setup(app)

	// get port
	port := config.Environment("PORT")

	// run
	app.Listen(":" + port)

}
