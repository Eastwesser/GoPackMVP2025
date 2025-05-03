package main

import "fmt"

//Полиморфизм позволяет использовать разные типы через общий интерфейс.
//Например, персонажи могут использовать свои способности, но каждый делает это по-своему.

// Интерфейс для способностей
type SkillUser interface {
	UseSkill()
}

// Персонаж Дилюк
type Diluc struct{}

func (d Diluc) UseSkill() {
	fmt.Println("Дилюк использует пиро навык: Огненный взрыв!")
}

// Персонаж Венти
type Venti struct{}

func (v Venti) UseSkill() {
	fmt.Println("Венти использует анемо навык: Небесный выстрел!")
}

func main() {
	characters := []SkillUser{Diluc{}, Venti{}}

	for _, character := range characters {
		character.UseSkill()
	}
}
