package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ikhlashmulya/golang-api-note/config"
	"github.com/ikhlashmulya/golang-api-note/controller"
	"github.com/ikhlashmulya/golang-api-note/repository/impl"
	"github.com/ikhlashmulya/golang-api-note/service/impl"

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

	//setup user
	userRepository := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepository, []byte("secret key"))

	userController := controller.NewUserController(userService)

	//setup fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	// app.Use(config.NewFiberKeyAuthConfig())

	//setup route
	noteController.Route(app)
	userController.Route(app)

	//start app
	app.Listen(":3000")
}
