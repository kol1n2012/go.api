package models

import "time"

type Users struct {
	Collection []User `json:"collection"`
}

type User struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
}
