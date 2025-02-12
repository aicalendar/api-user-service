package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `gorm:"unique" json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

func RegisterUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		log.Println("Failed to bind JSON")
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}
