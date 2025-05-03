package main

import (
	"fmt"
	"sync"
)

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
	// Создаем семафор с буфером 2 (можно изменить на нужное количество горутин)
	sem := make(chan struct{}, 2)

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Список персонажей
	characters := []*Tribius{
		{
			HSRCharacter: HSRCharacter{
				id:          "1",
				name:        "Трибби",
				elementType: "квантовый",
				role:        "Гармония",
			},
			date: "26 февраля 2025 года",
		},
		{
			HSRCharacter: HSRCharacter{
				id:          "2",
				name:        "Светлячок",
				elementType: "огненный",
				role:        "Разрушение",
			},
			date: "19 июня 2024 года",
		},
		{
			HSRCharacter: HSRCharacter{
				id:          "3",
				name:        "Гепард",
				elementType: "ледяной",
				role:        "Сохранение",
			},
			date: "26 апреля 2023 года",
		},
	}

	// Запускаем горутины для каждого персонажа
	for _, char := range characters {
		wg.Add(1) // Увеличиваем счетчик WaitGroup
		go func(c *Tribius) {
			defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении горутины

			sem <- struct{}{}        // Захватываем место в семафоре
			defer func() { <-sem }() // Освобождаем место в семафоре

			var nameShower NameShower
			nameShower = c
			fmt.Println(nameShower.ShowName())
		}(char)
	}

	// Ждем завершения всех горутин
	wg.Wait()
}
