package main

import (
	"log"

	"/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/api/getUsers", handlers.getUsers)
	r.GET("/api/getToken", basicAuth, getToken)

	authorized := r.Group("/api")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(AuthRequired())
	{
		authorized.GET("/getUser/:id", handlers.getUser)
		authorized.POST("/addUser", handlers.addUser)
		authorized.DELETE("/deleteUser", handlers.deleteUser)
	}

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func basicAuth(c *gin.Context) {
	// Get the Basic Authentication credentials
	user, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && user == "testuser" && password == "testpass" {
		log.WithFields(log.Fields{
			"user": user,
		}).Info("User authenticated")
	} else {
		c.Abort()
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		return
	}
}

// Authenticate User
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Get token and e-mail from header
		token := c.Request.Header.Get("AuthToken")
		email := c.Request.Header.Get("AuthEmail")

		//check to see if email & token were provided
		if len(token) == 0 || len(email) == 0 {
		}
		//Find email in database
		//Compare stored token with token provided in header
		//Return - Authentication was success or fail
	}
}

func getToken(c *gin.Context) {
}
