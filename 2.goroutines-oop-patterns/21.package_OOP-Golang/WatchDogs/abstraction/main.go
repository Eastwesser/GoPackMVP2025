package main

import "fmt"

//Абстракция позволяет скрыть сложность и показать только необходимые детали.
//В Watch Dogs это можно представить как общий интерфейс для всех хакеров, которые могут использовать свои навыки.

// Абстракция: интерфейс Hacker
type Hacker interface {
	UseSkill()
	UseUltimate()
}

// Структура для Aiden Pearce
type AidenPearce struct {
	Name string
}

func (a AidenPearce) UseSkill() {
	fmt.Println(a.Name, "использует навык: Взлом камеры!")
}

func (a AidenPearce) UseUltimate() {
	fmt.Println(a.Name, "использует ульту: Отключение электросети!")
}

// Структура для другого хакера
type RaymondKenney struct {
	Name string
}

func (r RaymondKenney) UseSkill() {
	fmt.Println(r.Name, "использует навык: Взлом банкомата!")
}

func (r RaymondKenney) UseUltimate() {
	fmt.Println(r.Name, "использует ульту: Массовая деактивация устройств!")
}

func main() {
	var hacker Hacker

	hacker = AidenPearce{Name: "Эйден Пирс"}
	hacker.UseSkill()
	hacker.UseUltimate()

	hacker = RaymondKenney{Name: "Рэймонд Кенни"}
	hacker.UseSkill()
	hacker.UseUltimate()
}
