package exception

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/ikhlashmulya/golang-api-note/model"
	"gorm.io/gorm"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(404).JSON(model.WebResponse{
			Code:    404,
			Status:  "NOT_FOUND",
			Message: err.Error(),
		})
	}

	if errors.Is(err, keyauth.ErrMissingOrMalformedAPIKey) {
		return ctx.Status(401).JSON(model.WebResponse{
			Code:    401,
			Status:  "UNAUTHORIZED",
			Message: err.Error(),
		})
	}

	if err.Error() == "Method Not Allowed" {
		return ctx.Status(405).JSON(model.WebResponse{
			Code:    405,
			Status:  "METHOD_NOT_ALLOWED",
			Message: err.Error(),
		})
	}

	_, ok := err.(validator.ValidationErrors)
	if ok {
		return ctx.Status(400).JSON(model.WebResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	return ctx.Status(500).JSON(model.WebResponse{
		Code:    500,
		Status:  "INTERNAL_SERVER_ERROR",
		Message: err.Error(),
	})
}
