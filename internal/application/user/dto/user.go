package dto

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        int       `json:"id,omitempty"`
	UserID    uuid.UUID `json:"user_id,omitempty"`
	Username  string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
