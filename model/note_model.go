package model

import (
	"strings"
	"time"

	"github.com/ikhlashmulya/noteapp-resful-api/entity"
)

//model for request and response in service

type CreateNoteRequest struct {
	Title string   `validate:"required"`
	Tags  []string `validate:"required"`
	Body  string   `validate:"required"`
}

type CreateNoteResponse struct {
	Id string `json:"id"`
}

type UpdateNoteRequest struct {
	Id    string   `validate:"required"`
	Title string   `validate:"required"`
	Tags  []string `validate:"required"`
	Body  string   `validate:"required"`
}

type UpdateNoteResponse struct {
	Title string   `json:"title"`
	Tags  []string `json:"tags"`
	Body  string   `json:"body"`
}

func ToUpdateNoteResponse(note entity.Note) UpdateNoteResponse {
	return UpdateNoteResponse{
		Title: note.Title,
		Tags:  strings.Split(note.Tags, ","),
		Body:  note.Body,
	}
}

type FindNoteResponse struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tags      []string  `json:"tags"`
	Body      string    `json:"body"`
}

func ToFindNoteResponse(note entity.Note) FindNoteResponse {
	return FindNoteResponse{
		Id:        note.Id,
		Title:     note.Title,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
		Tags:      strings.Split(note.Tags, ","),
		Body:      note.Body,
	}
}
