package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/ikhlashmulya/golang-api-note/exception"
	"github.com/ikhlashmulya/golang-api-note/middleware"
)

//fiber configuration

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}

func NewFiberKeyAuthConfig() fiber.Handler {
	return keyauth.New(keyauth.Config{
		SuccessHandler: func(ctx *fiber.Ctx) error {
			return ctx.Next()
		},
		ErrorHandler: exception.ErrorHandler,
		KeyLookup:    "header:x-api-key",
		Validator:    middleware.ValidateApiKey,
	})
}
