package config

import (
	"context"
	"fmt"
	"time"

	"github.com/ikhlashmulya/golang-api-note/entity"
	"github.com/ikhlashmulya/golang-api-note/exception"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// gorm configuration

func NewGormDB(configuration Config) *gorm.DB {
	// data source name
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Local",
		configuration.Get("DB_USERNAME"),
		configuration.Get("DB_PASSWORD"),
		configuration.Get("DB_HOST"),
		configuration.Get("DB_PORT"),
		configuration.Get("DB_NAME"),
	)

	//open connnection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	exception.PanicIfErr(err)

	//set connection pool
	sqlDB, err := db.DB()
	exception.PanicIfErr(err)

	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(50 * time.Minute)
	sqlDB.SetMaxOpenConns(100)

	//migration table
	db.AutoMigrate(&entity.Note{})

	return db
}

// context
func NewGormDBContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	return ctx, cancel
}
