package main

import "fmt"

/*
3. Стек

Пример:
Система отмены действий (например, отмена последнего выстрела).

Объяснение:
Стек позволяет отменить последнее действие, например, если игрок случайно выстрелил.
*/

func main() {
	actions := []string{} // Стек действий
	actions = append(actions, "Move", "Shoot", "Reload")
	fmt.Println("Last Action:", actions[len(actions)-1])
	actions = actions[:len(actions)-1] // Отмена последнего действия
	fmt.Println("Actions after undo:", actions)
}
