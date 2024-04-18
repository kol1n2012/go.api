package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kol1n2012/go.api/models"
	log "github.com/sirupsen/logrus"
)

type Token struct {
	Token string `json:"token"`
}

func GetToken(c *gin.Context) {

	result := models.Token{}

	result.Token = GenerateJWT()

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Успешно", "result": result})
}

func GenerateJWT() string {

	// Структура пользовательских утверждений
	type MyClaims struct {
		jwt.RegisteredClaims
		Username string `json:"username"`
		Admin    bool   `json:"admin"`
	}

	// Создание токена с пользовательскими утверждениями
	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{},
		Username:         os.Getenv("API_LOGIN"),
		Admin:            true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Секретный ключ для подписи
	var mySigningKey = []byte(os.Getenv("API_TOKEN"))

	// Генерация токена в строковом формате
	strToken, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Fatalf("Произошла ошибка: %v", err)
	}

	return strToken

}
