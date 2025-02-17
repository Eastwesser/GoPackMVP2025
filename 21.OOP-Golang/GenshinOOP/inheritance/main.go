package main

import "fmt"

//В Go нет классического наследования, но можно использовать встраивание (embedding).
//Например, базовый класс "Персонаж" и специализированный класс "Пиро персонаж".

// Базовый класс
type Character struct {
	Name string
}

func (c Character) Attack() {
	fmt.Println(c.Name, "атакует!")
}

// Специализированный класс
type PyroCharacter struct {
	Character
	Element string
}

func (p PyroCharacter) UsePyroSkill() {
	fmt.Println(p.Name, "использует пиро навык!")
}

func main() {
	diluc := PyroCharacter{
		Character: Character{Name: "Дилюк"},
		Element:   "Пиро",
	}

	diluc.Attack()
	diluc.UsePyroSkill()
}
