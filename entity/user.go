package entity

type User struct {
	Id       int `gorm:"primaryKey;column:id"`
	Name     string
	Username string
	Password string
}
