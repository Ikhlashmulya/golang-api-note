package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ikhlashmulya/golang-api-note/entity"
	"github.com/ikhlashmulya/golang-api-note/exception"
	"github.com/ikhlashmulya/golang-api-note/model"
	"github.com/ikhlashmulya/golang-api-note/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRespository repository.UserRespository
	SigningKey      []byte
}

func NewUserService(userRepository repository.UserRespository, signingKey []byte) *UserServiceImpl {
	return &UserServiceImpl{UserRespository: userRepository, SigningKey: signingKey}
}

func (service *UserServiceImpl) Login(ctx context.Context, input model.LoginInput) (response model.LoginResponse) {
	//get user
	user, err := service.UserRespository.GetUser(ctx, input.Username)
	exception.PanicIfErr(err)

	//compare password from user and from request
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	exception.PanicIfErr(err)

	//jwt token
	claims := jwt.MapClaims{
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//sign token
	signedToken, err := token.SignedString(service.SigningKey)
	exception.PanicIfErr(err)

	response.Token = signedToken
	return response
}

func (service *UserServiceImpl) Register(ctx context.Context, request model.RegisterRequest) model.RegisterResponse {
	//hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	exception.PanicIfErr(err)

	//make user
	user := entity.User{
		Name:     request.Name,
		Username: request.Password,
		Password: string(hashedPassword),
	}

	user = service.UserRespository.CreateUser(ctx, user)

	return model.ToRegisterResponse(user)
}
