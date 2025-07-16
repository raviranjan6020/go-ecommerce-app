package handlers

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)
type UserHandler struct{
	//user service
	svc service.UserService
}
func SetupUserRoutes(rh *rest.RestHandler){
	app :=rh.App
	// create an instance of user service and inject to handler
	svc := service.UserService{}
	handler := UserHandler{
		svc:  svc,
	}
	// public endpoints
	app.Post("/register",handler.Register)
	app.Post("/login",handler.Register)
	// private endpoints
	app.Get("/verify",handler.GetVerificationCode)
	app.Post("/verify",handler.Verify)

	app.Get("/profile",handler.Profile)
	app.Post("/profile",handler.CreateProfile)

	app.Get("/cart",handler.GetCart)
	app.Post("/cart",handler.AddToCart)

	app.Get("/order",handler.GetOrders)
	app.Get("/order/:id",handler.GetOrder)

	app.Post("/become-seller",handler.BecomeSeller)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignup{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide valid inputs",
		})
	}

	token, err := h.svc.Signup(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error on signup",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register",
		"token":   token,
	})
}
func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
	})
}
func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verify",
	})
}
func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verify",
	})
}
func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "profile-created",
	})
}
func (h *UserHandler) Profile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "profile",
	})
}
func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "add-to-cart",
	})
}
func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get-cart",
	})
}
func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get-order by id",
	})
}
func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get all orders",
	})
}
func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "become-seller",
	})
}