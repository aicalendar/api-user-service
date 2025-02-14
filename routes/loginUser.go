package routes

import (
	"api-user-service/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func LoginUser(c *gin.Context) {

	var JSONUserData User
	// Binds request to newJSONUser variable
	if err := c.BindJSON(&JSONUserData); err != nil {
		log.Error().Msg("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, "Failed to bind JSON")
		return
	}

	// Query for existing user based on type username
	userQuery := User{}
	queryResult := database.DB.Where(&User{Name: JSONUserData.Name}, "name").Find(&userQuery)
	if queryResult.Error != nil {
		log.Error().Msg("Error occurred while querying existing user!")
		c.JSON(http.StatusInternalServerError, "Error occurred while querying existing user!")
		return
	}

	if queryResult.RowsAffected == 0 {

		// If no user with matching name was found
		log.Warn().Msg("Invalid username!")
		c.JSON(http.StatusConflict, "Invalid username!")
		return

	} else if queryResult.RowsAffected > 0 {

		// If user was found
		log.Info().Msg("User found checking password!")
		// TODO: check password match, assign session token
	}
}
