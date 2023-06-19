package service

import (
	"context"

	"github.com/ikhlashmulya/golang-api-note/model"
)

type UserService interface {
	Login(ctx context.Context, input model.LoginInput) string
	Register(ctx context.Context, request model.RegisterRequest) model.RegisterResponse
}