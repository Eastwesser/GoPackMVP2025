package models

type Baristas *[]struct{}

type Barista struct {
	Name        string
	Age         int
	Temperament string
	Grade       string
	Skill       string
}

// Это реализовать через конструкторы = создание нового баристы
type Barista1 struct {
	Bob string
}

type Barista2 struct {
	Alice string
}
