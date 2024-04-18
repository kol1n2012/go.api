package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kol1n2012/go.api/api/sourse"
	"github.com/kol1n2012/go.api/models"
)

func GetUsers(c *gin.Context) {

	var users = models.Users{}

	status, message := sourse.SetCollectionFromFile(&users, "users")

	if err := c.ShouldBindWith(&users, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": status, "message": message, "result": users})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error(), "result": "[]"})
	}
}

func GetUser(c *gin.Context) {
}

func AddUser(c *gin.Context) {
}

func DeleteUser(c *gin.Context) {
}
