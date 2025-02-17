package main

import "fmt"

// Абстракция: интерфейс для хакеров
type Hacker interface {
	Attack()
	UseSkill()
}

// Базовый класс (Инкапсуляция)
type BaseHacker struct {
	Name  string
	Level int
}

func (b BaseHacker) GetLevel() int {
	return b.Level
}

func (b *BaseHacker) LevelUp() {
	b.Level++
	fmt.Println(b.Name, "повысил уровень до", b.Level)
}

// Композиция: хакер с гаджетом
type Gadget struct {
	Name  string
	Power int
}

// Наследование (встраивание) и полиморфизм
type EliteHacker struct {
	BaseHacker
	Gadget Gadget
}

func (e EliteHacker) Attack() {
	fmt.Println(e.Name, "атакует с помощью", e.Gadget.Name, "и мощностью", e.Gadget.Power)
}

func (e EliteHacker) UseSkill() {
	fmt.Println(e.Name, "использует навык: Взлом инфраструктуры!")
}

func main() {
	// Создаем хакера
	aiden := EliteHacker{
		BaseHacker: BaseHacker{Name: "Эйден Пирс", Level: 10},
		Gadget:     Gadget{Name: "Смартфон", Power: 100},
	}

	// Используем методы
	fmt.Println("Уровень", aiden.Name, ":", aiden.GetLevel())
	aiden.Attack()
	aiden.UseSkill()
	aiden.LevelUp()
}

/*
	Абстракция: Интерфейс Hacker скрывает детали реализации.
	Композиция: Хакер содержит гаджет и оружие.

	Инкапсуляция: Уровень хакера скрыт, но доступен через метод.
	Наследование: EliteHacker встраивает BaseHacker.
	Полиморфизм: Разные хакеры могут использовать свои навыки через общий интерфейс
*/
