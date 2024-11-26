package models

import "github.com/google/uuid"

type Message struct {
	Base
	UserID   uuid.UUID
	Content  string
	IsEdited bool
}
