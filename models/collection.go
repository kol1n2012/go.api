package models

//That is OOP? where is my abstract methods. WTF :)

type Collections interface {
	SetCollection()
	GetCollection()
}

type Collection struct {
	Collections
}

type Params struct {
	Filter map[string]any
	Order  map[string]string
	Limit  int
	Select []string
}
