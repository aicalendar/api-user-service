package main

import (
	"api-user-service/routes"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	routeGroup := r.Group("/user-service")
	{
		routeGroup.GET("/test", routes.Test)
	}
}
