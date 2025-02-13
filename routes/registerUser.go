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

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	PasswordHash string    `json:"passwordHash"`
	HashSalt     string    `json:"passwordSalt"`
	CreatedAt    time.Time `json:"createdAt"`
}

func RegisterUser(c *gin.Context) {
	// Create / modify table based on schema
	if err := database.DB.AutoMigrate(&User{}); err != nil {
		log.Error().Msg("Error occured while migrating scheme")
	}

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
		return
	} else if queryResult.RowsAffected == 0 {
		log.Info().Msg("No duplicate found, registering new user!")

		passwordHash, hashSalt, err := passwords.HashPassword(newJSONUser.PasswordHash)
		if err != nil {
			log.Error().Msg("Failed hashing password!")
			c.JSON(http.StatusInternalServerError, "Failed hashing password!")
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
			log.Error().Msg("Error occurred while inserting to database!")
			c.JSON(http.StatusInternalServerError, "Error occurred while inserting to database!")
			return
		}
		if dbCreateResult.RowsAffected > 0 {
			log.Info().Msg("Registered new user " + newDbUser.Name + " with id: " + newDbUser.ID)
			c.JSON(http.StatusCreated, "Registered new user "+newDbUser.Name+" with id: "+newDbUser.ID)
			return
		}
	}
}
