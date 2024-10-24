package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssr0016/ecommmerse-app/config"
	"github.com/ssr0016/ecommmerse-app/internal/api/rest"
	"github.com/ssr0016/ecommmerse-app/internal/api/rest/handlers"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	rh := &rest.RestHandler{
		App: app,
	}

	setupRoutes(rh)

	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	// user handler
	handlers.SetupUserRoutes(rh)

	// transactions

	// catalog
}
