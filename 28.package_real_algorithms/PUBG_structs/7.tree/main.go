package main

import "fmt"

/*
7. Деревья

Пример:
Иерархия команд (например, отряд игроков).

Объяснение:
Дерево позволяет организовать иерархию команд.
*/

type Player struct {
	name     string
	children []*Player
}

func (p *Player) addChild(child *Player) {
	p.children = append(p.children, child)
}

func main() {
	leader := &Player{name: "Leader"}
	member1 := &Player{name: "Member1"}
	member2 := &Player{name: "Member2"}
	leader.addChild(member1)
	leader.addChild(member2)
	fmt.Println("Team Hierarchy:")
	fmt.Println(leader.name, "->", leader.children[0].name, ",", leader.children[1].name)
}
