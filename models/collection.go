package models

//That is OOP? where is my abstract methods. WTF :)

type Collections interface {
	SetCollection()
	GetCollection()
}

type Collection struct {
	Collections
}
