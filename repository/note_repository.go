package repository

import (
	"github.com/ikhlashmulya/golang-api-note/entity"
)

// contract note repository
type NoteRepository interface {
	Create(note entity.Note)
	Update(note entity.Note)
	Delete(note entity.Note)
	FindById(noteId string) (note entity.Note, err error)
	FindAll() (notes []entity.Note)
}
