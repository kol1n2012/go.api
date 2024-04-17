package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kol1n2012/go.api/api"
	"github.com/kol1n2012/go.api/handlers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("/api/getUsers", handlers.GetUsers)
	r.GET("/api/getToken", api.GetBasicAuth, handlers.GetToken)

	authorized := r.Group("/api")

	authorized.Use(api.GetTokenAuth())
	{
		authorized.GET("/getUser/:id", handlers.GetUser)
		authorized.POST("/addUser", handlers.AddUser)
		authorized.DELETE("/deleteUser", handlers.DeleteUser)
	}

	r.Run(":8081")
}
