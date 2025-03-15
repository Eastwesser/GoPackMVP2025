package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	ID   int
	Name string
	Age  int
}

// Функция быстрой сортировки
func quicksort(slice []User, less func(a, b User) bool) {
	if len(slice) <= 1 {
		return
	}
	// Выбираем опорный элемент (pivot) pi = pivot index
	pivotIndex := partition(slice, less)
	// Рекурсивно сортируем левую и правую части
	quicksort(slice[:pivotIndex], less)
	quicksort(slice[pivotIndex+1:], less)
}

// Функция разделения (partition)
func partition(slice []User, less func(a, b User) bool) int {
	// Выбираем последний элемент как опорный
	pivot := slice[len(slice)-1]
	i := 0
	// Перемещаем элементы меньше опорного влево
	for j := 0; j < len(slice)-1; j++ {
		if less(slice[j], pivot) {
			slice[i], slice[j] = slice[j], slice[i]
			i++
		}
	}
	// Перемещаем опорный элемент на правильное место
	slice[i], slice[len(slice)-1] = slice[len(slice)-1], slice[i]
	return i
}

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Создаем слайс структур User
	users := []User{
		{3, "Nahida", 35},
		{1, "Raiden", 45},
		{2, "Zhongli", 57},
		{4, "Furina", 32},
		{5, "Venti", 20},
	}

	fmt.Println("Before sorting:", users)

	// Сортировка по ID
	quicksort(users, func(i, j User) bool {
		return i.ID < j.ID
	})
	fmt.Println("After sorting by ID:", users)

	// Сортировка по Age
	quicksort(users, func(i, j User) bool {
		return i.Age < j.Age
	})
	fmt.Println("After sorting by Age:", users)
}
