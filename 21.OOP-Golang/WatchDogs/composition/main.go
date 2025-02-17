package main

import "fmt"

//Композиция позволяет создавать сложные объекты, комбинируя более простые.
//Например, хакер может иметь гаджеты и оружие.

// Гаджет
type Gadget struct {
	Name  string
	Power int
}

// Оружие
type Weapon struct {
	Name   string
	Damage int
}

// Хакер
type Hacker struct {
	Name   string
	Gadget Gadget
	Weapon Weapon
}

func main() {
	aiden := Hacker{
		Name: "Эйден Пирс",
		Gadget: Gadget{
			Name:  "Смартфон",
			Power: 100,
		},
		Weapon: Weapon{
			Name:   "Пистолет",
			Damage: 50,
		},
	}

	fmt.Println("Хакер:", aiden.Name)
	fmt.Println("Гаджет:", aiden.Gadget.Name, "с мощностью", aiden.Gadget.Power)
	fmt.Println("Оружие:", aiden.Weapon.Name, "с уроном", aiden.Weapon.Damage)
}
