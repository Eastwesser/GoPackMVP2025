package main

import "fmt"

//Абстракция позволяет скрыть сложность и показать только необходимые детали.
//В Genshin Impact это можно представить как общий интерфейс для всех персонажей, которые могут использовать свои способности.

// Абстракция: интерфейс Character
type Character interface {
	UseSkill()
	UseBurst()
}

// Структура для персонажа Diluc
type Diluc struct {
	Name string
}

func (d Diluc) UseSkill() {
	fmt.Println(d.Name, "использует навык: Огненный взрыв!")
}

func (d Diluc) UseBurst() {
	fmt.Println(d.Name, "использует ульту: Феникс!")
}

// Структура для персонажа Venti
type Venti struct {
	Name string
}

func (v Venti) UseSkill() {
	fmt.Println(v.Name, "использует навык: Небесный выстрел!")
}

func (v Venti) UseBurst() {
	fmt.Println(v.Name, "использует ульту: Великий Одорадо!")
}

func main() {
	var character Character

	character = Diluc{Name: "Дилюк"}
	character.UseSkill()
	character.UseBurst()

	character = Venti{Name: "Венти"}
	character.UseSkill()
	character.UseBurst()
}
