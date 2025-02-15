package main

import (
	"api-user-service/database"
	"api-user-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	var err error

	// Connect to db
	if err = database.ConnectToDB(); err != nil {
		log.Panic().Err(err)
	} else {
		log.Info().Msg("Successfully connected to database!")
	}

	// Connect to redis
	if err = database.ConnectToRedis(); err != nil {
		log.Panic().Err(err)
	} else {
		log.Info().Msg("Successfully connected to redis!")
	}

	// Default gin engine instance
	r := gin.Default()

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
