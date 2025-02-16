package main

import "fmt"

type Character struct {
	name  string
	level int
}

func (c Character) GetLevel() int {
	return c.level
}

func (c *Character) LevelUp() {
	c.level++
	fmt.Println(c.name, "повысил уровень до", c.level)
}

func main() {
	diluc := Character{name: "Дилюк", level: 80}
	fmt.Println("Уровень", diluc.name, ":", diluc.GetLevel())

	diluc.LevelUp()
}
