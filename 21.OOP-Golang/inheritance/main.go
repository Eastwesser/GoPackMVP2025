package main

import "fmt"

type Animal struct {
	name string
}

func (a Animal) Speak() {
	fmt.Println("Я животное и умею издавать звуки.")
}

type Dog struct {
	Animal
	breed string
}

func (d Dog) Speak() {
	fmt.Println("Я собака и говорю 'Гав-гав!'")
}

func main() {
	animal := Animal{name: "Животное"}
	animal.Speak()

	dog := Dog{Animal: Animal{name: "Собака"}, breed: "Лабрадор"}
	dog.Speak()
}
