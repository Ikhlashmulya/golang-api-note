package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
	"github.com/ikhlashmulya/golang-api-note/config"
	"github.com/ikhlashmulya/golang-api-note/controller"
	"github.com/ikhlashmulya/golang-api-note/entity"
	"github.com/ikhlashmulya/golang-api-note/middleware"
	"github.com/ikhlashmulya/golang-api-note/model"
	"github.com/ikhlashmulya/golang-api-note/repository"
	"github.com/ikhlashmulya/golang-api-note/service"
	"github.com/stretchr/testify/assert"
)

var configuration = config.NewConfig("../.env.test")
var db = config.NewGormDB(configuration)
var noteRepository = repository.NewNoteRepository(db)
var validate = validator.New()
var noteService = service.NewNoteService(noteRepository, validate)
var noteController = controller.NewNoteController(noteService)
var app = setupTestApp()

func setupTestApp() *fiber.App {
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Use(middleware.AuthMiddleware())
	noteController.Route(app)
	return app
}

func TestCreateSuccess(t *testing.T) {
	note := model.CreateNoteRequest{
		Title: "Test Create Data",
		Tags:  []string{"tag1", "tag2"},
		Body:  "ini adalah test",
	}
	requestBody, _ := json.Marshal(note)
	request := httptest.NewRequest("POST", "/api/notes", bytes.NewBuffer(requestBody))
	request.Header.Set("content-type", "application/json")
	request.Header.Set("x-api-key", "secret")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "CREATED", webResponse.Status)
	assert.Equal(t, "success create new note", webResponse.Message)

	webResponseData := webResponse.Data.(map[string]any)
	assert.NotNil(t, webResponseData["id"])

	noteRepository.Delete(entity.Note{Id: webResponseData["id"].(string)})
}

func TestCreateBadRequest(t *testing.T) {
	note := model.CreateNoteRequest{
		Title: "",
		Tags:  []string{"tag1", "tag2"},
		Body:  "ini adalah test",
	}
	requestBody, _ := json.Marshal(note)
	request := httptest.NewRequest("POST", "/api/notes", bytes.NewBuffer(requestBody))
	request.Header.Set("content-type", "application/json")
	request.Header.Set("x-api-key", "secret")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "BAD_REQUEST", webResponse.Status)
}

func TestUpdateSuccess(t *testing.T) {
	id := uuid.New().String()
	noteRepository.Create(entity.Note{
		Id:    id,
		Title: "Test Create Data",
		Tags:  "tags1,tags2",
		Body:  "ini adalah test",
	})

	updateNoteRequest := model.UpdateNoteRequest{
		Title: "Test Create Data edited",
		Tags:  []string{"tag1", "tag2"},
		Body:  "ini adalah test",
	}
	requestBody, _ := json.Marshal(updateNoteRequest)

	request := httptest.NewRequest("PUT", "/api/notes/"+id, bytes.NewBuffer(requestBody))
	request.Header.Set("content-type", "application/json")
	request.Header.Set("x-api-key", "secret")

	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)
	assert.Equal(t, "success updated note", webResponse.Message)

	webResponseData := webResponse.Data.(map[string]any)
	assert.Equal(t, updateNoteRequest.Title, webResponseData["title"].(string))
	assert.Equal(t, updateNoteRequest.Body, webResponseData["body"].(string))

	for i := 0; i < len(updateNoteRequest.Tags); i++ {
		assert.Equal(t, updateNoteRequest.Tags[i], webResponseData["tags"].([]interface{})[i].(string))
	}

	noteRepository.Delete(entity.Note{Id: id})
}

func TestUpdateBadRequest(t *testing.T) {
	id := uuid.New().String()
	noteRepository.Create(entity.Note{
		Id:    id,
		Title: "Test Create Data",
		Tags:  "tags1,tags2",
		Body:  "ini adalah test",
	})

	updateNoteRequest := model.UpdateNoteRequest{
		Title: "",
		Tags:  []string{"tag1", "tag2"},
		Body:  "ini adalah test",
	}
	requestBody, _ := json.Marshal(updateNoteRequest)

	request := httptest.NewRequest("PUT", "/api/notes/"+id, bytes.NewBuffer(requestBody))
	request.Header.Set("content-type", "application/json")
	request.Header.Set("x-api-key", "secret")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "BAD_REQUEST", webResponse.Status)

	noteRepository.Delete(entity.Note{Id: id})
}

func TestDeleteSuccess(t *testing.T) {
	id := uuid.New().String()
	noteRepository.Create(entity.Note{
		Id:    id,
		Title: "Test Create Data",
		Tags:  "tags1,tags2",
		Body:  "ini adalah test",
	})

	request := httptest.NewRequest("DELETE", "/api/notes/"+id, nil)
	request.Header.Set("x-api-key", "secret")

	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)
	assert.Equal(t, "success deleted note", webResponse.Message)
}

func TestDeleteNotFound(t *testing.T) {
	request := httptest.NewRequest("DELETE", "/api/notes/"+"1234", nil)
	request.Header.Set("x-api-key", "secret")

	response, _ := app.Test(request)
	assert.Equal(t, 404, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, 404, webResponse.Code)
	assert.Equal(t, "NOT_FOUND", webResponse.Status)
}

func TestFindById(t *testing.T) {
	id := uuid.New().String()
	note := entity.Note{
		Id:    id,
		Title: "Test Create Data",
		Tags:  "tags1,tags2",
		Body:  "ini adalah test",
	}
	noteRepository.Create(note)

	request := httptest.NewRequest("GET", "/api/notes/"+id, nil)
	request.Header.Set("x-api-key", "secret")

	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)
	assert.Equal(t, "success get note", webResponse.Message)

	webResponseData := webResponse.Data.(map[string]any)
	assert.Equal(t, note.Title, webResponseData["title"].(string))
	assert.Equal(t, note.Body, webResponseData["body"].(string))

	noteTags := strings.Split(note.Tags, ",")
	for i := 0; i < len(noteTags); i++ {
		assert.Equal(t, noteTags[i], webResponseData["tags"].([]interface{})[i].(string))
	}

	noteRepository.Delete(entity.Note{Id: id})
}

func TestFindByIdNotFound(t *testing.T) {
	id := uuid.New().String()

	request := httptest.NewRequest("GET", "/api/notes/"+id, nil)
	request.Header.Set("x-api-key", "secret")

	response, _ := app.Test(request)
	assert.Equal(t, 404, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, 404, webResponse.Code)
	assert.Equal(t, "NOT_FOUND", webResponse.Status)
}

func TestFindAll(t *testing.T) {
	id1 := uuid.New().String()
	note1 := entity.Note{
		Id:    id1,
		Title: "Test Create Data",
		Tags:  "tags1,tags2",
		Body:  "ini adalah test",
	}
	noteRepository.Create(note1)
	id2 := uuid.New().String()
	note2 := entity.Note{
		Id:    id2,
		Title: "Test Create Data",
		Tags:  "tags1,tags2",
		Body:  "ini adalah test",
	}
	noteRepository.Create(note2)

	request := httptest.NewRequest("GET", "/api/notes", nil)
	request.Header.Set("x-api-key", "secret")

	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)
	assert.Equal(t, "success get all note", webResponse.Message)

	webResponseData := webResponse.Data.([]any)
	webResponseData1 := webResponseData[0].(map[string]any)
	webResponseData2 := webResponseData[1].(map[string]any)

	assert.Equal(t, note1.Title, webResponseData1["title"].(string))
	assert.Equal(t, note2.Title, webResponseData2["title"].(string))
	assert.Equal(t, note1.Body, webResponseData1["body"].(string))
	assert.Equal(t, note2.Body, webResponseData2["body"].(string))

	noteTags1 := strings.Split(note1.Tags, ",")
	for i := 0; i < len(noteTags1); i++ {
		assert.Equal(t, noteTags1[i], webResponseData1["tags"].([]any)[i].(string))
	}

	noteTags2 := strings.Split(note1.Tags, ",")
	for i := 0; i < len(noteTags2); i++ {
		assert.Equal(t, noteTags2[i], webResponseData2["tags"].([]any)[i].(string))
	}

	noteRepository.Delete(entity.Note{Id: id1})
	noteRepository.Delete(entity.Note{Id: id2})
}

func TestUnauthorized(t *testing.T) {
	id := uuid.New().String()

	request := httptest.NewRequest("GET", "/api/notes/"+id, nil)

	response, _ := app.Test(request)
	assert.Equal(t, 401, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, 401, webResponse.Code)
	assert.Equal(t, "UNAUTHORIZED", webResponse.Status)
}

func TestMethodNotAllowed(t *testing.T) {
	request := httptest.NewRequest("PUT", "/api/notes", nil)
	request.Header.Set("x-api-key", "secret")

	response, _ := app.Test(request)
	assert.Equal(t, 405, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, 405, webResponse.Code)
	assert.Equal(t, "METHOD_NOT_ALLOWED", webResponse.Status)
}
