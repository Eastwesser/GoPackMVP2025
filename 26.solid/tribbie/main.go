package main

import "fmt"

// есть персонаж honkai star rail, у него есть id, имя, стихия, роль
type HSRCharacter struct {
	id          string
	name        string
	elementType string
	role        string
}

// здесь мы лишь добавляем порядок их выхода 2025 год
type Tribius struct {
	HSRCharacter
	date string
}

// 	Tribbie string
//	Trianna string
//	Trionna string

type NameShower interface {
	ShowName() string
}

func (t *Tribius) ShowName() string {
	return fmt.Sprintf(
		"Привет! Я %s, и меня зовут %s! Я имею %s тип, и я следую пути %s. В игре с %s. Приятно познакомиться!",
		t.id,
		t.name,
		t.elementType,
		t.role,
		t.date,
	)
}

func main() {
	name := &Tribius{
		HSRCharacter: HSRCharacter{
			id:          "1",
			name:        "Трибби",
			elementType: "квантовый",
			role:        "Гармония",
		},
		date: "26 февраля 2025 года",
	}

	var nameShower NameShower
	nameShower = name
	fmt.Println(nameShower.ShowName())
}
