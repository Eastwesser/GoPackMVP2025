package main

import "fmt"

type Player interface {
	Play()
}

type FootballPlayer struct {
	name string
}

func (f FootballPlayer) Play() {
	fmt.Println("Футболист", f.name, "играет в футбол.")
}

type BasketballPlayer struct {
	name string
}

func (b BasketballPlayer) Play() {
	fmt.Println("Баскетболист", b.name, "играет в баскетбол.")
}

func main() {
	players := []Player{
		FootballPlayer{name: "Роналду"},
		BasketballPlayer{name: "Джордан"},
	}

	for _, player := range players {
		player.Play()
	}
}
