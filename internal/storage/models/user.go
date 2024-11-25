package models

import "github.com/google/uuid"

type User struct {
	Base
	UserID   uuid.UUID
	Username string
}
