package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikhlashmulya/golang-api-note/exception"
)

//fiber configuration

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
