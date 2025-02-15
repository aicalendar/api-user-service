package main

import (
	"api-user-service/database"
	"api-user-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	// Default gin engine instance
	r := gin.Default()

	// Connect to db
	database.ConnectToDB()

	// Connect to redis
	database.ConnectToRedis()

	// Setup all api endpoints
	SetupRoutes(r)

	// Run gin server
	if err := r.Run(":80"); err != nil {
		log.Panic().Msg("PANICIG failed to start gin server")
	}
}

func SetupRoutes(r *gin.Engine) {
	routeGroup := r.Group("/api-user-service")
	{
		routeGroup.GET("/test", routes.Test)
		routeGroup.POST("/registerUser", routes.RegisterUser)
		routeGroup.POST("/loginUser", routes.LoginUser)
	}
}
