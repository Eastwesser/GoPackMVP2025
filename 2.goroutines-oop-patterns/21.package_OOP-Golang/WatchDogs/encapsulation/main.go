package main

import "fmt"

//Инкапсуляция позволяет скрыть внутренние детали объекта.
//Например, уровень хакера может быть скрыт, но доступен через метод.

type Hacker struct {
	name  string
	level int
}

func (h Hacker) GetLevel() int {
	return h.level
}

func (h *Hacker) LevelUp() {
	h.level++
	fmt.Println(h.name, "повысил уровень до", h.level)
}

func main() {
	aiden := Hacker{name: "Эйден Пирс", level: 10}
	fmt.Println("Уровень хакера", aiden.name, ":", aiden.GetLevel())

	aiden.LevelUp()
}
