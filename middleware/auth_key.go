package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

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
