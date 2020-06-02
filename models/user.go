package models

type User struct {
	ID        uint `gorm:"primary_key"`
	Username string
	Password string
}
