package main

import "fmt"

// Абстракция: интерфейс для персонажей
type Character interface {
	Attack()
	UseSkill()
}

// Базовый класс (Инкапсуляция)
type BaseCharacter struct {
	Name  string
	Level int
}

func (b BaseCharacter) GetLevel() int {
	return b.Level
}

func (b *BaseCharacter) LevelUp() {
	b.Level++
	fmt.Println(b.Name, "повысил уровень до", b.Level)
}

// Композиция: персонаж с оружием
type Weapon struct {
	Name  string
	Power int
}

// Наследование (встраивание) и полиморфизм
type PyroCharacter struct {
	BaseCharacter
	Weapon Weapon
}

func (p PyroCharacter) Attack() {
	fmt.Println(p.Name, "атакует с силой", p.Weapon.Power)
}

func (p PyroCharacter) UseSkill() {
	fmt.Println(p.Name, "использует пиро навык: Огненный взрыв!")
}

func main() {
	// Создаем персонажа
	diluc := PyroCharacter{
		BaseCharacter: BaseCharacter{Name: "Дилюк", Level: 80},
		Weapon:        Weapon{Name: "Волчья погибель", Power: 100},
	}

	// Используем методы
	fmt.Println("Уровень", diluc.Name, ":", diluc.GetLevel())
	diluc.Attack()
	diluc.UseSkill()
	diluc.LevelUp()
}
