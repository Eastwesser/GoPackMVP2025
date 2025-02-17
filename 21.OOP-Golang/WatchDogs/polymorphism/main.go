package main

import "fmt"

//Полиморфизм позволяет использовать разные типы через общий интерфейс.
//Например, хакеры могут использовать свои навыки, но каждый делает это по-своему.

// Интерфейс для навыков
type SkillUser interface {
	UseSkill()
}

// Хакер Эйден Пирс
type AidenPearce struct{}

func (a AidenPearce) UseSkill() {
	fmt.Println("Эйден Пирс использует навык: Взлом камеры!")
}

// Хакер Рэймонд Кенни
type RaymondKenney struct{}

func (r RaymondKenney) UseSkill() {
	fmt.Println("Рэймонд Кенни использует навык: Взлом банкомата!")
}

func main() {
	hackers := []SkillUser{AidenPearce{}, RaymondKenney{}}

	for _, hacker := range hackers {
		hacker.UseSkill()
	}
}
