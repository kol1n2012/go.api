package models

import (
	"encoding/json"
	"fmt"
	"os"
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

func New() *Users {
	users := &Users{}
	users.SetCollection()
	return users
}

func (c *Users) GetCollection() []User {
	return c.collection
}

func (c *Users) SetCollection() {

	pwd, _ := os.Getwd()
	// Чтение содержимого файла
	fileData, err := os.ReadFile(pwd + string(os.PathSeparator) + "users" + ".json")
	if err != nil {
		fmt.Println(err)
	}

	_ = json.Unmarshal([]byte(string(fileData)), &c.collection)
}
