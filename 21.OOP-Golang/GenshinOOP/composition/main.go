package main

import "fmt"

//Композиция позволяет создавать сложные объекты, комбинируя более простые.
//Например, персонаж может иметь оружие и артефакты.

// Оружие
type Weapon struct {
	Name  string
	Power int
}

// Артефакт
type Artifact struct {
	Name  string
	Bonus string
}

// Персонаж
type Character struct {
	Name     string
	Weapon   Weapon
	Artifact Artifact
}

func main() {
	diluc := Character{
		Name: "Дилюк",
		Weapon: Weapon{
			Name:  "Волчья погибель",
			Power: 100,
		},
		Artifact: Artifact{
			Name:  "Пылающая алая ведьма",
			Bonus: "Увеличивает урон огнем",
		},
	}

	fmt.Println("Персонаж:", diluc.Name)
	fmt.Println("Оружие:", diluc.Weapon.Name, "с силой", diluc.Weapon.Power)
	fmt.Println("Артефакт:", diluc.Artifact.Name, "с бонусом:", diluc.Artifact.Bonus)
}
