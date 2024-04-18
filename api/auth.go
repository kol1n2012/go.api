package api

import (
	"github.com/gin-gonic/gin"
	//log "github.com/sirupsen/logrus"
)

// func GetBasicAuth(c *gin.Context) {
// 	// Get the Basic Authentication credentials
// 	user, password, hasAuth := c.Request.BasicAuth()
// 	if hasAuth && user == "testuser" && password == "testpass" {
// 		log.WithFields(log.Fields{
// 			"user": user,
// 		}).Info("User authenticated")
// 	} else {
// 		c.Abort()
// 		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
// 		return
// 	}
// }

// Authenticate User
func GetTokenAuth() gin.HandlerFunc {
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
