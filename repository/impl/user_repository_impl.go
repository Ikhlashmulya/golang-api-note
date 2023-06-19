package repository

import (
	"context"

	"github.com/ikhlashmulya/golang-api-note/entity"
	"github.com/ikhlashmulya/golang-api-note/exception"
	"gorm.io/gorm"
)

type UserRespositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRespositoryImpl {
	return &UserRespositoryImpl{DB: db}
}

func (repository *UserRespositoryImpl) CreateUser(ctx context.Context, user entity.User) entity.User {
	err := repository.DB.WithContext(ctx).Create(&user).Error
	exception.PanicIfErr(err)
	return user
}

func (repository *UserRespositoryImpl) GetUser(ctx context.Context, username string) (response entity.User, err error) {
	err = repository.DB.WithContext(ctx).First(&response, "username = ?", username).Error
	return response, err
}
