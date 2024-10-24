package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ssr0016/ecommmerse-app/internal/api/rest"
	"github.com/ssr0016/ecommmerse-app/internal/dto"
	"github.com/ssr0016/ecommmerse-app/internal/service"
)

type UserHandler struct {
	// user service
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	// create an instance of user handler & inject to handler
	svc := service.UserService{}
	handler := &UserHandler{
		svc: svc,
	}

	// Public endpoints
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Private endpoints
	app.Get("/verify", handler.GetVerificationCode)
	app.Post("/verify", handler.Verify)
	app.Post("/profile", handler.CreateProfile)
	app.Get("/profile", handler.GetProfile)

	app.Post("/cart", handler.AddToCart)
	app.Get("/cart", handler.GetCart)
	app.Get("/order", handler.GetOrders)
	app.Get("/order/:id", handler.GetOrder)

	app.Post("/become-seller", handler.BecomeSeller)

}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	// to create user
	user := dto.UserSignup{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "please provide valid inputs",
		})
	}

	token, err := h.svc.Signup(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "error on signup",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": token,
	})

}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "login",
	})

}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "get verification code",
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "verify",
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "create profile",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "get profile",
	})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "add to cart",
	})
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "get cart",
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "get orders",
	})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "get order",
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "become seller",
	})
}
