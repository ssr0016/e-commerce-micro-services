package api

import "github.com/gofiber/fiber/v2"

func StartServer() {

	app := fiber.New()

	app.Listen("localhost:9000")
}
