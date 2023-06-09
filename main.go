package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ikhlashmulya/noteapp-resful-api/config"
	"github.com/ikhlashmulya/noteapp-resful-api/controller"
	"github.com/ikhlashmulya/noteapp-resful-api/middleware"
	"github.com/ikhlashmulya/noteapp-resful-api/repository"
	"github.com/ikhlashmulya/noteapp-resful-api/service"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	//setup configuration
	configuration := config.NewConfig()

	//setup db
	db := config.NewGormDB(configuration)

	//setup validator
	validate := validator.New()

	//setup note repository
	noteRepository := repository.NewNoteRepository(db)

	//setup note service
	noteService := service.NewNoteService(noteRepository, validate)

	//setup note controller
	noteController := controller.NewNoteController(noteService)

	//setup fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Use(middleware.AuthMiddleware())

	//setup route
	noteController.Route(app)

	//start app
	app.Listen(":3000")
}
