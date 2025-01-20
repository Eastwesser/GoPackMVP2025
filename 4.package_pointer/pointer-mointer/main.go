package main

import "fmt"

func pointerMointer(afloat *float64) *float64 {
	*afloat += 1

	return afloat
}

func main() {
	var a float64
	fmt.Println("Введите ваше число в консоль: ")
	fmt.Scanln(&a)
	fmt.Println(*pointerMointer(&a))
}

// Напиши функцию которая принимает указатель на float, увеличивает значение числа на 1,
// с помощью инкремента и возвращает указатель на указатель на это число. Вызови эту функцию, в Print разыменуй указатель
