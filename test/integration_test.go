package test

import (
	"bytes"
	"context"
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
	"github.com/ikhlashmulya/golang-api-note/model"
	repository "github.com/ikhlashmulya/golang-api-note/repository/impl"
	service "github.com/ikhlashmulya/golang-api-note/service/impl"
	"github.com/stretchr/testify/assert"
)

var configuration = config.NewConfig("../.env.test")
var db = config.NewGormDB(configuration)
var noteRepository = repository.NewNoteRepository(db)
var validate = validator.New()
var noteService = service.NewNoteService(noteRepository, validate)
var noteController = controller.NewNoteController(noteService)
var userRepository = repository.NewUserRepository(db)
var userService = service.NewUserService(userRepository, []byte("secret key"))
var userController = controller.NewUserController(userService)
var app = setupTestApp()

func setupTestApp() *fiber.App {
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	// app.Use(config.NewFiberKeyAuthConfig())
	noteController.Route(app)
	userController.Route(app)
	return app
}

func TestRegister(t *testing.T) {
	requestBody := strings.NewReader(`{"name":"test", "username":"test", "password":"test"}`)
	request := httptest.NewRequest(fiber.MethodPost, "/api/auth/register", requestBody)
	request.Header.Add("content-type", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "register success", responseBody["message"])
}

func TestCreateNote(t *testing.T) {
	var token string
	t.Run("login", func(t *testing.T) {
		data := model.LoginInput{
			Username: "test",
			Password: "test",
		}
		loginInput, _ := json.Marshal(data)
		// loginInput := strings.NewReader(`{"username:"test", "password":"test"}`)
		request := httptest.NewRequest(fiber.MethodPost, "/api/auth/login", bytes.NewBuffer(loginInput))
		request.Header.Add("content-type", "application/json")

		response, _ := app.Test(request)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)

		body, _ := io.ReadAll(response.Body)
		var responseBody model.WebResponse
		json.Unmarshal(body, &responseBody)

		token = responseBody.Data.(map[string]any)["token"].(string)
		assert.NotNil(t, token)
	})

	t.Run("create note success", func(t *testing.T) {
		note := model.CreateNoteRequest{
			Title: "catatan1",
			Tags:  []string{"tag1", "tag2"},
			Body:  "ini adalah test",
		}
		requestBody, _ := json.Marshal(note)
		request := httptest.NewRequest("POST", "/api/notes", bytes.NewBuffer(requestBody))
		request.Header.Set("content-type", "application/json")
		request.Header.Set("Authorization", "Bearer "+token)

		response, _ := app.Test(request)
		assert.Equal(t, fiber.StatusCreated, response.StatusCode)

		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]any
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, fiber.StatusCreated, int(responseBody["code"].(float64)))
		assert.Equal(t, "CREATED", responseBody["status"].(string))
		assert.Equal(t, "success create new note", responseBody["message"].(string))
		assert.NotNil(t, responseBody["data"].(map[string]any)["id"])

		noteRepository.Delete(context.TODO(), entity.Note{
			Id: responseBody["data"].(map[string]any)["id"].(string),
		})
	})

	t.Run("create note bad request", func(t *testing.T) {
		note := model.CreateNoteRequest{
			Title: "",
			Tags:  []string{"tag1", "tag2"},
			Body:  "ini adalah test",
		}
		requestBody, _ := json.Marshal(note)
		request := httptest.NewRequest("POST", "/api/notes", bytes.NewBuffer(requestBody))
		request.Header.Set("content-type", "application/json")
		request.Header.Set("Authorization", "Bearer "+token)

		response, _ := app.Test(request)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)

		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]any
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, fiber.StatusBadRequest, int(responseBody["code"].(float64)))
		assert.Equal(t, "BAD_REQUEST", responseBody["status"].(string))
	})
}

func TestUpdateNote(t *testing.T) {
	var token string
	t.Run("login", func(t *testing.T) {
		data := model.LoginInput{
			Username: "test",
			Password: "test",
		}
		loginInput, _ := json.Marshal(data)
		// loginInput := strings.NewReader(`{"username:"test", "password":"test"}`)
		request := httptest.NewRequest(fiber.MethodPost, "/api/auth/login", bytes.NewBuffer(loginInput))
		request.Header.Add("content-type", "application/json")

		response, _ := app.Test(request)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)

		body, _ := io.ReadAll(response.Body)
		var responseBody model.WebResponse
		json.Unmarshal(body, &responseBody)

		token = responseBody.Data.(map[string]any)["token"].(string)
		assert.NotNil(t, token)
	})

	t.Run("test update note request success", func(t *testing.T) {
		id := uuid.New().String()
		noteRepository.Create(context.TODO(), entity.Note{
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
		request.Header.Set("Authorization", "Bearer "+token)

		response, _ := app.Test(request)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)

		responseBody, _ := io.ReadAll(response.Body)
		webResponse := model.WebResponse{}
		json.Unmarshal(responseBody, &webResponse)

		assert.Equal(t, fiber.StatusOK, webResponse.Code)
		assert.Equal(t, "OK", webResponse.Status)
		assert.Equal(t, "success updated note", webResponse.Message)

		webResponseData := webResponse.Data.(map[string]any)
		assert.Equal(t, updateNoteRequest.Title, webResponseData["title"].(string))
		assert.Equal(t, updateNoteRequest.Body, webResponseData["body"].(string))

		for i := 0; i < len(updateNoteRequest.Tags); i++ {
			assert.Equal(t, updateNoteRequest.Tags[i], webResponseData["tags"].([]interface{})[i].(string))
		}

		noteRepository.Delete(context.TODO(), entity.Note{Id: id})
	})

	t.Run("test update note bad request", func(t *testing.T) {
		id := uuid.New().String()
		noteRepository.Create(context.TODO(), entity.Note{
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
		request.Header.Set("Authorization", "Bearer "+token)

		response, _ := app.Test(request)
		assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)

		responseBody, _ := io.ReadAll(response.Body)
		webResponse := model.WebResponse{}
		json.Unmarshal(responseBody, &webResponse)

		assert.Equal(t, fiber.StatusBadRequest, webResponse.Code)
		assert.Equal(t, "BAD_REQUEST", webResponse.Status)

		noteRepository.Delete(context.TODO(), entity.Note{Id: id})
	})
}

func TestDeleteNote(t *testing.T) {
	var token string
	t.Run("login", func(t *testing.T) {
		data := model.LoginInput{
			Username: "test",
			Password: "test",
		}
		loginInput, _ := json.Marshal(data)
		// loginInput := strings.NewReader(`{"username:"test", "password":"test"}`)
		request := httptest.NewRequest(fiber.MethodPost, "/api/auth/login", bytes.NewBuffer(loginInput))
		request.Header.Add("content-type", "application/json")

		response, _ := app.Test(request)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)

		body, _ := io.ReadAll(response.Body)
		var responseBody model.WebResponse
		json.Unmarshal(body, &responseBody)

		token = responseBody.Data.(map[string]any)["token"].(string)
		assert.NotNil(t, token)
	})

	t.Run("test delete note success", func(t *testing.T) {
		id := uuid.New().String()
		noteRepository.Create(context.TODO(), entity.Note{
			Id:    id,
			Title: "Test Create Data",
			Tags:  "tags1,tags2",
			Body:  "ini adalah test",
		})

		request := httptest.NewRequest("DELETE", "/api/notes/"+id, nil)
		request.Header.Set("Authorization", "Bearer "+token)

		response, _ := app.Test(request)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)

		responseBody, _ := io.ReadAll(response.Body)
		webResponse := model.WebResponse{}
		json.Unmarshal(responseBody, &webResponse)

		assert.Equal(t, fiber.StatusOK, webResponse.Code)
		assert.Equal(t, "OK", webResponse.Status)
		assert.Equal(t, "success deleted note", webResponse.Message)
	})

	t.Run("test delete note not found", func(t *testing.T) {
		request := httptest.NewRequest("DELETE", "/api/notes/"+"1234", nil)
		request.Header.Set("Authorization", "Bearer "+token)

		response, _ := app.Test(request)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)

		responseBody, _ := io.ReadAll(response.Body)
		webResponse := model.WebResponse{}
		json.Unmarshal(responseBody, &webResponse)

		assert.Equal(t, fiber.StatusNotFound, webResponse.Code)
		assert.Equal(t, "NOT_FOUND", webResponse.Status)
	})
}

func TestFindByIdNote(t *testing.T) {
	var token string
	t.Run("login", func(t *testing.T) {
		data := model.LoginInput{
			Username: "test",
			Password: "test",
		}
		loginInput, _ := json.Marshal(data)
		// loginInput := strings.NewReader(`{"username:"test", "password":"test"}`)
		request := httptest.NewRequest(fiber.MethodPost, "/api/auth/login", bytes.NewBuffer(loginInput))
		request.Header.Add("content-type", "application/json")

		response, _ := app.Test(request)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)

		body, _ := io.ReadAll(response.Body)
		var responseBody model.WebResponse
		json.Unmarshal(body, &responseBody)

		token = responseBody.Data.(map[string]any)["token"].(string)
		assert.NotNil(t, token)
	})

	t.Run("test find by id note success", func(t *testing.T) {
		id := uuid.New().String()
		note := entity.Note{
			Id:    id,
			Title: "Test Create Data",
			Tags:  "tags1,tags2",
			Body:  "ini adalah test",
		}
		noteRepository.Create(context.TODO(), note)

		request := httptest.NewRequest("GET", "/api/notes/"+id, nil)
		request.Header.Set("Authorization", "Bearer "+token)

		response, _ := app.Test(request)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)

		responseBody, _ := io.ReadAll(response.Body)
		webResponse := model.WebResponse{}
		json.Unmarshal(responseBody, &webResponse)

		assert.Equal(t, fiber.StatusOK, webResponse.Code)
		assert.Equal(t, "OK", webResponse.Status)
		assert.Equal(t, "success get note", webResponse.Message)

		webResponseData := webResponse.Data.(map[string]any)
		assert.Equal(t, note.Title, webResponseData["title"].(string))
		assert.Equal(t, note.Body, webResponseData["body"].(string))

		noteTags := strings.Split(note.Tags, ",")
		for i := 0; i < len(noteTags); i++ {
			assert.Equal(t, noteTags[i], webResponseData["tags"].([]interface{})[i].(string))
		}

		noteRepository.Delete(context.TODO(), entity.Note{Id: id})
	})

	t.Run("test find by id note not found", func(t *testing.T) {
		id := uuid.New().String()

		request := httptest.NewRequest("GET", "/api/notes/"+id, nil)
		request.Header.Set("Authorization", "Bearer "+token)

		response, _ := app.Test(request)
		assert.Equal(t, fiber.StatusNotFound, response.StatusCode)

		responseBody, _ := io.ReadAll(response.Body)
		webResponse := model.WebResponse{}
		json.Unmarshal(responseBody, &webResponse)

		assert.Equal(t, fiber.StatusNotFound, webResponse.Code)
		assert.Equal(t, "NOT_FOUND", webResponse.Status)
	})
}

func TestFindAllNote(t *testing.T) {
	var token string
	t.Run("login", func(t *testing.T) {
		data := model.LoginInput{
			Username: "test",
			Password: "test",
		}
		loginInput, _ := json.Marshal(data)
		request := httptest.NewRequest(fiber.MethodPost, "/api/auth/login", bytes.NewBuffer(loginInput))
		request.Header.Add("content-type", "application/json")

		response, _ := app.Test(request)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)

		body, _ := io.ReadAll(response.Body)
		var responseBody model.WebResponse
		json.Unmarshal(body, &responseBody)

		token = responseBody.Data.(map[string]any)["token"].(string)
		assert.NotNil(t, token)
	})

	t.Run("test find all note", func(t *testing.T) {
		id1 := uuid.New().String()
		note1 := entity.Note{
			Id:    id1,
			Title: "Test Create Data",
			Tags:  "tags1,tags2",
			Body:  "ini adalah test",
		}
		noteRepository.Create(context.TODO(), note1)
		id2 := uuid.New().String()
		note2 := entity.Note{
			Id:    id2,
			Title: "Test Create Data",
			Tags:  "tags1,tags2",
			Body:  "ini adalah test",
		}
		noteRepository.Create(context.TODO(), note2)

		request := httptest.NewRequest("GET", "/api/notes", nil)
		request.Header.Set("Authorization", "Bearer "+token)

		response, _ := app.Test(request)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)

		responseBody, _ := io.ReadAll(response.Body)
		webResponse := model.WebResponse{}
		json.Unmarshal(responseBody, &webResponse)

		assert.Equal(t, fiber.StatusOK, webResponse.Code)
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

		noteRepository.Delete(context.TODO(), entity.Note{Id: id1})
		noteRepository.Delete(context.TODO(), entity.Note{Id: id2})
	})
}

func TestUnauthorized(t *testing.T) {
	id := uuid.New().String()

	request := httptest.NewRequest("GET", "/api/notes/"+id, nil)

	response, _ := app.Test(request)
	assert.Equal(t, fiber.StatusUnauthorized, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, fiber.StatusUnauthorized, webResponse.Code)
	assert.Equal(t, "UNAUTHORIZED", webResponse.Status)
}

func TestMethodNotAllowed(t *testing.T) {
	var token string
	t.Run("login", func(t *testing.T) {
		data := model.LoginInput{
			Username: "test",
			Password: "test",
		}
		loginInput, _ := json.Marshal(data)
		// loginInput := strings.NewReader(`{"username:"test", "password":"test"}`)
		request := httptest.NewRequest(fiber.MethodPost, "/api/auth/login", bytes.NewBuffer(loginInput))
		request.Header.Add("content-type", "application/json")

		response, _ := app.Test(request)
		assert.Equal(t, fiber.StatusOK, response.StatusCode)

		body, _ := io.ReadAll(response.Body)
		var responseBody model.WebResponse
		json.Unmarshal(body, &responseBody)

		token = responseBody.Data.(map[string]any)["token"].(string)
		assert.NotNil(t, token)
	})

	t.Run("test method not allowed", func(t *testing.T) {
		request := httptest.NewRequest("PUT", "/api/notes", nil)
		request.Header.Set("Authorization", "Bearer "+token)

		response, _ := app.Test(request)
		assert.Equal(t, fiber.StatusMethodNotAllowed, response.StatusCode)

		responseBody, _ := io.ReadAll(response.Body)
		webResponse := model.WebResponse{}
		json.Unmarshal(responseBody, &webResponse)

		assert.Equal(t, fiber.StatusMethodNotAllowed, webResponse.Code)
		assert.Equal(t, "METHOD_NOT_ALLOWED", webResponse.Status)
	})
}

func TestMain(m *testing.M) {
	db, _ := db.DB()
	db.Exec("TRUNCATE notes")
	db.Exec("TRUNCATE users")
	m.Run()
	db.Exec("TRUNCATE notes")
	db.Exec("TRUNCATE users")
}
