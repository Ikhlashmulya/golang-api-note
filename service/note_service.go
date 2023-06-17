package service

import (
	"context"

	"github.com/ikhlashmulya/golang-api-note/model"
)

// contract note service
type NoteService interface {
	Create(ctx context.Context, request model.CreateNoteRequest) (response model.CreateNoteResponse)
	Update(ctx context.Context, request model.UpdateNoteRequest) (response model.UpdateNoteResponse)
	Delete(ctx context.Context, noteId string)
	FindById(ctx context.Context, noteId string) (response model.FindNoteResponse)
	FindAll(ctx context.Context) (responses []model.FindNoteResponse)
}
