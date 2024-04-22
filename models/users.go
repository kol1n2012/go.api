package models

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
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

}
