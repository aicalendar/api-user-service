package routes

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
	CreatedAt time.Time `json:"createdAt"`
}

type Session struct {
	Token      string  `json:"token"`
	Expiration float64 `json:"expiration"`
}
