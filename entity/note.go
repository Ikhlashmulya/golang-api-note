package entity

import "time"

//entity / representation data table in database
type Note struct {
	Id        string `gorm:"primaryKey;column:id"`
	Title     string 
	Tags      string 
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
