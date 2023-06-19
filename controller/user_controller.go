package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikhlashmulya/golang-api-note/exception"
	"github.com/ikhlashmulya/golang-api-note/model"
	"github.com/ikhlashmulya/golang-api-note/service"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (controller *UserController) Route(app *fiber.App) {
	auth := app.Group("/api/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/register", controller.Register)
}

func (controller *UserController) Login(ctx *fiber.Ctx) error {
	var loginInput model.LoginInput
	err := ctx.BodyParser(&loginInput)
	exception.PanicIfErr(err)

	token := controller.UserService.Login(ctx.Context(), loginInput)

	return ctx.JSON(model.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "OK",
		Message: "success login",
		Data: model.LoginResponse{
			Token: token,
		},
	})
}

func (controller *UserController) Register(ctx *fiber.Ctx) error {
	var registerRequest model.RegisterRequest
	err := ctx.BodyParser(&registerRequest)
	exception.PanicIfErr(err)

	response := controller.UserService.Register(ctx.Context(), registerRequest)

	return ctx.JSON(model.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "OK",
		Message: "register success",
		Data:    response,
	})
}
