package repository

import (
	"context"

	"github.com/ikhlashmulya/golang-api-note/entity"
)

type UserRespository interface {
	CreateUser(ctx context.Context, user entity.User) entity.User
	GetUser(ctx context.Context, username string) (response entity.User, err error)
}
