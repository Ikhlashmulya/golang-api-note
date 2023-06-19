package exception

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ikhlashmulya/golang-api-note/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(model.WebResponse{
			Code:    fiber.StatusNotFound,
			Status:  "NOT_FOUND",
			Message: err.Error(),
		})
	}

	if err.Error() == "Method Not Allowed" {
		return ctx.Status(fiber.StatusMethodNotAllowed).JSON(model.WebResponse{
			Code:    fiber.StatusMethodNotAllowed,
			Status:  "METHOD_NOT_ALLOWED",
			Message: err.Error(),
		})
	}

	if err.Error() == "missing or malformed JWT" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse{
			Code:    fiber.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "Invalid or expired JWT",
		})
	}

	_, ok := err.(validator.ValidationErrors)
	if ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:    fiber.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:    fiber.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Invalid password",
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse{
		Code:    fiber.StatusInternalServerError,
		Status:  "INTERNAL_SERVER_ERROR",
		Message: err.Error(),
	})
}
