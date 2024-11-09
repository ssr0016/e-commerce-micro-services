package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ssr0016/ecommmerse-app/config"
	"github.com/ssr0016/ecommmerse-app/internal/api/rest"
	"github.com/ssr0016/ecommmerse-app/internal/api/rest/handlers"
	"github.com/ssr0016/ecommmerse-app/internal/domain"
	"github.com/ssr0016/ecommmerse-app/internal/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	// connect to database
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error %v\n", err)
	}

	log.Println("database connected successfully!")

	// run migration
	db.AutoMigrate(&domain.User{})

	auth := helper.SetupAuth(config.AppSecret)

	rh := &rest.RestHandler{
		App:    app,
		DB:     db,
		Auth:   auth,
		Config: config,
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
