package repository

import (
	"context"

	"github.com/ikhlashmulya/golang-api-note/entity"
)

// contract note repository
type NoteRepository interface {
	Create(ctx context.Context, note entity.Note)
	Update(ctx context.Context, note entity.Note)
	Delete(ctx context.Context, note entity.Note)
	FindById(ctx context.Context, noteId string) (note entity.Note, err error)
	FindAll(ctx context.Context) (notes []entity.Note)
}
