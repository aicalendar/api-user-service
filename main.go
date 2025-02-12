package main

import (
	"api-user-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	SetupRoutes(r)

	r.Run(":8000")
}

func SetupRoutes(r *gin.Engine) {
	routeGroup := r.Group("/user-service")
	{
		routeGroup.GET("/test", routes.Test)
	}
}
