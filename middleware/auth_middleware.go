package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikhlashmulya/golang-api-note/model"
)

func AuthMiddleware() fiber.Handler {

	allowedAPIKey := []string{
		"secret",
	}

	return func(ctx *fiber.Ctx) error {
		apiKey := ctx.Get("x-api-key")

		if !isValidAPIKey(apiKey, allowedAPIKey) {
			return ctx.Status(401).JSON(model.WebResponse{
				Code: 401,
				Status: "UNAUTHORIZED",
				Message: "Status Unauthorized",
			})
		}

		return ctx.Next()
	}
}

func isValidAPIKey(apiKey string, allowedAPIKey []string) bool {
	for _, key := range allowedAPIKey {
		if key == apiKey {
			return true
		}
	}

	return false
}
