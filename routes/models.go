package routes

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
	CreatedAt time.Time `json:"createdAt"`
}

type Session struct {
	Token      string  `json:"token"`
	Expiration float64 `json:"expiration"`
}
