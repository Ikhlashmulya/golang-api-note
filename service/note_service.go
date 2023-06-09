package service

import "github.com/ikhlashmulya/noteapp-resful-api/model"

// contract note service
type NoteService interface {
	Create(request model.CreateNoteRequest) (response model.CreateNoteResponse)
	Update(request model.UpdateNoteRequest) (response model.UpdateNoteResponse)
	Delete(noteId string)
	FindById(noteId string) (response model.FindNoteResponse)
	FindAll() (responses []model.FindNoteResponse)
}
