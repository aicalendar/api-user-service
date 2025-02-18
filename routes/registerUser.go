package routes

import (
	"api-user-service/database"
	"api-user-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func RegisterUser(c *gin.Context) {
	// Create / modify table based on schema
	if err := database.DB.AutoMigrate(&User{}); err != nil {
		c.JSON(500, err.Error())
		return
	}

	// Binds request to newJSONUser variable
	var newJSONUser User
	if err := c.BindJSON(&newJSONUser); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := utils.UsernameValidation(newJSONUser.Name); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := utils.PasswordValidation(newJSONUser.Password); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Query for existing user
	userQuery := User{}
	queryResult := database.DB.Where("name = ?", newJSONUser.Name).Find(&userQuery)

	// Return on error
	if queryResult.Error != nil {
		c.JSON(500, queryResult.Error.Error())
		return
	}

	// If duplicate was found return
	if queryResult.RowsAffected > 0 {
		c.JSON(400, gin.H{
			"error": "User with this name already exists!",
		})
		return
	}

	// If no duplicate was found try to hash password
	if queryResult.RowsAffected == 0 {
		passwordHash, hashSalt, err := utils.HashPassword(newJSONUser.Password)

		// Return on error
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		// User to insert into database
		newDbUser := User{
			ID:        uuid.New().String(),
			Name:      newJSONUser.Name,
			Password:  passwordHash,
			Salt:      hashSalt,
			CreatedAt: time.Now(),
		}

		// Trying to insert into database
		dbCreateResult := database.DB.Create(&newDbUser)
		if dbCreateResult.Error != nil {
			c.JSON(500, dbCreateResult.Error.Error())
			return
		}

		// Return result
		if dbCreateResult.RowsAffected > 0 {
			c.JSON(201, gin.H{
				"status": "success",
				"error":  nil,
				"user":   newDbUser,
			})
			return
		}
	}
}
