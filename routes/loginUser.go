package routes

import (
	"api-user-service/database"
	"api-user-service/utils"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {

	// Binds request to newJSONUser variable
	var JSONUserData User
	if err := c.BindJSON(&JSONUserData); err != nil {
		c.JSON(400, err)
		return
	}

	// Query for existing user based on type username
	userQuery := User{}
	queryResult := database.DB.Where(&User{Name: JSONUserData.Name}, "name").Find(&userQuery)

	// Return on error
	if queryResult.Error != nil {
		c.JSON(500, queryResult.Error)
		return
	}

	// If no user with matching name was found
	if queryResult.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"error": "Wrong username!",
		})
		return
	}

	// If user was found compare passwords
	if queryResult.RowsAffected > 0 {
		passwordMatch, err := utils.ComparePasswords(JSONUserData.PasswordHash, userQuery.PasswordHash, userQuery.HashSalt)

		// Return on error
		if err != nil {
			c.JSON(500, err)
			return
		}

		// Return if passwords doesn't match
		if !passwordMatch {
			c.JSON(400, gin.H{
				"error": "Wrong password!",
			})
			return
		}

		if passwordMatch {

			// Generate random session token
			salt, err := utils.GenerateSalt(32)
			if err != nil {
				c.JSON(500, err)
				return
			}

			// Preapare variables for redis
			sessionToken := base64.StdEncoding.EncodeToString(salt)
			key := fmt.Sprintf("user:%s:sessions", userQuery.ID)
			expiration := 24 * time.Hour

			// Insert session token in redis Hash
			if err := database.REDIS.SetEx(c, key+":"+sessionToken, sessionToken, expiration).Err(); err != nil {
				c.JSON(500, err)
				return
			}

			// Return session id
			c.JSON(200, gin.H{
				"sessionToken": sessionToken,
			})
			return
		}
	}
}
