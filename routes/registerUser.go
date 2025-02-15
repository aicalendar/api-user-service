package routes

import (
	"api-user-service/database"
	"api-user-service/passwords"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func RegisterUser(c *gin.Context) {
	// Create / modify table based on schema
	if err := database.DB.AutoMigrate(&User{}); err != nil {
		log.Error().Err(err)
	}

	// Binds request to newJSONUser variable
	var newJSONUser User
	if err := c.BindJSON(&newJSONUser); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Query for existing user
	userQuery := User{}
	queryResult := database.DB.Where("name = ?", newJSONUser.Name).Find(&userQuery)
	if queryResult.Error != nil {
		c.JSON(http.StatusInternalServerError, queryResult.Error)
		return
	}
	if queryResult.RowsAffected > 0 {
		// If duplicate was found return
		c.JSON(http.StatusConflict, "Cannot register user, username duplicate!")
		return
	} else if queryResult.RowsAffected == 0 {
		// If no duplicate was found try to hash password
		passwordHash, hashSalt, err := passwords.HashPassword(newJSONUser.PasswordHash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		// User to insert into database
		newDbUser := User{
			ID:           uuid.New().String(),
			Name:         newJSONUser.Name,
			PasswordHash: passwordHash,
			HashSalt:     hashSalt,
			CreatedAt:    time.Now(),
		}

		// Trying to insert into database
		dbCreateResult := database.DB.Create(&newDbUser)
		if dbCreateResult.Error != nil {
			c.JSON(http.StatusInternalServerError, "Error occurred while inserting to database!")
			return
		}
		// Return result
		if dbCreateResult.RowsAffected > 0 {
			c.JSON(http.StatusCreated, "Registered new user "+newDbUser.Name+" with id: "+newDbUser.ID)
			return
		}
	}
}
