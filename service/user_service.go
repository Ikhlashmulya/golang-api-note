package service

import (
	"context"

	"github.com/ikhlashmulya/golang-api-note/model"
)

// contract note service
type UserService interface {
	Login(ctx context.Context, input model.LoginInput) (response model.LoginResponse)
	Register(ctx context.Context, request model.RegisterRequest) model.RegisterResponse
}