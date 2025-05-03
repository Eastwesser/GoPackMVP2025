package main

import (
	"container/list"
	"fmt"
)

/*
2. Связанный список

Пример: История перемещений игрока.

Объяснение:
Связанный список позволяет хранить историю перемещений игрока,
чтобы можно было "откатиться" назад.
*/

func main() {
	history := list.New() // История перемещений
	history.PushBack("Location1")
	history.PushBack("Location2")
	history.PushBack("Location3")
	fmt.Println("Movement History:")
	for e := history.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
