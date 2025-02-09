package main

import (
	"context"
	"log"
	"time"
)

type Element string

const (
	Hydro   Element = "Hydro"
	Pyro    Element = "Pyro"
	Cryo    Element = "Cryo"
	Electro Element = "Electro"
	Anemo   Element = "Anemo"
	Geo     Element = "Geo"
	Dendro  Element = "Dendro"
)

type Character struct {
	Name    string
	Element Element
}

func (c *Character) ActivateElementalEffect(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second): // Симуляция действия эффекта стихии
		log.Printf("%s: %s elemental effect has ended.", c.Name, c.Element)
	case <-ctx.Done():
		log.Printf("%s: %s elemental effect was interrupted.", c.Name, c.Element)
	}
}

func main() {
	// Создаем контекст с тайм-аутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Создаем персонажей с разными стихиями
	characters := []Character{
		{"Amber", Pyro},
		{"Barbara", Hydro},
		{"Chongyun", Cryo},
		{"Electro", Electro},
		{"Venti", Anemo},
		{"Zhongli", Geo},
		{"Tighnari", Dendro},
	}

	// Для каждого персонажа активируем эффект стихии
	for _, character := range characters {
		go character.ActivateElementalEffect(ctx)
	}

	// Ожидаем завершения всех горутин
	time.Sleep(6 * time.Second)
}
