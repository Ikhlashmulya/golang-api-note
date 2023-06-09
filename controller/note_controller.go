package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikhlashmulya/golang-api-note/exception"
	"github.com/ikhlashmulya/golang-api-note/model"
	"github.com/ikhlashmulya/golang-api-note/service"
)

// note controller
type NoteController struct {
	noteService service.NoteService
}

func NewNoteController(service service.NoteService) *NoteController {
	return &NoteController{service}
}

func (controller *NoteController) Route(app *fiber.App) {
	app.Post("/api/notes", controller.Create)
	app.Get("/api/notes", controller.FindAll)
	app.Get("/api/notes/:noteId", controller.FindById)
	app.Put("/api/notes/:noteId", controller.Update)
	app.Delete("/api/notes/:noteId", controller.Delete)
}

func (controller *NoteController) Create(ctx *fiber.Ctx) error {
	// get data from request body
	createNoteRequest := model.CreateNoteRequest{}
	err := ctx.BodyParser(&createNoteRequest)
	exception.PanicIfErr(err)

	// send data to service and get response
	response := controller.noteService.Create(createNoteRequest)

	//response body
	return ctx.Status(201).JSON(model.WebResponse{
		Code:    201,
		Status:  "CREATED",
		Message: "success create new note",
		Data:    response,
	})
}

func (controller *NoteController) Update(ctx *fiber.Ctx) error {
	// get data from request body
	updateNoteRequest := model.UpdateNoteRequest{}
	err := ctx.BodyParser(&updateNoteRequest)
	exception.PanicIfErr(err)

	//get path parameter
	updateNoteRequest.Id = ctx.Params("noteId")

	// send data to service and get response
	response := controller.noteService.Update(updateNoteRequest)

	//response body
	return ctx.JSON(model.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "success updated note",
		Data:    response,
	})
}

func (controller *NoteController) Delete(ctx *fiber.Ctx) error {
	//get path parameter
	noteId := ctx.Params("noteId")

	// send data to service
	controller.noteService.Delete(noteId)

	//response body
	return ctx.JSON(model.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "success deleted note",
	})
}

func (controller *NoteController) FindById(ctx *fiber.Ctx) error {
	//get path parameter
	noteId := ctx.Params("noteId")

	//find data
	response := controller.noteService.FindById(noteId)

	//response body
	return ctx.JSON(model.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "success get note",
		Data:    response,
	})
}

func (controller *NoteController) FindAll(ctx *fiber.Ctx) error {
	//get all data
	response := controller.noteService.FindAll()

	//response body
	return ctx.JSON(model.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "success get all note",
		Data:    response,
	})
}
