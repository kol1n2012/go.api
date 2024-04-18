package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kol1n2012/go.api/api"
	"github.com/kol1n2012/go.api/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файл")
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	authorized := r.Group("/api", gin.BasicAuth(gin.Accounts{
		os.Getenv("API_LOGIN"): os.Getenv("API_PASSWORD"),
	}))

	//authorized.GET("/getToken", api.GetBasicAuth, handlers.GetToken)
	authorized.GET("/getToken", handlers.GetToken)

	authorized.Use(api.GetTokenAuth())
	{
		authorized.GET("/getUser/:id", handlers.GetUser)
		authorized.POST("/addUser", handlers.AddUser)
		authorized.DELETE("/deleteUser", handlers.DeleteUser)
	}

	r.GET("/api/getUsers", handlers.GetUsers)

	r.Run(":8081")
}
