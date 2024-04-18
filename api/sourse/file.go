package sourse

import (
	"encoding/json"
	"os"

	"github.com/kol1n2012/go.api/models"
	log "github.com/sirupsen/logrus"
)

func SetCollectionFromFile(u *models.Users, file string) (bool, string) {
	message := "Успешно"
	status := true

	pwd, _ := os.Getwd()
	// Чтение содержимого файла
	fileData, err := os.ReadFile(pwd + string(os.PathSeparator) + file + ".json")

	if err != nil {
		message = "Ошибка чтения файла"
		status = false

		log.WithFields(log.Fields{
			"message": "ERR_FILE. " + message,
		}).Error("User list")
	} else {
		err = nil

		err = json.Unmarshal([]byte(string(fileData)), &u)

		if err != nil {
			message = "Ошибка распознавания json"
			status = false

			log.WithFields(log.Fields{
				"message": "ERR_JSON. " + message,
			}).Error("User list")
		}
	}

	return status, message
}
