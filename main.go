package main

import (
	"api-user-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	SetupRoutes(r)

	r.Run(":80")
}

func SetupRoutes(r *gin.Engine) {
	routeGroup := r.Group("/api-user-service")
	{
		routeGroup.GET("/test", routes.Test)
	}
}
