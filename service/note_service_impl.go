package service

import (
	"context"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ikhlashmulya/golang-api-note/entity"
	"github.com/ikhlashmulya/golang-api-note/exception"
	"github.com/ikhlashmulya/golang-api-note/model"
	"github.com/ikhlashmulya/golang-api-note/repository"
)

// note service implementation
type NoteServiceImpl struct {
	noteRepository repository.NoteRepository
	validate       *validator.Validate
}

func NewNoteService(repository repository.NoteRepository, validate *validator.Validate) NoteService {
	return &NoteServiceImpl{repository, validate}
}

func (service *NoteServiceImpl) Create(ctx context.Context, request model.CreateNoteRequest) (response model.CreateNoteResponse) {
	//validation
	err := service.validate.Struct(request)
	exception.PanicIfErr(err)

	//create data
	note := entity.Note{
		Id:        uuid.New().String(),
		Title:     request.Title,
		Tags:      strings.Join(request.Tags, ","),
		Body:      request.Body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	//create data to repository
	service.noteRepository.Create(ctx, note)

	response.Id = note.Id
	return response
}

func (service *NoteServiceImpl) Update(ctx context.Context, request model.UpdateNoteRequest) (response model.UpdateNoteResponse) {
	//validation
	err := service.validate.Struct(request)
	exception.PanicIfErr(err)

	//data validation in repository
	note, err := service.noteRepository.FindById(ctx, request.Id)
	//panic if there is no data in database
	exception.PanicIfErr(err)

	// update data
	note.Id = request.Id
	note.Title = request.Title
	note.Body = request.Body
	note.Tags = strings.Join(request.Tags, ",")
	note.UpdatedAt = time.Now()

	// update data to repository
	service.noteRepository.Update(ctx, note)

	//create data response
	response = model.ToUpdateNoteResponse(note)

	return response
}

func (service *NoteServiceImpl) Delete(ctx context.Context, noteId string) {
	//data validation in repository
	note, err := service.noteRepository.FindById(ctx, noteId)
	//panic if there is no data in database
	exception.PanicIfErr(err)

	//delete data
	service.noteRepository.Delete(ctx, note)
}

func (service *NoteServiceImpl) FindById(ctx context.Context, noteId string) (response model.FindNoteResponse) {
	//find data in repository
	note, err := service.noteRepository.FindById(ctx, noteId)
	exception.PanicIfErr(err)

	//create data response
	response = model.ToFindNoteResponse(note)

	return response
}

func (service *NoteServiceImpl) FindAll(ctx context.Context) (responses []model.FindNoteResponse) {
	//find all data in repository
	notes := service.noteRepository.FindAll(ctx)

	//data responses
	for _, note := range notes {
		responses = append(responses, model.ToFindNoteResponse(note))
	}

	return responses
}
