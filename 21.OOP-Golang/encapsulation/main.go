package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func (p Person) GetAge() int {
	return p.age
}

func main() {
	person := Person{Name: "Иван", age: 25}

	fmt.Println("Имя:", person.Name)
	fmt.Println("Возраст:", person.GetAge())
}
