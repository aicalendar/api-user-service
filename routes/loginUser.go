package routes

import (
	"api-user-service/database"
	"api-user-service/passwords"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

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

	// Return on error
	if queryResult.Error != nil {
		c.JSON(http.StatusInternalServerError, "Error occurred while querying existing user!")
		return
	}

	// If no user with matching name was found
	if queryResult.RowsAffected == 0 {
		c.JSON(http.StatusConflict, "Invalid username!")
		return
	}

	// If user was found compare passwords
	if queryResult.RowsAffected > 0 {
		passwordMatch, err := passwords.ComparePasswords(JSONUserData.PasswordHash, userQuery.PasswordHash, userQuery.HashSalt)

		// Return on error
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		// Return if passwords doesn't match
		if !passwordMatch {
			c.JSON(http.StatusConflict, "Password doesn't match!")
			return
		}

		// TODO: generate session token
		if passwordMatch {

			// Generate random session token
			salt, err := passwords.GenerateSalt(32)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				return
			}

			// Preapare variables for redis
			sessionToken := base64.StdEncoding.EncodeToString(salt)
			key := fmt.Sprintf("user:%s:sessions", userQuery.ID)
			expiration := 24 * time.Hour

			// Insert session token in redis Hash
			if err := database.REDIS.SetEx(c, key+":"+sessionToken, sessionToken, expiration).Err(); err != nil {
				c.JSON(http.StatusInternalServerError, err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"session token": sessionToken,
			})
			return
		}
	}
}
