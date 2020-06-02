package dto

import "time"

type UserDto struct {
	ID        uint
	Username string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
