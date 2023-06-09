package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikhlashmulya/noteapp-resful-api/exception"
)

//fiber configuration

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
