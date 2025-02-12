package routes

import (
	"api-user-service/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

func RegisterUser(c *gin.Context) {
	var newJSONUser User

	if err := c.BindJSON(&newJSONUser); err != nil {
		log.Println("Failed to bind JSON")
		return
	}

	newDbUser := User{
		ID:        uuid.New().String(),
		Name:      newJSONUser.Name,
		Password:  newJSONUser.Password,
		CreatedAt: time.Now(),
	}

	dbCreateResult := database.DB.Create(&newDbUser)
	if dbCreateResult.Error != nil {
		log.Println("Error occurred while inserting to database!")
		return
	}

	if dbCreateResult.RowsAffected > 0 {
		log.Println("Successfully register new user!")
	}

	c.IndentedJSON(http.StatusCreated, newDbUser)
}
