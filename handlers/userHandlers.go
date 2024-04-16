package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kol1n2012/go.api/models"
)

func getUsers(c *gin.Context) {
	var users models.Users
	if err := c.ShouldBindWith(&users, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Успешно", "result": users})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error(), "result": "[]"})
	}
}

func getUser(c *gin.Context) {
}

func addUser(c *gin.Context) {
}

func deleteUser(c *gin.Context) {
}
