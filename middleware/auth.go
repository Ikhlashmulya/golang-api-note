package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/ikhlashmulya/golang-api-note/config"
)

func Protected() fiber.Handler {
	return jwtware.New(config.NewFiberAuthConfig())
}
