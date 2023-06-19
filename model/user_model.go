package model

import "github.com/ikhlashmulya/golang-api-note/entity"

type LoginInput struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Name     string `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type RegisterResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

func ToRegisterResponse(user entity.User) RegisterResponse {
	return RegisterResponse{
		Name:     user.Name,
		Username: user.Username,
	}
}
