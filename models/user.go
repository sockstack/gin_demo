package models

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	Username string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
