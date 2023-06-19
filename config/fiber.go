package config

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/ikhlashmulya/golang-api-note/exception"
)

//fiber configuration

var signingKey = []byte("secret key")

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}

func NewFiberAuthConfig() jwtware.Config {
	return jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: signingKey},
		ErrorHandler: exception.ErrorHandler,
	}
}
