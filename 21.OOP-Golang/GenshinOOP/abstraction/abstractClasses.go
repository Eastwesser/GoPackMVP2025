package main

import "fmt"

// Абстракция: интерфейс Character
type CharacterAbstraction interface {
	Attack()
	GetStats() string
}

// #####################################################################################################################

// Структура для мечника
type Swordsman struct {
	Name        string
	AttackPower int
}

func (s *Swordsman) Attack() {
	fmt.Printf("%s атакует мечом с силой %d!\n", s.Name, s.AttackPower)
}

func (s *Swordsman) GetStats() string {
	return fmt.Sprintf("Мечник: %s, Сила атаки: %d", s.Name, s.AttackPower)
}

// #####################################################################################################################

// Структура для лучника
type Archer struct {
	Name        string
	AttackPower int
	ArrowCount  int
}

func (a *Archer) Attack() {
	fmt.Printf("%s стреляет из лука с силой %d! Осталось стрел: %d\n", a.Name, a.AttackPower, a.ArrowCount)
}

func (a *Archer) GetStats() string {
	return fmt.Sprintf("Лучник: %s, Сила атаки: %d, Стрелы: %d", a.Name, a.AttackPower, a.ArrowCount)
}

// #####################################################################################################################

// Структура для мага
type Mage struct {
	Name       string
	SpellPower int
	Mana       int
}

func (m *Mage) Attack() {
	fmt.Printf("%s произносит заклинание с силой %d! Осталось маны: %d\n", m.Name, m.SpellPower, m.Mana)
}

func (m *Mage) GetStats() string {
	return fmt.Sprintf("Маг: %s, Сила заклинания: %d, Мана: %d", m.Name, m.SpellPower, m.Mana)
}

// #####################################################################################################################

func main() {
	var character CharacterAbstraction

	character = &Swordsman{Name: "Diluc", AttackPower: 100}
	character.Attack()
	fmt.Println(character.GetStats())

	character = &Archer{Name: "Venti", AttackPower: 80, ArrowCount: 10}
	character.Attack()
	fmt.Println(character.GetStats())

	character = &Mage{Name: "Lisa", SpellPower: 120, Mana: 50}
	character.Attack()
	fmt.Println(character.GetStats())
}
