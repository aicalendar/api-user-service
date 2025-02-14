package routes

import (
	"time"
)

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	PasswordHash string    `json:"passwordHash"`
	HashSalt     string    `json:"passwordSalt"`
	CreatedAt    time.Time `json:"createdAt"`
}
