package main

import "fmt"

//В Go нет классического наследования, но можно использовать встраивание (embedding).
//Например, базовый класс "Хакер" и специализированный класс "Элитный хакер".

// Базовый класс
type Hacker struct {
	Name string
}

func (h Hacker) Attack() {
	fmt.Println(h.Name, "атакует!")
}

// Специализированный класс
type EliteHacker struct {
	Hacker
	SpecialSkill string
}

func (e EliteHacker) UseSpecialSkill() {
	fmt.Println(e.Name, "использует специальный навык:", e.SpecialSkill)
}

func main() {
	aiden := EliteHacker{
		Hacker:       Hacker{Name: "Эйден Пирс"},
		SpecialSkill: "Взлом инфраструктуры",
	}

	aiden.Attack()
	aiden.UseSpecialSkill()
}
