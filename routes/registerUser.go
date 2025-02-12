package routes

import (
	"api-user-service/database"
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
	var newJSONUser User

	if err := c.BindJSON(&newJSONUser); err != nil {
		log.Println("Failed to bind JSON")
		return
	}

	newDbUser := User{
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
