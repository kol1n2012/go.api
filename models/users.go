package models

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

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
	params     Params
}

func NewUsers(params *Params) *Users {

	users := &Users{}

	var sourse = "users"

	users.setLimit(params.Limit)

	users.setFilter(params.Filter)

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

	f := c.params.Filter

	l := c.params.Limit

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

		if len(f) > 0 {

			filtered := []User{}

			for _, user := range c.collection {

				result := false

				for k, v := range f {
					var reflectValue reflect.Value = reflect.ValueOf(&user)

					if reflect.Indirect(reflectValue).FieldByName(k).Interface() == v {
						result = true
					}
				}

				if result {
					filtered = append(filtered, user)
				}
			}

			c.collection = filtered
		}

		if l > 0 && len(c.collection) > 0 {
			c.collection = c.collection[:l]
		}

		if err != nil {
			log.WithFields(log.Fields{
				"message": "ERR_JSON. Ошибка распознавания json",
			}).Error("User list")
		}
	}
}

func (c *Users) setFilter(filter map[string]any) {
	c.params.Filter = filter
}

func (c *Users) setLimit(limit int) {
	c.params.Limit = limit
}

func (c *Users) SetCollectionFromMysql(sourses string) {

	if len(sourses) == 0 {
		return
	}

	f := c.params.Filter

	l := c.params.Limit

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_LOGIN"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.Model(&User{})

	if l > 0 {
		db.Limit(l)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"message": "ERR_MYSQL. Ошибка подключения к Базе данных",
		}).Error("User list")
	} else {
		if len(f) > 0 {

			filter := map[string]interface{}{}

			for k, v := range f {
				filter[k] = v
			}

			db.Where(filter).Find(&c.collection)
		} else {
			db.Find(&c.collection)
		}

	}
}
