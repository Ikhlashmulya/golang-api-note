package service

import (
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ikhlashmulya/noteapp-resful-api/entity"
	"github.com/ikhlashmulya/noteapp-resful-api/exception"
	"github.com/ikhlashmulya/noteapp-resful-api/model"
	"github.com/ikhlashmulya/noteapp-resful-api/repository"
)

// note service implementation
type NoteServiceImpl struct {
	noteRepository repository.NoteRepository
	validate       *validator.Validate
}

func NewNoteService(repository repository.NoteRepository, validate *validator.Validate) NoteService {
	return &NoteServiceImpl{repository, validate}
}

func (service *NoteServiceImpl) Create(request model.CreateNoteRequest) (response model.CreateNoteResponse) {
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
	service.noteRepository.Create(note)

	response.Id = note.Id
	return response
}

func (service *NoteServiceImpl) Update(request model.UpdateNoteRequest) (response model.UpdateNoteResponse) {
	//validation
	err := service.validate.Struct(request)
	exception.PanicIfErr(err)

	//data validation in repository
	note, err := service.noteRepository.FindById(request.Id)
	//panic if there is no data in database
	exception.PanicIfErr(err)

	// update data
	note.Id = request.Id
	note.Title = request.Title
	note.Body = request.Body
	note.Tags = strings.Join(request.Tags, ",")
	note.UpdatedAt = time.Now()

	// update data to repository
	service.noteRepository.Update(note)

	//create data response
	response = model.ToUpdateNoteResponse(note)

	return response
}

func (service *NoteServiceImpl) Delete(noteId string) {
	//data validation in repository
	note, err := service.noteRepository.FindById(noteId)
	//panic if there is no data in database
	exception.PanicIfErr(err)

	//delete data
	service.noteRepository.Delete(note)
}

func (service *NoteServiceImpl) FindById(noteId string) (response model.FindNoteResponse) {
	//find data in repository
	note, err := service.noteRepository.FindById(noteId)
	exception.PanicIfErr(err)

	//create data response
	response = model.ToFindNoteResponse(note)

	return response
}

func (service *NoteServiceImpl) FindAll() (responses []model.FindNoteResponse) {
	//find all data in repository
	notes := service.noteRepository.FindAll()

	//data responses
	for _, note := range notes {
		responses = append(responses, model.ToFindNoteResponse(note))
	}

	return responses
}
