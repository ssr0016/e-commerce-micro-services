package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ssr0016/ecommmerse-app/config"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Get("/health", healthCheck)

	app.Listen(config.ServerPort)
}

func healthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "I'm  am  alive"})
}
