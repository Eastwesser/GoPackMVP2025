package main

import (
	"fmt"
	"strings"
)

func main() {
	var bobuilder strings.Builder // Создаём Builder для эффективного построения строк.

	// Добавляем строки по очереди.
	bobuilder.WriteString("Go")                      // Добавляем "Go".
	bobuilder.WriteString(" ")                       // Добавляем пробел.
	bobuilder.WriteString("is awesome!")             // Добавляем "is awesome!".
	fmt.Println("Built string:", bobuilder.String()) // Выводим: Built string: Go is awesome!

	// Получаем текущую длину строки в Builder.
	fmt.Println("Length:", bobuilder.Len()) // Длина: 13.

	// Очищаем содержимое Builder.
	bobuilder.Reset()
	fmt.Println("After reset:", bobuilder.String()) // После очистки: пустая строка.

	// Добавляем форматированную строку.
	fmt.Fprintf(&bobuilder, "Formatted number: %d", 42)
	fmt.Println("Formatted string:", bobuilder.String()) // Выводим: Formatted string: Formatted number: 42.
}
