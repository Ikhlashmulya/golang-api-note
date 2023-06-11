package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

// func AuthMiddleware() fiber.Handler {

// 	allowedAPIKey := []string{
// 		"secret",
// 	}

// 	return func(ctx *fiber.Ctx) error {
// 		apiKey := ctx.Get("x-api-key")

// 		if !isValidAPIKey(apiKey, allowedAPIKey) {
// 			return ctx.Status(401).JSON(model.WebResponse{
// 				Code: 401,
// 				Status: "UNAUTHORIZED",
// 				Message: "Status Unauthorized",
// 			})
// 		}

// 		return ctx.Next()
// 	}
// }

var allowedAPIKey []string = []string{
	"Mawar itu biru Violet itu merah",
	"secret",
}

func ValidateApiKey(ctx *fiber.Ctx, key string) (bool, error) {
	apiKey := ctx.Get("x-api-key")
	
	if !isValidAPIKey(apiKey, allowedAPIKey) {
		return false, keyauth.ErrMissingOrMalformedAPIKey
	}

	return true, nil
}

func isValidAPIKey(apiKey string, allowedAPIKey []string) bool {
	for _, key := range allowedAPIKey {
		if key == apiKey {
			return true
		}
	}

	return false
}
