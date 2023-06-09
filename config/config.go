package config

import (
	"os"

	"github.com/ikhlashmulya/golang-api-note/exception"
	"github.com/joho/godotenv"
)

//configuration .env

type Config interface {
	Get(key string) string
}

type ConfigImpl struct {
}

func (config *ConfigImpl) Get(key string) string {
	return os.Getenv(key)
}

func NewConfig(fileNames ...string) Config {
	err := godotenv.Load(fileNames...)
	exception.PanicIfErr(err)

	return &ConfigImpl{}
}
