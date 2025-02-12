package main

import (
	"api-user-service/database"
	"api-user-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Default gin engine instance
	r := gin.Default()

	// Connect to db
	database.ConnectToDB()

	// Setup all api endpoints
	SetupRoutes(r)

	// Run gin server
	r.Run(":80")
}

func SetupRoutes(r *gin.Engine) {
	routeGroup := r.Group("/api-user-service")
	{
		routeGroup.GET("/test", routes.Test)
		routeGroup.POST("/registerUser", routes.RegisterUser)
	}
}
