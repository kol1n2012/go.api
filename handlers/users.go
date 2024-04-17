package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kol1n2012/go.api/models"
	log "github.com/sirupsen/logrus"
)

func GetUsers(c *gin.Context) {

	var users = models.Users{}

	setUsersFromFile(&users, "users.json")

	if err := c.ShouldBindWith(&users, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Успешно", "result": users})
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

func setUsersFromFile(u *models.Users, file string) {

	pwd, _ := os.Getwd()
	// Чтение содержимого файла
	fileData, err := os.ReadFile(pwd + string(os.PathSeparator) + file)

	if err != nil {
		log.Fatal("ERR_FILE. Ошибка чтения файла: ", err)
	}

	err = nil

	err = json.Unmarshal([]byte(string(fileData)), &u)

	if err != nil {
		log.Fatal("ERR_JSON. ошибка распознавания json: ", err)
	}
}
