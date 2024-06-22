package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tst/backend/database"
	"github.com/tst/backend/route"
)

func DB() {
	database.DatabaseInit() // Connect to database SqlServer
	// database.DatabaseInitMysql() // Connect to  Database MySql
}

func handleRequest() {

	DB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: true,
	}))

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowMethods: strings.Join([]string{
	// 		fiber.MethodGet,
	// 		fiber.MethodPost,
	// 		fiber.MethodHead,
	// 		fiber.MethodPut,
	// 		fiber.MethodDelete,
	// 		fiber.MethodPatch,
	// 	}, ","),
	// 	AllowCredentials: true,
	// }))

	route.RouteInit(app)
	app.Listen(":3000")
}

func main() {

	handleRequest()
}
