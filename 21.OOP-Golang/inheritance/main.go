package main

import "fmt"

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
