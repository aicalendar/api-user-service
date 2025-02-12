package routes

import (
	"api-user-service/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
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
	// Create / modify table based on schema
	database.DB.AutoMigrate(&User{})

	var newJSONUser User
	// Binds request to newJSONUser variable
	if err := c.BindJSON(&newJSONUser); err != nil {
		log.Error().Msg("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, "Failed to bind JSON")
		return
	}

	// Query for existing user
	userQuery := User{}
	queryResult := database.DB.Where("name = ?", newJSONUser.Name).Find(&userQuery)
	if queryResult.Error != nil {
		log.Error().Msg("Error occurred while querying existing user!")
		c.JSON(http.StatusInternalServerError, "Error occurred while querying existing user!")
		return
	}
	if queryResult.RowsAffected > 0 {
		log.Warn().Msg("Cannot register user, username duplicate!")
		c.JSON(http.StatusConflict, "Cannot register user, username duplicate!")
	} else if queryResult.RowsAffected == 0 {
		log.Info().Msg("No duplicate found, registering new user!")

		// User to insert into database
		newDbUser := User{
			ID:        uuid.New().String(),
			Name:      newJSONUser.Name,
			Password:  newJSONUser.Password,
			CreatedAt: time.Now(),
		}

		// Trying to insert into database
		dbCreateResult := database.DB.Create(&newDbUser)
		if dbCreateResult.Error != nil {
			log.Error().Msg("Error occurred while inserting to database!")
			c.IndentedJSON(http.StatusInternalServerError, "Error occurred while inserting to database!")
			return
		}
		if dbCreateResult.RowsAffected > 0 {
			log.Info().Msg("Registered new user " + newDbUser.Name + " with id: " + newDbUser.ID)
			c.IndentedJSON(http.StatusCreated, newDbUser)
		}
	}
}
