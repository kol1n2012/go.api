package models

import (
	"encoding/json"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Users struct {
	Collection
	collection []User
}

func NewUser() *Users {

	users := &Users{}

	var sourse = "users"

	switch os.Getenv("DATA_SOURSE_DRIVER") {
	case "file":
		users.SetCollectionFromFile(sourse)
	case "mysql":
		users.SetCollectionFromMysql(sourse)
	}
	return users
}

func (c *Users) GetCollection() []User {
	return c.collection
}

func (c *Users) SetCollectionFromFile(sourses string) {

	if len(sourses) == 0 {
		return
	}

	pwd, _ := os.Getwd()

	// Чтение содержимого файла
	fileData, err := os.ReadFile(pwd + string(os.PathSeparator) + string(sourses) + ".json")
	if err != nil {
		log.WithFields(log.Fields{
			"message": "ERR_FILE. Ошибка чтения файла",
		}).Error("User list")
	} else {
		err = nil

		err = json.Unmarshal([]byte(string(fileData)), &c.collection)

		if err != nil {
			log.WithFields(log.Fields{
				"message": "ERR_JSON. Ошибка распознавания json",
			}).Error("User list")
		}
	}
}

func (c *Users) SetCollectionFromMysql(sourses string) {

	if len(sourses) == 0 {
		return
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_LOGIN"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.WithFields(log.Fields{
			"message": "ERR_MYSQL. Ошибка подключения к Базе данных",
		}).Error("User list")
	} else {
		db.Find(&c.collection)
	}
}
