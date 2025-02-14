package routes

import (
	"api-user-service/database"
	"api-user-service/passwords"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {

	// Binds request to newJSONUser variable
	var JSONUserData User
	if err := c.BindJSON(&JSONUserData); err != nil {
		c.JSON(http.StatusBadRequest, "Failed to bind JSON")
		return
	}

	// Query for existing user based on type username
	userQuery := User{}
	queryResult := database.DB.Where(&User{Name: JSONUserData.Name}, "name").Find(&userQuery)
	if queryResult.Error != nil {
		c.JSON(http.StatusInternalServerError, "Error occurred while querying existing user!")
		return
	}

	if queryResult.RowsAffected == 0 {
		// If no user with matching name was found
		c.JSON(http.StatusConflict, "Invalid username!")
		return

	} else if queryResult.RowsAffected > 0 {
		// If user was found compare passwords
		passwordMatch := passwords.ComparePasswords(JSONUserData.PasswordHash, userQuery.PasswordHash, userQuery.HashSalt)
		if passwordMatch {
			c.JSON(http.StatusOK, "Password matches!")
			return
		} else {
			c.JSON(http.StatusConflict, "Password doesn't match!")
			return
		}

	}
}
