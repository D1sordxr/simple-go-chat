package dto

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
