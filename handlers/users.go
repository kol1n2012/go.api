package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kol1n2012/go.api/models"
)

func GetUsers(c *gin.Context) {

	var params = new(models.Params)

	var users = models.NewUsers(params)

	if err := c.ShouldBindWith(&users, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "успех", "result": users.GetCollection()})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error(), "result": "[]"})
	}
}

func GetUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var params = new(models.Params)

	params.Filter = map[string]any{
		"id": id,
	}

	params.Limit = 1

	var users = models.NewUsers(params)

	if err := c.ShouldBindWith(&users, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "успех", "result": users.GetCollection()})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error(), "result": "[]"})
	}
}

func AddUser(c *gin.Context) {
}

func DeleteUser(c *gin.Context) {
}
